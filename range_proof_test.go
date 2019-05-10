package pgc

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestRangeProof(t *testing.T) {
	v := new(big.Int).SetUint64(10)
	curve := S256()
	n := curve.Params().N
	size := 8

	// compute g * r + h * v
	r, err := rand.Int(rand.Reader, n)
	if err != nil {
		t.Error(err)
		return
	}

	g := NewECPoint(curve.Params().Gx, curve.Params().Gy, curve)
	vb := NewRandomVectorBase(curve, size)
	h := vb.GetH()

	// commit point p.
	p := new(ECPoint).ScalarMult(g, v)
	p.Add(p, new(ECPoint).ScalarMult(h, r))

	prover := RangeProver{}
	proof, err := prover.GenerateRangeProof(vb, v, r)
	if err != nil {
		t.Error(err)
		return
	}

	verifier := RangeProofVerifier{}
	verifier.g = g.Copy()
	verifier.h = h.Copy()

	if !verifier.VerifyRangeProof(vb, p, proof) {
		t.Error("failed")
	}
}
