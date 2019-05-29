package pgc

import (
	"crypto/elliptic"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
)

var ether = new(big.Int).SetUint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000)

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

//
func TestPGCFlow(t *testing.T) {
	url := "http://127.0.0.1:8545"
	rpcClient, _ := Dial(url)
	client, _ := ethclient.Dial(url)

	senderKey, _ := crypto.GenerateKey()
	sender := bind.NewKeyedTransactor(senderKey)
	sender.GasLimit = 8000000
	accounts, _ := rpcClient.GetAccounts()

	amount := new(big.Int).SetUint64(10000)
	_, err := rpcClient.SendETH(accounts[0], sender.From, amount)
	if err != nil {
		t.Error(err)
		return
	}

	cs := DeployPGCContracts(sender, client)

	// generate alice, bob account.
	aliceInitBalance := new(big.Int).SetUint64(100)
	alice := CreateTestAccount("alice", aliceInitBalance)
	sender.Value = new(big.Int).Mul(ether, aliceInitBalance)
	alicePK := [2]*big.Int{alice.sk.PublicKey.X, alice.sk.PublicKey.Y}
	_, err = cs.PGC.DepositAccount(sender, alicePK)
	if err != nil {
		t.Error(err)
		return
	}
	// check for alice's encrypted balance.
	aliceEncryptB, _ := cs.PGC.GetUserBalance(nil, alice.sk.PublicKey.X, alice.sk.PublicKey.Y)
	aliceDecryptB := Decrypt(alice.sk, arrayToCT(aliceEncryptB, alice.sk.Curve))
	if aliceInitBalance.Cmp(new(big.Int).SetBytes(aliceDecryptB)) != 0 {
		t.Error("encrypt on chain not same alice")
		return
	}
	// keep balance same with chain.
	alice.balance = arrayToCT(aliceEncryptB, alice.sk.Curve)

	// generate bob.
	bobInitBalance := new(big.Int).SetUint64(500)
	bob := CreateTestAccount("bob", bobInitBalance)
	sender.Value = new(big.Int).Mul(ether, bobInitBalance)
	sender.Nonce.Add(sender.Nonce, one)
	bobPK := [2]*big.Int{bob.sk.PublicKey.X, bob.sk.PublicKey.Y}
	_, err = cs.PGC.DepositAccount(sender, bobPK)
	if err != nil {
		t.Error(err)
		return
	}
	sender.Nonce.Add(sender.Nonce, one)
	// check for bob's encrypted balance.
	bobEncryptB, _ := cs.PGC.GetUserBalance(nil, bob.sk.PublicKey.X, bob.sk.PublicKey.Y)
	bobDecryptB := Decrypt(bob.sk, arrayToCT(bobEncryptB, bob.sk.Curve))
	if bobInitBalance.Cmp(new(big.Int).SetBytes(bobDecryptB)) != 0 {
		t.Error("encrypt on chain not same bob")
		return
	}
	// keep balance same with chain.
	bob.balance = arrayToCT(bobEncryptB, bob.sk.Curve)

	transferAmount := new(big.Int).SetUint64(30)
	ctx, err := CreateCTX(alice, bob, transferAmount)
	if err != nil {
		t.Error(err)
		return
	}

	// send transfer tx to contract on chain.
	tx := ctxToTransferTx(ctx)
	sender.Value = nil
	transferTxHash, err := cs.PGC.Transfer(sender, tx.points, tx.scalar, tx.rpoints, tx.l, tx.r)
	if err != nil {
		t.Error(err)
		return
	}
	// cost 6444057 gas
	t.Log("transfer from alice to bob", transferAmount, transferTxHash.Hash().Hex())

	// check alice/bob balance.
	aliceEncBalanceAfter, _ := cs.PGC.GetUserBalance(nil, alicePK[0], alicePK[1])
	aliceBalanceAfter := Decrypt(alice.sk, arrayToCT(aliceEncBalanceAfter, alice.sk.Curve))
	aliceExceptBalance := new(big.Int).Sub(aliceInitBalance, transferAmount)
	if aliceExceptBalance.Cmp(new(big.Int).SetBytes(aliceBalanceAfter)) != 0 {
		t.Error("alice balance after not right")
		return
	}
	t.Log("alice'balance after is", new(big.Int).SetBytes(aliceBalanceAfter))

	// check for bob's balance.
	bobEncryptBalanceAfter, _ := cs.PGC.GetUserBalance(nil, bobPK[0], bobPK[1])
	bobBalanceAfter := Decrypt(bob.sk, arrayToCT(bobEncryptBalanceAfter, bob.sk.Curve))
	bobExceptBalance := new(big.Int).Add(bobInitBalance, transferAmount)
	if bobExceptBalance.Cmp(new(big.Int).SetBytes(bobBalanceAfter)) != 0 {
		t.Error("bob balance after not right")
		return
	}
	t.Log("bob's balance after is ", new(big.Int).SetBytes(bobBalanceAfter))

}

type transferTx struct {
	points  [28]*big.Int
	scalar  [14]*big.Int
	rpoints [16]*big.Int
	l       [16]*big.Int
	r       [16]*big.Int
}

func newTransferTx() *transferTx {
	tx := transferTx{}
	tx.points = [28]*big.Int{}
	tx.scalar = [14]*big.Int{}
	tx.rpoints = [16]*big.Int{}
	tx.l = [16]*big.Int{}
	tx.r = [16]*big.Int{}

	return &tx
}

func arrayToCT(p [4]*big.Int, curve elliptic.Curve) *CTEncPoint {
	c := CTEncPoint{}
	c.X = NewECPoint(p[0], p[1], curve)
	c.Y = NewECPoint(p[2], p[3], curve)

	return &c
}

func ctxToTransferTx(ctx *CTX) *transferTx {
	tx := newTransferTx()
	tx.points[0] = ctx.pk1.X
	tx.points[1] = ctx.pk1.Y
	tx.points[2] = ctx.c1.X.X
	tx.points[3] = ctx.c1.X.Y
	tx.points[4] = ctx.c1.Y.X
	tx.points[5] = ctx.c1.Y.Y
	tx.points[6] = ctx.pk2.X
	tx.points[7] = ctx.pk2.Y
	tx.points[8] = ctx.c2.X.X
	tx.points[9] = ctx.c2.X.Y
	tx.points[10] = ctx.c2.Y.X
	tx.points[11] = ctx.c2.Y.Y
	tx.points[12] = ctx.equalityProof.A1.X
	tx.points[13] = ctx.equalityProof.A1.Y
	tx.points[14] = ctx.equalityProof.A2.X
	tx.points[15] = ctx.equalityProof.A2.Y
	tx.points[16] = ctx.equalityProof.B1.X
	tx.points[17] = ctx.equalityProof.B1.Y
	tx.points[18] = ctx.equalityProof.B2.X
	tx.points[19] = ctx.equalityProof.B2.Y
	tx.points[20] = ctx.refreshUpdatedBalance.X.X
	tx.points[21] = ctx.refreshUpdatedBalance.X.Y
	tx.points[22] = ctx.refreshUpdatedBalance.Y.X
	tx.points[23] = ctx.refreshUpdatedBalance.Y.Y
	tx.points[24] = ctx.dleProof.A1.X
	tx.points[25] = ctx.dleProof.A1.Y
	tx.points[26] = ctx.dleProof.A2.X
	tx.points[27] = ctx.dleProof.A2.Y

	// scalar.
	tx.scalar[0] = ctx.equalityProof.z1
	tx.scalar[1] = ctx.equalityProof.z2
	tx.scalar[2] = ctx.equalityProof.z3
	tx.scalar[3] = ctx.dleProof.Z
	tx.scalar[4] = ctx.rangeProof1.t
	tx.scalar[5] = ctx.rangeProof1.tx
	tx.scalar[6] = ctx.rangeProof1.u
	tx.scalar[7] = ctx.rangeProof1.ipProof.a
	tx.scalar[8] = ctx.rangeProof1.ipProof.b
	tx.scalar[9] = ctx.rangeProof2.t
	tx.scalar[10] = ctx.rangeProof2.tx
	tx.scalar[11] = ctx.rangeProof2.u
	tx.scalar[12] = ctx.rangeProof2.ipProof.a
	tx.scalar[13] = ctx.rangeProof2.ipProof.b

	//
	tx.rpoints[0] = ctx.rangeProof1.A.X
	tx.rpoints[1] = ctx.rangeProof1.A.Y
	tx.rpoints[2] = ctx.rangeProof1.S.X
	tx.rpoints[3] = ctx.rangeProof1.S.Y
	tx.rpoints[4] = ctx.rangeProof1.T1.X
	tx.rpoints[5] = ctx.rangeProof1.T1.Y
	tx.rpoints[6] = ctx.rangeProof1.T2.X
	tx.rpoints[7] = ctx.rangeProof1.T2.Y
	tx.rpoints[8] = ctx.rangeProof2.A.X
	tx.rpoints[9] = ctx.rangeProof2.A.Y
	tx.rpoints[10] = ctx.rangeProof2.S.X
	tx.rpoints[11] = ctx.rangeProof2.S.Y
	tx.rpoints[12] = ctx.rangeProof2.T1.X
	tx.rpoints[13] = ctx.rangeProof2.T1.Y
	tx.rpoints[14] = ctx.rangeProof2.T2.X
	tx.rpoints[15] = ctx.rangeProof2.T2.Y

	//
	base := len(ctx.rangeProof1.ipProof.L) * 2
	for i, v := range ctx.rangeProof1.ipProof.L {
		tx.l[2*i] = v.X
		tx.l[2*i+1] = v.Y
		tx.r[2*i] = ctx.rangeProof1.ipProof.R[i].X
		tx.r[2*i+1] = ctx.rangeProof1.ipProof.R[i].Y

		tx.l[2*i+base] = ctx.rangeProof2.ipProof.L[i].X
		tx.l[2*i+1+base] = ctx.rangeProof2.ipProof.L[i].Y
		tx.r[2*i+base] = ctx.rangeProof2.ipProof.R[i].X
		tx.r[2*i+1+base] = ctx.rangeProof2.ipProof.R[i].Y
	}

	return tx
}
