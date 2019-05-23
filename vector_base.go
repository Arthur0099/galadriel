package pgc

import (
	"crypto/elliptic"
	"math/big"
)

// VectorBase represent public vector params used in protocol.
type VectorBase struct {
	gv, hv  *GeneratorVector
	g, h, u *ECPoint
}

// NewVecotrBase creates new instance.
func NewVecotrBase(gv, hv *GeneratorVector, g, h *ECPoint) *VectorBase {
	vb := VectorBase{}
	vb.gv = gv
	vb.hv = hv
	vb.h = h
	vb.g = g

	return &vb
}

// NewRandomVectorBase creates random vector base.
func NewRandomVectorBase(curve elliptic.Curve, size int) *VectorBase {
	vb := VectorBase{}
	vb.gv = NewRandomGeneratorVector(curve, size)
	vb.hv = NewRandomGeneratorVector(curve, size)
	vb.h = NewRandomECPoint(curve)
	vb.g = NewRandomECPoint(curve)
	vb.u = NewRandomECPoint(curve)

	return &vb
}

// NewDefaultVectorBase creates default vector base.
func NewDefaultVectorBase(curve elliptic.Curve, size int) *VectorBase {
	vb := VectorBase{}
	vb.gv = NewDefaultGV(curve, size)
	vb.hv = NewDefaultHV(curve, size)
	g := "g generator of twisted elg"
	vb.g = NewECPointByString(g, curve)

	h := "h generator of twisted elg"
	vb.h = NewECPointByString(h, curve)

	vb.u = vb.g.Copy()

	return &vb
}

// GetH returns public h point.
func (vb *VectorBase) GetH() *ECPoint {
	return vb.h
}

// GetG returns public g point.
func (vb *VectorBase) GetG() *ECPoint {
	return vb.g
}

// GetU returns public u point.
func (vb *VectorBase) GetU() *ECPoint {
	return vb.u
}

// GetHV returns h generator vector.
func (vb *VectorBase) GetHV() *GeneratorVector {
	return vb.hv
}

// GetGV returns g generator vector.
func (vb *VectorBase) GetGV() *GeneratorVector {
	return vb.gv
}

// GetVectorSize returns size of underlying g, h vector.
func (vb *VectorBase) GetVectorSize() int {
	return vb.gv.Size()
}

// GetCurve returns curve used.
func (vb *VectorBase) GetCurve() elliptic.Curve {
	return vb.h.Curve
}

// CommitTwoFieldVector compute .
func (vb *VectorBase) CommitTwoFieldVector(al, ar *FieldVector, alpha *big.Int) *ECPoint {
	return vb.CommitTwoVector(al.GetVector(), ar.GetVector(), alpha)
}

// CommitTwoVector compute.
func (vb *VectorBase) CommitTwoVector(al, ar []*big.Int, alpha *big.Int) *ECPoint {
	commit := vb.gv.Commit(al)
	commit.Add(commit, vb.hv.Commit(ar))
	hAlpha := new(ECPoint).ScalarMult(vb.h, alpha)
	commit.Add(commit, hAlpha)

	return commit
}
