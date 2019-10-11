package pgc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// CTEncPoint respresents encrypted ct tx point on curve.
type CTEncPoint struct {
	// X=pk*r; Y=g*r + h*v.(r is randomness, v is the msg)
	X, Y *ECPoint
}

// Sub .
func (c *CTEncPoint) Sub(first, second *CTEncPoint) *CTEncPoint {
	ecPoints := CTEncPoint{}

	ecPoints.X = new(ECPoint).Sub(first.X, second.X)
	ecPoints.Y = new(ECPoint).Sub(first.Y, second.Y)
	return &ecPoints
}

// Copy .
func (c *CTEncPoint) Copy() *CTEncPoint {
	ecPoints := CTEncPoint{}

	ecPoints.X = c.X.Copy()
	ecPoints.Y = c.Y.Copy()
	return &ecPoints
}

// PublicKey represents a public key used in twisted elgamal.
type PublicKey struct {
	elliptic.Curve

	// Point on curve.
	X, Y *big.Int
}

// PrivateKey represents a private key used in twisted elgamal.
type PrivateKey struct {
	// sk
	D *big.Int

	PublicKey
}

// MRTwistedELGamalCTPub represent public points in MRTwistedELGamalCT tx.
type MRTwistedELGamalCTPub struct {
	// X1=pk1*r; X2=pk2*r.
	X1, X2 *ECPoint
	// Y = g*r + h*m.
	Y *ECPoint
}

func (mrp *MRTwistedELGamalCTPub) First() *CTEncPoint {
	return &CTEncPoint{
		X: mrp.X1,
		Y: mrp.Y,
	}
}

// MRTwistedELGamalCT defines the structure of 2R1M ciphertext (MR denotes multiple recipients).
type MRTwistedELGamalCT struct {
	MRTwistedELGamalCTPub
	R      *big.Int
	EncMsg []byte
}

func (mr *MRTwistedELGamalCT) Pub() *MRTwistedELGamalCTPub {
	return &mr.MRTwistedELGamalCTPub
}

// TwistedELGamalCT respresents a encrypted transaction encoded in twisted-elgamal format.
type TwistedELGamalCT struct {
	// X, Y
	CTEncPoint
	// Random number used in encryption.
	R *big.Int
	// encrypted msg.
	EncMsg []byte
}

// CopyPublicPoint copys public encrypted ec point.
func (ct *TwistedELGamalCT) CopyPublicPoint() *CTEncPoint {
	ecPoints := CTEncPoint{}

	ecPoints.X = ct.X.Copy()
	ecPoints.Y = ct.Y.Copy()
	return &ecPoints
}

// Sub return new instance = first - second.
func (ct *TwistedELGamalCT) Sub(first, second *TwistedELGamalCT) *TwistedELGamalCT {
	ct.X = new(ECPoint).Sub(first.X, second.X)
	ct.Y = new(ECPoint).Sub(first.Y, second.Y)
	ct.R = nil
	ct.EncMsg = []byte{}

	return ct
}

// Copy returns new instance.
func (ct *TwistedELGamalCT) Copy() *TwistedELGamalCT {
	newCT := TwistedELGamalCT{}
	newCT.X = ct.X.Copy()
	newCT.Y = ct.Y.Copy()
	if ct.R != nil {
		newCT.R = new(big.Int).Set(ct.R)
	}
	if len(ct.EncMsg) > 0 {
		newCT.EncMsg = make([]byte, len(ct.EncMsg))
		copy(newCT.EncMsg, ct.EncMsg)
	}

	return &newCT
}

// copy from crypto/ecdsa.
// randFieldElement returns a random element of the field underlying the given
// curve using the procedure given in [NSA] A.2.1.
func randFieldElement(c elliptic.Curve, rand io.Reader) (k *big.Int, err error) {
	params := c.Params()
	b := make([]byte, params.BitSize/8+8)
	_, err = io.ReadFull(rand, b)
	if err != nil {
		return
	}

	k = new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)
	return
}

// MustGenerateKey generates a key pair and panic if err.
// Warn: test purpose only.
func MustGenerateKey() *ecdsa.PrivateKey {
	key, err := GenerateKey()
	if err != nil {
		panic(err)
	}

	return key
}

// GenerateKey generates key pair.
// Warn: The h point from global params is used for generating key pair, not original
// g base point from curve.
func GenerateKey() (priv *ecdsa.PrivateKey, err error) {
	params := Params()
	curve := params.Curve()
	h := params.GetH()

	k, err := randFieldElement(curve, rand.Reader)
	if err != nil {
		return
	}

	priv = new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = k
	// Warng: x, y == h * sk.
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarMult(h.X, h.Y, k.Bytes())
	return
}

// HexToKey returns key pair by hex.
func HexToKey(hexKey string) (priv *ecdsa.PrivateKey, err error) {
	b, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}

	params := Params()
	h := params.GetH()
	priv = new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = params.Curve()
	priv.D = new(big.Int).SetBytes(b)
	priv.PublicKey.X, priv.PublicKey.Y = params.Curve().ScalarMult(h.X, h.Y, b)
	return
}

// EncryptOnChain encrypts msg with random 0.
func EncryptOnChain(pk *ecdsa.PublicKey, msg []byte) (*TwistedELGamalCT, error) {
	params := Params()
	msgBit := new(big.Int).SetBytes(msg)
	log.Debug("msg", "v", msgBit, "string(b)", string(msg))

	if msgBit.BitLen() > params.BitSizeLimit() {
		return nil, errors.New("msg reach size limit")
	}

	// Create ct instance.
	ct := new(TwistedELGamalCT)
	ct.X = NewEmptyECPoint(pk.Curve)
	ct.Y = NewEmptyECPoint(pk.Curve)

	// set curve
	curve := pk.Curve

	// set random 0.
	r := new(big.Int).SetUint64(0)

	// for sigma proof purpose.
	ct.R = new(big.Int).Set(r)
	ct.EncMsg = make([]byte, len(msg))
	copy(ct.EncMsg, msg)

	// compute pk * r.(pk ^ r)
	ct.X.SetFromPublicKey(pk)
	ct.X.ScalarMult(ct.X, r)

	// compute g * m.(g ^ m)
	g := params.GetG()
	ct.Y.ScalarMult(g, msgBit)
	log.Debug("encrypt msg(g^m)", "x", ct.Y.X, "y", ct.Y.Y)
	// compute h * r.(h ^ r)
	h := params.GetH()
	s2X, s2Y := curve.ScalarMult(h.X, h.Y, r.Bytes())
	// compute g * m + h * r.
	ct.Y.Add(ct.Y, NewECPoint(s2X, s2Y, curve))

	return ct, nil

}

// EncryptTransfer encrypts msg by two different pk but with same random.
func EncryptTransfer(sender, receiver *ecdsa.PublicKey, msg []byte) (*MRTwistedELGamalCT, error) {
	curve := sender.Curve
	// get random.
	r, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}
	twSender, err := EncryptWithRandom(sender, msg, r)
	if err != nil {
		return nil, err
	}
	twReceiver, err := EncryptWithRandom(receiver, msg, r)
	if err != nil {
		return nil, err
	}

	if !twReceiver.Y.Equal(twSender.Y) {
		return nil, errors.New("twisted elgamal y point not equal")
	}

	return &MRTwistedELGamalCT{
		MRTwistedELGamalCTPub: MRTwistedELGamalCTPub{
			X1: twSender.X,
			X2: twReceiver.X,
			Y:  twSender.Y,
		},
		R:      r,
		EncMsg: msg,
	}, nil
}

// Encrypt encrypts msg.
func Encrypt(pk *ecdsa.PublicKey, msg []byte) (*TwistedELGamalCT, error) {
	curve := pk.Curve
	// get random.
	r, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}

	return EncryptWithRandom(pk, msg, r)
}

// EncryptWithRandom encrypts msg in twisted elgamal way.
func EncryptWithRandom(pk *ecdsa.PublicKey, msg []byte, r *big.Int) (*TwistedELGamalCT, error) {
	params := Params()
	msgBit := new(big.Int).SetBytes(msg)

	if msgBit.BitLen() > params.BitSizeLimit() {
		return nil, errors.New("msg reach size limit")
	}

	// Create ct instance.
	ct := new(TwistedELGamalCT)
	ct.X = NewEmptyECPoint(pk.Curve)
	ct.Y = NewEmptyECPoint(pk.Curve)

	// set curve
	curve := pk.Curve

	// for sigma proof purpose.
	ct.R = new(big.Int).Set(r)
	ct.EncMsg = make([]byte, len(msg))
	copy(ct.EncMsg, msg)

	// compute pk * r.(pk ^ r)
	ct.X.SetFromPublicKey(pk)
	ct.X.ScalarMult(ct.X, r)

	// compute g * m.(g ^ m)
	g := params.GetG()
	ct.Y.ScalarMult(g, msgBit)
	// compute h * r.(h ^ r)
	h := params.GetH()
	s2X, s2Y := curve.ScalarMult(h.X, h.Y, r.Bytes())
	// compute g * m + h * r.
	ct.Y.Add(ct.Y, NewECPoint(s2X, s2Y, curve))

	return ct, nil
}

// Decrypt decrypts msg in twisted elgamal way.
func Decrypt(sk *ecdsa.PrivateKey, ct *CTEncPoint) []byte {
	encodedMsg := getEncryptedMsg(sk, ct.Copy())
	return decryptEncodedMsg(encodedMsg)
}

func getEncryptedMsg(sk *ecdsa.PrivateKey, ct *CTEncPoint) *ECPoint {
	// get public curve.
	curve := sk.PublicKey.Curve
	// compute the inverse of sk.
	skInverse := new(big.Int).ModInverse(sk.D, curve.Params().N)
	// compute X * skInverse(X ^ skInverse) == h * r.
	ct.X.ScalarMult(ct.X, skInverse)
	// use ct.Y - ct.X Point to get encoded msg.
	// get Affine negation formulas of ct.Y.
	encodedMsg := ct.Y.Sub(ct.Y, ct.X)

	return encodedMsg
}

// decryptencodedmsg decrypts and returns original bytes of msg.
// encodeMsg = g * m
func decryptEncodedMsg(encodeMsg *ECPoint) []byte {
	params := Params()
	bit := uint64(params.BitSizeLimit())
	upperLimit := new(big.Int).Exp(two, new(big.Int).SetUint64(bit), nil)
	g := params.GetG()

	// todo: uint64 may not enough if bit size bigger than 64.
	for i := uint64(0); i < upperLimit.Uint64(); i++ {
		point := new(ECPoint).ScalarMult(g, new(big.Int).SetUint64(i))
		if point.Equal(encodeMsg) {
			return new(big.Int).SetUint64(i).Bytes()
		}
	}

	return []byte{}
}

// Refresh Ciphertext refreshing algorithm: output a fresh ciphertext for the message encrypted in CT.
func Refresh(sk *ecdsa.PrivateKey, ct *CTEncPoint) (*TwistedELGamalCT, error) {
	params := Params()
	// get encrypted msg.
	encodedMsg := getEncryptedMsg(sk, ct.Copy())
	// encrypt.
	pk := sk.PublicKey
	refreshCT := new(TwistedELGamalCT)
	refreshCT.X = NewEmptyECPoint(pk.Curve)
	refreshCT.Y = NewEmptyECPoint(pk.Curve)

	// set curve
	curve := pk.Curve

	// get random.
	r, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}
	// for sigma proof purpose.
	refreshCT.R = new(big.Int).Set(r)
	// don't set the msg.

	// compute pk * r.(pk ^ r)
	refreshCT.X.SetFromPublicKey(&pk)
	refreshCT.X.ScalarMult(refreshCT.X, r)

	// set g * m.(g ^ m)
	refreshCT.Y = encodedMsg.Copy()
	// compute h * r.(h ^ r)
	h := params.GetH()
	s2X, s2Y := curve.ScalarMult(h.X, h.Y, r.Bytes())
	// compute g * m + h * r.
	refreshCT.Y.Add(refreshCT.Y, NewECPoint(s2X, s2Y, curve))

	return refreshCT, nil
}
