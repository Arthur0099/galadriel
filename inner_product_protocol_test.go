package pgc

import (
	"math/big"
	"testing"
)

func TestInnerProductProtocol(t *testing.T) {
	// create public params for test.
	curve := S256()
	n := 8
	order := curve.Params().N
	g := NewRandomGeneratorVector(curve, n)
	h := NewRandomGeneratorVector(curve, n)

	a := NewRandomFieldVector(order, n)
	b := NewRandomFieldVector(order, n)

	c := a.InnerProduct(b)
	pa := g.Commit(a.GetVector())
	pb := h.Commit(b.GetVector())
	px, py := curve.Add(pa.X, pa.Y, pb.X, pb.Y)
	p := PointToPublicKey(px, py, curve)

	random := new(big.Int).SetUint64(1000)
	ux, uy := curve.ScalarBaseMult(random.Bytes())
	u := PointToPublicKey(ux, uy, curve)

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
