package curve

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/pgc/utils"
	"github.com/stretchr/testify/assert"
)

func TestBN128Curve(t *testing.T) {
	c := BN128{}
	n := c.Params().N
	x, y := c.ScalarBaseMult(n.Bytes())

	assert.Equal(t, x, y)
	assert.Equal(t, x.Cmp(utils.Zero), 0)

	assert.True(t, c.IsOnCurve(c.Params().Gx, c.Params().Gy))
}

func TestForFindYForX(t *testing.T) {
	c := BN128{}

	n := c.Params().N
	p := c.Params().P
	b := c.Params().B
	curveA := new(big.Int).Div(new(big.Int).Add(p, big.NewInt(1)), big.NewInt(4))

	for i := 0; i < 100; i++ {
		s, _ := rand.Int(rand.Reader, n)
		x, y_ := c.ScalarBaseMult(s.Bytes())
		xxx := new(big.Int).Mul(x, x)
		xxx.Mod(xxx, p)
		xxx.Mul(xxx, x)
		xxx.Mod(xxx, p)
		beta := new(big.Int).Add(xxx, b)

		y := new(big.Int).Exp(beta, curveA, p)
		// check y2 == beta
		y2 := new(big.Int).Mul(y, y)
		y2.Mod(y2, p)
		if y2.Cmp(beta) != 0 {
			t.Log(y.String())
			t.Log(y_.String())
			t.Fail()
		}

		if y.Cmp(y_) != 0 {
			t.Log(y.String())
			t.Log(y_.String())
			y_.Sub(p, y_)
			y_.Mod(y_, p)
			if y.Cmp(y_) != 0 {
				t.Fail()
			}
		}
	}

}
