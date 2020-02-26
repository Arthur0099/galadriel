package pgcsys

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
	"github.com/pgc/client"
	"github.com/pgc/contracts"
	"github.com/pgc/deployer"
	"github.com/pgc/proof"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/require"
)

var (
	precision = new(big.Int).SetUint64(100)
)

func TestPGCSystemContractETHLocal(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH()

	testPGCSystemContract(t, false, auth, ethclient)
}

func TestPGCSystemContractTokenLocal(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH()

	testPGCSystemContract(t, true, auth, ethclient)
}

func testPGCSystemContract(t *testing.T, tokenTest bool, auth *bind.TransactOpts, ethclient *ethclient.Client) {
	addrs, pgc := deployer.DeployPGCSystemAllContract(auth, ethclient)
	params := proof.DAggRangeProofParamsWithBitsize(64)

	token := common.Address{}
	if tokenTest {
		token = setForToken(t, addrs, auth, ethclient)
	}

	// alice
	aliceAmount := new(big.Int).SetUint64(500)
	name := "alice"
	alice := initTestAccount(t, params, token, aliceAmount, name, auth, ethclient, pgc)

	// bob.
	bobAmount := new(big.Int).SetUint64(500)
	name = "bob"
	bob := initTestAccount(t, params, token, bobAmount, name, auth, ethclient, pgc)

	transferAmount := new(big.Int).SetUint64(100)
	auth.GasLimit = 100000000
	aggTransfer(t, params, alice, bob, token, transferAmount, auth, ethclient, pgc)
	transferAmount = new(big.Int).SetUint64(50)
	aggTransfer(t, params, alice, bob, token, transferAmount, auth, ethclient, pgc)
	aggTransfer(t, params, alice, bob, token, transferAmount, auth, ethclient, pgc)

	receiver := auth.From
	burn(t, params, alice, receiver, token, auth, ethclient, pgc)
	burn(t, params, bob, receiver, token, auth, ethclient, pgc)
}

func setForToken(t *testing.T, addrs []common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client) common.Address {
	token, tokenCon := deployer.DeployToken(auth, ethclient)

	pgcAddr := addrs[len(addrs)-1]
	approveAmount := new(big.Int).SetUint64(100000)

	_, err := tokenCon.Approve(auth, pgcAddr, approveAmount)
	require.Nil(t, err, "approve for token failed", err)
	auth.Nonce.Add(auth.Nonce, utils.One)

	tokenConverter, err := contracts.NewTokenconverter(addrs[5], ethclient)
	require.Nil(t, err, "get token converter failed", err)

	_, err = tokenConverter.AddToken(auth, token, utils.One, "")
	require.Nil(t, err, "add token failed", err)
	auth.Nonce.Add(auth.Nonce, utils.One)

	return token
}

func initTestAccount(t *testing.T, params proof.AggRangeParams, token common.Address, amount *big.Int, name string, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *contracts.Pgc) *Account {
	alice := CreateTestAccount(params, name, amount)
	alicePK := [2]*big.Int{alice.sk.PublicKey.X, alice.sk.PublicKey.Y}

	var aliceTx *types.Transaction
	var err error
	if isToken(token) {
		auth.Value = new(big.Int).Mul(utils.Ether, amount)
		auth.Value.Div(auth.Value, precision)
		aliceTx, err = pgc.DepositAccountETH(auth, alicePK)

	} else {
		aliceTx, err = pgc.DepositAccount(auth, alicePK, token, amount)
	}

	require.Nil(t, err, "deposit contract failed")
	receipt := client.WaitForTx(ethclient, aliceTx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check for alice's encrypted balance.
	aliceEncryptB, _ := pgc.GetUserBalance(utils.CallOpt(), alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	alice.UpdateBalance(aliceEncryptB.Nonce, aliceEncryptB.Ct)
	require.Equal(t, amount.Bytes(), alice.m.Bytes(), "account'balance on chain not same with local", "expect", amount, "actual", alice.m)

	log.Info("deposit account success", "token", token.Hash().Hex(), "who", name, "amount", amount, "gas", receipt.GasUsed)
	auth.Value = nil

	return alice
}

func aggTransfer(t *testing.T, params proof.AggRangeParams, from, to *Account, token common.Address, amount *big.Int, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *contracts.Pgc) {
	ctx, err := CreateConfidentialTx(params, from, &to.sk.PublicKey, amount)
	require.Nil(t, err, "generate condidential tx failed", err)
	input := ctx.ToSolidityInput()

	var tx *types.Transaction
	if isToken(token) {
		tx, err = pgc.AggTransferETH(auth, input.points, input.scalars, input.l, input.r)
	} else {
		tx, err = pgc.AggTransfer(auth, input.points, input.scalars, token.Hash().Big(), input.l, input.r)
	}

	require.Nil(t, err, "create agg transfer tx failed", err)
	receipt := client.WaitForTx(ethclient, tx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check sender's balance.
	before := new(big.Int).Set(from.m)
	shouldBe := before.Sub(before, amount)
	encryptB, _ := pgc.GetUserBalance(utils.CallOpt(), from.sk.PublicKey.X, from.sk.PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, shouldBe.Bytes(), from.m.Bytes(), "sender's balance on chain invalid", "expect", shouldBe, "acutal", from.m)

	// check receiver's balance.
	before = new(big.Int).Set(to.m)
	shouldBe = before.Add(before, amount)
	encryptB, _ = pgc.GetUserBalance(utils.CallOpt(), to.sk.PublicKey.X, to.sk.PublicKey.Y, token)
	to.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, shouldBe.Bytes(), to.m.Bytes(), "receiver's balance on chain invalid", "expect", shouldBe, "acutal", to.m)

	log.Info("agg transfer", "token", token.Hash().Hex(), "amount", amount, "gas", receipt.GasUsed)
}

func burn(t *testing.T, params proof.AggRangeParams, from *Account, receiver, token common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *contracts.Pgc) {
	tx, err := CreateBurnETHTx(params, from, receiver, token)
	require.Nil(t, err, "generate burn eth tx failed")

	input := tx.ToSolidityInput()

	var btx *types.Transaction
	if isToken(token) {
		btx, err = pgc.BurnETH(auth, receiver, input.Amount, input.PubKey, input.Proof, input.Z)
	} else {
		btx, err = pgc.Burn(auth, receiver, token.Hash().Big(), input.Amount, input.PubKey, input.Proof, input.Z)
	}

	require.Nil(t, err, "create burn tx failed")
	receipt := client.WaitForTx(ethclient, btx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check balance.
	encryptB, _ := pgc.GetUserBalance(utils.CallOpt(), from.sk.PublicKey.X, from.sk.PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, uint64(0), from.m.Uint64(), "receiver's balance on chain invalid", "expect", 0, "acutal", from.m)

	log.Info("burn tx", "token", token.Hash().Hex(), "amount", input.Amount, "gas", receipt.GasUsed)
}

func isToken(token common.Address) bool {
	return token.Hash().Big().Cmp(utils.Zero) == 0
}
