package pgc

import (
	"testing"
)

func TestSig(t *testing.T) {
	key, _ := GenerateKey()

	data := []byte("hash")
	sig, err := Sign(key, data)
	if err != nil {
		t.Error(err)
		return
	}

	if !VerifySig(&key.PublicKey, data, sig) {
		t.Error("verify sig failed")
	}

	newKey, _ := GenerateKey()
	if VerifySig(&newKey.PublicKey, data, sig) {
		t.Error("invalid key verify success")
	}
}
