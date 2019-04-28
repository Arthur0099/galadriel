package pgc

import (
	"math/big"
	"testing"
)

func TestCT(t *testing.T) {
	// generate key.
	key, err := GenerateKey()
	if err != nil {
		panic("generate key failed")
	}

	// Encrypt msg.
	msg := []byte("h")
	// Just use a certain r for test.
	ct, err := Encrypt(&key.PublicKey, msg)
	if err != nil {
		t.Error("encrypt data failed", err)
	}

	Decrypt(key, ct)
}

func TestCurve(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		panic("generate key failed")
	}

	curve := key.PublicKey.Curve
	if !curve.IsOnCurve(key.X, key.Y) {
		t.Error("public key not on curve")
	}

	skX, skY := curve.ScalarBaseMult(key.D.Bytes())
	if skX.Cmp(key.X) != 0 || skY.Cmp(key.Y) != 0 {
		t.Error("sk * Gpoint not equal public key")
	}

	r := new(big.Int).SetUint64(199999999)
	pkRX, pkRY := curve.ScalarMult(key.X, key.Y, r.Bytes())
	skInverse := new(big.Int).ModInverse(key.D, curve.Params().N)
	dstX, dstY := curve.ScalarMult(pkRX, pkRY, skInverse.Bytes())
	if dstX.Cmp(curve.Params().Gx) != 0 || dstY.Cmp(curve.Params().Gy) != 0 {
	}
}
