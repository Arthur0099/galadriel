package pgc

import (
	"crypto/rand"
	"testing"
)

func BenchmarkCommit(b *testing.B) {
	vb := Params().GetVB()
	n := vb.GetCurve().Params().N
	vectorSize := vb.GetVectorSize()
	sl := NewRandomFieldVector(n, vectorSize)
	sr := NewRandomFieldVector(n, vectorSize)

	r, _ := rand.Int(rand.Reader, n)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		vb.CommitTwoFieldVector(sl, sr, r)
	}
}
