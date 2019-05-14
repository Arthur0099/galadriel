package pgc

import (
	"bytes"
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
	msg := []byte("h")
	// Just use a certain r for test.
	ct, err := sys.Encrypt(&key.PublicKey, msg)
	if err != nil {
		t.Error("encrypt data failed", err)
	}

	newMsg := sys.Decrypt(key, ct)
	if !bytes.Equal(msg, newMsg) {
		t.Error("encrypt/decrypt msg not equal")
	}
}
