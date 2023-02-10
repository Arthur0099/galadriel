package curve

import (
	"crypto/elliptic"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
)

// S256 returns curve used in btc.
func S256() elliptic.Curve {
	return crypto.S256()
}

// NoCGOS256 returns curve.
func NoCGOS256() elliptic.Curve {
	return btcec.S256()
}

// BN256 returns curve alt bn128.
func BN256() elliptic.Curve {
	return &BN128{}
}

// GetCurve returns curve by name
func GetCurve(name string) (elliptic.Curve, error) {
	switch name {
	case bn256Name:
		return BN256(), nil

	default:
		return nil, fmt.Errorf("Unsupported curve: %s", name)
	}
}
