package pgc

import (
	"bytes"
	"math/big"
	"testing"
)

func TestCT(t *testing.T) {
	sys := NewTwistedELGamalSystem()
	// generate key.
	key, err := sys.GenerateKey()
	if err != nil {
		panic("generate key failed")
	}

	// Encrypt msg.
	msg := new(big.Int).SetUint64(2).Bytes()
	// Just use a certain r for test.
	ct, err := sys.Encrypt(&key.PublicKey, msg)
	if err != nil {
		t.Error("encrypt data failed", err)
	}

	newMsg := sys.Decrypt(key, ct.CopyPublicPoint())
	if !bytes.Equal(msg, newMsg) {
		t.Error("encrypt/decrypt msg not equal")
	}
}
