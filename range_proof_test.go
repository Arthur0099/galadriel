package pgc

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"testing"
)

func TestRangeProof(t *testing.T) {
	v := new(big.Int).SetUint64(10)
	params := Params()

	// compute g * r + h * v
	r, err := rand.Int(rand.Reader, params.Curve().Params().N)
	if err != nil {
		t.Error(err)
		return
	}

	vb := params.GetVB()
	h := vb.GetH()
	g := vb.GetG()

	// commit point p.
	p := new(ECPoint).ScalarMult(g, v)
	p.Add(p, new(ECPoint).ScalarMult(h, r))

	proof, err := GenerateRangeProof(vb, v, r)
	if err != nil {
		t.Error(err)
		return
	}

	if !VerifyRangeProof(vb, p, proof) {
		t.Error("failed")
		return
	}

	newJSON := struct {
		Proof *RangeProof
		P     *ECPoint
	}{
		Proof: proof,
		P:     p,
	}

	data, err := json.Marshal(&newJSON)
	if err != nil {
		panic(err)
	}
	path := "solidity/proofs/rangeProof"
	WriteToFile(data, path)
}
