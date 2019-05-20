package pgc

import (
	"math/big"
	"testing"
)

func TestPGC(t *testing.T) {
	aliceBalance := new(big.Int).SetUint64(100)
	alice := CreateTestAccount("alice", aliceBalance)

	bobBalance := new(big.Int).SetUint64(200)
	bob := CreateTestAccount("bob", bobBalance)

	transferBalance := new(big.Int).SetUint64(20)

	ctx, err := CreateCTX(alice, bob, transferBalance)
	if err != nil {
		t.Error(err)
		return
	}

	if !VerifyCTX(ctx) {
		t.Error("verify valid tx failed")
		return
	}
}
