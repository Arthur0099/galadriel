package pgc

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/inconshreveable/log15"
)

func TestTT(t *testing.T) {
	s := "owpass$2019$#&#!@&"

	r := Keccak256([]byte(s))
	ss := common.BytesToHash(r).Hex()
	t.Log(ss)
}

func TestCode(t *testing.T) {
	c := BN128{}
	x := []*big.Int{
		new(big.Int).SetUint64(1),
		new(big.Int).SetUint64(4),
		new(big.Int).SetUint64(3),
		new(big.Int).SetUint64(2345345),
		new(big.Int).SetUint64(99),
	}
	x2 := make([]*big.Int, len(x))
	for i, xx := range x {
		x2[i] = new(big.Int).Mul(xx, xx)
		x2[i].Mod(x2[i], c.Params().N)
	}
	n := 5
	bitSize := 32

	l := make([]*big.Int, bitSize)
	ll := make([]*big.Int, bitSize)

	for i := 0; i < bitSize; i++ {
		for j := 0; j < n; j++ {
			tmp := new(big.Int).SetUint64(0)
			if smallParseBinary(i, j, n) {
				tmp = x[j]
			} else {
				tmp = new(big.Int).ModInverse(x[j], c.Params().N)
			}

			if j == 0 {
				l[i] = new(big.Int).Set(tmp)
			} else {
				l[i].Mul(l[i], tmp)
				l[i].Mod(l[i], c.Params().N)
			}
		}

		log.Debug("before", "i", i, "l", l[i])
		if i == 0 {
			for j := 0; j < n; j++ {
				tmp := new(big.Int).SetUint64(0)
				tmp = x[j]

				if j == 0 {
					ll[i] = new(big.Int).Set(tmp)
				} else {
					ll[i].Mul(ll[i], tmp)
					ll[i].Mod(ll[i], c.Params().N)
				}
			}
			ll[i].ModInverse(ll[i], c.Params().N)
		} else {
			k := getBiggestPos(i, n)
			kk := i - pow(k-1)
			ll[i] = new(big.Int).Mul(ll[kk], x2[5-k])
			ll[i].Mod(ll[i], c.Params().N)
		}
		log.Debug("after", "i", i, "ll", ll[i])

		if l[i].Cmp(ll[i]) != 0 {
			t.Fatal("failed")
			log.Error("failed", "i", i, "l", l[i], "ll", ll[i])
		}
	}
}

func TestCheckV(t *testing.T) {
	c := BN128{}
	r := new(big.Int).SetUint64(234)

	r2 := new(big.Int).Mul(r, r)
	r2.Mod(r2, c.Params().N)

	base := new(big.Int).SetUint64(2134234)
	v1 := new(big.Int).Mul(base, r)
	v1.Mod(v1, c.Params().N)
	v1.Mul(v1, r)
	v1.Mod(v1, c.Params().N)

	v2 := new(big.Int).Mul(base, r2)
	v2.Mod(v2, c.Params().N)
	log.Debug("test", "v1", v1, "v2", v2)
}

func calA() {
	//
}

func pow(n int) int {
	if n == 0 {
		return 1
	}

	res := 1
	for n > 0 {
		res *= 2
		n--
	}

	return res
}

func smallParseBinary(t, j, size int) bool {
	w := 1 << (size - 1)
	for i := 0; i < j; i++ {
		w = w >> 1
	}
	if (t & w) != 0 {
		return true
	}

	return false
}

func getBiggestPos(i, s int) int {
	l := 1 << s
	calTimes := 0
	for i < l && l > 0 {
		l = l >> 1
		calTimes++
	}

	return 1 + s - calTimes
}
