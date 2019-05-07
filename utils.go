package pgc

import (
	"math/big"

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
