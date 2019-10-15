package pgc

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pgc/contracts"
)

var (
	// GWEI is the uint of gas price
	GWEI         = new(big.Int).SetUint64(1000 * 1000 * 1000)
	testGasLimit = uint64(18000000)
	one          = new(big.Int).SetUint64(1)
	two          = new(big.Int).SetUint64(2)
	zero         = new(big.Int).SetUint64(0)
)

var parsed abi.ABI

// ComputeChallengeByECPoints computes challenge like (ecpoint.X, ecpoint.Y, ...).
func ComputeChallengeByECPoints(order *big.Int, points ...*ECPoint) (*big.Int, error) {
	datas := make([]interface{}, 0)

	for _, point := range points {
		datas = append(datas, point.X)
		datas = append(datas, point.Y)
	}

	return ComputeChallenge(order, datas...)
}

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

// HashTransfer returns the hash of transfer input.
func HashTransfer(tx *transferTx, token common.Address) ([]byte, error) {
	return hash(parsed, "transfer", tx.points, tx.scalar, tx.rpoints, tx.l, tx.r, new(big.Int).SetBytes(token.Bytes()), tx.nonce)
}

// HashBurn returns the hash of burn input.
func HashBurn(receiver, token common.Address, tx *burnTx) ([]byte, error) {
	return hash(parsed, "verifyBurnSig", new(big.Int).SetBytes(receiver.Bytes()), new(big.Int).SetBytes(token.Bytes()), tx.amount, tx.pk, tx.proof, tx.z, tx.nonce)
}

// HashBurnPart returns hash of burnPart.
func HashBurnPart(receiver, token common.Address, tx *burnPartTx) ([]byte, error) {
	return hash(parsed, "verifyBurnPartSig", new(big.Int).SetBytes(receiver.Bytes()), new(big.Int).SetBytes(token.Bytes()), tx.amount, tx.points, tx.scalar, tx.rpoints, tx.l, tx.r, tx.nonce)
}

// HashOpenPending returns hash of openPending.
func HashOpenPending(x, y, epoch, nonce *big.Int) ([]byte, error) {
	return hash(parsed, "openPending", x, y, epoch, nonce)
}

// HashClosePending returns hash of closePending.
func HashClosePending(x, y, nonce *big.Int) ([]byte, error) {
	return hash(parsed, "closePending", x, y, nonce)
}

func hash(parsedABI abi.ABI, name string, params ...interface{}) ([]byte, error) {
	method, exist := parsed.Methods[name]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", name)
	}

	if len(method.Inputs) <= 1 {
		return nil, fmt.Errorf("no enough inputs for compute hash for '%s", name)
	}
	arguments := method.Inputs[0 : len(method.Inputs)-1]
	data, err := arguments.Pack(params...)
	if err != nil {
		return nil, err
	}

	return Keccak256(data), nil
}

// Sha256Hash returns hash bytes using sha256.
func Sha256Hash(data []byte) []byte {
	h := sha256.New()
	// Hash.Write never returns an error per godoc
	h.Write(data)
	return h.Sum(nil)
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

// MultBitVector returns vector appending bitVector of vi.
func MultBitVector(v []*big.Int, n int) ([]*big.Int, error) {
	bitVector := make([]*big.Int, 0)
	for i := 0; i < len(v); i++ {
		tmp, err := BitVector(v[i], n)
		if err != nil {
			return nil, err
		}

		bitVector = append(bitVector, tmp...)
	}

	return bitVector, nil
}

// Delta represents (z - z^2) * <1^n, y^n> - z^3 * <1^n, 2^n>.
func Delta(y, z, order *big.Int, n int) *big.Int {
	zSquare := new(big.Int).Exp(z, two, order)
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

// DeltaMN represents (z - z^2) * <1^mn, y^mn> - z^(j+1) * <1^n, 2^n>. (j is [1, m])
func DeltaMN(y, z, order *big.Int, m, n int) *big.Int {
	zSquare := new(big.Int).Exp(z, two, order)
	res := new(big.Int).Sub(z, zSquare)
	res.Mod(res, order)

	ymn := PowVector(y, order, n*m)
	res.Mul(res, ymn.Sum())
	res.Mod(res, order)

	n2Sum := PowVector(two, order, n).Sum()
	tmp := new(big.Int)
	for j := 1; j <= m; j++ {
		zjtmp := new(big.Int).Exp(z, new(big.Int).SetUint64(uint64(j+2)), order)
		zjtmp.Mul(zjtmp, n2Sum)
		tmp.Add(tmp, zjtmp)
	}
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

// MustGetRandomMsg returns a random msg less than 2^bitSize.
// Warn: test purpose only.
func MustGetRandomMsg(bitSize int) *big.Int {
	n := new(big.Int).Exp(new(big.Int).SetUint64(2), new(big.Int).SetUint64(uint64(bitSize)), nil)
	r, err := rand.Int(rand.Reader, n)
	if err != nil {
		panic(err)
	}

	return r
}

// MustGetRandom returns a r.
func MustGetRandom(n *big.Int) *big.Int {
	r, err := rand.Int(rand.Reader, n)
	if err != nil {
		panic(err)
	}

	return r
}

func init() {
	parsed, _ = abi.JSON(strings.NewReader(contracts.PgcABI))
}
