package pgc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// CTEncPoint respresents encrypted ct tx point on curve.
type CTEncPoint struct {
	X, Y *ECPoint
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

// GenerateKey generates key pair using btcec s256 curve.
func GenerateKey() (key *ecdsa.PrivateKey, err error) {
	key, err = ecdsa.GenerateKey(S256(), rand.Reader)
	return
}

// Encrypt encrypts msg by in twisted elgamal way.
func Encrypt(pk *ecdsa.PublicKey, msg []byte) (*TwistedELGamalCT, error) {
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
	// compute g * r.(g ^ r)
	ct.Y.ScalarBaseMult(r)
	// compute h * m.(h ^ m)
	pubParams := Params()
	s2X, s2Y := curve.ScalarMult(pubParams.ElgGenerator.X, pubParams.ElgGenerator.Y, msg)
	log.Debug("encrypt msg(h^m)", "x", s2X, "y", s2Y)

	// compute g * r + h * m.
	ct.Y.Add(ct.Y, NewECPoint(s2X, s2Y, curve))

	return ct, nil
}

// Decrypt decrypts msg in twisted elgamal way.
func Decrypt(sk *ecdsa.PrivateKey, ct *TwistedELGamalCT) []byte {
	// get public curve.
	curve := sk.PublicKey.Curve
	// compute the inverse of sk.
	skInverse := new(big.Int).ModInverse(sk.D, curve.Params().N)
	// compute X * skInverse(X ^ skInverse) == g * r.
	ct.X.ScalarMult(ct.X, skInverse)
	// use ct.Y - ct.X Point to get encoded msg.
	// get Affine negation formulas of ct.Y.
	encodedMsg := ct.Y.Sub(ct.Y, ct.X)
	log.Debug("decrypt msg(h^m)", "x", encodedMsg.X, "y", encodedMsg.Y)
	_ = encodedMsg

	// todo: decrypt encoded msg.
	return []byte{}
}

// DecryptEncodedMsg decrypts and returns original bytes of msg.
func DecryptEncodedMsg() []byte {
	//
	return []byte{}
}

// SubECPoint computes x - y.
func SubECPoint(x, y *ecdsa.PublicKey) *ecdsa.PublicKey {
	curve := x.Curve

	// create instance of new point.
	newPoint := new(ecdsa.PublicKey)
	newPoint.Curve = curve

	// get negation of x.
	negation := negation(y)
	// calculate negation of new point.
	newPointNegaX, newPointNegaY := curve.Add(negation.X, negation.Y, x.X, x.Y)
	newPoint.X = newPointNegaX
	// set new point y to symmetry
	newPoint.Y = newPointNegaY

	return newPoint
}

// negation returns negation of ecpoint x on curve.
func negation(x *ecdsa.PublicKey) *ecdsa.PublicKey {
	negation := new(ecdsa.PublicKey)
	negation.Curve = x.Curve

	negation.X = new(big.Int).Set(x.X)
	negation.Y = new(big.Int).Set(x.Y)
	// y' = -y mod p.
	negation.Y = negation.Y.Neg(negation.Y)
	negation.Y = negation.Y.Mod(negation.Y, x.Curve.Params().P)
	//negation.Y = negation.Y.Mod(negation.Y, negation.Curve.Params().N)
	if !negation.Curve.IsOnCurve(negation.X, negation.Y) {
		panic("negation of x, y is not on curve")
	}

	return negation
}
