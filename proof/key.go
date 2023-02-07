package proof

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

// MustGenerateKey generates a key pair and panic if err.
// Warn: test purpose only.
func MustGenerateKey(params KeyParams) *ecdsa.PrivateKey {
	key, err := GenerateKey(params)
	if err != nil {
		panic(err)
	}

	return key
}

// GenerateKey generates key pair.
// Warn: The h point from global params is used for generating key pair, not original
// g base point from curve.
func GenerateKey(params KeyParams) (priv *ecdsa.PrivateKey, err error) {
	h := params.H()
	curve := h.Curve

	k, err := randFieldElement(curve, rand.Reader)
	if err != nil {
		return
	}

	priv = new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = k
	// Waring: x, y == h * sk.
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarMult(h.X, h.Y, k.Bytes())
	return
}

// HexToKey returns key pair by hex.
func HexToKey(hexKey string, params KeyParams) (priv *ecdsa.PrivateKey, err error) {
	b, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}

	h := params.H()
	curve := h.Curve
	priv = new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = new(big.Int).SetBytes(b)
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarMult(h.X, h.Y, b)
	return
}
