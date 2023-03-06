package utils

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBiggestInt(t *testing.T) {
	cases := []struct {
		size   int
		expect *big.Int
	}{
		{
			size:   2,
			expect: new(big.Int).SetUint64(3),
		},
		{
			size:   3,
			expect: new(big.Int).SetUint64(7),
		},
		{
			size:   4,
			expect: new(big.Int).SetUint64(15),
		},
		{
			size:   5,
			expect: new(big.Int).SetUint64(31),
		},
	}

	for _, c := range cases {
		actual := BiggestInt(c.size)
		assert.Equal(t, true, actual.Cmp(c.expect) == 0, fmt.Sprintf("expect %s, actual %s", c.expect, actual))
	}
}

func TestMurmurHash64A(t *testing.T) {
	key := []byte("This is a murmur hash test")
	another := []byte("This is an murmur hash test")

	toKey := make([]byte, len(key))
	copy(toKey, key)
	h := MurmurHash64AWithFixedSalt([]byte(key))
	require.True(t, bytes.Equal(key, toKey))
	ah := MurmurHash64AWithFixedSalt([]byte(another))
	require.True(t, h != ah)
}
