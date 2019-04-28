package pgc

import (
	"crypto/ecdsa"
	"math/big"
)

// PublicParams includes global public params used in PGC and bullet proof.
type PublicParams struct {
	// pgc public params.
	ElgGenerator *ecdsa.PublicKey

	// bullet proof params.
}

var params PublicParams

func Params() *PublicParams {
	return &params
}

// init public params.
func init() {
	// result is e7f38f4a9d49b6b6f46d18d6ae54be8a9e4df63e6b291c2e89b6915d5f45235e.
	h := Keccak256([]byte("h generator"))
	curve := S256()

	// make sure h is less than N in curve.
	if new(big.Int).SetBytes(h).Cmp(curve.Params().N) >= 0 {
		panic("h generator is bigger than N in curve")
	}

	elgGenerator := new(ecdsa.PublicKey)
	elgGenerator.Curve = curve
	elgGenerator.X, elgGenerator.Y = curve.ScalarBaseMult(h)
	params.ElgGenerator = elgGenerator
}
