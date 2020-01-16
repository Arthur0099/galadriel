package pgc

import (
	"crypto/elliptic"
	"math/big"
)

// VectorBase represent public vector params used in protocol.
type VectorBase struct {
	gv, hv  *GeneratorVector
	g, h, u *ECPoint

	// bit size of pgc system, also the default lenght of a g/h vector.
	bitSize int
	// aggreate proof size, number of g/h vector should be aggreateSize*bitSize.
	aggreateSize int

	fakeTest bool
}

// NewVecotrBase creates new instance.
func NewVecotrBase(gv, hv *GeneratorVector, g, h, u *ECPoint, bitSize, aggreateSize int, fake bool) *VectorBase {
	vb := VectorBase{}
	vb.gv = gv
	vb.hv = hv
	vb.h = h.Copy()
	vb.g = g.Copy()
	vb.u = u.Copy()

	vb.bitSize = bitSize
	vb.aggreateSize = aggreateSize
	vb.fakeTest = fake

	return &vb
}

// NewRandomVectorBase creates random vector base.
func NewRandomVectorBase(curve elliptic.Curve, bitSize, aggreateSize int) *VectorBase {
	vb := VectorBase{}
	vb.gv = NewRandomGeneratorVector(curve, bitSize*aggreateSize)
	vb.hv = NewRandomGeneratorVector(curve, bitSize*aggreateSize)
	vb.h = NewRandomECPoint(curve)
	vb.g = NewRandomECPoint(curve)
	vb.u = NewRandomECPoint(curve)

	vb.bitSize = bitSize
	vb.aggreateSize = aggreateSize

	return &vb
}

// NewDefaultVectorBase creates default vector base.
func NewDefaultVectorBase(curve elliptic.Curve, bitSize, aggreateSize int) *VectorBase {
	vb := VectorBase{}

	g := "g generator of twisted elg"
	vb.g = NewECPointByString(g, curve)

	h := "h generator of twisted elg"
	vb.h = NewECPointByString(h, curve)

	vb.gv = NewDefaultGV(curve, bitSize*aggreateSize)
	vb.hv = NewDefaultHV(curve, bitSize*aggreateSize)

	u := "u generator of innerproduct"
	vb.u = NewECPointByString(u, curve)

	vb.bitSize = bitSize
	vb.aggreateSize = aggreateSize

	return &vb
}

// BitSizeVB .
func (vb *VectorBase) BitSizeVB() *VectorBase {
	if vb.aggreateSize == 1 {
		return vb
	}

	return NewVecotrBase(vb.GetBitSizeGV(), vb.GetBitSizeHV(), vb.g, vb.h, vb.u, vb.bitSize, 1, false)
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

// GetBitSizeHV returns hv with len bitsize.
func (vb VectorBase) GetBitSizeHV() *GeneratorVector {
	if vb.aggreateSize == 1 {
		return vb.GetHV()
	}

	return vb.GetHV().SubVector(0, vb.bitSize)
}

// GetHV returns h generator vector.
func (vb *VectorBase) GetHV() *GeneratorVector {
	return vb.hv
}

// GetBitSizeGV returns gv with len bitsize.
func (vb *VectorBase) GetBitSizeGV() *GeneratorVector {
	if vb.aggreateSize == 1 {
		return vb.GetGV()
	}

	return vb.GetGV().SubVector(0, vb.bitSize)
}

// GetGV returns g generator vector.
func (vb *VectorBase) GetGV() *GeneratorVector {
	return vb.gv
}

// GetBitSize returns bitsize.
func (vb *VectorBase) GetBitSize() int {
	return vb.bitSize
}

// GetAggreateSize returns size of aggreated proof.
func (vb *VectorBase) GetAggreateSize() int {
	return vb.aggreateSize
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

// FakeTest .
func (vb *VectorBase) FakeTest() bool {
	return vb.fakeTest
}
