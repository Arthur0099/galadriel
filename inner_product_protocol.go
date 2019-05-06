package pgc

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
)

// IPProof represents inner product proof.
type IPProof struct {
	L, R []*ecdsa.PublicKey

	a, b *big.Int
}

// NewIPProof creates instance of inner product proof.
func NewIPProof(l, r []*ecdsa.PublicKey, a, b *big.Int) *IPProof {
	proof := IPProof{}
	proof.L = l
	proof.R = r
	proof.a = a
	proof.b = b

	return &proof
}

// IPProver .
type IPProver struct {
	// fix point in protocol.
	U *ecdsa.PublicKey
}

// GenerateIPProof generates proof using protocol 1 in bulletproof inner-product argument.
// p = g * a + h * b; c = <a, b>
func (ipProver *IPProver) GenerateIPProof(g, h *GeneratorVector, p *ecdsa.PublicKey, c *big.Int, a, b *FieldVector) (*IPProof, error) {
	if !(g.Size() == h.Size() && h.Size() == a.Size() && a.Size() == b.Size()) {
		return nil, errors.New("g, h, a, b size not equal")
	}
	// pick challenge x.
	curve := p.Curve
	n := curve.Params().N
	// todo:
	e, err := ComputeChallenge(n, p.X, p.Y, c)
	if err != nil {
		return nil, err
	}

	// compute newPoint.
	// p' = p + u * (e * c)
	ueX, ueY := curve.ScalarMult(ipProver.U.X, ipProver.U.Y, e.Bytes())
	uecX, uecY := curve.ScalarMult(ueX, ueY, c.Bytes())
	newPointX, newPointY := curve.Add(p.X, p.Y, uecX, uecY)

	l, r := make([]*ecdsa.PublicKey, 0), make([]*ecdsa.PublicKey, 0)

	// call protocol 2.
	return ipProver.generateIPProof(g, h, PointToPublicKey(ueX, ueY, curve), PointToPublicKey(newPointX, newPointY, curve), a, b, l, r)
}

// generateIPProof generates proof for inner product proof.
// g, h are two public vector generator used in bullet proof.
// u = u * e(u represents a fix point in protocol);
// p = g * a + h * b + u * c.
func (ipProver *IPProver) generateIPProof(g, h *GeneratorVector, u, p *ecdsa.PublicKey, a, b *FieldVector, l, r []*ecdsa.PublicKey) (*IPProof, error) {
	if g.Size() == 1 {
		return NewIPProof(l, r, a.First(), b.First()), nil
	}

	// todo: make sure size is power of 2.

	gLeft := g.HalfLeft()
	gRight := g.HalfRight()
	hLeft := h.HalfLeft()
	hRight := h.HalfRight()

	aLeft := a.HalfLeft()
	aRight := a.HalfRight()
	bLeft := b.HalfLeft()
	bRight := b.HalfRight()

	cL := aLeft.InnerProduct(bRight)
	cR := aRight.InnerProduct(bLeft)

	// compute L = gRight * aLeft + hLeft * bRight + u * cL.
	curve := u.Curve
	l1 := gRight.Commit(aLeft.GetVector())
	l2 := hLeft.Commit(bRight.GetVector())
	lx, ly := curve.ScalarMult(u.X, u.Y, cL.Bytes())
	lx, ly = curve.Add(l1.X, l1.Y, lx, ly)
	lx, ly = curve.Add(l2.X, l2.Y, lx, ly)
	l = append(l, PointToPublicKey(lx, ly, curve))

	// compute R = gLeft * aRight + hRight * bLeft + u * cR.
	r1 := gLeft.Commit(aRight.GetVector())
	r2 := hRight.Commit(bLeft.GetVector())
	rx, ry := curve.ScalarMult(u.X, u.Y, cR.Bytes())
	rx, ry = curve.Add(r1.X, r1.Y, rx, ry)
	rx, ry = curve.Add(r2.X, r2.Y, rx, ry)
	r = append(r, PointToPublicKey(rx, ry, curve))

	// compute challenge x base on l, r.
	// todo: may be change the challenge compution.
	n := curve.Params().N
	e, err := ComputeChallenge(curve.Params().N, lx, ly, rx, ry)
	if err != nil {
		return nil, err
	}
	eInverse := new(big.Int).ModInverse(e, n)

	gPrime := gLeft.Hadamard(eInverse).AddGeneratorVector(gRight.Hadamard(e))
	hPrime := hLeft.Hadamard(e).AddGeneratorVector(hRight.Hadamard(eInverse))

	aPrime := aLeft.Times(e).AddFieldVector(aRight.Times(eInverse))
	bPrime := bLeft.Times(eInverse).AddFieldVector(bRight.Times(e))

	// compute e ^ 2.
	eSquare := new(big.Int).Mul(e, e)
	eSquare.Mod(eSquare, n)
	// compute e ^ -2 = eInverse ^ 2.
	eInverseSquare := new(big.Int).Mul(eInverse, eInverse)
	eInverseSquare.Mod(eInverseSquare, n)

	// compute p' = l * (x ^ 2) + p + r * (x ^ -2).
	pPrimeX, pPrimeY := curve.ScalarMult(lx, ly, eSquare.Bytes())
	pPrimeX, pPrimeY = curve.Add(p.X, p.Y, pPrimeX, pPrimeY)
	rTmpX, rTmpY := curve.ScalarMult(rx, ry, eInverseSquare.Bytes())
	pPrimeX, pPrimeY = curve.Add(rTmpX, rTmpY, pPrimeX, pPrimeY)
	pNew := PointToPublicKey(pPrimeX, pPrimeY, curve)

	return ipProver.generateIPProof(gPrime, hPrime, u, pNew, aPrime, bPrime, l, r)
}

// IPVerifier .
type IPVerifier struct {
	// fix point in protocol.
	U *ecdsa.PublicKey
}

// VerifyIPProof validates inner product proof.
func (ipVerifier *IPVerifier) VerifyIPProof(g, h *GeneratorVector, p *ecdsa.PublicKey, c *big.Int, proof *IPProof) bool {
	curve := p.Curve
	n := curve.Params().N
	// todo:
	e, err := ComputeChallenge(n, p.X, p.Y, c)
	if err != nil {
		log.Printf("compute challenge for ipproof failed")
		return false
	}

	// compute newPoint.
	// p' = p + u * (e * c)
	ueX, ueY := curve.ScalarMult(ipVerifier.U.X, ipVerifier.U.Y, e.Bytes())
	uecX, uecY := curve.ScalarMult(ueX, ueY, c.Bytes())
	newPointX, newPointY := curve.Add(p.X, p.Y, uecX, uecY)

	return ipVerifier.verifyIPProof(g, h, PointToPublicKey(ueX, ueY, curve), PointToPublicKey(newPointX, newPointY, curve), proof)
}

// verifyIPProof validates inner product proof.
func (ipVerifier *IPVerifier) verifyIPProof(g, h *GeneratorVector, u, p *ecdsa.PublicKey, proof *IPProof) bool {
	curve := u.Curve
	n := curve.Params().N

	for i, l := range proof.L {
		gLeft := g.HalfLeft()
		gRight := g.HalfRight()
		hLeft := h.HalfLeft()
		hRight := h.HalfRight()

		r := proof.R[i]

		e, err := ComputeChallenge(n, l.X, l.Y, r.X, r.Y)
		if err != nil {
			log.Printf("compute challenge for ipproof failed")
			return false
		}

		eInverse := new(big.Int).ModInverse(e, n)

		gPrime := gLeft.Hadamard(eInverse).AddGeneratorVector(gRight.Hadamard(e))
		hPrime := hLeft.Hadamard(e).AddGeneratorVector(hRight.Hadamard(eInverse))

		// Compute e ^ 2.
		eSquare := new(big.Int).Mul(e, e)
		eSquare.Mod(eSquare, n)
		// compute e ^ -2 = eInverse ^ 2.
		eInverseSquare := new(big.Int).Mul(eInverse, eInverse)
		eInverseSquare.Mod(eInverseSquare, n)

		pPrimeX, pPrimeY := curve.ScalarMult(l.X, l.Y, eSquare.Bytes())
		pPrimeX, pPrimeY = curve.Add(p.X, p.Y, pPrimeX, pPrimeY)
		rTmpX, rTmpY := curve.ScalarMult(r.X, r.Y, eInverseSquare.Bytes())
		pPrimeX, pPrimeY = curve.Add(rTmpX, rTmpY, pPrimeX, pPrimeY)

		// set new params.
		p = PointToPublicKey(pPrimeX, pPrimeY, curve)
		g = gPrime
		h = hPrime
	}

	if g.Size() != 1 {
		log.Printf("g generator size wrong != 1")
		return false
	}

	if h.Size() != 1 {
		log.Printf("h generator size wrong != 1")
		return false
	}

	c := new(big.Int).Mul(proof.a, proof.b)
	c.Mod(c, n)

	wantX, wantY := curve.ScalarMult(u.X, u.Y, c.Bytes())
	// compute g * a.
	gFirst := g.Get(0)
	gaX, gaY := curve.ScalarMult(gFirst.X, gFirst.Y, proof.a.Bytes())
	// compute h * b.
	hFirst := h.Get(0)
	haX, haY := curve.ScalarMult(hFirst.X, hFirst.Y, proof.b.Bytes())
	// compute g * a + h * b + u * c.
	tmpX, tmpY := curve.Add(gaX, gaY, haX, haY)
	tmpX, tmpY = curve.Add(tmpX, tmpY, wantX, wantY)
	if p.X.Cmp(tmpX) != 0 || p.Y.Cmp(tmpY) != 0 {
		log.Printf(fmt.Sprintf("p point not equal, want x %x, y %x, actual x %x, y %x", p.X, p.Y, tmpX, tmpY))
		return false
	}

	return true
}
