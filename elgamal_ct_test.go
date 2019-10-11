package pgc

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCT(t *testing.T) {
	// generate key.
	key, err := GenerateKey()
	if err != nil {
		panic("generate key failed")
	}

	// Encrypt msg.
	msg := new(big.Int).SetUint64(2).Bytes()
	// Just use a certain r for test.
	ct, err := Encrypt(&key.PublicKey, msg)
	if err != nil {
		t.Error("encrypt data failed", err)
	}

	newMsg := Decrypt(key, ct.CopyPublicPoint())
	if !bytes.Equal(msg, newMsg) {
		t.Error("encrypt/decrypt msg not equal")
	}
}

func BenchmarkMultEncrypt(b *testing.B) {
	key, err := GenerateKey()
	if err != nil {
		panic("generate key failed")
	}

	// Encrypt msg.
	n := Params().Max()
	// Just use a certain r for test.
	for i := 0; i < b.N; i++ {
		msg, _ := rand.Int(rand.Reader, n)
		_, err := Encrypt(&key.PublicKey, msg.Bytes())
		if err != nil {
			b.Fatal(err)
		}
	}

}

func TestEncryptTransfer(t *testing.T) {
	alice, err := GenerateKey()
	require.Nil(t, err, "generate key failed")
	bob, err := GenerateKey()
	require.Nil(t, err, "generate key failed")
	msg, err := rand.Int(rand.Reader, Params().Max())
	require.Nil(t, err, "generate msg failed", err)

	ct, err := EncryptTransfer(&alice.PublicKey, &bob.PublicKey, msg.Bytes())
	require.Nil(t, err, "encrypte ct failed", err)

	// check pk * r == x.
	aliceX := new(ECPoint).ScalarMult(new(ECPoint).SetFromPublicKey(&alice.PublicKey), ct.R)
	require.Equal(t, true, aliceX.Equal(ct.X1), "alice pk*r not equal")

	bobX := new(ECPoint).ScalarMult(new(ECPoint).SetFromPublicKey(&bob.PublicKey), ct.R)
	require.Equal(t, true, bobX.Equal(ct.X2), "bob pk*r not equal")
}
