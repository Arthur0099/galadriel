package pgcsys

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
	"github.com/pgc/client"
	pgcm "github.com/pgc/contracts/pgc"
	tokenpkg "github.com/pgc/contracts/token"
	"github.com/pgc/contracts/tokenconverter"
	"github.com/pgc/contracts/verifier"
	"github.com/pgc/deployer"
	"github.com/pgc/proof"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	precision = new(big.Int).SetUint64(100)
)

func skipRopsten(t *testing.T) {
	if os.Getenv("ropsten") != "true" {
		t.Skip("skip tesing for ropsten")
	}
}

func skipMainnet(t *testing.T) {
	if os.Getenv("mainnet") != "true" {
		t.Skip("skip deploy for mainnet")
	}
}

func TestDeployPGCToMainnet(t *testing.T) {
	skipMainnet(t)

	mainnetURL := os.Getenv("mainnetURL")
	deployerKey := os.Getenv("deployerKey")

	ethclient := client.GetClient(mainnetURL)
	auth := client.GetAccountWithKey(deployerKey, ethclient)
	auth.GasLimit = 7500000
	addrs, _ := deployer.DeployPGCSystemAllContract(auth, ethclient)

	// set for tokens.
	token := setForToken(t, addrs, auth, ethclient)
	log.Info("token contract", "addr", token.Hex())
}

func TestPGCSystemTokenOnScroll(t *testing.T) {
	ethclient := client.GetScrollClientL2()
	// replace with account.
	auth := client.GetAccountWithKey("74175140cb1ed8129dfeba9df037ba68981d91608109da4c602e7fb30ee36601", ethclient)

	testPGCSystemContract(t, true, auth, ethclient)
}

func TestPGCSystemContractETHRopsten(t *testing.T) {
	skipRopsten(t)

	ethclient := client.GetRopstenInfura()
	// replace with account.
	auth := client.GetRopstenAccount(ethclient)

	testPGCSystemContract(t, false, auth, ethclient)
}

func TestPGCSystemContractETHLocal(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH(ethclient)

	testPGCSystemContract(t, false, auth, ethclient)
}

func TestPGCSystemContractTokenLocal(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH(ethclient)

	testPGCSystemContract(t, true, auth, ethclient)
}

func testPGCSystemContract(t *testing.T, tokenTest bool, auth *bind.TransactOpts, ethclient *ethclient.Client) {
	addrs, pgc := deployer.DeployPGCSystemAllContract(auth, ethclient)
	params := proof.DAggRangeProofParams32()
	proof.LoadMap(32, 7)
	deployer.InitVector(auth, ethclient, addrs[2], 32)

	token := common.Address{}
	if tokenTest {
		token = setForToken(t, addrs, auth, ethclient)
	}

	fmt.Scanln()
	// snap joke lazy talent prepare expect fence draw label delay balance erode
	authAlice := client.GetAccountWithKey("69f1a5e97304fd38cf0840182a988b51c3e566aae1862264443c985a1a5af682", ethclient)
	log.Info("", "alice", authAlice.From.Hex())
	client.SetNonce(authAlice, ethclient)
	authBob := client.GetAccountWithKey("681a477af4f60c5068f4d279e697d6741205e353dc4f24bb0dbdf79c8be40107", ethclient)
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
	alice := CreateTestAccount(params, name, aliceAmount)
	fmt.Println("")
	fmt.Println("New pgc account created")
	fmt.Println("Account name:", name)
	fmt.Printf("Account pubkey: %x\n", elliptic.Marshal(alice.sk.Curve, alice.sk.X, alice.sk.Y))

	// bob.
	bobAmount := new(big.Int).SetUint64(256)
	name = "Bob"
	bob := CreateTestAccount(params, name, bobAmount)
	fmt.Println("")
	fmt.Println("New pgc account created")
	fmt.Println("Account name:", name)
	fmt.Printf("Account pubkey: %x\n", elliptic.Marshal(bob.sk.Curve, bob.sk.X, bob.sk.Y))
	fmt.Println("")

	fmt.Scanln()
	initTestAccount(t, alice, params, token, aliceAmount, name, authAlice, ethclient, pgc)
	initTestAccount(t, bob, params, token, bobAmount, name, authBob, ethclient, pgc)

	fmt.Println("")
	fmt.Println("Deposit succeeds")
	fmt.Println("Alice's current balance", alice.m)
	fmt.Println("Bob's current balance", bob.m)
	fmt.Scanln()

	transferAmount := new(big.Int).SetUint64(128)
	auth.GasLimit = 8000000
	ctx := aggTransfer(t, params, alice, bob, token, transferAmount, auth, ethclient, pgc)

	fmt.Println("")
	fmt.Println("CTx transfer succeeds")
	fmt.Println("Alice's current balance", alice.m)
	fmt.Println("Bob's current balance", bob.m)
	fmt.Scanln()

	// 审计
	fmt.Println("")
	fmt.Println("Audit Begins")
	fmt.Println("Alice says that the amount in this CTx is 128, And she generates a proof for this CTx")
	encryptedTransfer := ctx.transfer.First().Copy()
	xx := elliptic.Marshal(bob.sk.Curve, encryptedTransfer.X.X, encryptedTransfer.X.Y)
	yy := elliptic.Marshal(bob.sk.Curve, encryptedTransfer.Y.X, encryptedTransfer.Y.Y)
	fmt.Printf("The encrypted amount in the transfer is: X: %x, Y: %x\n", xx, yy)
	proofct, err := proof.GenerateEqualProof(params, transferAmount, ctx.transfer.First().Copy(), alice.sk)
	require.Nil(t, err)
	A1 := elliptic.Marshal(bob.sk.Curve, proofct.A1.X, proofct.A1.Y)
	A2 := elliptic.Marshal(bob.sk.Curve, proofct.A2.X, proofct.A2.Y)
	z := proofct.Z.Bytes()
	fmt.Printf("The proof is(anyone can verify this proof with alice's pk): A1: %x, A2: %x, z: %x\n", A1, A2, z)
	r := proof.VerifyEqualProof(params, encryptedTransfer, transferAmount, new(utils.ECPoint).SetFromPublicKey(&alice.sk.PublicKey), proofct)
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

func initTestAccount(t *testing.T, alice *Account, params proof.AggRangeParams, token common.Address, amount *big.Int, name string, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *pgcm.Pgc) *Account {
	alicePK := new(utils.ECPoint).SetFromPublicKey(&alice.sk.PublicKey)

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
	aliceEncryptB, _ := pgc.GetUserBalance(utils.CallOpt(), alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	alice.UpdateBalance(aliceEncryptB.Nonce, aliceEncryptB.Ct)
	require.Equal(t, amount.Bytes(), alice.m.Bytes(), "account'balance on chain not same with local", "expect", amount, "actual", alice.m)

	log.Info("deposit account success", "token", token.Hash().Hex(), "who", name, "amount", amount, "gas", receipt.GasUsed, "tx", aliceTx.Hash().Hex())
	auth.Value = nil

	return alice
}

func aggTransfer(t *testing.T, params proof.AggRangeParams, from, to *Account, token common.Address, amount *big.Int, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *pgcm.Pgc) *ConfidentialTx {
	ctx, err := CreateConfidentialTx(params, from, &to.sk.PublicKey, amount, token.Hash().Big())
	require.Nil(t, err, "generate condidential tx failed", err)
	input := ctx.ToSolidityInput()

	var tx *types.Transaction
	if isToken(token) {
		tx, err = pgc.AggTransferETH(auth, input.points, input.scalars, input.lr)
	} else {
		tx, err = pgc.AggTransfer(auth, input.points, input.scalars, token.Hash().Big(), input.lr)
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

	log.Info("agg transfer", "token", token.Hash().Hex(), "amount", amount, "gas", receipt.GasUsed, "tx", tx.Hash().Hex())

	return ctx
}

func burn(t *testing.T, params proof.AggRangeParams, from *Account, receiver, token common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client, pgc *pgcm.Pgc) {
	tx, err := CreateBurnETHTx(params, from, receiver, token)
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
	encryptB, _ := pgc.GetUserBalance(utils.CallOpt(), from.sk.PublicKey.X, from.sk.PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, uint64(0), from.m.Uint64(), "receiver's balance on chain invalid", "expect", 0, "acutal", from.m)

	log.Info("burn tx", "token", token.Hash().Hex(), "amount", input.Amount, "gas", receipt.GasUsed, "tx", btx.Hash().Hex())
}

func isToken(token common.Address) bool {
	return token.Hash().Big().Cmp(utils.Zero) == 0
}

func TestAggTransferByV2Local(t *testing.T) {
	amount := big.NewInt(25)
	params := proof.DAggRangeProofParamsWithBitsize(64)
	from := CreateTestAccount(params, "alice", big.NewInt(100))
	to := CreateTestAccount(params, "bob", big.NewInt(100))

	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH(ethclient)

	paramsAddr, _ := deployer.DeployParams(auth, ethclient)
	addr, _, con, _ := verifier.DeployVerifier(auth, ethclient, paramsAddr)
	deployer.InitVector(auth, ethclient, addr, 64)

	testAggTransferByV2(t, true, amount, params, from, to, con)
}

func testAggTransferByV2(t *testing.T, expect bool, amount *big.Int, params proof.AggRangeParams, from, to *Account, con *verifier.Verifier) {
	ctx, err := CreateConfidentialTx(params, from, &to.sk.PublicKey, amount, big.NewInt(0))
	require.Nil(t, err, "generate condidential tx failed", err)

	proof, state := ToSolidityV2(ctx)

	actual, err := con.VerifyAggTransfer(utils.CallOpt(), proof, state)
	require.Nil(t, err, "call agg range verifier failed", err)
	assert.Equal(t, expect, actual)
}

// ToSolidityV2 format data to solidity input v2.
func ToSolidityV2(ctx *ConfidentialTx) (verifier.VerifierTransferProof, verifier.VerifierTransferStatement) {
	proof := verifier.VerifierTransferProof{}
	state := verifier.VerifierTransferStatement{}

	// set proof.
	proof.PteProof = verifier.VerifierPTEProof{}
	proof.PteProof.A1 = toSolidityPoint(ctx.sigmaPTEqualityProof.A1)
	proof.PteProof.A2 = toSolidityPoint(ctx.sigmaPTEqualityProof.A2)
	proof.PteProof.B = toSolidityPoint(ctx.sigmaPTEqualityProof.B)
	proof.PteProof.Z1 = ctx.sigmaPTEqualityProof.Z1
	proof.PteProof.Z2 = ctx.sigmaPTEqualityProof.Z2

	proof.DleProof = verifier.VerifierDLEProof{}
	proof.DleProof.A1 = toSolidityPoint(ctx.sigmaDlogeqProof.A1)
	proof.DleProof.A2 = toSolidityPoint(ctx.sigmaDlogeqProof.A2)
	proof.DleProof.Z = ctx.sigmaDlogeqProof.Z

	proof.CtvProof = verifier.VerifierCTVProof{}
	proof.CtvProof.A = toSolidityPoint(ctx.sigmaCTValidProof.A)
	proof.CtvProof.B = toSolidityPoint(ctx.sigmaCTValidProof.B)
	proof.CtvProof.Z1 = ctx.sigmaCTValidProof.Z1
	proof.CtvProof.Z2 = ctx.sigmaCTValidProof.Z2

	proof.AggProof = verifier.VerifierAggProof{}
	proof.AggProof.A = toSolidityPoint(ctx.bulletProof.A)
	proof.AggProof.S = toSolidityPoint(ctx.bulletProof.S)
	proof.AggProof.T1 = toSolidityPoint(ctx.bulletProof.T1)
	proof.AggProof.T2 = toSolidityPoint(ctx.bulletProof.T2)
	proof.AggProof.T = ctx.bulletProof.T()
	proof.AggProof.Txx = ctx.bulletProof.TX()
	proof.AggProof.U = ctx.bulletProof.U()
	proof.AggProof.Ap = ctx.bulletProof.AIP()
	proof.AggProof.Bp = ctx.bulletProof.BIP()
	proof.AggProof.L = toSolidityPoints(ctx.bulletProof.L())
	proof.AggProof.R = toSolidityPoints(ctx.bulletProof.R())

	// set state.
	state.Pk1 = toSolidityPoint(ctx.pk1)
	state.Pk2 = toSolidityPoint(ctx.pk2)
	state.Refresh = verifier.VerifierCT{
		X: toSolidityPoint(ctx.refreshBalance.X),
		Y: toSolidityPoint(ctx.refreshBalance.Y),
	}
	state.Updated = verifier.VerifierCT{
		X: toSolidityPoint(ctx.updatedBalance.X),
		Y: toSolidityPoint(ctx.updatedBalance.Y),
	}
	state.Mrct = verifier.VerifierMRCT{
		X1: toSolidityPoint(ctx.transfer.X1),
		X2: toSolidityPoint(ctx.transfer.X2),
		Y:  toSolidityPoint(ctx.transfer.Y),
	}
	state.Custom = ctx.Custom()

	return proof, state
}

func toSolidityPoints(ps []*utils.ECPoint) []verifier.BN128G1Point {
	nps := make([]verifier.BN128G1Point, 0)
	for _, p := range ps {
		nps = append(nps, toSolidityPoint(p))
	}

	return nps
}

func toSolidityPoint(p *utils.ECPoint) verifier.BN128G1Point {
	np := verifier.BN128G1Point{}
	np.X = new(big.Int).Set(p.X)
	np.Y = new(big.Int).Set(p.Y)

	return np
}
