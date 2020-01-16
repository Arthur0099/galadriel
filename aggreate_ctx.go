package pgc

import (
	"crypto/ecdsa"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// PGCCtx is a tx for pgc transfer system(using aggreate bulletproof).
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

func (pctx *PGCCtx) logPk() {
	log.Info("pk2", "x", pctx.pk2.X, "y", pctx.pk2.Y)
}

func (pctx *PGCCtx) logPteProof() {
	log.Info("A1", "x", pctx.sigmaPTEqualityProof.A1.X, "y", pctx.sigmaPTEqualityProof.A1.Y)
	log.Info("A2", "x", pctx.sigmaPTEqualityProof.A2.X, "y", pctx.sigmaPTEqualityProof.A2.Y)
	log.Info("B", "x", pctx.sigmaPTEqualityProof.B.X, "y", pctx.sigmaPTEqualityProof.B.Y)
	log.Info("z1", "v", pctx.sigmaPTEqualityProof.Z1)
	log.Info("z2", "v", pctx.sigmaPTEqualityProof.Z2)
}

type solidityPGCInput struct {
	points [36]*big.Int

	l, r [12]*big.Int

	scalars [11]*big.Int
}

func (tx *PGCCtx) ExportToSolidityInput() *solidityPGCInput {
	// c := BN128{}
	// points := [4]*big.Int{}
	// points[0] = new(big.Int).SetBytes(c.Marshal(tx.balance.X.X, tx.balance.X.Y))
	// use point .

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
	for i := 0; i < len(tx.bulletProof.ipProof.L); i++ {
		input.l[i*2] = tx.bulletProof.ipProof.L[i].X
		input.l[i*2+1] = tx.bulletProof.ipProof.L[i].Y

		input.r[i*2] = tx.bulletProof.ipProof.R[i].X
		input.r[i*2+1] = tx.bulletProof.ipProof.R[i].Y
	}

	// scalar
	input.scalars[0] = tx.nonce
	input.scalars[1] = tx.sigmaPTEqualityProof.Z1
	input.scalars[2] = tx.sigmaPTEqualityProof.Z2
	input.scalars[3] = tx.sigmaCTValidProof.Z1
	input.scalars[4] = tx.sigmaCTValidProof.Z2
	input.scalars[5] = tx.sigmaDlogeqProof.Z
	// range proof.
	input.scalars[6] = tx.bulletProof.t
	input.scalars[7] = tx.bulletProof.tx
	input.scalars[8] = tx.bulletProof.u
	// inner proof.
	input.scalars[9] = tx.bulletProof.ipProof.a
	input.scalars[10] = tx.bulletProof.ipProof.b

	return &input
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
