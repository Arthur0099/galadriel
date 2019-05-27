package pgc

import (
	"encoding/json"
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

	balanceAfter := new(big.Int).SetBytes(Decrypt(alice.sk, alice.balance))
	except := new(big.Int).Sub(aliceBalance, transferBalance)
	if except.Cmp(balanceAfter) != 0 {
		t.Error("balance after not correct")
	}
}

// depoist/burn account should be same.
type DepositBurnTest struct {
	Deposit []DepositTest `json:"deposit"`
	Burn    []*BurnTx     `json:"burn"`
}

type DepositTest struct {
	Account *ECPoint
	Balance *big.Int
}

func (d *DepositTest) MarshalJSON() ([]byte, error) {
	newJSON := struct {
		Account *ECPoint `json:"account"`
		Balance string   `json:"balance"`
	}{
		Account: d.Account,
		Balance: d.Balance.String(),
	}

	return json.Marshal(&newJSON)
}

// to test on chain.
func TestDepositBurn(t *testing.T) {
	test := DepositBurnTest{}
	// deposit to an account.
	aliceInitBalance := new(big.Int).SetUint64(10)
	alice := CreateTestAccount("alice", aliceInitBalance)
	// same with chain.
	aliceChain, _ := EncryptOnChain(&alice.sk.PublicKey, aliceInitBalance.Bytes())
	alice.balance = aliceChain.CopyPublicPoint()

	// deposit to bob.
	bobInitBalance := new(big.Int).SetUint64(20)
	bob := CreateTestAccount("bob", bobInitBalance)
	bobChain, _ := EncryptOnChain(&bob.sk.PublicKey, bobInitBalance.Bytes())
	bob.balance = bobChain.CopyPublicPoint()

	test.Deposit = make([]DepositTest, 0)
	// deposit alice.
	depositAlice := DepositTest{}
	depositAlice.Account = new(ECPoint).SetFromPublicKey(&alice.sk.PublicKey)
	depositAlice.Balance = aliceInitBalance
	test.Deposit = append(test.Deposit, depositAlice)

	// deposit bob.
	depositBob := DepositTest{}
	depositBob.Account = new(ECPoint).SetFromPublicKey(&bob.sk.PublicKey)
	depositBob.Balance = bobInitBalance
	test.Deposit = append(test.Deposit, depositBob)

	// burn an account.
	burnAlice, err := CreateBurnTx(alice, aliceInitBalance)
	if err != nil {
		t.Error(err)
		return
	}

	params := Params()
	g := params.GetG()
	h := params.GetH()
	aliceY := new(ECPoint).Sub(alice.balance.Y, new(ECPoint).ScalarMult(g, aliceInitBalance))
	// check for dle sigma proof.
	if !verifyDLESigmaProof(aliceY, alice.balance.X, h, new(ECPoint).SetFromPublicKey(&alice.sk.PublicKey), burnAlice.Proof) {
		t.Error("alice burn proof invalid")
		return
	}
	burnBob, err := CreateBurnTx(bob, bobInitBalance)
	bobY := new(ECPoint).Sub(bob.balance.Y, new(ECPoint).ScalarMult(g, bobInitBalance))
	if !verifyDLESigmaProof(bobY, bob.balance.X, h, new(ECPoint).SetFromPublicKey(&bob.sk.PublicKey), burnBob.Proof) {
		t.Error("bob burn proof invalid")
		return
	}
	if err != nil {
		t.Error(err)
		return
	}

	test.Burn = make([]*BurnTx, 0)
	test.Burn = append(test.Burn, burnAlice)
	test.Burn = append(test.Burn, burnBob)

	data, err := json.Marshal(test)
	if err != nil {
		t.Error(err)
		return
	}

	path := "solidity/proofs/depositBurn"
	WriteToFile(data, path)
}
