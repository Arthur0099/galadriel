package pgc

import (
	"crypto/rand"
	"encoding/json"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// RangeProver .
type RangeProver struct {
}

// RangeProof contains everything to prove a v in certain range.
type RangeProof struct {
	A, S *ECPoint

	T1, T2 *ECPoint

	t, tx, u *big.Int

	ipProof *IPProof
}

// MarshalJSON defines custom way to json.
func (rangeProof *RangeProof) MarshalJSON() ([]byte, error) {
	newJSON := struct {
		A          *ECPoint `json:"A"`
		S          *ECPoint `json:"S"`
		T1         *ECPoint `json:"T1"`
		T2         *ECPoint `json:"T2"`
		T          string   `json:"t"`
		TX         string   `json:"tx"`
		U          string   `json:"u"`
		IPProofIns *IPProof `json:"ipProof"`
	}{
		A:          rangeProof.A,
		S:          rangeProof.S,
		T1:         rangeProof.T1,
		T2:         rangeProof.T2,
		T:          rangeProof.t.String(),
		TX:         rangeProof.tx.String(),
		U:          rangeProof.u.String(),
		IPProofIns: rangeProof.ipProof,
	}

	return json.Marshal(&newJSON)
}

// GenerateRangeProof generates proof to prove v in certain range without revealing it.
func GenerateRangeProof(vb *VectorBase, v, random *big.Int) (*RangeProof, error) {
	size := vb.GetVectorSize()
	n := vb.GetCurve().Params().N

	alVector, err := BitVector(v, size)
	if err != nil {
		return nil, err
	}

	// <al, 2 ^ n> == v; ar == al - 1 ^ n.
	al := NewFieldVector(alVector, n)
	ar := al.AllItemsSubOne()
	// pick a random number alpha.
	alpha, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	// compute commitment to al and ar.
	// commitA = g vector * al vector + h vector * ar vector + h point * alpha.
	commitA := vb.CommitTwoFieldVector(al, ar, alpha)

	// pick binding vector sl, sr.
	sl := NewRandomFieldVector(n, size)
	sr := NewRandomFieldVector(n, size)
	// pick another random number rho.
	rho, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	// computation same with commitA.
	commitB := vb.CommitTwoFieldVector(sl, sr, rho)

	// compute challenge y, z.
	y, err := ComputeChallenge(n, commitA.X, commitA.Y, commitB.X, commitB.Y)
	if err != nil {
		return nil, err
	}
	z, err := ComputeChallenge(n, y)
	if err != nil {
		return nil, err
	}

	// 2^n.
	n2 := PowVector(new(big.Int).SetUint64(2), n, size)
	// y^n.
	yn := PowVector(y, n, size)
	// sr hadamard y^n.
	sryn := sr.Hadamard(yn)

	// compute t0, t1, t2.
	zSquare := new(big.Int).Mul(z, z)
	zSquare.Mod(zSquare, n)

	// compute t0 for test.
	t0 := new(big.Int).Set(v)
	t0.Mul(t0, zSquare)
	t0.Mod(t0, n)
	delta := Delta(y, z, n, size)
	t0.Add(t0, delta)

	// compute t1; t1 == <sl, y^n innerproduct (ar + z*1^n) + z*z*2^n> + <al - z * 1^n, sr hadamard y^n>.
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	// compute <al-z*1^n, sr hadamard y^n>.
	tmpField := al.AddFieldVector(RepeatItemVector(zNeg, n, size))
	t1 := tmpField.InnerProduct(sryn)
	// compute <sl, y^n hadamard (ar + z * 1^n) + z^2 * 2^n>
	tmpField2 := ar.AddFieldVector(RepeatItemVector(z, n, size))
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
	h := vb.GetH()
	g := vb.GetG()
	T1 := new(ECPoint).ScalarMult(h, r1)
	T1.Add(T1, new(ECPoint).ScalarMult(g, t1))
	T2 := new(ECPoint).ScalarMult(h, r2)
	T2.Add(T2, new(ECPoint).ScalarMult(g, t2))

	// compute challenge x.
	x, err := ComputeChallenge(n, T1.X, T1.Y, T2.X, T2.Y)
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

	// tt := new(big.Int).Set(t0)
	// tt1 := new(big.Int).Mul(t1, x)
	// tt1.Mod(tt1, n)
	// tt2 := new(big.Int).Mul(t2, x2)
	// tt2.Mod(tt2, n)
	// tt.Add(tt, tt1)
	// tt.Mod(tt, n)
	// tt.Add(tt, tt2)
	// tt.Mod(tt, n)
	// // This is just for check.
	// if t.Cmp(tt) != 0 {
	// 	return nil, errors.New("t0 + t1*x + t2*x^2 != <l,r>")
	// }

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
	hPrime := vb.GetHV().Hadamard(yn.ModInverse().GetVector())
	// compute p point. p = A + S + g*-z + h'*(z*y^n + z^2 * 2^n).
	// p := commitA.Copy()
	// p.Add(p, new(ECPoint).ScalarMult(commitB, x))
	// p.Add(p, new(ECPoint).ScalarMult(vb.GetGV().Sum(), zNeg))
	// tmpExp := yn.Times(z).AddFieldVector(n2.Times(zSquare))
	// p.Add(p, hPrime.Commit(tmpExp.GetVector()))

	// // compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
	// newP := p.Sub(p, new(ECPoint).ScalarMult(vb.GetH(), u))
	// just for check p' == g*l + h'*r.
	actual := vb.GetGV().Commit(l.GetVector())
	tmphr := hPrime.Commit(r.GetVector())
	actual.Add(actual, tmphr)
	// if !newP.Equal(actual) {
	// 	return nil, errors.New("newp not equal g*l+h'*r")
	// }

	ipProver := IPProver{}
	ipProver.U = vb.GetU()
	ipProof, err := ipProver.GenerateIPProof(vb.GetGV(), hPrime, actual, t, l, r)
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

// VerifyRangeProof validates a range proof.
func VerifyRangeProof(vb *VectorBase, v *ECPoint, proof *RangeProof) bool {
	curve := vb.GetCurve()
	n := curve.Params().N
	size := vb.GetVectorSize()
	y, err := ComputeChallenge(n, proof.A.X, proof.A.Y, proof.S.X, proof.S.Y)
	if err != nil {
		log.Warn("compute challenge y failed", "error", err)
		return false
	}
	yn := PowVector(y, n, size)

	z, err := ComputeChallenge(n, y)
	if err != nil {
		log.Warn("compute challenge z failed", "error", err)
		return false
	}
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	zSquare := new(big.Int).Exp(z, new(big.Int).SetUint64(2), n)
	n2 := PowVector(new(big.Int).SetUint64(2), n, size)

	hv := vb.GetHV()
	hPrime := hv.Hadamard(yn.ModInverse().GetVector())

	x, err := ComputeChallenge(n, proof.T1.X, proof.T1.Y, proof.T2.X, proof.T2.Y)
	if err != nil {
		log.Warn("compute challenge x failed", "error", err)
		return false
	}
	x2 := new(big.Int).Exp(x, new(big.Int).SetUint64(2), n)

	h := vb.GetH()
	g := vb.GetG()

	// check g*tx + h*t ?= v*z^2 + g*dleta + T1*x + T2*x^2.
	expect := new(ECPoint).ScalarMult(v, zSquare)
	expect.Add(expect, new(ECPoint).ScalarMult(proof.T1, x))
	expect.Add(expect, new(ECPoint).ScalarMult(proof.T2, x2))
	dleta := Delta(y, z, n, size)

	expect.Add(expect, new(ECPoint).ScalarMult(g, dleta))
	actual := new(ECPoint).ScalarMult(g, proof.t)
	actual.Add(actual, new(ECPoint).ScalarMult(h, proof.tx))
	if !expect.Equal(actual) {
		log.Warn("point not equal", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	// compute p point. p = A + S*x + gv*-z + h'*(z*y^n + z^2 * 2^n).
	p := proof.A.Copy()
	p.Add(p, new(ECPoint).ScalarMult(proof.S, x))
	p.Add(p, new(ECPoint).ScalarMult(vb.GetGV().Sum(), zNeg))
	tmpExp := yn.Times(z).AddFieldVector(n2.Times(zSquare))
	p.Add(p, hPrime.Commit(tmpExp.GetVector()))

	// compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
	newP := p.Sub(p, new(ECPoint).ScalarMult(vb.GetH(), proof.u))

	ipVerifier := IPVerifier{}
	ipVerifier.U = vb.GetU()

	return ipVerifier.VerifyIPProof(vb.GetGV(), hPrime, newP, proof.t, proof.ipProof)
}
