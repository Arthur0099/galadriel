package pgc

import (
	"math/big"
	"testing"
)

func TestVectorBits(t *testing.T) {
	v := new(big.Int).SetUint64(3)
	v2 := new(big.Int).SetUint64(4)

	t.Log(BitVector(v, 4))
	t.Log(BitVector(v2, 4))
}
