package proof

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/pgc/curve"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRangeProof(t *testing.T) {
	cases := []struct {
		curve   elliptic.Curve
		bitsize int
		v       *big.Int
		expect  bool
	}{
		{
			curve:   curve.BN256(),
			bitsize: 32,
			v:       new(big.Int).SetUint64(0),
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 32,
			v:       utils.BiggestInt(32),
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 32,
			v:       new(big.Int).SetInt64(-1),
			expect:  false,
		},
		{
			curve:   curve.BN256(),
			bitsize: 64,
			v:       new(big.Int).SetUint64(0),
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 64,
			v:       utils.BiggestInt(64),
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 64,
			v:       new(big.Int).SetInt64(-1),
			expect:  false,
		},
		{
			curve:   curve.BN256(),
			bitsize: 16,
			v:       new(big.Int).SetUint64(0),
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 16,
			v:       utils.BiggestInt(16),
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 16,
			v:       new(big.Int).SetInt64(-1),
			expect:  false,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       new(big.Int).SetUint64(0),
			expect:  true,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       utils.BiggestInt(16),
			expect:  true,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       new(big.Int).SetInt64(-1),
			expect:  false,
		},
	}

	for _, c := range cases {
		testRangeProof(t, c.curve, c.bitsize, c.v, c.expect)
	}
}

func testRangeProof(t *testing.T, curve elliptic.Curve, bitsize int, v *big.Int, expect bool) {
	params := newRandomRangeParams(curve, bitsize)
	p, r := newRandomCommitmentsRangeProof(params, v)

	proof, err := GenerateRangeProof(params, v, r)
	require.Nil(t, err, "generate range proof failed")

	assert.Equal(t, expect, VerifyRangeProof(params, p, proof), "range proof verify not expect")
	assert.Equal(t, expect, OptimizedVerifyRangeProof(params, p, proof), "optimized range proof not expect")

	//for simple fake proof.
	if expect {
		proof.t.Sub(proof.t, utils.One)
		assert.Equal(t, false, VerifyRangeProof(params, p, proof), "invalid proof pass normal verify")
		assert.Equal(t, false, OptimizedVerifyRangeProof(params, p, proof), "invalid range proof pass optimized verify")
	}
}

func newRandomCommitmentsRangeProof(params RangeParams, v *big.Int) (*utils.ECPoint, *big.Int) {
	g := params.G()
	h := params.H()

	r, err := rand.Int(rand.Reader, params.Curve().Params().N)
	if err != nil {
		panic(err)
	}

	p := new(utils.ECPoint).ScalarMult(g, v)
	p.Add(p, new(utils.ECPoint).ScalarMult(h, r))

	return p, r
}

func BenchmarkBN256RangeProofVerifyNormal(b *testing.B) {
	benchmarkRangeProofVerify(b, new(big.Int).SetUint64(100), curve.BN256(), 32, VerifyRangeProof)
}

func BenchmarkBN256RangeProofVerifyOptimized(b *testing.B) {
	benchmarkRangeProofVerify(b, new(big.Int).SetUint64(100), curve.BN256(), 32, OptimizedVerifyRangeProof)
}

type rangeProofVerifyFunc func(params RangeParams, v *utils.ECPoint, proof *RangeProof) bool

func benchmarkRangeProofVerify(b *testing.B, v *big.Int, curve elliptic.Curve, bitsize int, vf rangeProofVerifyFunc) {
	b.StopTimer()
	params := newRandomRangeParams(curve, bitsize)
	p, r := newRandomCommitmentsRangeProof(params, v)

	proof, err := GenerateRangeProof(params, v, r)
	require.Nil(b, err, "generate range proof failed")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !vf(params, p, proof) {
			b.Fatal("verify failed")
		}
	}
}
