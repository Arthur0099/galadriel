package pgc

import (
	"testing"
)

func TestInnerProductProtocol(t *testing.T) {
	// create public params for test.
	curve := S256()
	n := 32
	order := curve.Params().N
	g := NewDefaultGV(curve, n)
	h := NewDefaultHV(curve, n)
	u := NewRandomECPoint(curve)

	a := NewRandomFieldVector(order, n)
	b := NewRandomFieldVector(order, n)

	c := a.InnerProduct(b)
	pa := g.Commit(a.GetVector())
	pb := h.Commit(b.GetVector())
	p := new(ECPoint).Add(pa, pb)

	prover := IPProver{}
	prover.U = u

	proof, err := prover.GenerateIPProof(g.Copy(), h.Copy(), p, c, a, b)
	if err != nil {
		t.Error("generator proof failed", err)
		return
	}

	verifier := IPVerifier{}
	verifier.U = u
	if !verifier.VerifyIPProof(g.Copy(), h.Copy(), p, c, proof) {
		t.Error("verify failed")
	}
}

func BenchmarkInnerProduct(b *testing.B) {
	order := Params().Curve().Params().N
	n := Params().GetVB().GetVectorSize()
	a := NewRandomFieldVector(order, n)
	aa := NewRandomFieldVector(order, n)
	g := Params().GetVB().GetGV()
	h := Params().GetVB().GetHV()

	c := a.InnerProduct(aa)
	pa := g.Commit(a.GetVector())
	pb := h.Commit(aa.GetVector())
	p := new(ECPoint).Add(pa, pb)

	prover := IPProver{}
	prover.U = Params().GetU()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := prover.GenerateIPProof(g.Copy(), h.Copy(), p, c, a, aa)
		if err != nil {
			b.Fatal(err)
			return
		}
	}

	// 	verifier := IPVerifier{}
	// 	verifier.U = u
	// 	if !verifier.VerifyIPProof(g.Copy(), h.Copy(), p, c, proof) {
	// 		t.Error("verify failed")
	// 	}
}
