package pgc

import (
	"crypto/ecdsa"
	"math/big"
)

// Account represents the account used in pgc system.
type Account struct {
	name    string
	sk      *ecdsa.PrivateKey
	balance *CTEncPoint
	// balance of account. for speed up purpose only.
	m     *big.Int
	nonce uint64
}

// CreateTestAccount creates account for test pupose.
func CreateTestAccount(name string, balance *big.Int) *Account {
	a := Account{}
	a.name = name
	a.m = new(big.Int).Set(balance)
	a.nonce = 0
	key, err := GenerateKey()
	if err != nil {
		panic(err)
	}
	a.sk = key
	ct, err := Encrypt(&a.sk.PublicKey, balance.Bytes())
	if err != nil {
		panic(err)
	}
	a.balance = ct.CopyPublicPoint()

	return &a
}

// UpdateBalance updates user's encrypted balance and nonce.
func (a *Account) UpdateBalance(data CT) {
	a.balance.X = NewECPoint(data.Ct[0], data.Ct[1], a.sk.Curve)
	a.balance.Y = NewECPoint(data.Ct[2], data.Ct[3], a.sk.Curve)
	a.nonce = data.Nonce.Uint64()
}
