package utils

import (
	"testing"

	"github.com/pgc/curve"
	"github.com/stretchr/testify/require"
)

func TestCompress(t *testing.T) {
	cv := curve.BN256()

	rp := NewRandomECPoint(cv)
	compress := rp.Compress()
	dec, err := DecompressPointBytes(cv, compress[:])
	require.Nil(t, err)
	require.True(t, rp.Equal(dec))
}
