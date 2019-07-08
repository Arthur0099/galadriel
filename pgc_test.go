package pgc

import (
	"crypto/elliptic"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pgc/contracts"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var ether = new(big.Int).SetUint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000)

func TestPGC(t *testing.T) {
	// create alice account.
	aliceBalance := new(big.Int).SetUint64(12)
	alice := CreateTestAccount("alice", aliceBalance)

	// create bob account.
	bobBalance := new(big.Int).SetUint64(12)
	bob := CreateTestAccount("bob", bobBalance)

	transferBalance := new(big.Int).SetUint64(2)

	// create ctx.
	ctx, err := CreateCTX(alice, &bob.sk.PublicKey, transferBalance)
	require.Nil(t, err, "create ctx failed")
	require.True(t, VerifyCTX(ctx), "verify valid ctx failed")

	// check for balance.
	balanceAfter := new(big.Int).SetBytes(Decrypt(alice.sk, alice.balance))
	except := new(big.Int).Sub(aliceBalance, transferBalance)
	require.Equal(t, balanceAfter, except, "")
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
	require.Nil(t, err, "alice create burntx failed")

	params := Params()
	g := params.GetG()
	h := params.GetH()
	aliceY := new(ECPoint).Sub(alice.balance.Y, new(ECPoint).ScalarMult(g, aliceInitBalance))
	// check for dle sigma proof.
	require.True(t, verifyDLESigmaProof(aliceY, alice.balance.X, h, new(ECPoint).SetFromPublicKey(&alice.sk.PublicKey), burnAlice.Proof), "alice burn proof invalid")

	burnBob, err := CreateBurnTx(bob, bobInitBalance)
	bobY := new(ECPoint).Sub(bob.balance.Y, new(ECPoint).ScalarMult(g, bobInitBalance))
	require.Nil(t, err, "bob create burntx failed")
	require.True(t, verifyDLESigmaProof(bobY, bob.balance.X, h, new(ECPoint).SetFromPublicKey(&bob.sk.PublicKey), burnBob.Proof), "bob burn proof invalid")

	test.Burn = make([]*BurnTx, 0)
	test.Burn = append(test.Burn, burnAlice)
	test.Burn = append(test.Burn, burnBob)

	data, _ := json.Marshal(test)
	path := "solidity/proofs/depositBurn"
	WriteToFile(data, path)
}

func getAccountByHexKey(hexKey string) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		panic(err)
	}

	return bind.NewKeyedTransactor(key)
}

func TestLocal(t *testing.T) {
	rpcClient := GetLocalRPC()
	client := GetLocal()

	senderKey, _ := crypto.GenerateKey()
	sender := bind.NewKeyedTransactor(senderKey)
	sender.GasLimit = testGasLimit
	sender.GasPrice = new(big.Int).Mul(new(big.Int).SetUint64(10), GWEI)
	accounts, err := rpcClient.GetAccounts()
	require.Nil(t, err)

	amount := new(big.Int).SetUint64(10000)
	_, err = rpcClient.SendETH(accounts[0], sender.From, amount)
	require.Nil(t, err, "send eth failed")

	cs := DeployPGCContracts(sender, client)

	// test for eth.
	testPGCFlow(sender, client, t, nil, common.Address{}, cs)

	// test for token.
	tokenContract, tokenAddr := DeployToken(sender, client)
	testPGCFlow(sender, client, t, tokenContract, tokenAddr, cs)

	// test for pending capacity.
	testPGCPending(sender, rpcClient, client, t, tokenContract, tokenAddr, cs)
}

func TestRopsten(t *testing.T) {
	sender := GetRopstenAccount()
	sender.GasPrice = new(big.Int).SetUint64(10 * 1000 * 1000 * 1000)
	client := GetRopstenInfura()
	cs := DeployPGCContracts(sender, client)
	//testPGCFlow(sender, client, t, nil, common.Address{}, cs)

	// test for token.
	tokenContract, tokenAddr := DeployToken(sender, client)
	testPGCFlow(sender, client, t, tokenContract, tokenAddr, cs)

	// test for pending capacity.
	testPGCPending(sender, nil, client, t, tokenContract, tokenAddr, cs)
}

func testPGCPending(sender *bind.TransactOpts, rpcClient *Client, client *ethclient.Client, t *testing.T, tokenContract *contracts.Token, token common.Address, cs *Contracts) {
	var err error
	sender.GasLimit = testGasLimit

	// if test for token.
	// approve to contract first.
	if tokenContract != nil {
		approveAmount := new(big.Int).SetUint64(1000)

		approveTx, err := tokenContract.Approve(sender, cs.PGCAddress, approveAmount)
		require.Nil(t, err, "approve for token failed")
		waitFor(approveTx.Hash(), client)
		sender.Nonce.Add(sender.Nonce, one)

		added, _, _, _ := cs.TokenConverter.GetTokenInfo(nil, token)
		if !added {
			// add token contract to tokenconvert.
			addTokenTx, err := cs.TokenConverter.AddToken(sender, token, one, "")
			if err != nil {
				panic(err)
			}
			t.Log("add token to tokenconverter", "tx", addTokenTx.Hash().Hex())
			waitFor(addTokenTx.Hash(), client)
			sender.Nonce.Add(sender.Nonce, one)

		}

	}

	// generate alice, bob account.
	aliceInitBalance := new(big.Int).SetUint64(500)
	alice := CreateTestAccount("alice", aliceInitBalance)
	sender.Value = new(big.Int).Mul(ether, aliceInitBalance)
	sender.Value.Div(sender.Value, precision)
	if tokenContract != nil {
		sender.Value = nil
	}
	alicePK := [2]*big.Int{alice.sk.PublicKey.X, alice.sk.PublicKey.Y}
	aliceTx, err := cs.PGC.DepositAccount(sender, alicePK, token, aliceInitBalance)
	require.Nil(t, err, "deposit contract failed")
	waitFor(aliceTx.Hash(), client)
	sender.Nonce.Add(sender.Nonce, one)

	// check for alice's encrypted balance.
	aliceEncryptB, _ := cs.PGC.GetUserBalance(nil, alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	aliceDecryptB := Decrypt(alice.sk, arrayToCT(aliceEncryptB, alice.sk.Curve))
	require.Equal(t, aliceInitBalance.Bytes(), aliceDecryptB, "alice'balance on chain not same with local")

	// keep balance same with chain.
	alice.UpdateBalance(aliceEncryptB)

	// generate bob.
	bobInitBalance := new(big.Int).SetUint64(500)
	bob := CreateTestAccount("bob", bobInitBalance)
	// bob open pending capacity.
	bobEpoch := new(big.Int).SetUint64(50)
	bobOpenPendingHash, err := HashOpenPending(bob.sk.PublicKey.X, bob.sk.PublicKey.Y, bobEpoch)
	require.Nil(t, err, "hash open pending failed")
	bobOpenPendingSig, err := Sign(bob.sk, bobOpenPendingHash)
	require.Nil(t, err, "sig for open pending failed")
	openPendingTx, err := cs.PGC.OpenPending(sender, bob.sk.PublicKey.X, bob.sk.PublicKey.Y, bobEpoch, bobOpenPendingSig.ToInputs())
	require.Nil(t, err, "open pending tx failed")
	t.Log("open pending tx", openPendingTx.Hash().Hex())
	waitFor(openPendingTx.Hash(), client)
	sender.Value = new(big.Int).Mul(ether, bobInitBalance)
	sender.Value.Div(sender.Value, precision)
	if tokenContract != nil {
		sender.Value = nil
	}
	sender.Nonce.Add(sender.Nonce, one)
	bobPK := [2]*big.Int{bob.sk.PublicKey.X, bob.sk.PublicKey.Y}
	bobTx, err := cs.PGC.DepositAccount(sender, bobPK, token, bobInitBalance)
	require.Nil(t, err, "deposit to bob on chain failed")
	waitFor(bobTx.Hash(), client)
	sender.Nonce.Add(sender.Nonce, one)

	// check for bob's encrypted balance.
	bobEncryptB, _ := cs.PGC.GetUserBalance(nil, bob.sk.PublicKey.X, bob.sk.PublicKey.Y, token)
	bobDecryptB := Decrypt(bob.sk, arrayToCT(bobEncryptB, bob.sk.Curve))
	// because bob open pending capacity, so it won't update bob's balance(can spend) directly, pending it in epoch it arrives.
	require.Equal(t, new(big.Int).SetUint64(0).Bytes(), bobDecryptB, "bob's balance updated in pending epoch(wrong)")

	// advance block number after block.
	if rpcClient != nil {
		_ = rpcClient.Mine(int(bobEpoch.Uint64()))
	} else {
		waitForBlocks(bobEpoch.Uint64(), client)
	}

	// at this epoch, bob can use pgc sent before current epoch.
	bobEncryptB, _ = cs.PGC.GetUserBalance(nil, bob.sk.PublicKey.X, bob.sk.PublicKey.Y, token)
	bobDecryptB = Decrypt(bob.sk, arrayToCT(bobEncryptB, bob.sk.Curve))
	require.Equal(t, bobInitBalance.Bytes(), bobDecryptB, "bob's balance not same with excepted after pending epoch")
	// keep balance same with chain.
	bob.UpdateBalance(bobEncryptB)

	// alice open pending capacity.
	aliceEpoch := new(big.Int).SetUint64(50)
	aliceOpenPendingHash, _ := HashOpenPending(alice.sk.PublicKey.X, alice.sk.PublicKey.Y, aliceEpoch)
	aliceOpenPendingSig, _ := Sign(alice.sk, aliceOpenPendingHash)
	aliceOpenPendingTx, err := cs.PGC.OpenPending(sender, alice.sk.PublicKey.X, alice.sk.PublicKey.Y, aliceEpoch, aliceOpenPendingSig.ToInputs())
	require.Nil(t, err, "alice open pending capacity tx failed")
	t.Log("open pending tx", aliceOpenPendingTx.Hash().Hex())
	waitFor(aliceOpenPendingTx.Hash(), client)
	sender.Nonce.Add(sender.Nonce, one)

	// bob transfer 50 amount to alice.
	transferAmount := new(big.Int).SetUint64(50)
	ctx, err := CreateCTX(bob, &alice.sk.PublicKey, transferAmount)
	require.Nil(t, err, "create transfer tx from bob to alice failed")
	transferTx := ctxToTransferTx(ctx)
	transferHash, err := HashTransfer(transferTx, token)
	require.Nil(t, err, "hash transfer data failed")
	transferSig, err := Sign(bob.sk, transferHash)
	require.Nil(t, err, "sig for transfer tx failed")
	sender.Value = nil
	transferTxHash, err := cs.PGC.Transfer(sender, transferTx.points, transferTx.scalar, transferTx.rpoints, transferTx.l, transferTx.r, new(big.Int).SetBytes(token.Bytes()), transferTx.nonce, transferSig.ToInputs())
	require.Nil(t, err, "transfer tx failed")
	waitFor(transferTxHash.Hash(), client)
	t.Log("transfer from bob to alice", transferAmount, transferTxHash.Hash().Hex())
	sender.Nonce.Add(sender.Nonce, one)

	// check for bob's and alice's balance.
	bobEncryptBalanceAfter, _ := cs.PGC.GetUserBalance(nil, bob.sk.PublicKey.X, bob.sk.PublicKey.Y, token)
	bobExceptBalance := new(big.Int).Sub(bobInitBalance, transferAmount)
	bobBalanceAfter := Decrypt(bob.sk, arrayToCT(bobEncryptBalanceAfter, bob.sk.Curve))
	require.Equal(t, bobExceptBalance.Bytes(), bobBalanceAfter, "bob balance after transfer is not correct")
	aliceEncrypteBalanceAfter, _ := cs.PGC.GetUserBalance(nil, alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	aliceExpectBalance := aliceInitBalance
	aliceBalanceAfter := Decrypt(alice.sk, arrayToCT(aliceEncrypteBalanceAfter, alice.sk.Curve))
	require.Equal(t, aliceExpectBalance.Bytes(), aliceBalanceAfter, "alice's balance not correct after transfer")

	// alice close pending capacity.
	closePendingHash, _ := HashClosePending(alice.sk.PublicKey.X, alice.sk.PublicKey.Y)
	closePendingSig, _ := Sign(alice.sk, closePendingHash)
	closePendingTx, err := cs.PGC.ClosePending(sender, alice.sk.PublicKey.X, alice.sk.PublicKey.Y, closePendingSig.ToInputs())
	require.Nil(t, err, "alice close pending failed")
	t.Log("alice close pending tx", closePendingTx.Hash().Hex())
	waitFor(closePendingTx.Hash(), client)
	sender.Nonce.Add(sender.Nonce, one)

	// alice burn part amount.
	aliceBalanceEn, _ := cs.PGC.GetUserBalance(nil, alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	aliceBalanceEx := new(big.Int).Add(aliceInitBalance, transferAmount)
	aliceBalanceDe := Decrypt(alice.sk, arrayToCT(aliceBalanceEn, alice.sk.Curve))
	require.Equal(t, aliceBalanceEx.Bytes(), aliceBalanceDe, "alice's balance not correct after close pending capacity")
	alice.UpdateBalance(aliceBalanceEn)
	alice.UpdateBalanceSn(aliceBalanceEx)
	burnAmount := new(big.Int).SetUint64(520)
	burnPartTx, err := CreateBurnPartTx(alice, burnAmount)
	require.Nil(t, err, "create burn part tx failed")
	receiver := sender.From
	burnPartTTx := txToBurnPartTx(burnPartTx)
	burnHash, _ := HashBurnPart(receiver, token, burnPartTTx)
	burnSig, _ := Sign(alice.sk, burnHash)
	burnTx, err := cs.PGC.BurnPart(sender, receiver, new(big.Int).SetBytes(token.Bytes()), burnPartTTx.amount, burnPartTTx.points, burnPartTTx.scalar, burnPartTTx.rpoints, burnPartTTx.l, burnPartTTx.r, burnPartTTx.nonce, burnSig.ToInputs())
	t.Log("burn part tx", burnTx.Hash().Hex())
	waitFor(burnTx.Hash(), client)

	// check alice's balance
	aliceFinalEnB, _ := cs.PGC.GetUserBalance(nil, alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	aliceFinalExpect := new(big.Int).Sub(aliceBalanceEx, burnAmount)
	aliceFinalActual := Decrypt(alice.sk, arrayToCT(aliceFinalEnB, alice.sk.Curve))
	require.Equal(t, aliceFinalExpect.Bytes(), aliceFinalActual, "aclie' balance not correct")
}

func testPGCFlow(sender *bind.TransactOpts, client *ethclient.Client, t *testing.T, tokenContract *contracts.Token, token common.Address, cs *Contracts) {
	var err error
	sender.GasLimit = testGasLimit

	// if test for token.
	// approve to contract first.
	if tokenContract != nil {
		approveAmount := new(big.Int).SetUint64(1000)

		approveTx, err := tokenContract.Approve(sender, cs.PGCAddress, approveAmount)
		require.Nil(t, err, "approve for token failed")
		waitFor(approveTx.Hash(), client)
		sender.Nonce.Add(sender.Nonce, one)

		added, _, _, _ := cs.TokenConverter.GetTokenInfo(nil, token)
		if !added {
			// add token contract to tokenconvert.
			addTokenTx, err := cs.TokenConverter.AddToken(sender, token, one, "")
			if err != nil {
				panic(err)
			}
			t.Log("add token to tokenconverter", "tx", addTokenTx.Hash().Hex())
			waitFor(addTokenTx.Hash(), client)
			sender.Nonce.Add(sender.Nonce, one)

		}

	}

	// generate alice, bob account.
	aliceInitBalance := new(big.Int).SetUint64(500)
	alice := CreateTestAccount("alice", aliceInitBalance)
	sender.Value = new(big.Int).Mul(ether, aliceInitBalance)
	sender.Value.Div(sender.Value, precision)
	if tokenContract != nil {
		sender.Value = nil
	}
	alicePK := [2]*big.Int{alice.sk.PublicKey.X, alice.sk.PublicKey.Y}
	aliceTx, err := cs.PGC.DepositAccount(sender, alicePK, token, aliceInitBalance)
	require.Nil(t, err, "deposit contract failed")
	waitFor(aliceTx.Hash(), client)

	// check for alice's encrypted balance.
	aliceEncryptB, _ := cs.PGC.GetUserBalance(nil, alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	aliceDecryptB := Decrypt(alice.sk, arrayToCT(aliceEncryptB, alice.sk.Curve))
	require.Equal(t, aliceInitBalance.Bytes(), aliceDecryptB, "alice'balance on chain not same with local")

	// keep balance same with chain.
	alice.UpdateBalance(aliceEncryptB)

	// generate bob.
	bobInitBalance := new(big.Int).SetUint64(500)
	bob := CreateTestAccount("bob", bobInitBalance)
	sender.Value = new(big.Int).Mul(ether, bobInitBalance)
	sender.Value.Div(sender.Value, precision)
	if tokenContract != nil {
		sender.Value = nil
	}
	sender.Nonce.Add(sender.Nonce, one)
	bobPK := [2]*big.Int{bob.sk.PublicKey.X, bob.sk.PublicKey.Y}
	bobTx, err := cs.PGC.DepositAccount(sender, bobPK, token, bobInitBalance)
	require.Nil(t, err, "deposit to bob on chain failed")
	waitFor(bobTx.Hash(), client)
	sender.Nonce.Add(sender.Nonce, one)

	// check for bob's encrypted balance.
	bobEncryptB, _ := cs.PGC.GetUserBalance(nil, bob.sk.PublicKey.X, bob.sk.PublicKey.Y, token)
	bobDecryptB := Decrypt(bob.sk, arrayToCT(bobEncryptB, bob.sk.Curve))
	require.Equal(t, bobInitBalance.Bytes(), bobDecryptB, "bob's balance on chain not same with local")
	// keep balance same with chain.
	bob.UpdateBalance(bobEncryptB)

	transferAmount := new(big.Int).SetUint64(300)
	ctx, err := CreateCTX(alice, &bob.sk.PublicKey, transferAmount)
	require.Nil(t, err, "create transfer ctx failed")

	// send transfer tx to contract on chain.
	tx := ctxToTransferTx(ctx)
	transferHash, err := HashTransfer(tx, token)
	require.Nil(t, err, "hash transfer data failed")

	transferSig, err := Sign(alice.sk, transferHash)
	require.Nil(t, err, "alice sign for transfer tx failed")

	sender.Value = nil
	transferTxHash, err := cs.PGC.Transfer(sender, tx.points, tx.scalar, tx.rpoints, tx.l, tx.r, new(big.Int).SetBytes(token.Bytes()), tx.nonce, transferSig.ToInputs())
	require.Nil(t, err, "create transfer tx failed")

	sender.Nonce.Add(sender.Nonce, one)
	waitFor(transferTxHash.Hash(), client)
	// cost 6444057 gas
	t.Log("transfer from alice to bob", transferAmount, transferTxHash.Hash().Hex())

	// check alice/bob balance.
	aliceEncBalanceAfter, _ := cs.PGC.GetUserBalance(nil, alicePK[0], alicePK[1], token)
	aliceBalanceAfter := Decrypt(alice.sk, arrayToCT(aliceEncBalanceAfter, alice.sk.Curve))
	aliceExceptBalance := new(big.Int).Sub(aliceInitBalance, transferAmount)
	require.Equal(t, aliceExceptBalance.Bytes(), aliceBalanceAfter, "alice balance after transfer is invalid")
	t.Log("alice'balance after is", new(big.Int).SetBytes(aliceBalanceAfter))

	// check for bob's balance.
	bobEncryptBalanceAfter, _ := cs.PGC.GetUserBalance(nil, bobPK[0], bobPK[1], token)
	bobBalanceAfter := Decrypt(bob.sk, arrayToCT(bobEncryptBalanceAfter, bob.sk.Curve))
	bobExceptBalance := new(big.Int).Add(bobInitBalance, transferAmount)
	require.Equal(t, bobExceptBalance.Bytes(), bobBalanceAfter, "bob balance after transfer is invalid")
	t.Log("bob's balance after is ", new(big.Int).SetBytes(bobBalanceAfter))
	bob.UpdateBalance(bobEncryptBalanceAfter)

	bob.m = new(big.Int).Set(bobExceptBalance)
	// create burn tx for bob
	burnAmount := new(big.Int).SetUint64(50)
	burnPartTx, err := CreateBurnPartTx(bob, burnAmount)
	require.Nil(t, err)

	receiver := sender.From
	tmpInfo := txToBurnPartTx(burnPartTx)
	burnHash, err := HashBurnPart(receiver, token, tmpInfo)
	require.Nil(t, err)
	burnSig, err := Sign(bob.sk, burnHash)
	require.Nil(t, err)

	burnTxH, err := cs.PGC.BurnPart(sender, receiver, new(big.Int).SetBytes(token.Bytes()), tmpInfo.amount, tmpInfo.points, tmpInfo.scalar, tmpInfo.rpoints, tmpInfo.l, tmpInfo.r, tmpInfo.nonce, burnSig.ToInputs())
	require.Nil(t, err, "bob create burn part tx failed")
	t.Log("burn part tx", burnTxH.Hash().Hex(), "burn amount", burnAmount)
	waitFor(burnTxH.Hash(), client)
	sender.Nonce.Add(sender.Nonce, one)

	// check bob's balance.
	bobFinalEncBalance, _ := cs.PGC.GetUserBalance(nil, bobPK[0], bobPK[1], token)
	bobf := Decrypt(bob.sk, arrayToCT(bobFinalEncBalance, bob.sk.Curve))
	t.Log("bob's final balance", "amount", new(big.Int).SetBytes(bobf))
	finalExpect := new(big.Int).Sub(new(big.Int).SetBytes(bobBalanceAfter), burnAmount)
	require.Equal(t, finalExpect.Bytes(), bobf, "bob'balance after burn is invalid")
}

func newBurnTx() *burnTx {
	tx := burnTx{}
	tx.pk = [2]*big.Int{}
	tx.proof = [4]*big.Int{}

	return &tx
}

func burnTxToOb(b *BurnTx) *burnTx {
	r := newBurnTx()
	r.nonce = b.Nonce
	r.pk[0] = b.Account.X
	r.pk[1] = b.Account.Y
	r.amount = b.Amount
	r.proof[0] = b.Proof.A1.X
	r.proof[1] = b.Proof.A1.Y
	r.proof[2] = b.Proof.A2.X
	r.proof[3] = b.Proof.A2.Y
	r.z = b.Proof.Z

	return r
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

func newBurnPartTx() *burnPartTx {
	tx := burnPartTx{}
	tx.points = [18]*big.Int{}
	tx.scalar = [12]*big.Int{}
	tx.rpoints = [16]*big.Int{}
	tx.l = [16]*big.Int{}
	tx.r = [16]*big.Int{}

	return &tx
}

func arrayToCT(p CT, curve elliptic.Curve) *CTEncPoint {
	c := CTEncPoint{}
	c.X = NewECPoint(p.Ct[0], p.Ct[1], curve)
	c.Y = NewECPoint(p.Ct[2], p.Ct[3], curve)

	return &c
}

func txToBurnPartTx(tx *BurnPartTx) *burnPartTx {
	partTx := newBurnPartTx()

	partTx.amount = new(big.Int).Set(tx.amount)
	partTx.nonce = new(big.Int).SetUint64(tx.nonce)
	partTx.points[0] = tx.pk.X
	partTx.points[1] = tx.pk.Y
	partTx.points[2] = tx.ct.X.X
	partTx.points[3] = tx.ct.X.Y
	partTx.points[4] = tx.ct.Y.X
	partTx.points[5] = tx.ct.Y.Y
	partTx.points[6] = tx.refreshBalance.X.X
	partTx.points[7] = tx.refreshBalance.X.Y
	partTx.points[8] = tx.refreshBalance.Y.X
	partTx.points[9] = tx.refreshBalance.Y.Y
	partTx.points[10] = tx.dleProof1.A1.X
	partTx.points[11] = tx.dleProof1.A1.Y
	partTx.points[12] = tx.dleProof1.A2.X
	partTx.points[13] = tx.dleProof1.A2.Y
	partTx.points[14] = tx.dleProof2.A1.X
	partTx.points[15] = tx.dleProof2.A1.Y
	partTx.points[16] = tx.dleProof2.A2.X
	partTx.points[17] = tx.dleProof2.A2.Y

	partTx.scalar[0] = tx.dleProof1.Z
	partTx.scalar[1] = tx.dleProof2.Z

	partTx.scalar[2] = tx.proof1.t
	partTx.scalar[3] = tx.proof1.tx
	partTx.scalar[4] = tx.proof1.u
	partTx.scalar[5] = tx.proof1.ipProof.a
	partTx.scalar[6] = tx.proof1.ipProof.b

	partTx.scalar[7] = tx.proof2.t
	partTx.scalar[8] = tx.proof2.tx
	partTx.scalar[9] = tx.proof2.u
	partTx.scalar[10] = tx.proof2.ipProof.a
	partTx.scalar[11] = tx.proof2.ipProof.b

	partTx.rpoints[0] = tx.proof1.A.X
	partTx.rpoints[1] = tx.proof1.A.Y
	partTx.rpoints[2] = tx.proof1.S.X
	partTx.rpoints[3] = tx.proof1.S.Y
	partTx.rpoints[4] = tx.proof1.T1.X
	partTx.rpoints[5] = tx.proof1.T1.Y
	partTx.rpoints[6] = tx.proof1.T2.X
	partTx.rpoints[7] = tx.proof1.T2.Y

	partTx.rpoints[8] = tx.proof2.A.X
	partTx.rpoints[9] = tx.proof2.A.Y
	partTx.rpoints[10] = tx.proof2.S.X
	partTx.rpoints[11] = tx.proof2.S.Y
	partTx.rpoints[12] = tx.proof2.T1.X
	partTx.rpoints[13] = tx.proof2.T1.Y
	partTx.rpoints[14] = tx.proof2.T2.X
	partTx.rpoints[15] = tx.proof2.T2.Y

	base := len(tx.proof1.ipProof.L) * 2
	for i, v := range tx.proof1.ipProof.L {
		partTx.l[2*i] = v.X
		partTx.l[2*i+1] = v.Y
		partTx.r[2*i] = tx.proof1.ipProof.R[i].X
		partTx.r[2*i+1] = tx.proof1.ipProof.R[i].Y

		partTx.l[2*i+base] = tx.proof2.ipProof.L[i].X
		partTx.l[2*i+1+base] = tx.proof2.ipProof.L[i].Y
		partTx.r[2*i+base] = tx.proof2.ipProof.R[i].X
		partTx.r[2*i+1+base] = tx.proof2.ipProof.R[i].Y
	}

	return partTx
}

func ctxToTransferTx(ctx *CTX) *transferTx {
	tx := newTransferTx()
	tx.nonce = new(big.Int).SetUint64(ctx.nonce)
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
