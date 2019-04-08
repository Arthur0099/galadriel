package pgc

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
)

// TestEncryptDecrypt tests base func of encrypt/decrypt in twistedELG.
func TestEncryptDecrypt(t *testing.T) {
	key, err := GenTwistedELG()
	if err != nil {
		t.Error(err)
	}

	data := []byte("hello world")
	c1, c2, err := Encrypt(rand.Reader, &key.PublicKey, data)
	if err != nil {
		t.Errorf("error encrypting: %s", err)
	}
	data2, err := Decrypt(&key.PrivateKey, c1, c2)
	if err != nil {
		t.Errorf("error decrypting: %s", err)
	}
	if !bytes.Equal(data, data2) {
		t.Errorf("decryption failed, got: %x, want %x", data2, data)
	}

	p, _, h := GroupGen()
	m := new(big.Int).SetBytes(data)
	hExpM := new(big.Int).Exp(h, m, p)
	fmt.Printf("new h ^^ m is %x\n", hExpM)
}
