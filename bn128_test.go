package pgc

import (
	"crypto/rand"
	"math/big"
	"testing"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	log "github.com/inconshreveable/log15"
)

func TestG1Point(t *testing.T) {
	// get a random g1 point.
	_, p, err := bn256.RandomG1(rand.Reader)
	if err != nil {
		t.Error(err)
		return
	}

	x, y, err := G1ToPoint(p)
	if err != nil {
		t.Error(err)
		return
	}

	newP, err := PointToG1(x, y)
	if err != nil {
		t.Error(err)
		return
	}

	if p.String() != newP.String() {
		t.Error("convert/unconvert from/to g1 point failed")
	}

	c := BN128{}
	params := c.Params()
	gx, gy := params.Gx, params.Gy
	r, err := rand.Int(rand.Reader, params.N)
	if err != nil {
		t.Error(err)
		return
	}

	a, b := c.ScalarBaseMult(r.Bytes())
	aa, bb := c.ScalarMult(gx, gy, r.Bytes())
	if a.Cmp(aa) != 0 || b.Cmp(bb) != 0 {
		t.Error("curve error")
	}

	// get a neg of a point.
	pp := MustPointToG1(a, b)
	ppNeg := new(bn256.G1).Neg(pp)
	a, bNeg := MustG1ToPoint(ppNeg)

	pOrder := new(big.Int).Add(bNeg, b)
	if pOrder.Cmp(params.P) != 0 {
		t.Error("curve p error")
	}
}

func TestInverse(t *testing.T) {
	c := BN128{}
	r := fromString("11888242871839275222246405745257275088548364400416034343698204186575808495617")

	norx, nory := c.ScalarBaseMult(r.Bytes())
	rInverse := new(big.Int).ModInverse(r, c.Params().N)
	log.Debug("n", "n", c.Params().N)
	nn := new(big.Int).Add(r, rInverse)
	log.Debug("nn", "nn", nn)
	invx, invy := c.ScalarBaseMult(rInverse.Bytes())
	log.Debug("nor", "x", norx, "y", nory)
	log.Debug("inv", "x", invx, "y", invy)
	log.Debug("base", "x", c.Params().Gx, "y", c.Params().Gy)
	ggx, ggy := c.Add(norx, nory, invx, invy)
	log.Debug("gg", "x", ggx, "y", ggy)
}

func BenchmarkMult(b *testing.B) {
	r := fromString("11888242871839275222246405745257275088548364400416034343698204186575808495617")
	c := BN128{}

	for i := 0; i < b.N; i++ {
		c.ScalarBaseMult(r.Bytes())
	}
}

func TestBaseFun(t *testing.T) {
	r := new(big.Int).SetUint64(0)
	c := BN128{}
	x, y := c.ScalarBaseMult(r.Bytes())
	log.Info("base", "x", x, "y", y)
}

func BenchmarkMultS(b *testing.B) {
	c := S256()
	n := c.Params().N

	for i := 0; i < b.N; i++ {
		r, _ := rand.Int(rand.Reader, n)
		c.ScalarBaseMult(r.Bytes())
	}
}

func BenchmarkMultSNoCGO(b *testing.B) {
	c := NoCGOS256()
	n := c.Params().N

	for i := 0; i < b.N; i++ {
		r, _ := rand.Int(rand.Reader, n)
		c.ScalarBaseMult(r.Bytes())
	}
}
