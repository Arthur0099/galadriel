package pgc

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
)

const (
	// This is copied from https://github.com/golang/crypto/blob/master/openpgp/elgamal/elgamal_test.go
	// todo: The rationality should be discussed before usage.

	// pHex is global public parameter used in twistedELG.
	pHex = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371"
	// gHex is global public parameter used in twistedELG.
	gHex = "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"
	// H is global public parameter used in twistedELG as another random generator.
	// Just change last two hex value to "00".
	hHex = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4300"
	// Just for test.

)

// GroupGen generates/returns global public parameters used in twistedELG.
func GroupGen() (*big.Int, *big.Int, *big.Int) {
	p, ok := new(big.Int).SetString(pHex, 16)
	if !ok {
		panic("failed to parse hex number p")
	}

	g, ok := new(big.Int).SetString(gHex, 16)
	if !ok {
		panic("failed to parse hex number g")
	}

	h, ok := new(big.Int).SetString(hHex, 16)
	if !ok {
		panic("failed to parse hex number h")
	}

	if h.Cmp(p) >= 0 {
		panic("random generator h can't bigger than p")
	}

	return p, g, h
}

// The original implementation is in https://github.com/golang/crypto/blob/master/openpgp/elgamal/elgamal.go.
// The implementation here will change the encryption and decryption to implement twisting in pgc.
// https://eprint.iacr.org/2019/319.pdf

// PublicKey represents an ElGamal public key.
// Y represents public value y := g ** x mod p, and x is the private key.
type PublicKey struct {
	G, P, Y, H *big.Int
}

// PrivateKey represents an ElGamal private key.
// X is the private key in (1, p-1)
type PrivateKey struct {
	PublicKey
	X *big.Int
}

// TwistedELG respresents key pair described in pgc.
type TwistedELG struct {
	PrivateKey
}

// GenTwistedELG generates a twistedELG key pair using random priv key.
func GenTwistedELG() (*TwistedELG, error) {
	p, g, h := GroupGen()
	// generate a random private key.
	randReader := rand.Reader

	r, err := rand.Int(randReader, p)
	if err != nil {
		return nil, err
	}

	twistedELGKey := &TwistedELG{}
	twistedELGKey.X = new(big.Int).Set(r)

	// compute public key.
	twistedELGKey.Y = new(big.Int).Exp(g, r, p)

	// set global public parameters.
	twistedELGKey.G = new(big.Int).Set(g)
	twistedELGKey.P = new(big.Int).Set(p)
	twistedELGKey.H = new(big.Int).Set(h)

	return twistedELGKey, nil
}

// Encrypt encrypts the given message to the given public key. The result is a
// pair of integers. Errors can result from reading random, or because msg is
// too large to be encrypted to the public key.
func Encrypt(random io.Reader, pub *PublicKey, msg []byte) (c1, c2 *big.Int, err error) {
	// todo: limit the length of msg.

	m := new(big.Int).SetBytes(msg)
	// msg can't bigger than p.
	if m.Cmp(pub.P) >= 0 {
		err = errors.New("message too long")
		return
	}

	r, err := rand.Int(random, pub.P)
	if err != nil {
		return
	}

	c1 = new(big.Int).Exp(pub.Y, r, pub.P)
	s1 := new(big.Int).Exp(pub.G, r, pub.P)
	s2 := new(big.Int).Exp(pub.H, m, pub.P)
	//
	fmt.Printf("encrypt h^m is %x\n", s2)
	c2 = s1.Mul(s1, s2)
	c2.Mod(c2, pub.P)

	return
}

// Decrypt takes two integers, resulting from an ElGamal encryption, and
// returns the plaintext of the message. An error can result only if the
// ciphertext is invalid. Users should keep in mind that this is a padding
// oracle and thus, if exposed to an adaptive chosen ciphertext attack, can
// be used to break the cryptosystem.  See ``Chosen Ciphertext Attacks
// Against Protocols Based on the RSA Encryption Standard PKCS #1'', Daniel
// Bleichenbacher, Advances in Cryptology (Crypto '98),
func Decrypt(priv *PrivateKey, c1, c2 *big.Int) (msg []byte, err error) {

	// Seems some problem in it.todo: check and solve.

	// compute inverse of sk.(sk-1)
	s := new(big.Int).Set(priv.X)
	s.ModInverse(s, priv.P)
	// compute x ^ sk-1
	s.Exp(c1, s, priv.P)
	// compute inverse of x ^ sk-1.
	s.ModInverse(s, priv.P)
	// compute y/x^sk-1=y * inverse(x^sk-1)
	s.Mul(s, c2)
	// compute h ^ m.
	hM := s.Mod(s, priv.P)

	fmt.Printf("h ^ m is %x\n", hM)
	// This is just a correct test for decrypt.
	oriM := new(big.Int).SetBytes([]byte("hello world"))
	newHM := new(big.Int).Exp(priv.H, oriM, priv.P)
	fmt.Printf("hm value check %x\n", newHM)

	// Loop to find m.
	s = findM(priv.H, priv.P, hM)

	return s.Bytes(), nil
}

// findM try to find m by h ^ m mod p == res.
func findM(h, p, res *big.Int) *big.Int {
	m := new(big.Int).SetInt64(1)
	one := new(big.Int).SetInt64(1)
	// loop to find m.
	for h.Exp(h, m, p).Cmp(res) != 0 {
		m.Add(m, one)
	}

	return m
}
