package pgc

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// FieldVector respresents vector scalar array.
type FieldVector struct {
	items []*big.Int
	// order n.
	n *big.Int
}

// NewFieldVector creates instance of field vector.
func NewFieldVector(items []*big.Int, n *big.Int) *FieldVector {
	f := FieldVector{}
	f.items = items
	f.n = new(big.Int).Set(n)

	return &f
}

// NewRandomFieldVector generates field vector randomly.
// Warning: Just for test purpose.
func NewRandomFieldVector(order *big.Int, n int) *FieldVector {
	f := FieldVector{}
	f.n = new(big.Int).Set(order)
	f.items = make([]*big.Int, 0)
	for i := 0; i < n; i++ {
		tmp, err := rand.Int(rand.Reader, order)
		if err != nil {
			// no sense for test purpose.
			panic(err)
		}

		f.items = append(f.items, tmp)
	}

	return &f
}

// HalfLeft returns half of items on left.
func (f *FieldVector) HalfLeft() *FieldVector {
	return f.SubFieldVector(0, f.Size()/2)
}

// HalfRight returns half of items on right.
func (f *FieldVector) HalfRight() *FieldVector {
	size := f.Size()
	return f.SubFieldVector(size/2, size)
}

// Size returns len of underlying items.
func (f *FieldVector) Size() int {
	return len(f.items)
}

// SubFieldVector returns sub items by start/end index.
func (f *FieldVector) SubFieldVector(start, end int) *FieldVector {
	if start < 0 || end > f.Size() {
		panic(fmt.Sprintf("field vector index start %d, end %d out of range", start, end))
	}

	newItems := make([]*big.Int, end-start)
	copy(newItems, f.items[start:end])

	return NewFieldVector(newItems, f.n)
}

// First returns first item in field.
func (f *FieldVector) First() *big.Int {
	return f.Get(0)
}

// Get returns item by index.
func (f *FieldVector) Get(i int) *big.Int {
	return new(big.Int).Set(f.items[i])
}

// InnerProduct computes <items, other>.
func (f *FieldVector) InnerProduct(other *FieldVector) *big.Int {
	if f.Size() != other.Size() {
		panic(fmt.Sprintf("field vector size %d != %d", f.Size(), other.Size()))
	}

	res := new(big.Int)
	for i, item := range f.items {
		tmp := new(big.Int).Mul(item, other.Get(i))
		res.Add(res, tmp)
		res.Mod(res, f.n)
	}

	return res.Mod(res, f.n)
}

// GetVector returns underlying items.
func (f *FieldVector) GetVector() []*big.Int {
	newItems := make([]*big.Int, f.Size())
	copy(newItems, f.items)

	return newItems
}

// Times compute item * x and returns new instance.
func (f *FieldVector) Times(x *big.Int) *FieldVector {
	newItems := make([]*big.Int, 0)

	for _, i := range f.items {
		t := new(big.Int).Mul(i, x)
		t.Mod(t, f.n)
		newItems = append(newItems, t)
	}

	return NewFieldVector(newItems, f.n)
}

// AddFieldVector computes fi + otheri.
func (f *FieldVector) AddFieldVector(other *FieldVector) *FieldVector {
	if f.Size() != other.Size() {
		panic(fmt.Sprintf("filed vector size not equal %d != %d", f.Size(), other.Size()))
	}

	newItems := make([]*big.Int, 0)
	for i, item := range f.items {
		t := new(big.Int).Add(item, other.Get(i))
		t.Mod(t, f.n)
		newItems = append(newItems, t)
	}

	return NewFieldVector(newItems, f.n)
}
