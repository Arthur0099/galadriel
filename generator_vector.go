package pgc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

// GeneratorVector respresents ecpoints generatorvector used in bulletproof.
type GeneratorVector struct {
	// ecpoints as vector.
	vector []*ecdsa.PublicKey
}

// NewGeneratorVector creates GeneratorVector instance.
// Warning: change to original vector will also change vector in GeneratorVector.
func NewGeneratorVector(ecPoints []*ecdsa.PublicKey) *GeneratorVector {
	g := GeneratorVector{}
	g.vector = ecPoints

	return &g
}

// NewRandomGeneratorVector creates generator vector randomly.
// Warning: Just for test Purpose.
func NewRandomGeneratorVector(curve elliptic.Curve, n int) *GeneratorVector {
	g := GeneratorVector{}
	order := curve.Params().N
	g.vector = make([]*ecdsa.PublicKey, 0)

	for i := 0; i < n; i++ {
		tmp, err := rand.Int(rand.Reader, order)
		if err != nil {
			// no sense for test.
			panic(err)
		}

		x, y := curve.ScalarBaseMult(tmp.Bytes())
		g.vector = append(g.vector, PointToPublicKey(x, y, curve))
	}

	return &g
}

// Size returns len of underlying vector.
func (gv *GeneratorVector) Size() int {
	return len(gv.vector)
}

// HalfLeft returns half vector on left.
func (gv *GeneratorVector) HalfLeft() *GeneratorVector {
	return gv.SubVector(0, gv.Size()/2)
}

// HalfRight returns half vector on right.
func (gv *GeneratorVector) HalfRight() *GeneratorVector {
	size := gv.Size()
	return gv.SubVector(size/2, size)
}

// SubVector returns new sub vector instance by index of start and end.
func (gv *GeneratorVector) SubVector(start, end int) *GeneratorVector {
	if start < 0 || end > len(gv.vector) {
		panic(fmt.Sprintf("vector index start %d, end %d out of range", start, end))
	}

	newVector := make([]*ecdsa.PublicKey, 0)
	for _, point := range gv.vector[start:end] {
		newVector = append(newVector, PointToPublicKey(point.X, point.Y, point.Curve))
	}

	return NewGeneratorVector(newVector)
}

// Copy returns a new copy instance of generator vector.
func (gv *GeneratorVector) Copy() *GeneratorVector {
	return gv.SubVector(0, len(gv.vector))
}

// Commit computes res = gi * ai + ... + gn * an.
func (gv *GeneratorVector) Commit(a []*big.Int) *ecdsa.PublicKey {
	if len(gv.vector) != len(a) {
		panic(fmt.Sprintf("vector len %d != field vector len %d", len(gv.vector), len(a)))
	}

	// compute res.
	res := new(ecdsa.PublicKey)
	first := gv.vector[0]
	res.Curve = first.Curve
	res.X, res.Y = res.Curve.ScalarMult(first.X, first.Y, a[0].Bytes())
	for i := 1; i < len(a); i++ {
		cX, cY := res.Curve.ScalarMult(gv.vector[i].X, gv.vector[i].Y, a[i].Bytes())
		res.X, res.Y = res.Curve.Add(res.X, res.Y, cX, cY)
	}

	return res
}

// Hadamard computes gi * x + ... + gn * x.
func (gv *GeneratorVector) Hadamard(x *big.Int) *GeneratorVector {
	curve := gv.vector[0].Curve
	newVector := make([]*ecdsa.PublicKey, 0)
	for _, point := range gv.vector {
		px, py := curve.ScalarMult(point.X, point.Y, x.Bytes())
		newVector = append(newVector, PointToPublicKey(px, py, curve))
	}

	return NewGeneratorVector(newVector)
}

// Get returns point by index.
func (gv *GeneratorVector) Get(i int) *ecdsa.PublicKey {
	return gv.vector[i]
}

// AddGeneratorVector add two GeneratorVectors by gi + otheri.
func (gv *GeneratorVector) AddGeneratorVector(other *GeneratorVector) *GeneratorVector {
	if gv.Size() != other.Size() {
		panic(fmt.Sprintf("two generator vector size not equal %d != %d", gv.Size(), other.Size()))
	}

	curve := gv.vector[0].Curve

	// add vector.

	newVector := make([]*ecdsa.PublicKey, 0)
	for i, point := range gv.vector {
		otherPoint := other.Get(i)
		x, y := curve.Add(point.X, point.Y, otherPoint.X, otherPoint.Y)
		newVector = append(newVector, PointToPublicKey(x, y, curve))
	}

	return NewGeneratorVector(newVector)
}

// Print print all info of generator vector(test purpose)
func (gv *GeneratorVector) Print() {
	for i, p := range gv.vector {
		log.Printf("%d vector x %x, y %x", i, p.X, p.Y)
	}
}
