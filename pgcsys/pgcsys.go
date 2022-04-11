package pgcsys

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/inconshreveable/log15"
	"github.com/pgc/proof"
	"github.com/pgc/utils"
)

// ConfidentialTx is a tx for pgc transfer system(using aggreate bulletproof).
type ConfidentialTx struct {
	nonce, token *big.Int
	// in real system, this will be on chain and not inculded in ctx.
	balance  *proof.CTEncPoint
	pk1, pk2 *utils.ECPoint
	transfer *proof.MRTwistedELGamalCTPub

	refreshBalance *proof.CTEncPoint
	// only used for test, should be calculated in solidity.
	updatedBalance *proof.CTEncPoint

	// proof
	sigmaPTEqualityProof *proof.PTEqualityProof
	bulletProof          *proof.AggRangeProof
	sigmaCTValidProof    *proof.CTValidProof
	sigmaDlogeqProof     *proof.DLESigmaProof
}

type solidityPGCInput struct {
	// 36
	points []byte

	lr []byte

	scalars [10]*big.Int
}

// ToSolidityInput formats tx to solidity to verify contract
func (tx *ConfidentialTx) ToSolidityInput() *solidityPGCInput {
	input := solidityPGCInput{}
	input.points = make([]byte, 0)
	input.points = append(input.points, tx.pk1.Compress()...)
	input.points = append(input.points, tx.pk2.Compress()...)
	input.points = append(input.points, tx.transfer.X1.Compress()...)
	input.points = append(input.points, tx.transfer.X2.Compress()...)
	input.points = append(input.points, tx.transfer.Y.Compress()...)
	input.points = append(input.points, tx.sigmaPTEqualityProof.A1.Compress()...)
	input.points = append(input.points, tx.sigmaPTEqualityProof.A2.Compress()...)
	input.points = append(input.points, tx.sigmaPTEqualityProof.B.Compress()...)
	input.points = append(input.points, tx.refreshBalance.X.Compress()...)
	input.points = append(input.points, tx.refreshBalance.Y.Compress()...)
	input.points = append(input.points, tx.sigmaCTValidProof.A.Compress()...)
	input.points = append(input.points, tx.sigmaCTValidProof.B.Compress()...)
	input.points = append(input.points, tx.sigmaDlogeqProof.A1.Compress()...)
	input.points = append(input.points, tx.sigmaDlogeqProof.A2.Compress()...)
	// range proof.
	input.points = append(input.points, tx.bulletProof.A.Compress()...)
	input.points = append(input.points, tx.bulletProof.S.Compress()...)
	input.points = append(input.points, tx.bulletProof.T1.Compress()...)
	input.points = append(input.points, tx.bulletProof.T2.Compress()...)

	// L, R
	input.lr = make([]byte, 0)
	for i := 0; i < tx.bulletProof.Len(); i++ {
		input.lr = append(input.lr, tx.bulletProof.Li(i).Compress()...)
	}
	for i := 0; i < tx.bulletProof.Len(); i++ {
		input.lr = append(input.lr, tx.bulletProof.Ri(i).Compress()...)
	}

	// scalar
	input.scalars[0] = tx.sigmaPTEqualityProof.Z1
	input.scalars[1] = tx.sigmaPTEqualityProof.Z2
	input.scalars[2] = tx.sigmaCTValidProof.Z1
	input.scalars[3] = tx.sigmaCTValidProof.Z2
	input.scalars[4] = tx.sigmaDlogeqProof.Z
	// range proof.
	input.scalars[5] = tx.bulletProof.T()
	input.scalars[6] = tx.bulletProof.TX()
	input.scalars[7] = tx.bulletProof.U()
	// inner proof.
	input.scalars[8] = tx.bulletProof.AIP()
	input.scalars[9] = tx.bulletProof.BIP()

	return &input
}

// Custom returns custom field added to generate challenge point.
func (tx *ConfidentialTx) Custom() []*big.Int {
	customs := make([]*big.Int, 0)
	customs = append(customs, tx.nonce)
	customs = append(customs, tx.token)
	customs = append(customs, tx.pk1.X)
	customs = append(customs, tx.pk1.Y)
	customs = append(customs, tx.pk2.X)
	customs = append(customs, tx.pk2.Y)
	customs = append(customs, tx.transfer.X1.X)
	customs = append(customs, tx.transfer.X1.Y)
	customs = append(customs, tx.transfer.X2.X)
	customs = append(customs, tx.transfer.X2.Y)
	customs = append(customs, tx.transfer.Y.X)
	customs = append(customs, tx.transfer.Y.Y)

	return customs
}

// CreateConfidentialTx creates confidential transaction to transfer assets from alice to bob.
// alice和bob的值不要使用負數， 在進行加密時，會自動使用絕對值進行計算。
func CreateConfidentialTx(params proof.AggRangeParams, alice *Account, bob *ecdsa.PublicKey, v, token *big.Int) (*ConfidentialTx, error) {
	ctx := ConfidentialTx{}
	alicePublicKey := &alice.sk.PublicKey

	ctx.nonce = new(big.Int).SetUint64(alice.nonce)
	ctx.token = new(big.Int).Set(token)
	ctx.pk1 = new(utils.ECPoint).SetFromPublicKey(alicePublicKey)
	ctx.pk2 = new(utils.ECPoint).SetFromPublicKey(bob)

	transferEnc, err := proof.EncryptTransfer(params, alicePublicKey, bob, v.Bytes())
	if err != nil {
		return nil, err
	}
	ctx.transfer = transferEnc.Pub()
	ctx.sigmaPTEqualityProof, err = proof.GeneratePTEqualityProof(params, alicePublicKey, bob, transferEnc)
	if err != nil {
		return nil, err
	}
	ctx.balance = alice.balance
	updateBalanceCT := new(proof.CTEncPoint).Sub(ctx.balance, ctx.transfer.First())
	ctx.updatedBalance = updateBalanceCT.Copy()
	refreshBalanceCT, err := proof.Refresh(params, alice.sk, updateBalanceCT)
	// for speed up.
	refreshBalanceCT.EncMsg = new(big.Int).Sub(alice.m, v).Bytes()
	if err != nil {
		return nil, err
	}
	ctx.refreshBalance = refreshBalanceCT.CopyPublicPoint()
	customs := ctx.Custom()
	ctx.sigmaDlogeqProof, err = proof.GenerateDLESigmaProof(params, updateBalanceCT, ctx.refreshBalance,
		alice.sk, customs...)
	if err != nil {
		return nil, err
	}

	ctx.sigmaCTValidProof, err = proof.GenerateCTValidProof(params, alicePublicKey, refreshBalanceCT)
	if err != nil {
		return nil, err
	}

	vlist := make([]*big.Int, 0)
	vlist = append(vlist, new(big.Int).SetBytes(transferEnc.EncMsg))
	vlist = append(vlist, new(big.Int).SetBytes(refreshBalanceCT.EncMsg))
	random := make([]*big.Int, 0)
	random = append(random, transferEnc.R)
	random = append(random, refreshBalanceCT.R)
	ctx.bulletProof, err = proof.GenerateAggRangeProof(params, vlist, random)
	if err != nil {
		return nil, err
	}

	return &ctx, nil
}

// VerifyConfidentialTx .
func VerifyConfidentialTx(params proof.AggRangeParams, ctx *ConfidentialTx) bool {

	if !proof.VerifyPTEqualityProof(params, ctx.pk1.ToPublicKey(), ctx.pk2.ToPublicKey(), ctx.transfer, ctx.sigmaPTEqualityProof) {
		log.Warn("verify pte equality proof failed")
		return false
	}

	updatedBalance := new(proof.CTEncPoint).Sub(ctx.balance, ctx.transfer.First())

	customs := ctx.Custom()
	if !proof.VerifyDLESigmaProof(params, updatedBalance, ctx.refreshBalance, ctx.pk1.ToPublicKey(), ctx.sigmaDlogeqProof, customs...) {
		log.Warn("verify dle sigma proof failed")
		return false
	}

	if !proof.VerifyCTValidProof(params, ctx.pk1.ToPublicKey(), ctx.refreshBalance, ctx.sigmaCTValidProof) {
		log.Warn("verify ct valid proof failed")
		return false
	}

	vpoints := make([]*utils.ECPoint, 0)
	vpoints = append(vpoints, ctx.transfer.Y)
	vpoints = append(vpoints, ctx.refreshBalance.Y)
	if !proof.VerifyAggRangeProof(params, vpoints, ctx.bulletProof) {
		log.Warn("verify aggreate proof failed")
		return false
	}

	return true
}

// BurnETHTx includes info to burn an account and withdraw eth.
type BurnETHTx struct {
	Account  *utils.ECPoint       `json:"account"`
	Amount   *big.Int             `json:"amount"`
	Proof    *proof.DLESigmaProof `json:"proof"`
	Receiver common.Address
}

type burnEHTTxInput struct {
	Receiver common.Address
	Amount   *big.Int
	Points   []byte
	Z        *big.Int
}

func (btx *BurnETHTx) ToSolidityInput() *burnEHTTxInput {
	ps := make([]byte, 0)
	ps = append(ps, btx.Account.Compress()...)
	ps = append(ps, btx.Proof.A1.Compress()...)
	ps = append(ps, btx.Proof.A2.Compress()...)
	return &burnEHTTxInput{
		Receiver: btx.Receiver,
		Amount:   new(big.Int).Set(btx.Amount),
		Points:   ps,
		Z:        btx.Proof.Z,
	}
}

// CreateBurnETHTx creates tx to burn eth on chain.
func CreateBurnETHTx(params proof.AggRangeParams, alice *Account, receiver, token common.Address) (*BurnETHTx, error) {
	tx := BurnETHTx{}

	tx.Account = new(utils.ECPoint).SetFromPublicKey(&alice.sk.PublicKey)
	tx.Amount = new(big.Int).Set(alice.m)

	// generate proof to prove alice has the sk and the amount is indeed same with value encrypted.
	// alice's encrypted balance should be right set.
	proof, err := proof.GenerateEqualProof(params, tx.Amount, alice.balance, alice.sk, new(big.Int).SetUint64(alice.nonce),
		receiver.Hash().Big(), token.Hash().Big())
	if err != nil {
		return nil, err
	}
	tx.Proof = proof

	return &tx, nil
}

// VerifyBurnETHTx .
func VerifyBurnETHTx(params proof.AggRangeParams, nonce *big.Int, receiver common.Address, balance *proof.CTEncPoint, btx *BurnETHTx) bool {
	return proof.VerifyEqualProof(params, balance, btx.Amount, btx.Account, btx.Proof, nonce, receiver.Hash().Big())
}
