package pgc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// PTEqualityProof .
type PTEqualityProof struct {
	A1, A2, B *ECPoint

	Z1, Z2 *big.Int
}

// GeneratePTEqualityProof generates proof to prove msg in two plaintext is the same.
func GeneratePTEqualityProof(pk1, pk2 *ecdsa.PublicKey, ct *MRTwistedELGamalCT) (*PTEqualityProof, error) {
	curve := ct.X1.Curve
	n := curve.Params().N
	params := Params()

	// get random.
	a, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	b, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	proof := PTEqualityProof{}
	// A1=pk1^a; A2=pk2^a.
	proof.A1 = new(ECPoint).SetFromPublicKey(pk1)
	proof.A1.ScalarMult(proof.A1, a)
	proof.A2 = new(ECPoint).SetFromPublicKey(pk2)
	proof.A2.ScalarMult(proof.A2, a)

	e, err := ComputeChallengeByECPoints(curve.Params().N, proof.A1, proof.A2)
	if err != nil {
		return nil, err
	}

	r := new(big.Int).Set(ct.R)
	v := new(big.Int).SetBytes(ct.EncMsg)
	// Z1 = a + e * r.
	proof.Z1 = new(big.Int).Mul(e, r)
	proof.Z1.Mod(proof.Z1, n)
	proof.Z1.Add(proof.Z1, a)
	proof.Z1.Mod(proof.Z1, n)

	// Z2 = b + e * v.
	proof.Z2 = new(big.Int).Mul(e, v)
	proof.Z2.Mod(proof.Z2, n)
	proof.Z2.Add(proof.Z2, b)
	proof.Z2.Mod(proof.Z2, n)

	// B = g*b + h*a.
	g := params.GetG()
	h := params.GetH()
	proof.B = new(ECPoint).ScalarMult(g, b)
	proof.B.Add(proof.B, new(ECPoint).ScalarMult(h, a))

	return &proof, nil
}

// VerifyPTEqualityProof verify proof.
func VerifyPTEqualityProof(pk1, pk2 *ecdsa.PublicKey, ct *MRTwistedELGamalCTPub, proof *PTEqualityProof) bool {
	curve := pk1.Curve
	n := curve.Params().N
	// compute challenge.
	e, err := ComputeChallengeByECPoints(n, proof.A1, proof.A2)
	if err != nil {
		log.Warn("verify pteequality proof failed(compute challenge)", "err", err)
		return false
	}
	// check pk1 * z1 == A1 + X1 * e.
	if !checkTwistedELGamalCTX(new(ECPoint).SetFromPublicKey(pk1), proof.A1, ct.X1, proof.Z1, e) {
		return false
	}
	// check pk2 * z1 == A2 + X2 * e.
	if !checkTwistedELGamalCTX(new(ECPoint).SetFromPublicKey(pk2), proof.A2, ct.X2, proof.Z1, e) {
		return false
	}

	params := Params()
	h := params.GetH()
	g := params.GetG()
	// Check h * z1 + g * z2 == B + Y * e.
	if !checkTwistedELGamalCTY(g, h, proof.B, ct.Y, proof.Z1, proof.Z2, e) {
		return false
	}

	return true
}

// checkTwistedELGamalCTX checks pk * z == A + X * e or not(check X).
func checkTwistedELGamalCTX(pk, A, X *ECPoint, z, e *big.Int) bool {

	// compute x * e + A
	actual := new(ECPoint).ScalarMult(X, e)
	actual.Add(actual, A)
	// compute pk * z.
	expect := new(ECPoint).ScalarMult(pk, z)

	if !expect.Equal(actual) {
		log.Warn("pk * z != A + X * e", "expect x", expect.X, "actual x", actual.X, "expect y", expect.Y, "actual y", actual.Y)
		return false
	}

	return true
}

// checkTwistedELGamalCTY checks h*z1 + g*z2 == B + Y*e(check Y).
func checkTwistedELGamalCTY(g, h, B, Y *ECPoint, z1, z2, e *big.Int) bool {
	// compute g * z2 + h * z1.
	gzb := new(ECPoint).ScalarMult(g, z2)
	actual := new(ECPoint).ScalarMult(h, z1)
	actual.Add(actual, gzb)
	// compute B + Y * e.
	expect := new(ECPoint).ScalarMult(Y, e)
	expect.Add(expect, B)
	if !expect.Equal(actual) {
		log.Warn("g * z2 + h * z1 != B + Y * e", "expect x", expect.X, "actual x", actual.X, "expect y", expect.Y, "actual y", actual.Y)
		return false
	}

	return true
}
