package pgc

import (
	"errors"
	"math/big"
	"testing"
)

func TestPGCCtx(t *testing.T) {
	if err := testpgcctx(new(big.Int).SetUint64(100000000), new(big.Int).SetUint64(2000000000), new(big.Int).SetUint64(9999)); err != nil {
		t.Fatal(err)
	}

	params := Params()
	senderAmount := MustGetRandomMsg(params.BitSizeLimit())
	receiverAmount := MustGetRandomMsg(params.BitSizeLimit())
	sent := new(big.Int).Add(senderAmount, one)
	if err := testpgcctx(senderAmount, receiverAmount, sent); err == nil {
		t.Fatal("invalid transfer passed")
	}
}

func testpgcctx(senderAmount, receiver, amount *big.Int) error {
	alice := CreateTestAccount("alice", senderAmount)
	bob := CreateTestAccount("bob", receiver)

	ctx, err := CreatePGCCtx(alice, &bob.sk.PublicKey, amount)
	if err != nil {
		return err
	}

	if !VerifyPGCCtx(ctx) {
		return errors.New("verify pgc failed")
	}

	return nil
}

func BenchmarkGenerateCTX(b *testing.B) {
	senderAmount := new(big.Int).SetUint64(10000)
	receiver := new(big.Int).SetUint64(20000)
	amount := new(big.Int).SetUint64(9000)

	alice := CreateTestAccount("alice", senderAmount)
	bob := CreateTestAccount("bob", receiver)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := CreatePGCCtx(alice, &bob.sk.PublicKey, amount)
		if err != nil {
			b.Fatal(err)
		}
		// if !VerifyPGCCtx(ctx) {
		// 	b.Fatal(errors.New("failed"))
		// }
	}

}

func BenchmarkVerifyCTX(b *testing.B) {
	senderAmount := new(big.Int).SetUint64(10000)
	receiver := new(big.Int).SetUint64(20000)
	amount := new(big.Int).SetUint64(9000)

	alice := CreateTestAccount("alice", senderAmount)
	bob := CreateTestAccount("bob", receiver)
	ctx, err := CreatePGCCtx(alice, &bob.sk.PublicKey, amount)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !VerifyPGCCtx(ctx) {
			b.Fatal(errors.New("failed"))
		}
	}
}
