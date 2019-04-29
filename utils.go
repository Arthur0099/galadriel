package pgc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

// PointToPurblicKey converts point to public key object.
func PointToPublicKey(x, y *big.Int, curve elliptic.Curve) *ecdsa.PublicKey {
	key := new(ecdsa.PublicKey)

	key.X = x
	key.Y = y
	key.Curve = curve

	return key
}
