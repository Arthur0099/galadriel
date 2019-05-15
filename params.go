package pgc

import (
	"crypto/elliptic"
	"encoding/json"
)

// KeyBasePoint represents the base point used in curve to generate keys.
type KeyBasePoint interface {
	// GetH() returns base point used in curve.
	GetH() *ECPoint
}

// TwistedELGamalPublicParams contains params used in twisting elgamal encrypt protocol.
type TwistedELGamalPublicParams interface {
	// GetG return g generator point.
	GetG() *ECPoint

	KeyBasePoint

	// BitSizeLimit returns the limit of the bit size of msg to be encrypted.
	BitSizeLimit() int

	// Curve returns ec curve used in underlying field.
	Curve() elliptic.Curve
}

// PublicParams includes global public params used in PGC and bullet proof.
type PublicParams struct {
	//
	g, h *ECPoint

	//
	curve elliptic.Curve

	// bitSize is the size of msg.
	// msg < 2^BitSize.
	bitSize int
}

// GetG returns g point.
func (pp *PublicParams) GetG() *ECPoint {
	return pp.g.Copy()
}

// GetH returns h point.
func (pp *PublicParams) GetH() *ECPoint {
	return pp.h.Copy()
}

// BitSizeLimit returns the limit of the bit size of msg to be encrypted.
func (pp *PublicParams) BitSizeLimit() int {
	return pp.bitSize
}

// Curve returns ec curve used in underlying field.
func (pp *PublicParams) Curve() elliptic.Curve {
	return pp.curve
}

// MarshalJSON defines custom way to json.
func (pp *PublicParams) MarshalJSON() ([]byte, error) {
	newJSON := struct {
		G *ECPoint `json:"g"`
		H *ECPoint `json:"h"`
	}{
		G: pp.g,
		H: pp.h,
	}

	return json.Marshal(&newJSON)
}

var params PublicParams

// Params returns public params used in protocol.
func Params() *PublicParams {
	return &params
}

// init public params.
func init() {
	// todo: set config to switch curve.
	//curve := S256()
	curve := BN256()
	// to be compatible with sig curve.
	h := NewECPoint(curve.Params().Gx, curve.Params().Gy, curve)
	g := "g generator of twisted elg"

	params.g = NewECPointByString(g, curve)
	params.h = h
	params.curve = curve
	params.bitSize = 16
}
