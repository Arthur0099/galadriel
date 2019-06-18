package pgc

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"

	log "github.com/inconshreveable/log15"
)

var (
	decimal   = new(big.Int).SetUint64(2)
	base      = new(big.Int).SetUint64(10)
	precision = new(big.Int).Exp(base, decimal, nil)
)

// CTX represents a encrypted tx on chain.
type CTX struct {
	nonce uint64
	// public of two accounts.
	pk1, pk2                 *ecdsa.PublicKey
	rangeProof1, rangeProof2 *RangeProof
	c1, c2                   *CTEncPoint
	// proof proving two encryptions(ct1, ct2) of the same value.
	equalityProof *SigmaProof
	// proof proving DLESigmaProof.
	dleProof *DLESigmaProof

	//
	senderBalance, refreshUpdatedBalance *CTEncPoint
}

type transferTx struct {
	points  [28]*big.Int
	scalar  [14]*big.Int
	rpoints [16]*big.Int
	l       [16]*big.Int
	r       [16]*big.Int
	nonce   *big.Int
}

type burnTx struct {
	pk     [2]*big.Int
	amount *big.Int
	proof  [4]*big.Int
	z      *big.Int
	nonce  *big.Int
}

type burnPartTx struct {
	amount  *big.Int
	points  [18]*big.Int
	scalar  [12]*big.Int
	rpoints [16]*big.Int
	l       [16]*big.Int
	r       [16]*big.Int
	nonce   *big.Int
}

// MarshalJSON defines custom way to JSON.
func (ctx *CTX) MarshalJSON() ([]byte, error) {
	newJSON := struct {
		Pk1                    *ECPoint       `json:"pk1"`
		Pk2                    *ECPoint       `json:"pk2"`
		RangeProof1            *RangeProof    `json:"rangeProof1"`
		RangeProof2            *RangeProof    `json:"rangeProof2"`
		C1                     *CTEncPoint    `json:"c1"`
		C2                     *CTEncPoint    `json:"c2"`
		EqualityProof          *SigmaProof    `json:"equalityProof"`
		DleProof               *DLESigmaProof `json:"dleProof"`
		SenderBalance          *CTEncPoint    `json:"senderBalance"`
		RefreashUpdatedBalance *CTEncPoint    `json:"refreshUpdatedBalance"`
	}{
		Pk1:                    new(ECPoint).SetFromPublicKey(ctx.pk1),
		Pk2:                    new(ECPoint).SetFromPublicKey(ctx.pk2),
		RangeProof1:            ctx.rangeProof1,
		RangeProof2:            ctx.rangeProof2,
		EqualityProof:          ctx.equalityProof,
		DleProof:               ctx.dleProof,
		SenderBalance:          ctx.senderBalance,
		RefreashUpdatedBalance: ctx.refreshUpdatedBalance,
	}

	return json.Marshal(&newJSON)
}

// CreateCTX creates en encrypt tx for transfering v balance from
// alice account to bob account.
func CreateCTX(alice, bob *Account, v *big.Int) (*CTX, error) {
	log.Debug("create ctx")
	params := Params()

	ctx := CTX{}
	ctx.pk1 = &alice.sk.PublicKey
	ctx.pk2 = &bob.sk.PublicKey
	ctx.senderBalance = alice.balance
	ctx.nonce = alice.nonce

	// encrypt the transfer balance.
	c1, err := Encrypt(ctx.pk1, v.Bytes())
	c2, err := Encrypt(ctx.pk2, v.Bytes())
	ctx.c1 = c1.CopyPublicPoint()
	ctx.c2 = c2.CopyPublicPoint()

	// generate proof.
	equalityProof, err := GenerateSigmaProof(ctx.pk1, ctx.pk2, c1, c2)
	if err != nil {
		return nil, err
	}
	ctx.equalityProof = equalityProof
	// generate range proof v in range.
	rangeProof1, err := GenerateRangeProof(params.GetVB(), v, c1.R)
	if err != nil {
		return nil, err
	}
	ctx.rangeProof1 = rangeProof1

	// updated balance of alice.
	updatedBalance := new(CTEncPoint).Sub(alice.balance, c1.CopyPublicPoint())
	log.Debug("alice' updated balance", "x.x", updatedBalance.X.X, "x.y", updatedBalance.X.Y, "y.x", updatedBalance.Y.X, "y.y", updatedBalance.Y.Y)
	refreshUpdatedBalance, err := Refresh(alice.sk, updatedBalance)

	if err != nil {
		return nil, err
	}

	ctx.refreshUpdatedBalance = refreshUpdatedBalance.CopyPublicPoint()

	// generate proof to prove two ciphertexts encrypt the same value under same public key.
	dleProof, err := GenerateDLESigmaProof(updatedBalance, refreshUpdatedBalance.CopyPublicPoint(), alice.sk)
	if err != nil {
		return nil, err
	}
	ctx.dleProof = dleProof

	// generate proof to prove the refreshUpdatedBalance in right range.
	rangeProof2, err := GenerateRangeProof(params.GetVB(), new(big.Int).Sub(alice.m, v), refreshUpdatedBalance.R)
	if err != nil {
		return nil, err
	}
	ctx.rangeProof2 = rangeProof2

	// update info.
	// update alice balance and bob balance
	alice.balance = updatedBalance.Copy()
	alice.nonce = alice.nonce + 1

	return &ctx, nil
}

// BurnPartTx represent a tx to burn part/all balance of a pgc amount.
type BurnPartTx struct {
	nonce  uint64
	pk     *ecdsa.PublicKey
	amount *big.Int
	// encry amount.
	ct *CTEncPoint
	// proof 1 to prove amount is same with balance in ct.
	// proof 2 to prove account's updated balance is same with
	// refreshBalance.
	dleProof1, dleProof2 *DLESigmaProof
	refreshBalance       *CTEncPoint
	// proof 1 to prove amount in range.
	// proof 2 to prove account's balance - amount in range.
	proof1, proof2 *RangeProof
}

// CreateBurnPartTx creates tx to burn part or all balance of a pgc amount.
func CreateBurnPartTx(alice *Account, amount *big.Int) (*BurnPartTx, error) {
	burnPartTx := BurnPartTx{}
	burnPartTx.pk = &alice.sk.PublicKey
	burnPartTx.nonce = alice.nonce
	burnPartTx.amount = new(big.Int).Set(amount)

	// encrypt amount.
	ct, err := Encrypt(burnPartTx.pk, amount.Bytes())
	if err != nil {
		return nil, err
	}
	burnPartTx.ct = ct.CopyPublicPoint()

	// generate proof to prove amount is same with value in ct.
	dleProof1, err := GenerateEqualProof(amount, ct.CopyPublicPoint(), alice.sk)
	if err != nil {
		return nil, err
	}
	burnPartTx.dleProof1 = dleProof1

	// generate proof.
	proof1, err := GenerateRangeProof(params.GetVB(), amount, ct.R)
	if err != nil {
		return nil, err
	}
	burnPartTx.proof1 = proof1

	// update balace of alice.
	updatedBalance := new(CTEncPoint).Sub(alice.balance, ct.CopyPublicPoint())
	refreshBalance, err := Refresh(alice.sk, updatedBalance)
	if err != nil {
		return nil, err
	}
	burnPartTx.refreshBalance = refreshBalance.CopyPublicPoint()
	//
	dleProof2, err := GenerateDLESigmaProof(updatedBalance, refreshBalance.CopyPublicPoint(), alice.sk)
	if err != nil {
		return nil, err
	}
	burnPartTx.dleProof2 = dleProof2
	proof2, err := GenerateRangeProof(params.GetVB(), new(big.Int).Sub(alice.m, amount), refreshBalance.R)
	if err != nil {
		return nil, err
	}
	burnPartTx.proof2 = proof2

	// update info.
	alice.balance = updatedBalance.Copy()
	alice.nonce = alice.nonce + 1

	return &burnPartTx, nil
}

// VerifyCTX checks the ctx is valid or not.
func VerifyCTX(ctx *CTX) bool {
	params := Params()
	if !VerifySigmaProof(ctx.equalityProof) {
		log.Warn("two encrypted value not same")
		return false
	}

	if !VerifyRangeProof(params.GetVB(), ctx.c1.Y, ctx.rangeProof1) {
		log.Warn("range proof for transfer v failed")
		return false
	}

	updatedBalance := new(CTEncPoint).Sub(ctx.senderBalance, ctx.c1)

	if !VerifyDLESigmaProof(updatedBalance, ctx.refreshUpdatedBalance, ctx.pk1, ctx.dleProof) {
		log.Warn("dle proof failed")
		return false
	}

	if !VerifyRangeProof(params.GetVB(), ctx.refreshUpdatedBalance.Y, ctx.rangeProof2) {
		log.Warn("range proof for refreshed balance failed")
		return false
	}

	return true
}

// FundTX includes info to fund an account.
type FundTX struct {
	account *ECPoint
	amount  *big.Int
}

// CreateFundTX creates tx to fund an account.
func CreateFundTX(alice *Account, amount *big.Int) *FundTX {
	tx := FundTX{}
	tx.account = new(ECPoint).SetFromPublicKey(&alice.sk.PublicKey)
	limit := Params().Max()
	if amount.Cmp(limit) >= 0 {
		panic("out of limit")
	}

	tx.amount = new(big.Int).Set(amount)
	return &tx
}

// BurnTx includes info to burn an account and withdraw eth.
type BurnTx struct {
	Account *ECPoint       `json:"account"`
	Amount  *big.Int       `json:"amount"`
	Proof   *DLESigmaProof `json:"proof"`
	Nonce   *big.Int       `json:"nonce"`
}

// CreateBurnTx creates tx to burn an account.
func CreateBurnTx(alice *Account, amount *big.Int) (*BurnTx, error) {
	tx := BurnTx{}
	tx.Nonce = new(big.Int).SetUint64(alice.nonce)

	//
	tx.Account = new(ECPoint).SetFromPublicKey(&alice.sk.PublicKey)
	tx.Amount = new(big.Int).Set(amount)

	// generate proof to prove alice has the sk and the amount is indeed same with value encrypted.
	// alice's encrypted balance should be right set.
	proof, err := GenerateEqualProof(amount, alice.balance, alice.sk)
	if err != nil {
		return nil, err
	}

	tx.Proof = proof
	alice.nonce = alice.nonce + 1
	// todo: update alice's balance.

	return &tx, nil
}
