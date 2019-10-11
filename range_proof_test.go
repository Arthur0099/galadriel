package pgc

import (
	"crypto/rand"
	"math/big"
	"testing"
)

func TestRangeProof(t *testing.T) {
	v := MustGetRandomMsg(Params().BitSizeLimit())
	v = new(big.Int).SetUint64(5)
	params := Params()

	r, err := rand.Int(rand.Reader, params.Curve().Params().N)
	if err != nil {
		t.Error(err)
		return
	}

	vb := params.GetVB()
	h := vb.GetH()
	g := vb.GetG()

	// commit point p.
	// p = g*v + h*r.
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

	// v = params.Max()
	// v.Add(v, one)
	// p = new(ECPoint).ScalarMult(g, v)
	// p.Add(p, new(ECPoint).ScalarMult(h, r))
	// proof, err = GenerateRangeProof(vb, v, r)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if VerifyRangeProof(vb, p, proof) {
	// 	t.Fatal("failed")
	// }
}
