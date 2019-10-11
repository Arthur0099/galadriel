package pgc

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	log "github.com/inconshreveable/log15"
)

type AggreateBulletProof struct {
	RangeProof
}

// GenerateAggreateBulletProof aggreates many(2 current) bullet proof together to reduce proof size.
func GenerateAggreateBulletProof(vb *VectorBase, v, random []*big.Int) (*AggreateBulletProof, error) {
	bitSize := vb.GetBitSize()
	aggreateSize := vb.GetAggreateSize()
	vectorSize := bitSize * aggreateSize
	n := vb.GetCurve().Params().N

	if len(v) != aggreateSize {
		return nil, fmt.Errorf("witness v len %d not equal aggreate size %d", len(v), aggreateSize)
	}

	if len(random) != aggreateSize {
		return nil, fmt.Errorf("witness r len %d not equal aggreate size %d", len(random), aggreateSize)
	}

	// todo: add more assert.
	alVector, err := MultBitVector(v, bitSize)
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
	sl := NewRandomFieldVector(n, vectorSize)
	sr := NewRandomFieldVector(n, vectorSize)
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

	// 2^mn.
	mn2 := PowVector(two, n, vectorSize)
	// y^mn.
	ymn := PowVector(y, n, vectorSize)
	// sr hadamard y^mn.
	rr1 := sr.Hadamard(ymn)
	// compute y^mn (ar + z * 1^mn) + z^(1+j) * (0^(j-1)*n || 2^n || 0^(m-j)*n). (j=[1, m])
	rr0 := ar.AddFieldVector(RepeatItemVector(z, n, vectorSize))
	rr0 = rr0.Hadamard(ymn)
	n2 := mn2.SubFieldVector(0, bitSize)
	for j := 1; j <= aggreateSize; j++ {
		// 0^((j-1)*n)
		tmpz := RepeatItemVector(zero, n, bitSize*(j-1))
		// 2^n
		tmpz = tmpz.Append(n2)
		// 0^((m-j)*n)
		tmpz = tmpz.Append(RepeatItemVector(zero, n, (aggreateSize-j)*bitSize))
		if tmpz.Size() != vectorSize {
			return nil, fmt.Errorf("tmp z calcualte failed expect len %d, actual len %d", vectorSize, tmpz.Size())
		}

		// z^(1+j)
		j1 := new(big.Int).SetUint64(uint64(1 + j))
		zj := new(big.Int).Exp(z, j1, n)
		rr0 = rr0.AddFieldVector(tmpz.Times(zj))
	}

	// compute t0, t1, t2.
	zSquare := new(big.Int).Mul(z, z)
	zSquare.Mod(zSquare, n)

	// <lx, rx> = (ll0 + ll1*X) * (rr0 + rr1*X)
	// => t0=<ll0, rr0>; t1 = <ll0, rr1> + <ll1, rro>; t2 = <ll1, rr1>
	// ll0 = al - z*1^mn;
	// ll1 = sl
	// rr0 = y^mn (ar + z*1^mn) + ...
	// rr1 = y^mn (sr)

	// compute t0(for check only).
	// t0 = <ll0, rr0>
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	ll0 := al.AddFieldVector(RepeatItemVector(zNeg, n, vectorSize))
	t0 := ll0.InnerProduct(rr0)
	t0.Mod(t0, n)

	// compute t1.
	// t1 = <ll0, rr1> + <ll1, rro>;
	t1 := ll0.InnerProduct(rr1)
	t1.Add(t1, sl.InnerProduct(rr0))
	t1.Mod(t1, n)

	// t2 = <sl, sr hadamard y^mn>
	// t2 = <ll1, rr1>
	t2 := sl.InnerProduct(rr1)
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
	x2 := new(big.Int).Exp(x, two, n)

	// compute l, r...
	// l = al - z*1^mn + sl*x.
	l := sl.Times(x)
	l = l.AddFieldVector(ll0)
	// r.
	r := rr0.AddFieldVector(rr1.Times(x))
	t := l.InnerProduct(r)
	t.Mod(t, n)

	// just for test check.
	if vb.FakeTest() {
		tt := new(big.Int).Set(t0)
		tt1 := new(big.Int).Mul(t1, x)
		tt1.Mod(tt1, n)
		tt2 := new(big.Int).Mul(t2, x2)
		tt2.Mod(tt2, n)
		tt.Add(tt, tt1)
		tt.Mod(tt, n)
		tt.Add(tt, tt2)
		tt.Mod(tt, n)
		// This is just for check.
		if t.Cmp(tt) != 0 {
			return nil, errors.New("t0 + t1*x + t2*x^2 != <lx,rx>")
		}
	}

	// compute r2 * x^2 + r1 * x + ....
	bindingX := new(big.Int).Mul(r2, x2)
	bindingX.Mod(bindingX, n)
	bindingX.Add(bindingX, new(big.Int).Mul(r1, x))
	bindingX.Mod(bindingX, n)
	for j := 1; j <= aggreateSize; j++ {
		j1 := new(big.Int).SetUint64(uint64(1 + j))
		zj := new(big.Int).Exp(z, j1, n)
		zj.Mul(zj, random[j-1])
		bindingX.Add(bindingX, zj)
	}
	bindingX.Mod(bindingX, n)

	// alpha and rho blind A, S.
	u := new(big.Int).Mul(x, rho)
	u.Mod(u, n)
	u.Add(u, alpha)
	u.Mod(u, n)

	// compute new h generator vector; h' = h * (y ^ -mn).
	hPrime := vb.GetHV().Hadamard(ymn.ModInverse().GetVector())

	// compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
	newP := vb.GetGV().Commit(l.GetVector())
	tmpP := hPrime.Commit(r.GetVector())
	newP.Add(newP, tmpP)
	if vb.FakeTest() {
		//compute p point. p = A + S*x + gv*-z + h'*(z*y^mn) + hj'^(z^j+1 * z^n). (hj=h'[(j-1)*n:j*n-1], j=[1, m])
		p := commitA.Copy()
		p.Add(p, new(ECPoint).ScalarMult(commitB, x))
		p.Add(p, new(ECPoint).ScalarMult(vb.GetGV().Sum(), zNeg))
		p.Add(p, hPrime.Commit(ymn.Times(z).GetVector()))

		for j := 1; j <= aggreateSize; j++ {
			htmp := hPrime.SubVector((j-1)*bitSize, j*bitSize)
			if htmp.Size() != bitSize {
				return nil, errors.New("sub hprime vector not euqal bitsize")
			}
			zj := new(big.Int).Exp(z, new(big.Int).SetUint64(uint64(j+1)), n)
			zjn2 := n2.Times(zj)
			p.Add(p, htmp.Commit(zjn2.GetVector()))
		}

		// compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
		expect := p.Sub(p, new(ECPoint).ScalarMult(vb.GetH(), u))
		if !expect.Equal(newP) {
			return nil, fmt.Errorf("generate proof failed, new p not equal g*l+h'*r")
		}
	}

	ipProver := IPProver{}
	ipProver.U = vb.GetU()
	ipProof, err := ipProver.GenerateIPProof(vb.GetGV(), hPrime, newP, t, l, r)
	if err != nil {
		return nil, err
	}

	proof := AggreateBulletProof{}
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

// VerifyAggreateBulletProof verify agrreate range proofs.
func VerifyAggreateBulletProof(vb *VectorBase, v []*ECPoint, proof *AggreateBulletProof) bool {
	curve := vb.GetCurve()
	n := curve.Params().N
	m := vb.GetAggreateSize()
	size := vb.GetVectorSize()
	bitSize := vb.GetBitSize()
	y, err := ComputeChallenge(n, proof.A.X, proof.A.Y, proof.S.X, proof.S.Y)
	if err != nil {
		log.Warn("compute challenge y failed", "error", err)
		return false
	}
	ymn := PowVector(y, n, size)

	z, err := ComputeChallenge(n, y)
	if err != nil {
		log.Warn("compute challenge z failed", "error", err)
		return false
	}
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	zSquare := new(big.Int).Exp(z, two, n)

	x, err := ComputeChallenge(n, proof.T1.X, proof.T1.Y, proof.T2.X, proof.T2.Y)
	if err != nil {
		log.Warn("compute challenge x failed", "error", err)
		return false
	}
	x2 := new(big.Int).Exp(x, two, n)

	h := vb.GetH()
	g := vb.GetG()

	// check g*tx + h*t ?= v*(z^2 * z^m) + g*dleta + T1*x + T2*x^2. (z^m is a vector)
	zm := PowVector(z, n, m).Times(zSquare)
	expect := NewGeneratorVector(v).Commit(zm.GetVector())

	expect.Add(expect, new(ECPoint).ScalarMult(proof.T1, x))
	expect.Add(expect, new(ECPoint).ScalarMult(proof.T2, x2))
	dleta := DeltaMN(y, z, n, m, bitSize)
	expect.Add(expect, new(ECPoint).ScalarMult(g, dleta))

	actual := new(ECPoint).ScalarMult(g, proof.t)
	actual.Add(actual, new(ECPoint).ScalarMult(h, proof.tx))
	if !expect.Equal(actual) {
		log.Warn("point not equal", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	hPrime := vb.GetHV().Hadamard(ymn.ModInverse().GetVector())
	// compute p point. p = A + S*x + gv*-z + h'*(z*y^mn) + hj'^(z^j+1 * z^n). (hj=h'[(j-1)*n:j*n-1], j=[1, m])
	p := proof.A.Copy()
	p.Add(p, new(ECPoint).ScalarMult(proof.S, x))
	p.Add(p, new(ECPoint).ScalarMult(vb.GetGV().Sum(), zNeg))
	p.Add(p, hPrime.Commit(ymn.Times(z).GetVector()))

	n2 := PowVector(new(big.Int).SetUint64(2), n, bitSize)
	for j := 1; j <= m; j++ {
		htmp := hPrime.SubVector((j-1)*bitSize, j*bitSize)
		zj := new(big.Int).Exp(z, new(big.Int).SetUint64(uint64(j+1)), n)
		zjn2 := n2.Times(zj)
		p.Add(p, htmp.Commit(zjn2.GetVector()))
	}

	// compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
	newP := p.Sub(p, new(ECPoint).ScalarMult(vb.GetH(), proof.u))

	ipVerifier := IPVerifier{}
	ipVerifier.U = vb.GetU()

	return ipVerifier.VerifyIPProof(vb.GetGV(), hPrime, newP, proof.t, proof.ipProof)
}
