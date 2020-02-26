package proof

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/pgc/curve"
	"github.com/pgc/utils"
)

// RangeParams contains every param to generate/verify normal range proof.
// bitszie == len(gv) == len(hv)
type RangeParams interface {
	Bitsize() int
	Curve() elliptic.Curve
	GV() *utils.GeneratorVector
	HV() *utils.GeneratorVector
	U() *utils.ECPoint
	G() *utils.ECPoint
	H() *utils.ECPoint
}

type rangeParams struct {
	gv, hv  *utils.GeneratorVector
	u, g, h *utils.ECPoint
	bitsize int
}

// DRangeProofParams returns params for pgc/solidity system to genrate
// verify single range proofs.
func DRangeProofParams() RangeParams {
	curve := curve.BN256()
	bitsize := 32
	g := "g generator of twisted elg"
	gpoint := utils.NewECPointByString(g, curve)

	h := "h generator of twisted elg"
	hpoint := utils.NewECPointByString(h, curve)

	gv := utils.NewDefaultGV(curve, bitsize)
	hv := utils.NewDefaultHV(curve, bitsize)

	u := "u generator of innerproduct"
	upoint := utils.NewECPointByString(u, curve)

	return NewRangeParams(gv, hv, upoint, gpoint, hpoint, bitsize)
}

// NewRangeParams creates a new instance for single range proof.
func NewRangeParams(gv, hv *utils.GeneratorVector, u, g, h *utils.ECPoint, bitsize int) RangeParams {
	rp := rangeParams{}
	rp.hv = hv.Copy()
	rp.gv = gv.Copy()
	rp.u = u.Copy()
	rp.g = g.Copy()
	rp.h = h.Copy()
	rp.bitsize = bitsize

	return &rp
}

func newRandomRangeParams(curve elliptic.Curve, bitsize int) RangeParams {
	rp := rangeParams{}
	rp.gv = utils.NewRandomGeneratorVector(curve, bitsize)
	rp.hv = utils.NewRandomGeneratorVector(curve, bitsize)
	rp.u = utils.NewRandomECPoint(curve)
	rp.g = utils.NewRandomECPoint(curve)
	rp.h = utils.NewRandomECPoint(curve)
	rp.bitsize = bitsize

	return &rp
}

func (rp *rangeParams) Bitsize() int {
	return rp.bitsize
}

func (rp *rangeParams) Curve() elliptic.Curve {
	return rp.u.Curve
}

func (rp *rangeParams) GV() *utils.GeneratorVector {
	return rp.gv
}

func (rp *rangeParams) HV() *utils.GeneratorVector {
	return rp.hv
}

func (rp *rangeParams) U() *utils.ECPoint {
	return rp.u
}

func (rp *rangeParams) G() *utils.ECPoint {
	return rp.g
}

func (rp *rangeParams) H() *utils.ECPoint {
	return rp.h
}

// RangeProof contains everything to prove a v in certain range.
type RangeProof struct {
	A, S *utils.ECPoint

	T1, T2 *utils.ECPoint

	t, tx, u *big.Int

	ipProof *IPProof
}

type rangeProofInput struct {
	points [10]*big.Int
	scalar [5]*big.Int
	l, r   [10]*big.Int
	ll, rr []*big.Int
}

// ToSolidityInput formats data to solidity input to verify contract.
func (proof *RangeProof) ToSolidityInput() *rangeProofInput {
	input := rangeProofInput{}
	input.points[0] = proof.A.X
	input.points[1] = proof.A.Y
	input.points[2] = proof.S.X
	input.points[3] = proof.S.Y
	input.points[4] = proof.T1.X
	input.points[5] = proof.T1.Y
	input.points[6] = proof.T2.X
	input.points[7] = proof.T2.Y

	input.scalar[0] = proof.t
	input.scalar[1] = proof.tx
	input.scalar[2] = proof.u
	input.scalar[3] = proof.ipProof.a
	input.scalar[4] = proof.ipProof.b

	input.ll = make([]*big.Int, 2*len(proof.ipProof.l))
	input.rr = make([]*big.Int, 2*len(proof.ipProof.l))
	for i := 0; i < len(proof.ipProof.l); i++ {
		input.l[2*i] = proof.ipProof.l[i].X
		input.l[2*i+1] = proof.ipProof.l[i].Y
		input.ll[2*i] = proof.ipProof.l[i].X
		input.ll[2*i+1] = proof.ipProof.l[i].Y

		input.r[2*i] = proof.ipProof.r[i].X
		input.r[2*i+1] = proof.ipProof.r[i].Y
		input.rr[2*i] = proof.ipProof.r[i].X
		input.rr[2*i+1] = proof.ipProof.r[i].Y
	}

	return &input
}

// GenerateRangeProof generates proof to prove v in certain range without revealing it.
func GenerateRangeProof(params RangeParams, v, random *big.Int) (*RangeProof, error) {
	size := params.Bitsize()
	n := params.Curve().Params().N

	alVector, err := utils.BitVector(v, size)
	if err != nil {
		return nil, err
	}

	// <al, 2 ^ n> == v; ar == al - 1 ^ n.
	al := utils.NewFieldVector(alVector, n)
	ar := al.AllItemsSubOne()
	// pick a random number alpha.
	alpha, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	gv := params.GV()
	hv := params.HV()
	h := params.H()
	g := params.G()
	// compute commitment to al and ar.
	// commitA = g vector * al vector + h vector * ar vector + h point * alpha.
	commitA := utils.CommitTwoFieldVector(gv, hv, h, al, ar, alpha)

	// pick binding vector sl, sr.
	sl := utils.NewRandomFieldVector(n, size)
	sr := utils.NewRandomFieldVector(n, size)
	// pick another random number rho.
	rho, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	// computation same with commitA.
	commitB := utils.CommitTwoFieldVector(gv, hv, h, sl, sr, rho)

	// compute challenge y, z.
	y, err := utils.ComputeChallenge(n, commitA.X, commitA.Y, commitB.X, commitB.Y)
	if err != nil {
		return nil, err
	}
	z, err := utils.ComputeChallenge(n, commitB.X, commitB.Y, commitA.X, commitA.Y)
	if err != nil {
		return nil, err
	}

	// 2^n.
	n2 := utils.PowVector(new(big.Int).SetUint64(2), n, size)
	// y^n.
	yn := utils.PowVector(y, n, size)
	// sr hadamard y^n.
	sryn := sr.Hadamard(yn)

	// compute t0, t1, t2.
	zSquare := new(big.Int).Mul(z, z)
	zSquare.Mod(zSquare, n)

	// compute t0 for test.
	t0 := new(big.Int).Set(v)
	t0.Mul(t0, zSquare)
	t0.Mod(t0, n)
	delta := utils.Delta(y, z, n, size)
	t0.Add(t0, delta)

	// compute t1; t1 == <sl, y^n innerproduct (ar + z*1^n) + z*z*2^n> + <al - z * 1^n, sr hadamard y^n>.
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	// compute <al-z*1^n, sr hadamard y^n>.
	tmpField := al.AddFieldVector(utils.RepeatItemVector(zNeg, n, size))
	t1 := tmpField.InnerProduct(sryn)
	// compute <sl, y^n hadamard (ar + z * 1^n) + z^2 * 2^n>
	tmpField2 := ar.AddFieldVector(utils.RepeatItemVector(z, n, size))
	tmpField2 = tmpField2.Hadamard(yn)
	tmpField2 = tmpField2.AddFieldVector(n2.Times(zSquare))
	t1.Add(t1, sl.InnerProduct(tmpField2))
	t1.Mod(t1, n)

	// t2 = <sl, sr hadamard y^n>
	t2 := sl.InnerProduct(sryn)
	t2.Mod(t2, n)

	// commit to t1, t2.
	// pick two random number.
	r1, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	r2, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	T1 := new(utils.ECPoint).ScalarMult(h, r1)
	T1.Add(T1, new(utils.ECPoint).ScalarMult(g, t1))
	T2 := new(utils.ECPoint).ScalarMult(h, r2)
	T2.Add(T2, new(utils.ECPoint).ScalarMult(g, t2))

	// compute challenge x.
	x, err := utils.ComputeChallenge(n, T1.X, T1.Y, T2.X, T2.Y)
	if err != nil {
		return nil, err
	}
	x2 := new(big.Int).Exp(x, new(big.Int).SetUint64(2), n)

	// compute l, r...
	// l = al - z*1^n + sl*x.
	l := sl.Times(x)
	l = l.AddFieldVector(tmpField)
	// r = y^n hadamard (ar + z*1^n + sr*x) + z^2 * 2^n.
	// r = y^n hadamard (ar + z*1^n) + z^2 * 2^n + y^n hadamard sr*x.
	r := tmpField2.AddFieldVector(yn.Hadamard(sr.Times(x)))
	t := l.InnerProduct(r)
	t.Mod(t, n)

	// compute r2 * x^2 + r1 * x + z^2 * random.
	bindingX := new(big.Int).Mul(zSquare, random)
	bindingX.Mod(bindingX, n)
	bindingX.Add(bindingX, new(big.Int).Mul(r2, x2))
	bindingX.Mod(bindingX, n)
	bindingX.Add(bindingX, new(big.Int).Mul(r1, x))
	bindingX.Mod(bindingX, n)

	// alpha and rho blind A, S.
	u := new(big.Int).Mul(x, rho)
	u.Mod(u, n)
	u.Add(u, alpha)
	u.Mod(u, n)

	// compute new h generator vector; h' = h * (y ^ -n).
	hPrime := hv.Hadamard(yn.ModInverse().GetVector())
	// compute p point. p = A + S + gv*-z + h'*(z*y^n + z^2 * 2^n).
	p := commitA.Copy()
	p.Add(p, new(utils.ECPoint).ScalarMult(commitB, x))
	p.Add(p, new(utils.ECPoint).ScalarMult(gv.Sum(), zNeg))
	tmpExp := yn.Times(z).AddFieldVector(n2.Times(zSquare))
	p.Add(p, hPrime.Commit(tmpExp.GetVector()))

	// compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
	newP := p.Sub(p, new(utils.ECPoint).ScalarMult(h, u))

	ipparams := NewIPParams(gv, hPrime, params.U())
	ipProof, err := GenIPProof(ipparams, newP, t, l, r)
	if err != nil {
		return nil, err
	}

	proof := RangeProof{}
	proof.A = commitA
	proof.S = commitB
	proof.t = t
	proof.T1 = T1
	proof.T2 = T2
	proof.tx = bindingX
	proof.u = u
	proof.ipProof = ipProof
	return &proof, nil
}

// VerifyRangeProof validates a single range proof.
func VerifyRangeProof(params RangeParams, v *utils.ECPoint, proof *RangeProof) bool {
	n := params.Curve().Params().N
	size := params.Bitsize()
	y, err := utils.ComputeChallenge(n, proof.A.X, proof.A.Y, proof.S.X, proof.S.Y)
	if err != nil {
		log.Warn("compute challenge y failed", "error", err)
		return false
	}
	yn := utils.PowVector(y, n, size)

	z, err := utils.ComputeChallenge(n, proof.S.X, proof.S.Y, proof.A.X, proof.A.Y)
	if err != nil {
		log.Warn("compute challenge z failed", "error", err)
		return false
	}
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	zSquare := new(big.Int).Exp(z, new(big.Int).SetUint64(2), n)
	n2 := utils.PowVector(new(big.Int).SetUint64(2), n, size)

	hv := params.HV()
	gv := params.GV()
	hPrime := hv.Hadamard(yn.ModInverse().GetVector())

	x, err := utils.ComputeChallenge(n, proof.T1.X, proof.T1.Y, proof.T2.X, proof.T2.Y)
	if err != nil {
		log.Warn("compute challenge x failed", "error", err)
		return false
	}
	x2 := new(big.Int).Exp(x, new(big.Int).SetUint64(2), n)

	h := params.H()
	g := params.G()

	// check g*tx + h*t ?= v*z^2 + g*dleta + T1*x + T2*x^2.
	expect := new(utils.ECPoint).ScalarMult(v, zSquare)
	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T1, x))
	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T2, x2))
	dleta := utils.Delta(y, z, n, size)

	expect.Add(expect, new(utils.ECPoint).ScalarMult(g, dleta))
	actual := new(utils.ECPoint).ScalarMult(g, proof.t)
	actual.Add(actual, new(utils.ECPoint).ScalarMult(h, proof.tx))
	if !expect.Equal(actual) {
		log.Warn("point not equal", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	// compute p point. p = A + S*x + gv*-z + h'*(z*y^n + z^2 * 2^n).
	p := proof.A.Copy()
	p.Add(p, new(utils.ECPoint).ScalarMult(proof.S, x))
	p.Add(p, new(utils.ECPoint).ScalarMult(gv.Sum(), zNeg))
	tmpExp := yn.Times(z).AddFieldVector(n2.Times(zSquare))
	p.Add(p, hPrime.Commit(tmpExp.GetVector()))

	// compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
	newP := p.Sub(p, new(utils.ECPoint).ScalarMult(h, proof.u))

	return OptimizedVerifyIPProof(NewIPParams(gv, hPrime, params.U()), newP, proof.t, proof.ipProof)
}

// OptimizedVerifyRangeProof verify single range proof by optimized way.
func OptimizedVerifyRangeProof(params RangeParams, v *utils.ECPoint, proof *RangeProof) bool {
	n := params.Curve().Params().N
	size := params.Bitsize()
	y, err := utils.ComputeChallenge(n, proof.A.X, proof.A.Y, proof.S.X, proof.S.Y)
	if err != nil {
		log.Warn("compute challenge y failed", "error", err)
		return false
	}
	ynInverse := utils.PowVector(new(big.Int).ModInverse(y, n), n, size)

	z, err := utils.ComputeChallenge(n, proof.S.X, proof.S.Y, proof.A.X, proof.A.Y)
	if err != nil {
		log.Warn("compute challenge z failed", "error", err)
		return false
	}
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	zSquare := new(big.Int).Exp(z, utils.Two, n)

	x, err := utils.ComputeChallenge(n, proof.T1.X, proof.T1.Y, proof.T2.X, proof.T2.Y)
	if err != nil {
		log.Warn("compute challenge x failed", "error", err)
		return false
	}
	x2 := new(big.Int).Exp(x, utils.Two, n)

	h := params.H()
	g := params.G()

	// check g*tx + h*t ?= v*z^2 + g*dleta + T1*x + T2*x^2.
	// check g*(tx-dleta) + h*t ?= v*z^2 + T1*x + T2*x^2.
	expect := new(utils.ECPoint).ScalarMult(v, zSquare)
	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T1, x))
	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T2, x2))

	dleta := utils.Delta(y, z, n, size)
	txdleta := new(big.Int).Sub(proof.t, dleta)
	txdleta.Mod(txdleta, n)
	txdleta.Mod(txdleta, n)
	actual := new(utils.ECPoint).ScalarMult(g, txdleta)
	actual.Add(actual, new(utils.ECPoint).ScalarMult(h, proof.tx))
	if !expect.Equal(actual) {
		log.Warn("point not equal", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	right := new(utils.ECPoint).ScalarMult(proof.S, x)
	right.Add(right, proof.A)

	xj := make([]*big.Int, 0)
	xj2 := make([]*big.Int, 0)
	xj2Inv := make([]*big.Int, 0)

	for i := 0; i < len(proof.ipProof.l); i++ {
		l := proof.ipProof.l[i]
		r := proof.ipProof.r[i]
		tmpx, err := utils.ComputeChallenge(n, l.X, l.Y, r.X, r.Y)
		if err != nil {
			log.Warn("compute challenge for l, r failed", "err", err)
			return false
		}

		xj = append(xj, tmpx)
		tmpx2 := new(big.Int).Mul(tmpx, tmpx)
		tmpx2.Mod(tmpx2, n)
		xj2 = append(xj2, tmpx2)

		tmpx2Inv := new(big.Int).ModInverse(tmpx2, n)
		xj2Inv = append(xj2Inv, tmpx2Inv)

		tmpp := new(utils.ECPoint).ScalarMult(l, tmpx2)
		right.Add(right, tmpp)
		tmpp = new(utils.ECPoint).ScalarMult(r, tmpx2Inv)
		right.Add(right, tmpp)
	}

	// scalar mul, add.
	tl := make([]*big.Int, size)
	tr := make([]*big.Int, size)
	rl := make([]*big.Int, size)
	ll := make([]*big.Int, size)

	challengeLen := len(proof.ipProof.l)
	n2 := utils.PowVector(new(big.Int).SetUint64(2), n, size)
	for i := 0; i < size; i++ {
		if i == 0 {
			for j := 0; j < challengeLen; j++ {
				if j == 0 {
					tl[i] = new(big.Int).Set(xj[j])
				} else {
					tl[i].Mul(tl[i], xj[j])
					tl[i].Mod(tl[i], n)
				}

			}

			tr[i] = new(big.Int).Set(tl[i])
			tl[i] = tl[i].ModInverse(tl[i], n)
		} else {
			//todo:Computing scalars optimize.
			k := utils.GetBiggestPos(i, challengeLen) //;
			tl[i] = new(big.Int).Mul(tl[i-utils.Pow(k-1)], xj2[challengeLen-k])
			tl[i].Mod(tl[i], n)

			tr[i] = new(big.Int).Mul(tr[i-utils.Pow(k-1)], xj2Inv[challengeLen-k])
			tr[i].Mod(tr[i], n)
		}

		ll[i] = new(big.Int).Set(tl[i])
		rl[i] = new(big.Int).Set(tr[i])

		ll[i] = ll[i].Mul(ll[i], proof.ipProof.a)
		ll[i].Mod(ll[i], n)
		ll[i].Add(ll[i], z)
		ll[i].Mod(ll[i], n)

		rl[i] = rl[i].Mul(rl[i], proof.ipProof.b)
		rl[i].Mod(rl[i], n)

		z2n2 := new(big.Int).Mul(zSquare, n2.Get(i))
		z2n2.Mod(z2n2, n)
		rl[i].Sub(rl[i], z2n2)
		rl[i].Mod(rl[i], n)
		rl[i].Mul(rl[i], ynInverse.Get(i))
		rl[i].Mod(rl[i], n)
		rl[i].Sub(rl[i], z)
		rl[i].Mod(rl[i], n)
	}

	xu, err := utils.ComputeChallenge(n, proof.t)
	if err != nil {
		log.Warn("compute challenge for xu failed", "err", err)
		return false
	}

	gv := params.GV()
	hv := params.HV()

	left := gv.Commit(ll)
	left.Add(left, hv.Commit(rl))
	uBase := params.U()
	// xu(ab-t)
	xabt := new(big.Int).Mul(proof.ipProof.a, proof.ipProof.b)
	xabt.Mod(xabt, n)
	xabt.Sub(xabt, proof.t)
	xabt.Mod(xabt, n)
	xabt.Mul(xabt, xu)
	xabt.Mod(xabt, n)

	left.Add(left, new(utils.ECPoint).ScalarMult(uBase, xabt))
	left.Add(left, new(utils.ECPoint).ScalarMult(h, proof.u))

	return left.Equal(right)
}
