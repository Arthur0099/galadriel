package pgc

import (
	"crypto/ecdsa"
	"math/big"

	log "github.com/inconshreveable/log15"
)

//
type PGCCtx struct {
	nonce *big.Int
	// in real system, this will be on chain and not inculded in ctx.
	balance  *CTEncPoint
	pk1, pk2 *ECPoint
	transfer *MRTwistedELGamalCTPub

	// proof
	sigmaPTEqualityProof *PTEqualityProof
	bulletProof          *AggreateBulletProof
	refreshBalance       *CTEncPoint
	sigmaCTValidProof    *CTValidProof
	sigmaDlogeqProof     *DLESigmaProof
}

func CreatePGCCtx(alice *Account, bob *ecdsa.PublicKey, v *big.Int) (*PGCCtx, error) {
	ctx := PGCCtx{}
	alicePublicKey := &alice.sk.PublicKey
	ctx.nonce = new(big.Int).SetUint64(alice.nonce)
	ctx.pk1 = new(ECPoint).SetFromPublicKey(alicePublicKey)
	ctx.pk2 = new(ECPoint).SetFromPublicKey(bob)

	transferEnc, err := EncryptTransfer(alicePublicKey, bob, v.Bytes())
	if err != nil {
		return nil, err
	}
	ctx.transfer = transferEnc.Pub()
	ctx.sigmaPTEqualityProof, err = GeneratePTEqualityProof(alicePublicKey, bob, transferEnc)
	if err != nil {
		return nil, err
	}
	ctx.balance = alice.balance
	updateBalanceCT := new(CTEncPoint).Sub(ctx.balance, ctx.transfer.First())
	refreshBalanceCT, err := Refresh(alice.sk, updateBalanceCT)
	// for speed up.
	refreshBalanceCT.EncMsg = new(big.Int).Sub(alice.m, v).Bytes()
	if err != nil {
		return nil, err
	}
	ctx.refreshBalance = refreshBalanceCT.CopyPublicPoint()
	ctx.sigmaDlogeqProof, err = GenerateDLESigmaProof(updateBalanceCT, ctx.refreshBalance, alice.sk)
	if err != nil {
		return nil, err
	}

	ctx.sigmaCTValidProof, err = GenerateCTValidProof(alicePublicKey, refreshBalanceCT)
	if err != nil {
		return nil, err
	}

	vb := Params().GetVB()
	vlist := make([]*big.Int, 0)
	vlist = append(vlist, new(big.Int).SetBytes(transferEnc.EncMsg))
	vlist = append(vlist, new(big.Int).SetBytes(refreshBalanceCT.EncMsg))
	random := make([]*big.Int, 0)
	random = append(random, transferEnc.R)
	random = append(random, refreshBalanceCT.R)
	ctx.bulletProof, err = GenerateAggreateBulletProof(vb, vlist, random)
	if err != nil {
		return nil, err
	}

	return &ctx, nil
}

func VerifyPGCCtx(ctx *PGCCtx) bool {

	if !VerifyPTEqualityProof(ctx.pk1.ToPublicKey(), ctx.pk2.ToPublicKey(), ctx.transfer, ctx.sigmaPTEqualityProof) {
		log.Warn("verify pte equality proof failed")
		return false
	}

	updatedBalance := new(CTEncPoint).Sub(ctx.balance, ctx.transfer.First())

	if !VerifyDLESigmaProof(updatedBalance, ctx.refreshBalance, ctx.pk1.ToPublicKey(), ctx.sigmaDlogeqProof) {
		log.Warn("verify dle sigma proof failed")
		return false
	}

	if !VerifyCTValidProof(ctx.pk1.ToPublicKey(), ctx.refreshBalance, ctx.sigmaCTValidProof) {
		log.Warn("verify ct valid proof failed")
		return false
	}

	vb := Params().GetVB()
	vpoints := make([]*ECPoint, 0)
	vpoints = append(vpoints, ctx.transfer.Y)
	vpoints = append(vpoints, ctx.refreshBalance.Y)
	if !VerifyAggreateBulletProof(vb, vpoints, ctx.bulletProof) {
		log.Warn("verify aggreate proof failed")
		return false
	}

	return true
}
