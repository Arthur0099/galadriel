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
