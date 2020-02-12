package proof

import (
	"bytes"
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/pgc/curve"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyGenerate(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}

	for _, curve := range curves {
		params := newRandomKeyParams(curve)
		// generate key.
		key, err := GenerateKey(params)
		if err != nil {
			panic("generate key failed")
		}

		// check priv.D * h = pub
		x, y := curve.ScalarMult(params.H().X, params.H().Y, key.D.Bytes())
		assert.Equal(t, key.X, x, "key generate failed x not equal")
		assert.Equal(t, key.Y, y, "key generate failed y not equal")
	}

}

func TestCT(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}
	bitsize := 16

	msgs := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(-1),
	}

	for _, curve := range curves {
		for _, msg := range msgs {
			params := newRandomCTParams(curve, bitsize)
			// generate key.
			key, err := GenerateKey(params)
			require.Nil(t, err, "generate key failed")

			// Encrypt msg.
			// Just use a certain r for test.
			ct, err := Encrypt(params, &key.PublicKey, msg.Bytes())
			require.Nil(t, err, "encrypt data failed")

			newMsg := Decrypt(params, key, ct.CopyPublicPoint())
			if !bytes.Equal(msg.Bytes(), newMsg) {
				t.Error("encrypt/decrypt msg not equal")
			}

			nct, err := Refresh(params, key, ct.CopyPublicPoint())
			require.Nil(t, err, "refresh failed")
			assert.NotEqual(t, ct.R, nct.R, "randome equal after refresh")
		}
	}
}

func TestEncryptTransfer(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}
	bitsize := 16

	for _, curve := range curves {
		params := newRandomCTParams(curve, bitsize)
		alice, err := GenerateKey(params)
		require.Nil(t, err, "generate key failed")
		bob, err := GenerateKey(params)
		require.Nil(t, err, "generate key failed")
		msg := new(big.Int).SetUint64(100)
		require.Nil(t, err, "generate msg failed", err)

		ct, err := EncryptTransfer(params, &alice.PublicKey, &bob.PublicKey, msg.Bytes())
		require.Nil(t, err, "encrypte ct failed", err)

		// check pk * r == x.
		aliceX := new(utils.ECPoint).ScalarMult(new(utils.ECPoint).SetFromPublicKey(&alice.PublicKey), ct.R)
		require.Equal(t, true, aliceX.Equal(ct.X1), "alice pk*r not equal")

		bobX := new(utils.ECPoint).ScalarMult(new(utils.ECPoint).SetFromPublicKey(&bob.PublicKey), ct.R)
		require.Equal(t, true, bobX.Equal(ct.X2), "bob pk*r not equal")
	}

}
