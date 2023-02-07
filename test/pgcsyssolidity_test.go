package test

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
	"github.com/joho/godotenv"
	"github.com/pgc/client"
	pgcm "github.com/pgc/contracts/pgc"
	tokenpkg "github.com/pgc/contracts/token"
	"github.com/pgc/contracts/tokenconverter"
	"github.com/pgc/deployer"
	"github.com/pgc/pgcsys"
	"github.com/pgc/proof"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/require"
)

var (
	precision = new(big.Int).SetUint64(100)
)

func TestPGCSystemTokenOnScroll(t *testing.T) {
	ethclient := client.GetScrollClientL2()
	auth := client.GetAccountWithKey(os.Getenv("DeployerKey"), ethclient)

	testPGCSystemContract(t, true, auth, ethclient, pgcsys.DefaultPgcSysParams())
}

func TestPGCSystemContractETHLocal(t *testing.T) {
	ethclient := client.GetLocal()
	auth := client.GetAccountWithKey(os.Getenv("DeployerKey"), ethclient)

	testPGCSystemContract(t, false, auth, ethclient, pgcsys.DefaultPgcSysParams())
}

func TestPGCSystemContractTokenLocal(t *testing.T) {
	ethclient := client.GetLocal()
	auth := client.GetAccountWithKey(os.Getenv("DeployerKey"), ethclient)

	testPGCSystemContract(t, true, auth, ethclient, pgcsys.DefaultPgcSysParams())
}

func testPGCSystemContract(t *testing.T, tokenTest bool, auth *bind.TransactOpts, ethclient *ethclient.Client, params pgcsys.PgcSys) {
	addrs, pgc := deployer.DeployPGCSystemAllContract(auth, ethclient, params.Pub().X, params.Pub().Y)
	proof.BuildAndLoadMapIfNotExist(params.G(), 32, 7, 4)
	deployer.InitVector(auth, ethclient, addrs[2], 32)

	token := common.Address{}
	if tokenTest {
		token = setForToken(t, addrs, auth, ethclient)
	}

	fmt.Scanln()
	authAlice := client.GetAccountWithKey(os.Getenv("AliceKey"), ethclient)
	log.Info("", "alice", authAlice.From.Hex())
	client.SetNonce(authAlice, ethclient)
	authBob := client.GetAccountWithKey(os.Getenv("BobKey"), ethclient)
	client.SetNonce(authBob, ethclient)

	// mint token for alice and bob
	if tokenTest {
		tokenContract, _ := tokenpkg.NewToken(token, ethclient)
		_, err := tokenContract.Transfer(auth, authAlice.From, new(big.Int).Mul(utils.Ether, big.NewInt(60)))
		require.Nil(t, err)
		auth.Nonce.Add(auth.Nonce, utils.One)
		tx, err := tokenContract.Transfer(auth, authBob.From, new(big.Int).Mul(utils.Ether, big.NewInt(30)))
		require.Nil(t, err)
		auth.Nonce.Add(auth.Nonce, utils.One)
		// approve
		tokenContract.Approve(authAlice, addrs[len(addrs)-1], new(big.Int).Mul(utils.Ether, big.NewInt(1000)))
		tx, err = tokenContract.Approve(authBob, addrs[len(addrs)-1], new(big.Int).Mul(utils.Ether, big.NewInt(1000)))
		require.Nil(t, err)
		authAlice.Nonce.Add(authAlice.Nonce, utils.One)
		authBob.Nonce.Add(authBob.Nonce, utils.One)
		client.WaitForTx(ethclient, tx.Hash())
	}

	// alice
	aliceAmount := new(big.Int).SetUint64(512)
	name := "Alice"
	alice := pgcsys.CreateTestAccount(params, name, aliceAmount)
	fmt.Println("")
	fmt.Println("New pgc account created")
	fmt.Println("Account name:", name)
	fmt.Printf("Account pubkey: %x\n", elliptic.Marshal(alice.Priv().Curve, alice.Priv().X, alice.Priv().Y))

	// bob.
	bobAmount := new(big.Int).SetUint64(256)
	name = "Bob"
	bob := pgcsys.CreateTestAccount(params, name, bobAmount)
	fmt.Println("")
	fmt.Println("New pgc account created")
	fmt.Println("Account name:", name)
	fmt.Printf("Account pubkey: %x\n", elliptic.Marshal(bob.Priv().Curve, bob.Priv().X, bob.Priv().Y))
	fmt.Println("")

	fmt.Scanln()
	initTestAccount(t, alice, params, token, aliceAmount, name, authAlice, ethclient, pgc)
	initTestAccount(t, bob, params, token, bobAmount, name, authBob, ethclient, pgc)

	fmt.Println("")
	fmt.Println("Deposit succeeds")
	fmt.Println("Alice's current balance", alice.M())
	fmt.Println("Bob's current balance", bob.M())
	fmt.Scanln()

	transferAmount := new(big.Int).SetUint64(128)
	auth.GasLimit = 8000000
	ctx := aggTransfer(t, params, alice, bob, token, transferAmount, auth, ethclient, pgc)

	fmt.Println("")
	fmt.Println("CTx transfer succeeds")
	fmt.Println("Alice's current balance", alice.M())
	fmt.Println("Bob's current balance", bob.M())
	fmt.Scanln()

	// 审计
	fmt.Println("")
	fmt.Println("Audit Begins")
	fmt.Println("Alice says that the amount in this CTx is 128, And she generates a proof for this CTx")
	encryptedTransfer := ctx.Transfer().First().Copy()
	xx := elliptic.Marshal(bob.Priv().Curve, encryptedTransfer.X.X, encryptedTransfer.X.Y)
	yy := elliptic.Marshal(bob.Priv().Curve, encryptedTransfer.Y.X, encryptedTransfer.Y.Y)
	fmt.Printf("The encrypted amount in the transfer is: X: %x, Y: %x\n", xx, yy)
	proofct, err := proof.GenerateEqualProof(params, transferAmount, ctx.Transfer().First().Copy(), alice.Priv())
	require.Nil(t, err)
	A1 := elliptic.Marshal(bob.Priv().Curve, proofct.A1.X, proofct.A1.Y)
	A2 := elliptic.Marshal(bob.Priv().Curve, proofct.A2.X, proofct.A2.Y)
	z := proofct.Z.Bytes()
	fmt.Printf("The proof is(anyone can verify this proof with alice's pk): A1: %x, A2: %x, z: %x\n", A1, A2, z)
	r := proof.VerifyEqualProof(params, encryptedTransfer, transferAmount, new(utils.ECPoint).SetFromPublicKey(&alice.Priv().PublicKey), proofct)
	fmt.Println("Verify result: ", r)

	fmt.Scanln()
	fmt.Println("")
	receiver := authAlice.From
	burn(t, params, alice, receiver, token, auth, ethclient, pgc)
	receiver = authBob.From
	burn(t, params, bob, receiver, token, auth, ethclient, pgc)
}

func setForToken(t *testing.T, addrs []common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client) common.Address {
	token, tokenCon := deployer.DeployToken(auth, ethclient)

	pgcAddr := addrs[len(addrs)-1]
	approveAmount := new(big.Int).SetUint64(100000)

	_, err := tokenCon.Approve(auth, pgcAddr, approveAmount)
	require.Nil(t, err, "approve for token failed", err)
	auth.Nonce.Add(auth.Nonce, utils.One)

	tokenConverter, err := tokenconverter.NewTokenconverter(addrs[1], ethclient)
	require.Nil(t, err, "get token converter failed", err)

	_, err = tokenConverter.AddToken(auth, token, utils.One, "")
	require.Nil(t, err, "add token failed", err)
	auth.Nonce.Add(auth.Nonce, utils.One)

	return token
}

func initTestAccount(t *testing.T, alice *pgcsys.Account, params proof.AggRangeParams, token common.Address, amount *big.Int, name string, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *pgcm.Pgc) *pgcsys.Account {
	alicePK := new(utils.ECPoint).SetFromPublicKey(&alice.Priv().PublicKey)

	var aliceTx *types.Transaction
	var err error
	if isToken(token) {
		auth.Value = new(big.Int).Mul(utils.Ether, amount)
		auth.Value.Div(auth.Value, precision)
		aliceTx, err = pgc.DepositAccountETH(auth, alicePK.Compress())
	} else {
		aliceTx, err = pgc.DepositAccount(auth, alicePK.Compress(), token, amount)
	}

	require.Nil(t, err, "deposit contract failed")
	receipt := client.WaitForTx(ethclient, aliceTx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check for alice's encrypted balance.
	aliceEncryptB, _ := pgc.GetUserBalance(utils.CallOpt(), alice.Priv().PublicKey.X, alice.Priv().PublicKey.Y, token)
	alice.UpdateBalance(aliceEncryptB.Nonce, aliceEncryptB.Ct)
	require.Equal(t, amount.Bytes(), alice.M().Bytes(), "account'balance on chain not same with local", "expect", amount, "actual", alice.M())

	log.Info("deposit account success", "token", token.Hash().Hex(), "who", name, "amount", amount, "gas", receipt.GasUsed, "tx", aliceTx.Hash().Hex())
	auth.Value = nil

	return alice
}

func aggTransfer(t *testing.T, params pgcsys.PgcSys, from, to *pgcsys.Account, token common.Address, amount *big.Int, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *pgcm.Pgc) *pgcsys.ConfidentialTx {
	ctx, err := pgcsys.CreateConfidentialTx(params, from, &to.Priv().PublicKey, amount, token.Hash().Big())
	require.Nil(t, err, "generate confidential tx failed", err)
	input := ctx.ToSolidityInput()

	var tx *types.Transaction
	if isToken(token) {
		tx, err = pgc.AggTransferETH(auth, input.Points, input.Scalars, input.Lr)
	} else {
		tx, err = pgc.AggTransfer(auth, input.Points, input.Scalars, token.Hash().Big(), input.Lr)
	}

	require.Nil(t, err, "create agg transfer tx failed", err)
	receipt := client.WaitForTx(ethclient, tx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check sender's balance.
	before := new(big.Int).Set(from.M())
	shouldBe := before.Sub(before, amount)
	encryptB, _ := pgc.GetUserBalance(utils.CallOpt(), from.Priv().PublicKey.X, from.Priv().PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, shouldBe.Bytes(), from.M().Bytes(), "sender's balance on chain invalid", "expect", shouldBe, "actual", from.M())

	// check receiver's balance.
	before = new(big.Int).Set(to.M())
	shouldBe = before.Add(before, amount)
	encryptB, _ = pgc.GetUserBalance(utils.CallOpt(), to.Priv().PublicKey.X, to.Priv().PublicKey.Y, token)
	to.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, shouldBe.Bytes(), to.M().Bytes(), "receiver's balance on chain invalid", "expect", shouldBe, "acutal", to.M())

	log.Info("agg transfer", "token", token.Hash().Hex(), "amount", amount, "gas", receipt.GasUsed, "tx", tx.Hash().Hex())

	return ctx
}

func burn(t *testing.T, params pgcsys.PgcSys, from *pgcsys.Account, receiver, token common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *pgcm.Pgc) {
	tx, err := pgcsys.CreateBurnETHTx(params, from, receiver, token)
	require.Nil(t, err, "generate burn eth tx failed")

	input := tx.ToSolidityInput()

	var btx *types.Transaction
	if isToken(token) {
		btx, err = pgc.BurnETH(auth, receiver, input.Amount, input.Points, input.Z)
	} else {
		btx, err = pgc.Burn(auth, receiver, token.Hash().Big(), input.Amount, input.Points, input.Z)
	}

	require.Nil(t, err, "create burn tx failed")
	receipt := client.WaitForTx(ethclient, btx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check balance.
	encryptB, _ := pgc.GetUserBalance(utils.CallOpt(), from.Priv().PublicKey.X, from.Priv().PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, uint64(0), from.M().Uint64(), "receiver's balance on chain invalid", "expect", 0, "actual", from.M())

	log.Info("burn tx", "token", token.Hash().Hex(), "amount", input.Amount, "gas", receipt.GasUsed, "tx", btx.Hash().Hex())
}

func isToken(token common.Address) bool {
	return token.Hash().Big().Cmp(utils.Zero) == 0
}

func init() {
	godotenv.Load("../.env")
}
