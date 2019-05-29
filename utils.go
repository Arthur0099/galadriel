package pgc

import (
	"errors"
	"math/big"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ComputeChallenge computes challenge x using hash func(hash(pack(data))).
// todo: same with Keccak256(a1, a2, b1, b2) in solidity.
// use abi.Arguments.Pack(A1, A2, B1, B2)
// hash(bytes)
func ComputeChallenge(order *big.Int, data ...interface{}) (*big.Int, error) {
	uint256Type, _ := abi.NewType("uint256", nil)
	arguments := make(abi.Arguments, 0)
	for i := 0; i < len(data); i++ {
		argument := abi.Argument{}
		argument.Type = uint256Type

		arguments = append(arguments, argument)
	}

	packedData, err := arguments.Pack(data...)
	if err != nil {
		return nil, err
	}
	e := new(big.Int).SetBytes(Keccak256(packedData))
	e = e.Mod(e, order)

	return e, nil
}

// BitVector returns vector containing the bits of v.
// v = <al, 2 ^ n>. and all items in al are {0, 1}
func BitVector(v *big.Int, n int) ([]*big.Int, error) {
	if v.BitLen() > n {
		return nil, errors.New("v is out of range")
	}
	bitVector := make([]*big.Int, 0)

	for i := 0; i < n; i++ {
		bitVector = append(bitVector, new(big.Int).SetUint64(uint64(v.Bit(i))))
	}

	return bitVector, nil
}

// Delta represents (z - z^2) * <1^n, y^n> - z^3 * <1^n, 2^n>.
func Delta(y, z, order *big.Int, n int) *big.Int {
	zSquare := new(big.Int).Exp(z, new(big.Int).SetUint64(2), order)
	res := new(big.Int).Sub(z, zSquare)
	res.Mod(res, order)

	yn := PowVector(y, order, n)
	res.Mul(res, yn.Sum())
	res.Mod(res, order)

	n2 := PowVector(new(big.Int).SetUint64(2), order, n)
	zCubed := new(big.Int).Exp(z, new(big.Int).SetUint64(3), order)
	tmp := zCubed.Mul(zCubed, n2.Sum())
	tmp.Mod(tmp, order)

	res.Sub(res, tmp)
	res.Mod(res, order)

	return res
}

// LowerCaseFirst converts first character of string to lowercase.
func LowerCaseFirst(data string) string {
	// Do nothing with empty string.
	if len(data) == 0 {
		return data
	}

	runeStr := []rune(data)
	runeStr[0] = unicode.ToLower(runeStr[0])
	return string(runeStr)
}
