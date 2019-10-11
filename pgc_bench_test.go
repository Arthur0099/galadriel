package pgc

import (
	"errors"
	"math/big"
	"testing"
)

func BenchmarkGeneratePGCCTX(b *testing.B) {
	senderAmount := new(big.Int).SetUint64(10000)
	receiver := new(big.Int).SetUint64(20000)
	amount := new(big.Int).SetUint64(900)

	alice := CreateTestAccount("alice", senderAmount)
	bob := CreateTestAccount("bob", receiver)

	for i := 0; i < b.N; i++ {
		_, err := CreateCTX(alice, &bob.sk.PublicKey, amount)
		if err != nil {
			b.Fatal(err)
		}
		// if !VerifyCTX(ctx) {
		// 	b.Fatal(errors.New("failed"))
		// }
	}
}

func BenchmarkVerifyPGCCTX(b *testing.B) {
	b.StopTimer()
	senderAmount := new(big.Int).SetUint64(10000)
	receiver := new(big.Int).SetUint64(20000)
	amount := new(big.Int).SetUint64(10000)

	alice := CreateTestAccount("alice", senderAmount)
	bob := CreateTestAccount("bob", receiver)
	ctx, err := CreateCTX(alice, &bob.sk.PublicKey, amount)
	if err != nil {
		b.Fatal(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !VerifyCTX(ctx) {
			b.Fatal(errors.New("failed"))
		}
	}
}
