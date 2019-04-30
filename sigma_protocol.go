package pgc

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// SigmaProof includes items to prove two encryption of same message under two different public key.
type SigmaProof struct {
	//
	z1, z2, z3     *big.Int
	Pk1, Pk2       *ecdsa.PublicKey
	A1, A2, B1, B2 *ecdsa.PublicKey
	// encrypt ct point.
	ct1, ct2 *CTEncPoint
}

type SigmaProver struct {
	// TestFlag means generation is for a test purpose(if set true).
	TestFlag bool
}

// GenerateProof generates proof to prove that two encryptions(ct1, ct2) of the same
// message under pk1 and pk2.
func (prover *SigmaProver) GenerateProof(pk1, pk2 *ecdsa.PublicKey, ct1, ct2 *TwistedELGamalCT) (proof *SigmaProof, err error) {
	curve := pk1.Curve
	n := curve.Params().N

	// check msg encoded is same.
	// testFlag == true means generate a fake proof for test purpose.
	if !prover.TestFlag && !bytes.Equal(ct1.EncMsg, ct2.EncMsg) {
		err = errors.New("msg encoded in two ct isn't same")
		return
	}

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
	h := Params().ElgGenerator
	A1X, A1Y := curve.ScalarMult(pk1.X, pk1.Y, a1.Bytes())
	A1 := PointToPublicKey(A1X, A1Y, curve)
	A2X, A2Y := curve.ScalarMult(pk2.X, pk2.Y, a2.Bytes())
	A2 := PointToPublicKey(A2X, A2Y, curve)
	// compute g * a1 + h * b
	hbX, hbY := curve.ScalarMult(h.X, h.Y, b.Bytes())
	b1X, b1Y := curve.ScalarBaseMult(a1.Bytes())
	B1X, B1Y := curve.Add(hbX, hbY, b1X, b1Y)
	B1 := PointToPublicKey(B1X, B1Y, curve)
	// compute g * a2 + h * b
	b2X, b2Y := curve.ScalarBaseMult(a2.Bytes())
	B2X, B2Y := curve.Add(hbX, hbY, b2X, b2Y)
	B2 := PointToPublicKey(B2X, B2Y, curve)

	// compute challenge e.
	e, err := SigmaFiatShamir(A1, A2, B1, B2)
	if err != nil {
		return
	}
	// make sure e < n.
	if e.Cmp(n) >= 0 {
		err = errors.New("challenge e out of bound")
		return
	}

	// compute proof.
	// z1.
	z1 := new(big.Int).Mul(e, ct1.R)
	z1 = z1.Mod(z1, n)
	z1 = z1.Add(z1, a1)
	z1 = z1.Mod(z1, n)
	// z2.
	z2 := new(big.Int).Mul(e, ct2.R)
	z2 = z2.Mod(z2, n)
	z2 = z2.Add(z2, a2)
	z2 = z2.Mod(z2, n)
	// z3.
	z3 := new(big.Int).Mul(e, new(big.Int).SetBytes(ct1.EncMsg))
	z3 = z3.Mod(z3, n)
	z3 = z3.Add(z3, b)
	z3 = z3.Mod(z3, n)

	// set proof.
	proof = new(SigmaProof)
	proof.ct1 = ct1.CopyPublicPoint()
	proof.ct2 = ct2.CopyPublicPoint()
	// warning: change to proof.pk1 will alse change pk1.
	proof.Pk1 = pk1
	proof.Pk2 = pk2
	proof.z1 = z1
	proof.z2 = z2
	proof.z3 = z3
	proof.A1 = A1
	proof.A2 = A2
	proof.B1 = B1
	proof.B2 = B2

	return
}

// SigmaFiatShamir calculate challenge e.
// Don't change A1, A2, B1, B2
func SigmaFiatShamir(A1, A2, B1, B2 *ecdsa.PublicKey) (*big.Int, error) {
	// todo: same with Keccak256(a1, a2, b1, b2) in solidity.
	// use abi.Arguments.Pack(A1, A2, B1, B2)
	// hash(bytes)
	return sigmaFiatShamir(A1.X, A1.Y, A2.X, A2.Y, B1.X, B1.Y, B2.X, B2.Y)
}

// sigmaFiatShamir returns result of hash(pack(data))
// The type of data must be big.Int
// todo: check type of data is big.Int
func sigmaFiatShamir(data ...interface{}) (*big.Int, error) {
	uint256Type, _ := abi.NewType("uint256", nil)
	arguments := make(abi.Arguments, 0)
	for i := 0; i < len(data); i++ {
		argument := abi.Argument{}
		argument.Type = uint256Type

		arguments = append(arguments, argument)
	}

	packedData, err := arguments.Pack(data...)
	if err != nil {
		return nil, err
	}

	return new(big.Int).SetBytes(Keccak256(packedData)), nil
}

func VerifySigmaProof(proof *SigmaProof) bool {
	e, err := SigmaFiatShamir(proof.A1, proof.A2, proof.B1, proof.B2)
	if err != nil {
		return false
	}
	if e.Cmp(proof.Pk1.Curve.Params().N) >= 0 {
		log.Printf("challenge e is bigger than N(G point order)")
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

	h := Params().ElgGenerator
	// check g * z1 + h * z3 == B1 + Y1 * e.
	if !checkSigmaStep2(h, proof.B1, proof.ct1.Y, proof.z1, proof.z3, e) {
		return false
	}

	// check g * z2 + h * z3 == B2 + Y2 * e.
	if !checkSigmaStep2(h, proof.B2, proof.ct2.Y, proof.z2, proof.z3, e) {
		return false
	}
	return true
}

// checkSigmaStep1 checks pk * z == A + X * e or not.
func checkSigmaStep1(pk, A, X *ecdsa.PublicKey, z, e *big.Int) bool {
	curve := pk.Curve
	checkX, checkY := curve.ScalarMult(X.X, X.Y, e.Bytes())
	dstX, dstY := curve.Add(A.X, A.Y, checkX, checkY)
	srcX, srcY := curve.ScalarMult(pk.X, pk.Y, z.Bytes())
	if dstX.Cmp(srcX) != 0 || dstY.Cmp(srcY) != 0 {
		log.Printf("pk * z(x %x, y %x) != A + X * e(x %x, y %x)\n", srcX, srcY, dstX, dstY)
		return false
	}

	return true
}

// checkSigmaStep2 checks g * z + h * z' == B + Y * e.
func checkSigmaStep2(h, B, Y *ecdsa.PublicKey, za, zb, e *big.Int) bool {
	curve := h.Curve
	// compute g * za + h * zb.
	checkX, checkY := curve.ScalarBaseMult(za.Bytes())
	hX, hY := curve.ScalarMult(h.X, h.Y, zb.Bytes())
	checkX, checkY = curve.Add(checkX, checkY, hX, hY)
	// compute B + Y * e.
	dstX, dstY := curve.ScalarMult(Y.X, Y.Y, e.Bytes())
	dstX, dstY = curve.Add(B.X, B.Y, dstX, dstY)
	if dstX.Cmp(checkX) != 0 || dstY.Cmp(checkY) != 0 {
		log.Printf("g*z + h*z' (x %x, y %x) != B + Y * e(x %x, y %x)\n", checkX, checkY, dstX, dstY)
		return false
	}

	return true
}

// DLESigmaProver holds method to prove two ciphertexts encrypt the same value under same public key.
type DLESigmaProver struct {
	//
}

// DLESigmaProof includes items of zero-knowledge proof.
type DLESigmaProof struct {
	A1, A2 *ecdsa.PublicKey

	Z *big.Int

	g1, h1, g2, h2 *ecdsa.PublicKey
}

// GenerateProof generates zero knowledge proof to prove two ciphertexts encrypt the same value under same public key.
func (dleProver *DLESigmaProver) GenerateProof(ori, refresh *TwistedELGamalCT, sk *ecdsa.PrivateKey) (*DLESigmaProof, error) {
	// g1 = Y(fresh) - Y(ori)
	g1 := SubECPoint(refresh.Y, ori.Y)
	// h1 = X(fresh) - X(ori)
	h1 := SubECPoint(refresh.X, ori.X)
	// g2 = G base point.
	g2 := PointToPublicKey(sk.Curve.Params().Gx, sk.Curve.Params().Gy, sk.Curve)
	// h2 = pk.
	h2 := &sk.PublicKey
	// witness = sk.
	w := new(big.Int).Set(sk.D)
	return dleProver.generateProof(g1, h1, g2, h2, w)
}

// generateProof generates items to prove g1 ^ w == h1; g2 ^ w == h2; w==w.
func (dleProver *DLESigmaProver) generateProof(g1, h1, g2, h2 *ecdsa.PublicKey, w *big.Int) (*DLESigmaProof, error) {
	//
	curve := g1.Curve
	n := curve.Params().N
	a, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	A1X, A1Y := curve.ScalarMult(g1.X, g1.Y, a.Bytes())
	A2X, A2Y := curve.ScalarMult(g2.X, g2.Y, a.Bytes())
	// compute challenge e.
	e, err := sigmaFiatShamir(A1X, A1Y, A2X, A2Y)
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
	proof.A1 = PointToPublicKey(A1X, A1Y, curve)
	proof.A2 = PointToPublicKey(A2X, A2Y, curve)
	proof.Z = z
	proof.g1 = g1
	proof.h1 = h1
	proof.g2 = g2
	proof.h2 = h2

	return proof, nil
}

// VerifyDLESigmaProof verify the proof is valid or not.
func VerifyDLESigmaProof(proof *DLESigmaProof) bool {
	curve := proof.A1.Curve

	e, err := sigmaFiatShamir(proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y)
	if err != nil {
		return false
	}

	if e.Cmp(curve.Params().N) >= 0 {
		log.Printf("challenge e out of bound\n")
		return false
	}

	// check g1 * z == A1 + h1 * e.
	if !checkDLESigmaProof(proof.g1, proof.A1, proof.h1, proof.Z, e) {
		return false
	}
	// check g2 * z == A2 + h2 * e.
	if !checkDLESigmaProof(proof.g2, proof.A2, proof.h2, proof.Z, e) {
		return false
	}

	return true
}

// checkDLESigmaProof checks g * z == A + h * e.
func checkDLESigmaProof(g, A, H *ecdsa.PublicKey, z, e *big.Int) bool {
	curve := g.Curve
	checkX, checkY := curve.ScalarMult(g.X, g.Y, z.Bytes())
	dstX, dstY := curve.ScalarMult(H.X, H.Y, e.Bytes())
	dstX, dstY = curve.Add(A.X, A.Y, dstX, dstY)

	if dstX.Cmp(checkX) != 0 || dstY.Cmp(checkY) != 0 {
		log.Printf("g * z(x %x, y %x) != A + h * e(x %x, y %x)\n", checkX, checkY, dstX, dstY)
		return false
	}

	return true
}
