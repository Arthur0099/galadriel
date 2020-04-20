package proof

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/pgc/utils"
)

// SigmaProof includes items to prove two encryption of same message under two different public key.
type SigmaProof struct {
	//
	z1, z2, z3     *big.Int
	Pk1, Pk2       *utils.ECPoint
	A1, A2, B1, B2 *utils.ECPoint
	// encrypt ct point.
	ct1, ct2 *CTEncPoint
}

// GenerateSigmaProof generates proof to prove that two encryptions(ct1, ct2) of the same
// message under pk1 and pk2.
func GenerateSigmaProof(params BaseParams, pk1, pk2 *ecdsa.PublicKey, ct1, ct2 *TwistedELGamalCT) (proof *SigmaProof, err error) {
	curve := pk1.Curve
	n := curve.Params().N

	// pick random number.
	a1, err := rand.Int(rand.Reader, n)
	if err != nil {
		return
	}
	a2, err := rand.Int(rand.Reader, n)
	if err != nil {
		return
	}
	b, err := rand.Int(rand.Reader, n)
	if err != nil {
		return
	}

	// compute proof.
	h := params.H()
	g := params.G()

	// A2 = pk1 * a1; A2 = pk2 * a2.
	A1 := new(utils.ECPoint).SetFromPublicKey(pk1)
	A1.ScalarMult(A1, a1)
	A2 := new(utils.ECPoint).SetFromPublicKey(pk2)
	A2.ScalarMult(A2, a2)

	// compute B1 = g * b + h * a1.
	ha1 := new(utils.ECPoint).ScalarMult(h, a1)
	gb := new(utils.ECPoint).ScalarMult(g, b)
	B1 := new(utils.ECPoint).Add(gb, ha1)
	// compute B2 = g * b + h * a2.
	ha2 := new(utils.ECPoint).ScalarMult(h, a2)
	B2 := new(utils.ECPoint).Add(gb, ha2)

	// compute challenge e.
	e, err := utils.ComputeChallengeByECPoints(n, A1, A2, B1, B2)
	if err != nil {
		return
	}

	// compute proof.
	// z1 = a1 + e*r1.
	z1 := new(big.Int).Mul(e, ct1.R)
	z1 = z1.Mod(z1, n)
	z1 = z1.Add(z1, a1)
	z1 = z1.Mod(z1, n)
	// z2 = a2 + e*r2.
	z2 := new(big.Int).Mul(e, ct2.R)
	z2 = z2.Mod(z2, n)
	z2 = z2.Add(z2, a2)
	z2 = z2.Mod(z2, n)
	// z3 = b + e*m.
	z3 := new(big.Int).Mul(e, new(big.Int).SetBytes(ct1.EncMsg))
	z3 = z3.Mod(z3, n)
	z3 = z3.Add(z3, b)
	z3 = z3.Mod(z3, n)

	// set proof.
	proof = new(SigmaProof)
	proof.ct1 = ct1.CopyPublicPoint()
	proof.ct2 = ct2.CopyPublicPoint()
	// warning: change to proof.pk1 will alse change pk1.
	proof.Pk1 = new(utils.ECPoint).SetFromPublicKey(pk1)
	proof.Pk2 = new(utils.ECPoint).SetFromPublicKey(pk2)
	proof.z1 = z1
	proof.z2 = z2
	proof.z3 = z3
	proof.A1 = A1
	proof.A2 = A2
	proof.B1 = B1
	proof.B2 = B2

	return
}

// VerifySigmaProof validates proof.
func VerifySigmaProof(params BaseParams, proof *SigmaProof) bool {
	n := params.H().Curve.Params().N
	e, err := utils.ComputeChallengeByECPoints(n, proof.A1, proof.A2, proof.B1, proof.B2)
	if err != nil {
		return false
	}
	// check pk1 * z1 == A1 + X1 * e.
	if !checkSigmaStep1(proof.Pk1, proof.A1, proof.ct1.X, proof.z1, e) {
		return false
	}

	// check pk2 * z2 == A2 + X2 * e.
	if !checkSigmaStep1(proof.Pk2, proof.A2, proof.ct2.X, proof.z2, e) {
		return false
	}

	h := params.H()
	g := params.G()
	// Check h * z1 + g * z3 == B1 + Y1 * e.
	if !checkSigmaStep2(g, h, proof.B1, proof.ct1.Y, proof.z1, proof.z3, e) {
		return false
	}

	// check h * z2 + g * z3 == B2 + Y2 * e.
	if !checkSigmaStep2(g, h, proof.B2, proof.ct2.Y, proof.z2, proof.z3, e) {
		return false
	}
	return true
}

// checkSigmaStep1 checks pk * z == A + X * e or not.
func checkSigmaStep1(pk, A, X *utils.ECPoint, z, e *big.Int) bool {

	// compute x * e + A
	actual := new(utils.ECPoint).ScalarMult(X, e)
	actual.Add(actual, A)
	// compute pk * z.
	expect := new(utils.ECPoint).ScalarMult(pk, z)

	if !expect.Equal(actual) {
		log.Warn("pk * z != A + X * e", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	return true
}

// checkSigmaStep2 checks h*za + g*zb == B + Y*e.
func checkSigmaStep2(g, h, B, Y *utils.ECPoint, za, zb, e *big.Int) bool {
	// compute g * za + h * zb.
	gzb := new(utils.ECPoint).ScalarMult(g, zb)
	hza := new(utils.ECPoint).ScalarMult(h, za)
	actual := new(utils.ECPoint).Add(gzb, hza)
	// compute B + Y * e.
	ye := new(utils.ECPoint).ScalarMult(Y, e)
	expect := new(utils.ECPoint).Add(ye, B)
	if !expect.Equal(actual) {
		log.Warn("g * z + h * z' != B + Y * e", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	return true
}

// DLESigmaProof includes items of zero-knowledge proof.
type DLESigmaProof struct {
	A1 *utils.ECPoint
	A2 *utils.ECPoint

	Z *big.Int
}

// GenerateDLESigmaProof generates zero knowledge proof to prove two ciphertexts encrypt the same value under same public key.
func GenerateDLESigmaProof(params KeyParams, ori, refresh *CTEncPoint, sk *ecdsa.PrivateKey, custom ...*big.Int) (*DLESigmaProof, error) {
	// g1 = Y(fresh) - Y(ori)
	g1 := new(utils.ECPoint).Sub(refresh.Y, ori.Y)
	// h1 = X(fresh) - X(ori)
	h1 := new(utils.ECPoint).Sub(refresh.X, ori.X)
	// g2 = H base point.
	g2 := params.H()
	// h2 = pk.
	h2 := new(utils.ECPoint).SetFromPublicKey(&sk.PublicKey)
	// witness = sk.
	w := new(big.Int).Set(sk.D)
	return generateDLESimaProof(g1, h1, g2, h2, w, custom...)
}

// GenerateDLESigmaProofWithNonce generates zero knowledge proof to prove two ciphertexts encrypt the same value under same public key.
// nonce is a praram to calculate hash.

// GenerateEqualProof generates a proof to prove amount is same with value in encrypted ct.
func GenerateEqualProof(params BaseParams, amount *big.Int, ct *CTEncPoint, sk *ecdsa.PrivateKey, custom ...*big.Int) (*DLESigmaProof, error) {
	g1 := new(utils.ECPoint).Sub(ct.Y, new(utils.ECPoint).ScalarMult(params.G(), amount))
	h1 := ct.X
	g2 := params.H()
	h2 := new(utils.ECPoint).SetFromPublicKey(&sk.PublicKey)
	w := new(big.Int).Set(sk.D)

	return generateDLESimaProof(g1, h1, g2, h2, w, custom...)
}

// VerifyEqualProof verify equal proof.
func VerifyEqualProof(params BaseParams, ct *CTEncPoint, amount *big.Int, pk *utils.ECPoint, proof *DLESigmaProof, custom ...*big.Int) bool {
	//
	g1 := new(utils.ECPoint).Sub(ct.Y, new(utils.ECPoint).ScalarMult(params.G(), amount))
	h1 := ct.X
	g2 := params.H()
	h2 := pk.Copy()
	return verifyDLESigmaProof(g1, h1, g2, h2, proof, custom...)
}

// generateDLESigmaProof generates items to prove g1 ^ w == h1; g2 ^ w == h2; w==w.
func generateDLESimaProof(g1, h1, g2, h2 *utils.ECPoint, w *big.Int, custom ...*big.Int) (*DLESigmaProof, error) {
	//
	curve := g1.Curve
	n := curve.Params().N
	a, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	// A1 = g1 * a; A2 = g2 * a.
	A1 := new(utils.ECPoint).ScalarMult(g1, a)
	A2 := new(utils.ECPoint).ScalarMult(g2, a)
	// compute challenge e prime.
	eprime, _ := utils.ComputeChallenge(n, A1.X, A1.Y, A2.X, A2.Y)
	cinput := make([]interface{}, 0)
	for _, c := range custom {
		cinput = append(cinput, c)
	}
	// compute custom input hash.
	hcustom, _ := utils.ComputeChallenge(n, cinput...)
	// compute final challenge.
	e, err := utils.ComputeChallenge(n, eprime, hcustom)
	if err != nil {
		return nil, err
	}

	// compute z = a + e * w.
	z := new(big.Int).Mul(e, w)
	z = z.Mod(z, n)
	z = z.Add(z, a)
	z = z.Mod(z, n)

	// set proof
	proof := new(DLESigmaProof)
	proof.A1 = A1
	proof.A2 = A2
	proof.Z = z

	return proof, nil
}

// VerifyDLESigmaProof verify the proof is valid or not.
func VerifyDLESigmaProof(params KeyParams, ori, refresh *CTEncPoint, pk *ecdsa.PublicKey, proof *DLESigmaProof, custom ...*big.Int) bool {
	// g1 = Y(fresh) - Y(ori)
	g1 := new(utils.ECPoint).Sub(refresh.Y, ori.Y)
	// h1 = X(fresh) - X(ori)
	h1 := new(utils.ECPoint).Sub(refresh.X, ori.X)
	// g2 = h base point.
	g2 := params.H()
	// h2 = pk.
	h2 := new(utils.ECPoint).SetFromPublicKey(pk)

	return verifyDLESigmaProof(g1, h1, g2, h2, proof, custom...)
}

func verifyDLESigmaProof(g1, h1, g2, h2 *utils.ECPoint, proof *DLESigmaProof, custom ...*big.Int) bool {
	curve := proof.A1.Curve
	n := curve.Params().N

	data := []interface{}{proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y}
	for _, c := range custom {
		data = append(data, c)
	}
	// compute e prime challenge
	eprime, _ := utils.ComputeChallenge(n, proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y)
	cinput := make([]interface{}, 0)
	for _, c := range custom {
		cinput = append(cinput, c)
	}
	// compute custom input hash
	hcustom, _ := utils.ComputeChallenge(n, cinput...)
	// compute final challenge.
	e, err := utils.ComputeChallenge(n, eprime, hcustom)
	if err != nil {
		return false
	}

	// check g1 * z == A1 + h1 * e.
	if !checkDLESigmaProof(g1, proof.A1, h1, proof.Z, e) {
		return false
	}
	// check g2 * z == A2 + h2 * e.
	if !checkDLESigmaProof(g2, proof.A2, h2, proof.Z, e) {
		return false
	}

	return true
}

// checkDLESigmaProof checks g * z == A + h * e.
func checkDLESigmaProof(g, A, H *utils.ECPoint, z, e *big.Int) bool {
	// g * z.
	gz := new(utils.ECPoint).ScalarMult(g, z)
	// h * e + A.
	he := new(utils.ECPoint).ScalarMult(H, e)
	expect := new(utils.ECPoint).Add(A, he)

	if !expect.Equal(gz) {
		log.Warn("g * z != A + h * e", "expect x", expect.X, "expect y", expect.Y, "actual x", gz.X, "actual y", gz.Y)
		return false
	}

	return true
}
