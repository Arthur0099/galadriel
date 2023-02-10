package pgcsys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/inconshreveable/log15"
	"github.com/pgc/proof"
	"github.com/pgc/utils"
)

type PgcSys interface {
	proof.AggRangeParams
	Pub() *utils.ECPoint
}

type pgcSysParams struct {
	gv, hv           *utils.GeneratorVector
	u, g, h, pub     *utils.ECPoint
	bitsize, aggsize int
}

func NewPgcSysRandomParams(curve elliptic.Curve, bitsize, aggsize int) PgcSys {
	arp := pgcSysParams{}
	arp.gv = utils.NewRandomGeneratorVector(curve, bitsize*aggsize)
	arp.hv = utils.NewRandomGeneratorVector(curve, bitsize*aggsize)
	arp.u = utils.NewRandomECPoint(curve)
	arp.g = utils.NewRandomECPoint(curve)
	arp.h = utils.NewRandomECPoint(curve)
	key := proof.MustGenerateKey(&arp)
	arp.pub = new(utils.ECPoint).SetFromPublicKey(&key.PublicKey)
	arp.bitsize = bitsize
	arp.aggsize = aggsize

	return &arp
}

func (arp *pgcSysParams) Bitsize() int {
	return arp.bitsize
}

func (arp *pgcSysParams) Aggsize() int {
	return arp.aggsize
}

func (arp *pgcSysParams) Curve() elliptic.Curve {
	return arp.u.Curve
}

func (arp *pgcSysParams) GV() *utils.GeneratorVector {
	return arp.gv
}

func (arp *pgcSysParams) HV() *utils.GeneratorVector {
	return arp.hv
}

func (arp *pgcSysParams) U() *utils.ECPoint {
	return arp.u
}

func (arp *pgcSysParams) G() *utils.ECPoint {
	return arp.g
}

func (arp *pgcSysParams) H() *utils.ECPoint {
	return arp.h
}

func (arp *pgcSysParams) Pub() *utils.ECPoint {
	return arp.pub
}

// MarshalJSON define custom way to marshal json
func (arp *pgcSysParams) MarshalJSON() ([]byte, error) {
	newJSON := struct {
		GV, HV           *utils.GeneratorVector
		U, G, H, Pub     *utils.ECPoint
		Bitsize, Aggsize int
	}{
		GV:      arp.gv,
		HV:      arp.hv,
		U:       arp.u,
		G:       arp.g,
		H:       arp.h,
		Pub:     arp.pub,
		Bitsize: arp.bitsize,
		Aggsize: arp.aggsize,
	}

	return json.Marshal(&newJSON)
}

// UnmarshalJSON defines custom way to unmarshal
func (arp *pgcSysParams) UnmarshalJSON(bz []byte) error {
	var data struct {
		GV, HV           *utils.GeneratorVector
		U, G, H, Pub     *utils.ECPoint
		Bitsize, Aggsize int
	}

	if err := json.Unmarshal(bz, &data); err != nil {
		return err
	}

	*arp = pgcSysParams{
		gv:      data.GV,
		hv:      data.HV,
		u:       data.U,
		g:       data.G,
		h:       data.H,
		pub:     data.Pub,
		bitsize: data.Bitsize,
		aggsize: data.Aggsize,
	}
	return nil
}

// ConfidentialTx is a tx for pgc transfer system(using aggreate bulletproof).
type ConfidentialTx struct {
	nonce, token *big.Int
	// in real system, this will be on chain and not included in ctx.
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

func (ctx *ConfidentialTx) Transfer() *proof.MRTwistedELGamalCTPub {
	return ctx.transfer
}

func (ctx *ConfidentialTx) RefreshBalance() *proof.CTEncPoint {
	return ctx.refreshBalance
}

func (ctx *ConfidentialTx) AggRangeProof() *proof.AggRangeProof {
	return ctx.bulletProof
}

type SolidityPGCInput struct {
	Points []byte

	Lr []byte

	Scalars [10]*big.Int
}

func (sin *SolidityPGCInput) MarshalJSON() ([]byte, error) {
	out := struct {
		Points string

		Lr string

		Scalars []string
	}{
		Points: "0x" + hex.EncodeToString(sin.Points),
		Lr:     "0x" + hex.EncodeToString(sin.Lr),
	}

	s := make([]string, 0)
	for _, v := range sin.Scalars {
		s = append(s, v.String())
	}
	out.Scalars = s

	return json.Marshal(&out)
}

// ToSolidityInput formats tx to solidity to verify contract
func (tx *ConfidentialTx) ToSolidityInput() *SolidityPGCInput {
	input := SolidityPGCInput{}
	input.Points = make([]byte, 0)
	input.Points = append(input.Points, tx.pk1.Compress()...)
	input.Points = append(input.Points, tx.pk2.Compress()...)
	input.Points = append(input.Points, tx.transfer.X1.Compress()...)
	input.Points = append(input.Points, tx.transfer.X2.Compress()...)
	input.Points = append(input.Points, tx.transfer.X3.Compress()...)
	input.Points = append(input.Points, tx.transfer.Y.Compress()...)
	input.Points = append(input.Points, tx.sigmaPTEqualityProof.A1.Compress()...)
	input.Points = append(input.Points, tx.sigmaPTEqualityProof.A2.Compress()...)
	input.Points = append(input.Points, tx.sigmaPTEqualityProof.A3.Compress()...)
	input.Points = append(input.Points, tx.sigmaPTEqualityProof.B.Compress()...)
	input.Points = append(input.Points, tx.refreshBalance.X.Compress()...)
	input.Points = append(input.Points, tx.refreshBalance.Y.Compress()...)
	input.Points = append(input.Points, tx.sigmaCTValidProof.A.Compress()...)
	input.Points = append(input.Points, tx.sigmaCTValidProof.B.Compress()...)
	input.Points = append(input.Points, tx.sigmaDlogeqProof.A1.Compress()...)
	input.Points = append(input.Points, tx.sigmaDlogeqProof.A2.Compress()...)
	// range proof.
	input.Points = append(input.Points, tx.bulletProof.A.Compress()...)
	input.Points = append(input.Points, tx.bulletProof.S.Compress()...)
	input.Points = append(input.Points, tx.bulletProof.T1.Compress()...)
	input.Points = append(input.Points, tx.bulletProof.T2.Compress()...)

	// L, R
	input.Lr = make([]byte, 0)
	for i := 0; i < tx.bulletProof.Len(); i++ {
		input.Lr = append(input.Lr, tx.bulletProof.Li(i).Compress()...)
	}
	for i := 0; i < tx.bulletProof.Len(); i++ {
		input.Lr = append(input.Lr, tx.bulletProof.Ri(i).Compress()...)
	}

	// scalar
	input.Scalars[0] = tx.sigmaPTEqualityProof.Z1
	input.Scalars[1] = tx.sigmaPTEqualityProof.Z2
	input.Scalars[2] = tx.sigmaCTValidProof.Z1
	input.Scalars[3] = tx.sigmaCTValidProof.Z2
	input.Scalars[4] = tx.sigmaDlogeqProof.Z
	// range proof.
	input.Scalars[5] = tx.bulletProof.T()
	input.Scalars[6] = tx.bulletProof.TX()
	input.Scalars[7] = tx.bulletProof.U()
	// inner proof.
	input.Scalars[8] = tx.bulletProof.AIP()
	input.Scalars[9] = tx.bulletProof.BIP()

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
	customs = append(customs, tx.transfer.X3.X)
	customs = append(customs, tx.transfer.X3.Y)
	customs = append(customs, tx.transfer.Y.X)
	customs = append(customs, tx.transfer.Y.Y)

	return customs
}

// CreateConfidentialTx creates confidential transaction to transfer assets from alice to bob.
func CreateConfidentialTx(params PgcSys, alice *Account, bob *ecdsa.PublicKey, v, token *big.Int) (*ConfidentialTx, error) {
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
func VerifyConfidentialTx(params PgcSys, ctx *ConfidentialTx) bool {

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

type BurnEHTTxInput struct {
	Receiver common.Address
	Amount   *big.Int
	Points   []byte
	Z        *big.Int
}

func (bin *BurnEHTTxInput) MarshalJSON() ([]byte, error) {
	out := struct {
		Receiver common.Address
		Amount   string
		Points   string
		Z        string
	}{
		Receiver: bin.Receiver,
		Points:   "0x" + hex.EncodeToString(bin.Points),
		Amount:   bin.Amount.String(),
		Z:        bin.Z.String(),
	}

	return json.Marshal(&out)
}

func (btx *BurnETHTx) ToSolidityInput() *BurnEHTTxInput {
	ps := make([]byte, 0)
	ps = append(ps, btx.Account.Compress()...)
	ps = append(ps, btx.Proof.A1.Compress()...)
	ps = append(ps, btx.Proof.A2.Compress()...)
	return &BurnEHTTxInput{
		Receiver: btx.Receiver,
		Amount:   new(big.Int).Set(btx.Amount),
		Points:   ps,
		Z:        btx.Proof.Z,
	}
}

// CreateBurnETHTx creates tx to burn eth on chain.
func CreateBurnETHTx(params PgcSys, alice *Account, receiver, token common.Address) (*BurnETHTx, error) {
	tx := BurnETHTx{}

	tx.Account = new(utils.ECPoint).SetFromPublicKey(&alice.sk.PublicKey)
	tx.Amount = new(big.Int).Set(alice.m)
	tx.Receiver = receiver

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
func VerifyBurnETHTx(params PgcSys, nonce *big.Int, balance *proof.CTEncPoint, btx *BurnETHTx) bool {
	return proof.VerifyEqualProof(params, balance, btx.Amount, btx.Account, btx.Proof, nonce, btx.Receiver.Hash().Big())
}
