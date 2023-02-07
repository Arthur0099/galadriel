package proof

import (
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/pgc/curve"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/require"
)

func TestPTEqualityProof(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}
	bitsizes := []int{16, 32}
	for _, curve := range curves {
		for _, bitsize := range bitsizes {
			params := newRandomCTParams(curve, bitsize)
			// generate key pair for alice and bob.
			alice, err := GenerateKey(params)
			require.Nil(t, err, "generate key pair for alice failed", err)

			bob, err := GenerateKey(params)
			require.Nil(t, err, "generate key pair for bob failed", err)

			msg := new(big.Int).SetUint64(100)

			encryptedCT, err := EncryptTransfer(params, &alice.PublicKey, &bob.PublicKey, msg.Bytes())
			require.Nil(t, err, "encrypt transfer failed", err)

			proof, err := GeneratePTEqualityProof(params, &alice.PublicKey, &bob.PublicKey, encryptedCT)
			require.Nil(t, err, "generate pte quality proof failed", err)

			result := VerifyPTEqualityProof(params, &alice.PublicKey, &bob.PublicKey, encryptedCT.Pub(), proof)
			require.Equal(t, true, result, "verify pte quality proof failed")

			// fake proof should fail
			fakeMsg := new(big.Int).Sub(msg, utils.One)
			fakeCT, err := EncryptTransfer(params, &alice.PublicKey, &bob.PublicKey, fakeMsg.Bytes())
			require.Nil(t, err, "encrypt transfer failed", err)

			result = VerifyPTEqualityProof(params, &alice.PublicKey, &bob.PublicKey, fakeCT.Pub(), proof)
			require.Equal(t, false, result, "fake proof pass check")
		}
	}
}
