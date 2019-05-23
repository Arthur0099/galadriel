package pgc

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"

	log "github.com/inconshreveable/log15"
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

	sys := NewTwistedELGamalSystem()
	ctx := CTX{}
	ctx.pk1 = &alice.sk.PublicKey
	ctx.pk2 = &bob.sk.PublicKey
	ctx.senderBalance = alice.balance

	// encrypt the transfer balance.
	c1, err := sys.Encrypt(ctx.pk1, v.Bytes())
	c2, err := sys.Encrypt(ctx.pk2, v.Bytes())
	ctx.c1 = c1.CopyPublicPoint()
	ctx.c2 = c2.CopyPublicPoint()

	sigmaSys := NewSigmaSys()
	// generate proof.
	equalityProof, err := sigmaSys.GenerateProof(ctx.pk1, ctx.pk2, c1, c2)
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
	refreshUpdatedBalance, err := sys.Refresh(alice.sk, updatedBalance)

	if err != nil {
		return nil, err
	}

	ctx.refreshUpdatedBalance = refreshUpdatedBalance.CopyPublicPoint()
	// update alice balance and bob balance
	alice.balance = refreshUpdatedBalance.CopyPublicPoint()

	// generate proof to prove two ciphertexts encrypt the same value under same public key.
	dleProver := DLESigmaProver{}
	dleProver.params = Params()
	dleProof, err := dleProver.GenerateProof(updatedBalance, refreshUpdatedBalance.CopyPublicPoint(), alice.sk)
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

	return &ctx, nil
}

// VerifyCTX checks the ctx is valid or not.
func VerifyCTX(ctx *CTX) bool {
	params := Params()
	sigmaSys := NewSigmaSys()
	if !sigmaSys.VerifySigmaProof(ctx.equalityProof) {
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
