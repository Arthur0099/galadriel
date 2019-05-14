package pgc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"io"
	"math/big"

	log "github.com/inconshreveable/log15"
)

var one = new(big.Int).SetUint64(1)
var two = new(big.Int).SetUint64(2)

// CTEncPoint respresents encrypted ct tx point on curve.
type CTEncPoint struct {
	// X=pk*r; Y=g*r + h*v.(r is randomness, v is the msg)
	X, Y *ECPoint
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

// TwistedELGamalSystem represents twisted elgamal system for both encryption and decryption.
type TwistedELGamalSystem struct {
	params TwistedELGamalPublicParams
}

// NewTwistedELGamalSystem returns instance of TwistedELGamalSystem.
func NewTwistedELGamalSystem() *TwistedELGamalSystem {
	t := TwistedELGamalSystem{}
	t.params = Params()

	return &t
}

// GenerateKey generates key pair using btcec s256 curve.
func (twELGSys *TwistedELGamalSystem) GenerateKey() (priv *ecdsa.PrivateKey, err error) {
	curve := twELGSys.params.Curve()
	h := twELGSys.params.GetH()

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

// Encrypt encrypts msg by in twisted elgamal way.
func (twELGSys *TwistedELGamalSystem) Encrypt(pk *ecdsa.PublicKey, msg []byte) (*TwistedELGamalCT, error) {
	msgBit := new(big.Int).SetBytes(msg)
	log.Debug("msg", "v", msgBit, "string(b)", string(msg))

	if msgBit.BitLen() > twELGSys.params.BitSizeLimit() {
		return nil, errors.New("msg reach size limit")
	}

	// Create ct instance.
	ct := new(TwistedELGamalCT)
	ct.X = NewEmptyECPoint(pk.Curve)
	ct.Y = NewEmptyECPoint(pk.Curve)

	// set curve
	curve := pk.Curve

	// get random.
	r, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}
	// for sigma proof purpose.
	ct.R = new(big.Int).Set(r)
	ct.EncMsg = make([]byte, len(msg))
	copy(ct.EncMsg, msg)

	// compute pk * r.(pk ^ r)
	ct.X.SetFromPublicKey(pk)
	ct.X.ScalarMult(ct.X, r)

	// compute g * m.(g ^ m)
	g := twELGSys.params.GetG()
	ct.Y.ScalarMult(g, msgBit)
	log.Debug("encrypt msg(g^m)", "x", ct.Y.X, "y", ct.Y.Y)
	// compute h * r.(h ^ r)
	h := twELGSys.params.GetH()
	s2X, s2Y := curve.ScalarMult(h.X, h.Y, r.Bytes())
	// compute g * m + h * r.
	ct.Y.Add(ct.Y, NewECPoint(s2X, s2Y, curve))

	return ct, nil
}

// Decrypt decrypts msg in twisted elgamal way.
func (twELGSys *TwistedELGamalSystem) Decrypt(sk *ecdsa.PrivateKey, ct *TwistedELGamalCT) []byte {
	// get public curve.
	curve := sk.PublicKey.Curve
	// compute the inverse of sk.
	skInverse := new(big.Int).ModInverse(sk.D, curve.Params().N)
	// compute X * skInverse(X ^ skInverse) == h * r.
	ct.X.ScalarMult(ct.X, skInverse)
	// use ct.Y - ct.X Point to get encoded msg.
	// get Affine negation formulas of ct.Y.
	encodedMsg := ct.Y.Sub(ct.Y, ct.X)
	log.Debug("decrypt msg(g^m)", "x", encodedMsg.X, "y", encodedMsg.Y)

	return twELGSys.decryptEncodedMsg(encodedMsg)
}

// Decryptencodedmsg decrypts and returns original bytes of msg.
// encodeMsg = g * m
func (twELGSys *TwistedELGamalSystem) decryptEncodedMsg(encodeMsg *ECPoint) []byte {
	bit := uint64(twELGSys.params.BitSizeLimit())
	upperLimit := new(big.Int).Exp(two, new(big.Int).SetUint64(bit), nil)
	g := twELGSys.params.GetG()

	for i := uint64(1); i < upperLimit.Uint64(); i++ {
		point := new(ECPoint).ScalarMult(g, new(big.Int).SetUint64(i))
		if point.Equals(encodeMsg) {
			return new(big.Int).SetUint64(i).Bytes()
		}
	}

	return []byte{}
}
