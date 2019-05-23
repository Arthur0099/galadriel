package pgc

import (
	"math/big"
	"testing"
)

func TestPGC(t *testing.T) {
	aliceBalance := new(big.Int).SetUint64(12)
	alice := CreateTestAccount("alice", aliceBalance)

	bobBalance := new(big.Int).SetUint64(12)
	bob := CreateTestAccount("bob", bobBalance)

	transferBalance := new(big.Int).SetUint64(2)

	ctx, err := CreateCTX(alice, bob, transferBalance)
	if err != nil {
		t.Error(err)
		return
	}

	if !VerifyCTX(ctx) {
		t.Error("verify valid tx failed")
		return
	}

	sys := NewTwistedELGamalSystem()
	balanceAfter := new(big.Int).SetBytes(sys.Decrypt(alice.sk, alice.balance))
	except := new(big.Int).Sub(aliceBalance, transferBalance)
	if except.Cmp(balanceAfter) != 0 {
		t.Error("balance after not correct")
	}
}
