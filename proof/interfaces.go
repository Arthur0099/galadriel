package proof

import (
	"crypto/elliptic"

	"github.com/pgc/curve"
	"github.com/pgc/utils"
)

type all interface {
	Bitsize() int
	Aggsize() int
	Curve() elliptic.Curve
	GV() *utils.GeneratorVector
	HV() *utils.GeneratorVector
	U() *utils.ECPoint
	G() *utils.ECPoint
	H() *utils.ECPoint
	Pub() *utils.ECPoint
}

// KeyParams contains params to genrate key.
type KeyParams interface {
	H() *utils.ECPoint
}

type keyParams struct {
	h *utils.ECPoint
}

func (k *keyParams) H() *utils.ECPoint {
	return k.h
}

func newRandomKeyParams(curve elliptic.Curve) KeyParams {
	h := utils.NewRandomECPoint(curve)

	return &keyParams{h}
}

func DefaultKeyParams() KeyParams {
	curve := curve.BN256()
	h := "h generator of twisted elg"
	hpoint := utils.NewECPointByString(h, curve)

	return &keyParams{hpoint}
}
