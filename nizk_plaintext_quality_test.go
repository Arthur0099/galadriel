package pgc

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPTEqualityProof(t *testing.T) {
	// generate key pair for alice and bob.
	alice, err := GenerateKey()
	require.Nil(t, err, "generate key pair for alice failed", err)

	bob, err := GenerateKey()
	require.Nil(t, err, "generate key pair for bob failed", err)

	msg, err := rand.Int(rand.Reader, Params().Max())
	require.Nil(t, err, "generate msg failed", err)

	encryptedCT, err := EncryptTransfer(&alice.PublicKey, &bob.PublicKey, msg.Bytes())
	require.Nil(t, err, "encrypte transfer failed", err)

	proof, err := GeneratePTEqualityProof(&alice.PublicKey, &bob.PublicKey, encryptedCT)
	require.Nil(t, err, "generate ptequality proof failed", err)

	result := VerifyPTEqualityProof(&alice.PublicKey, &bob.PublicKey, encryptedCT.Pub(), proof)
	require.Equal(t, true, result, "verify ptequality proof failed")
}
