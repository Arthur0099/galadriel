package pgc

import (
	"encoding/json"
	"errors"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// IPProof represents inner product proof.
type IPProof struct {
	L, R []*ECPoint

	a, b *big.Int
}

// MarshalJSON defines custom way to json.
func (ipProof *IPProof) MarshalJSON() ([]byte, error) {
	newJSON := struct {
		L []*ECPoint `json:"l"`
		R []*ECPoint `json:"r"`
		A string     `json:"a"`
		B string     `json:"b"`
	}{
		L: ipProof.L,
		R: ipProof.R,
		A: ipProof.a.String(),
		B: ipProof.b.String(),
	}

	return json.Marshal(&newJSON)
}

// NewIPProof creates instance of inner product proof.
func NewIPProof(l, r []*ECPoint, a, b *big.Int) *IPProof {
	proof := IPProof{}
	proof.L = l
	proof.R = r
	proof.a = new(big.Int).Set(a)
	proof.b = new(big.Int).Set(b)

	return &proof
}

// IPProver .
type IPProver struct {
	// fix point in protocol.
	U *ECPoint
}

// GenerateIPProof generates proof using protocol 1 in bulletproof inner-product argument to prove that
// prover knows two vector a, b and p = g * a + h *b; and c = <a, b>;
// the g, h, p, h is public known by verifier.
// g and h are generator vectors.
func (ipProver *IPProver) GenerateIPProof(g, h *GeneratorVector, p *ECPoint, c *big.Int, a, b *FieldVector) (*IPProof, error) {
	if !(g.Size() == h.Size() && h.Size() == a.Size() && a.Size() == b.Size()) {
		return nil, errors.New("g, h, a, b size not equal")
	}

	// pick challenge x.
	curve := p.Curve
	n := curve.Params().N
	// todo:
	e, err := ComputeChallenge(n, c)
	if err != nil {
		return nil, err
	}
	// compute newPoint.
	// p' = p + u * (e * c)
	ue := new(ECPoint).ScalarMult(ipProver.U, e)
	newP := new(ECPoint).ScalarMult(ue, c)
	newP.Add(newP, p)

	l, r := make([]*ECPoint, 0), make([]*ECPoint, 0)

	// call protocol 2.
	return ipProver.generateIPProof(g, h, ue, newP, a, b, l, r)
}

// generateIPProof generates proof for inner product proof.
// g, h are two public vector generator used in bullet proof.
// u = u * e(u represents a fix point in protocol);
// p = g * a + h * b + u * c.
func (ipProver *IPProver) generateIPProof(g, h *GeneratorVector, u, p *ECPoint, a, b *FieldVector, l, r []*ECPoint) (*IPProof, error) {
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

	lp := new(ECPoint).ScalarMult(u, cL)
	lp.Add(lp, l1)
	lp.Add(lp, l2)
	l = append(l, lp)

	// compute R = gLeft * aRight + hRight * bLeft + u * cR.
	r1 := gLeft.Commit(aRight.GetVector())
	r2 := hRight.Commit(bLeft.GetVector())
	rp := new(ECPoint).ScalarMult(u, cR)
	rp.Add(rp, r1)
	rp.Add(rp, r2)
	r = append(r, rp)

	// compute challenge x base on l, r.
	// todo: may be change the challenge compution.
	n := curve.Params().N
	e, err := ComputeChallenge(n, lp.X, lp.Y, rp.X, rp.Y)
	if err != nil {
		return nil, err
	}
	eInverse := new(big.Int).ModInverse(e, n)

	gPrime := gLeft.HadamardScalar(eInverse).AddGeneratorVector(gRight.HadamardScalar(e))
	hPrime := hLeft.HadamardScalar(e).AddGeneratorVector(hRight.HadamardScalar(eInverse))

	aPrime := aLeft.Times(e).AddFieldVector(aRight.Times(eInverse))
	bPrime := bLeft.Times(eInverse).AddFieldVector(bRight.Times(e))

	// compute e ^ 2.
	eSquare := new(big.Int).Mul(e, e)
	eSquare.Mod(eSquare, n)
	// compute e ^ -2 = eInverse ^ 2.
	eInverseSquare := new(big.Int).Mul(eInverse, eInverse)
	eInverseSquare.Mod(eInverseSquare, n)

	// compute p' = l * (x ^ 2) + p + r * (x ^ -2).
	newP := new(ECPoint).ScalarMult(lp, eSquare)
	newP.Add(newP, p)
	rTmp := new(ECPoint).ScalarMult(rp, eInverseSquare)
	newP.Add(newP, rTmp)

	return ipProver.generateIPProof(gPrime, hPrime, u, newP, aPrime, bPrime, l, r)
}

// IPVerifier .
type IPVerifier struct {
	// fix point in protocol.
	U *ECPoint
}

// VerifyIPProof validates inner product proof.
func (ipVerifier *IPVerifier) VerifyIPProof(g, h *GeneratorVector, p *ECPoint, c *big.Int, proof *IPProof) bool {
	curve := p.Curve
	n := curve.Params().N
	// todo:
	e, err := ComputeChallenge(n, c)
	if err != nil {
		log.Warn("IPVerifier compute challenge failed in protocol 1", "error", err)
		return false
	}

	// compute newPoint.
	// p' = p + u * (e * c)
	ue := new(ECPoint).ScalarMult(ipVerifier.U, e)
	np := new(ECPoint).ScalarMult(ue, c)
	np.Add(np, p)

	return ipVerifier.verifyIPProof(g, h, ue, np, proof)
}

func (ipVerifier *IPVerifier) optVerifyIPProof(g, h *GeneratorVector, p *ECPoint, c *big.Int, proof *IPProof) bool {
	curve := p.Curve
	n := curve.Params().N
	// todo:
	e, err := ComputeChallenge(n, c)
	if err != nil {
		log.Warn("IPVerifier compute challenge failed in protocol 1", "error", err)
		return false
	}

	// compute newPoint.
	// p' = p + u * (e * c)
	ue := new(ECPoint).ScalarMult(ipVerifier.U, e)
	np := new(ECPoint).ScalarMult(ue, c)
	np.Add(np, p)

	return ipVerifier.optmizedVerifyIPProof(g, h, ue, np, proof)
}

// verifyIPProof validates inner product proof.
func (ipVerifier *IPVerifier) verifyIPProof(g, h *GeneratorVector, u, p *ECPoint, proof *IPProof) bool {
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
			log.Warn("IPVerifier compute challenge failed in protocol 2", "error", err)
			return false
		}
		eInverse := new(big.Int).ModInverse(e, n)

		gPrime := gLeft.HadamardScalar(eInverse).AddGeneratorVector(gRight.HadamardScalar(e))
		hPrime := hLeft.HadamardScalar(e).AddGeneratorVector(hRight.HadamardScalar(eInverse))

		// Compute e ^ 2.
		eSquare := new(big.Int).Mul(e, e)
		eSquare.Mod(eSquare, n)
		// compute e ^ -2 = eInverse ^ 2.
		eInverseSquare := new(big.Int).Mul(eInverse, eInverse)
		eInverseSquare.Mod(eInverseSquare, n)

		// update params.
		np := new(ECPoint).ScalarMult(l, eSquare)
		np.Add(np, p)
		rTmp := new(ECPoint).ScalarMult(r, eInverseSquare)
		np.Add(np, rTmp)

		// set new params.
		p = np.Copy()
		g = gPrime
		h = hPrime
	}

	if g.Size() != 1 {
		log.Warn("IPVerifier g generator size != 1")
		return false
	}

	if h.Size() != 1 {
		log.Warn("IPVerifier h generator size != 1")
		return false
	}

	c := new(big.Int).Mul(proof.a, proof.b)
	c.Mod(c, n)

	// compute u * c.
	want := new(ECPoint).ScalarMult(u, c)
	// compute g * a.
	ga := new(ECPoint).ScalarMult(g.Get(0), proof.a)
	// compute h * b.
	hb := new(ECPoint).ScalarMult(h.Get(0), proof.b)
	// compute g * a + h * b + u * c.
	want.Add(want, ga)
	want.Add(want, hb)

	if !p.Equal(want) {
		log.Warn("Verifier p != p1", "want x", p.X, "want y", p.Y, "actual x", want.X, "actual y", want.Y)
		return false
	}

	return true
}

//
func (ipVerifier *IPVerifier) optmizedVerifyIPProof(g, h *GeneratorVector, u, p *ECPoint, proof *IPProof) bool {
	curve := u.Curve
	n := curve.Params().N

	right := p.Copy()
	xjs := make([]*big.Int, 0)
	xjsInv := make([]*big.Int, 0)
	for i := 0; i < len(proof.L); i++ {
		xj, err := ComputeChallenge(n, proof.L[i].X, proof.L[i].Y, proof.R[i].X, proof.R[i].Y)
		if err != nil {
			log.Warn("compute challenge for optimize inner product failed", "err", err)
			return false
		}
		xjInv := new(big.Int).ModInverse(xj, n)
		xjs = append(xjs, xj)
		xjsInv = append(xjsInv, xjInv)

		xj2 := new(big.Int).Mul(xj, xj)
		xj2.Mod(xj2, n)

		xj2Inv := new(big.Int).Mul(xjInv, xjInv)
		xj2Inv.Mod(xj2Inv, n)

		right.Add(right, new(ECPoint).ScalarMult(proof.L[i], xj2))
		right.Add(right, new(ECPoint).ScalarMult(proof.R[i], xj2Inv))
	}

	left := u.Copy()
	ab := new(big.Int).Mul(proof.a, proof.b)
	ab.Mod(ab, n)
	left.ScalarMult(left, ab)

	s := make([]*big.Int, g.Size())

	for i := 0; i < g.Size(); i++ {
		for j := 0; j < len(proof.L); j++ {
			tmp := new(big.Int)
			if smallParseBinary(i, j, len(proof.L)) {
				tmp.Set(xjs[j])
			} else {
				tmp.Set(xjsInv[j])
			}

			if j == 0 {
				s[i] = tmp
			} else {
				s[i].Mul(s[i], tmp)
				s[i].Mod(s[i], n)
			}
		}
	}

	as := newFieldVector(s, n).Times(proof.a).GetVector()
	left.Add(left, g.Commit(as))
	bsinv := newFieldVector(s, n).ModInverse().Times(proof.b).GetVector()
	left.Add(left, h.Commit(bsinv))

	return left.Equal(right)
}
