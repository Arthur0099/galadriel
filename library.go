package pgc

import (
	"crypto/elliptic"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
)

// S256 returns curve used in pgc.
func S256() elliptic.Curve {
	return btcec.S256()
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	return crypto.Keccak256(data...)
}
