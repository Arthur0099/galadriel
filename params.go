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
	curve elliptic.Curve

	// bitSize is the size of msg.
	// msg < 2^BitSize.
	bitSize int

	// vector base.
	vb *VectorBase
}

// GetG returns g point.
func (pp *PublicParams) GetG() *ECPoint {
	return pp.vb.g.Copy()
}

// GetH returns h point.
func (pp *PublicParams) GetH() *ECPoint {
	return pp.vb.h.Copy()
}

// GetU returns u point.
func (pp *PublicParams) GetU() *ECPoint {
	return pp.vb.u.Copy()
}

// GetVB returns vector base.
func (pp *PublicParams) GetVB() *VectorBase {
	return pp.vb
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
		G  *ECPoint         `json:"g"`
		H  *ECPoint         `json:"h"`
		GV *GeneratorVector `json:"gv"`
		HV *GeneratorVector `json:"hv"`
		U  *ECPoint         `json:"u"`
	}{
		G:  pp.vb.g,
		H:  pp.vb.h,
		GV: pp.vb.gv,
		HV: pp.vb.hv,
		U:  pp.vb.u,
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

	params.curve = curve
	params.bitSize = 16

	params.vb = NewDefaultVectorBase(curve, params.bitSize)
}
