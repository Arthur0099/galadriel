package pgcsys

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/pgc/proof"
	"github.com/pgc/utils"
)

// Account represents the account used in pgc system.
type Account struct {
	name    string
	sk      *ecdsa.PrivateKey
	balance *proof.CTEncPoint
	// balance of account. for speed up purpose only.
	m      *big.Int
	nonce  uint64
	params proof.CTParams
}

// CreateTestAccount creates account for test pupose.
func CreateTestAccount(params proof.CTParams, name string, balance *big.Int) *Account {
	a := Account{}
	a.name = name
	a.m = new(big.Int).Set(balance)
	a.nonce = 0
	a.params = params
	key := proof.MustGenerateKey(params)

	a.sk = key
	ct, err := proof.Encrypt(params, &a.sk.PublicKey, balance.Bytes())
	if err != nil {
		panic(err)
	}
	a.balance = ct.CopyPublicPoint()

	return &a
}

// UpdateBalance updates user's encrypted balance and nonce.
func (a *Account) UpdateBalance(nonce *big.Int, ct [4]*big.Int) {
	a.balance.X = utils.NewECPoint(ct[0], ct[1], a.sk.Curve)
	a.balance.Y = utils.NewECPoint(ct[2], ct[3], a.sk.Curve)
	a.nonce = nonce.Uint64()

	// decrypt balance
	balance := proof.Decrypt(a.params, a.sk, a.balance)
	a.m = new(big.Int).SetBytes(balance)
}

// // UpdateBalanceSn .
// func (a *Account) UpdateBalanceSn(m *big.Int) {
// 	a.m.Set(m)
// }

// func (a *Account) Copy() *Account {
// 	b := Account{}
// 	b.balance = a.balance
// 	b.m = a.m
// 	b.name = a.name
// 	b.nonce = a.nonce
// 	b.sk = a.sk
// 	return &b
// }
