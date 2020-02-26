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
	nonce *big.Int
	// in real system, this will be on chain and not inculded in ctx.
	balance  *proof.CTEncPoint
	pk1, pk2 *utils.ECPoint
	transfer *proof.MRTwistedELGamalCTPub

	// proof
	sigmaPTEqualityProof *proof.PTEqualityProof
	bulletProof          *proof.AggRangeProof
	refreshBalance       *proof.CTEncPoint
	sigmaCTValidProof    *proof.CTValidProof
	sigmaDlogeqProof     *proof.DLESigmaProof
}

type solidityPGCInput struct {
	points [36]*big.Int

	l, r [proof.LRsize * 2]*big.Int

	scalars [10]*big.Int
}

// ToSolidityInput formats tx to solidity to verify contract
func (tx *ConfidentialTx) ToSolidityInput() *solidityPGCInput {
	input := solidityPGCInput{}
	input.points[0] = tx.pk1.X
	input.points[1] = tx.pk1.Y
	input.points[2] = tx.pk2.X
	input.points[3] = tx.pk2.Y
	input.points[4] = tx.transfer.X1.X
	input.points[5] = tx.transfer.X1.Y
	input.points[6] = tx.transfer.X2.X
	input.points[7] = tx.transfer.X2.Y
	input.points[8] = tx.transfer.Y.X
	input.points[9] = tx.transfer.Y.Y
	input.points[10] = tx.sigmaPTEqualityProof.A1.X
	input.points[11] = tx.sigmaPTEqualityProof.A1.Y
	input.points[12] = tx.sigmaPTEqualityProof.A2.X
	input.points[13] = tx.sigmaPTEqualityProof.A2.Y
	input.points[14] = tx.sigmaPTEqualityProof.B.X
	input.points[15] = tx.sigmaPTEqualityProof.B.Y
	input.points[16] = tx.refreshBalance.X.X
	input.points[17] = tx.refreshBalance.X.Y
	input.points[18] = tx.refreshBalance.Y.X
	input.points[19] = tx.refreshBalance.Y.Y
	input.points[20] = tx.sigmaCTValidProof.A.X
	input.points[21] = tx.sigmaCTValidProof.A.Y
	input.points[22] = tx.sigmaCTValidProof.B.X
	input.points[23] = tx.sigmaCTValidProof.B.Y
	input.points[24] = tx.sigmaDlogeqProof.A1.X
	input.points[25] = tx.sigmaDlogeqProof.A1.Y
	input.points[26] = tx.sigmaDlogeqProof.A2.X
	input.points[27] = tx.sigmaDlogeqProof.A2.Y
	// range proof.
	input.points[28] = tx.bulletProof.A.X
	input.points[29] = tx.bulletProof.A.Y
	input.points[30] = tx.bulletProof.S.X
	input.points[31] = tx.bulletProof.S.Y
	input.points[32] = tx.bulletProof.T1.X
	input.points[33] = tx.bulletProof.T1.Y
	input.points[34] = tx.bulletProof.T2.X
	input.points[35] = tx.bulletProof.T2.Y

	// L, R
	for i := 0; i < tx.bulletProof.Len(); i++ {
		input.l[i*2] = tx.bulletProof.Li(i).X
		input.l[i*2+1] = tx.bulletProof.Li(i).Y

		input.r[i*2] = tx.bulletProof.Ri(i).X
		input.r[i*2+1] = tx.bulletProof.Ri(i).Y
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

// CreateConfidentialTx creates confidential transaction to transfer assets from alice to bob.
// alice和bob的值不要使用負數， 在進行加密時，會自動使用絕對值進行計算。
func CreateConfidentialTx(params proof.AggRangeParams, alice *Account, bob *ecdsa.PublicKey, v *big.Int) (*ConfidentialTx, error) {
	ctx := ConfidentialTx{}
	alicePublicKey := &alice.sk.PublicKey

	ctx.nonce = new(big.Int).SetUint64(alice.nonce)
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
	refreshBalanceCT, err := proof.Refresh(params, alice.sk, updateBalanceCT)
	// for speed up.
	refreshBalanceCT.EncMsg = new(big.Int).Sub(alice.m, v).Bytes()
	if err != nil {
		return nil, err
	}
	ctx.refreshBalance = refreshBalanceCT.CopyPublicPoint()
	ctx.sigmaDlogeqProof, err = proof.GenerateDLESigmaProof(params, updateBalanceCT, ctx.refreshBalance, alice.sk, ctx.nonce)
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

	if !proof.VerifyDLESigmaProof(params, updatedBalance, ctx.refreshBalance, ctx.pk1.ToPublicKey(), ctx.sigmaDlogeqProof, ctx.nonce) {
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
	PubKey   [2]*big.Int
	Proof    [4]*big.Int
	Z        *big.Int
}

func (btx *BurnETHTx) ToSolidityInput() *burnEHTTxInput {
	return &burnEHTTxInput{
		Receiver: btx.Receiver,
		Amount:   new(big.Int).Set(btx.Amount),
		PubKey:   [2]*big.Int{btx.Account.X, btx.Account.Y},
		Proof:    [4]*big.Int{btx.Proof.A1.X, btx.Proof.A1.Y, btx.Proof.A2.X, btx.Proof.A2.Y},
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
