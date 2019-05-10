package pgc

import (
	"crypto/elliptic"
	"math/big"
)

// VectorBase represent public vector params used in protocol.
type VectorBase struct {
	gv, hv *GeneratorVector
	h      *ECPoint
}

// NewVecotrBase creates new instance.
func NewVecotrBase(gv, hv *GeneratorVector, h *ECPoint) *VectorBase {
	vb := VectorBase{}
	vb.gv = gv
	vb.hv = hv
	vb.h = h

	return &vb
}

// NewRandomVectorBase creates random vector base.
func NewRandomVectorBase(curve elliptic.Curve, size int) *VectorBase {
	vb := VectorBase{}
	vb.gv = NewRandomGeneratorVector(curve, size)
	vb.hv = NewRandomGeneratorVector(curve, size)
	vb.h = NewRandomECPoint(curve)

	return &vb
}

// GetH returns public h point.
func (vb *VectorBase) GetH() *ECPoint {
	return vb.h
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
