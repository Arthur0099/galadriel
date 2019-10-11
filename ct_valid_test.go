package pgc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCTValidProof(t *testing.T) {
	alice := MustGenerateKey()
	msg := MustGetRandomMsg(Params().BitSizeLimit())
	ct, err := Encrypt(&alice.PublicKey, msg.Bytes())
	require.Nil(t, err, "encrypt msg failed", err)

	proof, err := GenerateCTValidProof(&alice.PublicKey, ct)
	require.Nil(t, err, "generate ct valid proof failed", err)

	result := VerifyCTValidProof(&alice.PublicKey, ct.CopyPublicPoint(), proof)
	require.Equal(t, result, true, "veerify a valid ctvalid proof failed")
}
