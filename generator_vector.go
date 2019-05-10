package pgc

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// GeneratorVector respresents ecpoints generatorvector used in bulletproof.
type GeneratorVector struct {
	// ecpoints as vector.
	vector []*ECPoint
	// for convinence.
	Curve elliptic.Curve
}

// NewGeneratorVector creates GeneratorVector instance.
// Warning: change to original vector will also change vector in GeneratorVector.
func NewGeneratorVector(ecPoints []*ECPoint) *GeneratorVector {
	g := GeneratorVector{}
	if len(ecPoints) > 0 {
		g.vector = ecPoints
		g.Curve = ecPoints[0].Curve
	}

	return &g
}

// NewRandomGeneratorVector creates generator vector randomly.
// Warning: Just for test Purpose.
func NewRandomGeneratorVector(curve elliptic.Curve, n int) *GeneratorVector {
	g := GeneratorVector{}
	order := curve.Params().N
	g.vector = make([]*ECPoint, 0)
	g.Curve = curve

	for i := 0; i < n; i++ {
		tmp, err := rand.Int(rand.Reader, order)
		if err != nil {
			// no sense for test.
			panic(err)
		}

		x, y := curve.ScalarBaseMult(tmp.Bytes())
		g.vector = append(g.vector, NewECPoint(x, y, curve))
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

	newVector := make([]*ECPoint, 0)
	for _, point := range gv.vector[start:end] {
		newVector = append(newVector, NewECPoint(point.X, point.Y, point.Curve))
	}

	return NewGeneratorVector(newVector)
}

// Copy returns a new copy instance of generator vector.
func (gv *GeneratorVector) Copy() *GeneratorVector {
	return gv.SubVector(0, len(gv.vector))
}

// Sum compute gi + ... + gn.
func (gv *GeneratorVector) Sum() *ECPoint {
	res := gv.vector[0].Copy()

	for i := 1; i < gv.Size(); i++ {
		res.Add(res, gv.vector[i])
	}

	return res
}

// Commit computes res = gi * ai + ... + gn * an.
func (gv *GeneratorVector) Commit(a []*big.Int) *ECPoint {
	if len(gv.vector) != len(a) {
		panic(fmt.Sprintf("vector len %d != field vector len %d", len(gv.vector), len(a)))
	}

	// compute res.
	res := new(ECPoint).ScalarMult(gv.vector[0], a[0])
	for i := 1; i < len(a); i++ {
		tmpP := new(ECPoint).ScalarMult(gv.vector[i], a[i])
		res.Add(res, tmpP)
	}

	return res
}

// HadamardScalar computes gi * x + ... + gn * x.
func (gv *GeneratorVector) HadamardScalar(x *big.Int) *GeneratorVector {
	newVector := make([]*ECPoint, 0)
	for _, point := range gv.vector {
		p := new(ECPoint).ScalarMult(point, x)
		newVector = append(newVector, p)
	}

	return NewGeneratorVector(newVector)
}

// Hadamard computes gi*ai + ... + gn*an.
func (gv *GeneratorVector) Hadamard(exponent []*big.Int) *GeneratorVector {
	if gv.Size() != len(exponent) {
		panic("exponent len not equal with generator vector size")
	}

	newVector := make([]*ECPoint, 0)

	for i, point := range gv.vector {
		p := new(ECPoint).ScalarMult(point, exponent[i])
		newVector = append(newVector, p)
	}

	return NewGeneratorVector(newVector)
}

// Get returns point by index.
func (gv *GeneratorVector) Get(i int) *ECPoint {
	return gv.vector[i]
}

// AddGeneratorVector add two GeneratorVectors by gi + otheri.
func (gv *GeneratorVector) AddGeneratorVector(other *GeneratorVector) *GeneratorVector {
	if gv.Size() != other.Size() {
		panic(fmt.Sprintf("two generator vector size not equal %d != %d", gv.Size(), other.Size()))
	}

	// add vector.
	newVector := make([]*ECPoint, 0)
	for i, point := range gv.vector {
		otherPoint := other.Get(i)
		newP := new(ECPoint).Add(point, otherPoint)
		newVector = append(newVector, newP)
	}

	return NewGeneratorVector(newVector)
}

// Print print all info of generator vector(test purpose)
func (gv *GeneratorVector) Print() {
	for i, p := range gv.vector {
		log.Debug("generator vector", "index", i, "x", p.X, "y", p.Y)
	}
}
