package pgc

import (
	"encoding/json"
	"testing"
)

func TestInnerProductProtocol(t *testing.T) {
	// create public params for test.
	curve := BN256()
	n := Params().BitSizeLimit()
	order := curve.Params().N
	g := NewDefaultGV(curve, n)
	h := NewDefaultHV(curve, n)

	a := NewRandomFieldVector(order, n)
	b := NewRandomFieldVector(order, n)

	c := a.InnerProduct(b)
	pa := g.Commit(a.GetVector())
	pb := h.Commit(b.GetVector())
	p := new(ECPoint).Add(pa, pb)

	prover := IPProver{}
	prover.U = Params().GetU()

	proof, err := prover.GenerateIPProof(g.Copy(), h.Copy(), p, c, a, b)
	if err != nil {
		t.Error("generator proof failed", err)
		return
	}

	verifier := IPVerifier{}
	verifier.U = Params().GetU()
	if !verifier.VerifyIPProof(g.Copy(), h.Copy(), p, c, proof) {
		t.Error("verify failed")
	}

	newJSON := struct {
		P     *ECPoint `json:"p"`
		C     string   `json:"c"`
		Proof *IPProof `json:"proof"`
	}{
		P:     p.Copy(),
		C:     c.String(),
		Proof: proof,
	}

	data, err := json.Marshal(&newJSON)
	if err != nil {
		panic(err)
	}

	path := "solidity/proofs/ipProof"
	WriteToFile(data, path)
}
