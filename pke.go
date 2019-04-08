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

	// P is global public parameter used in twistedELG.
	P = "B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371"
	// G is global public parameter used in twistedELG.
	G = "A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5"
	// H is global public parameter used in twistedELG as another random generator.
	// Just for test.
	H = "44"
)

// GroupGen generates/returns global public parameters used in twistedELG.
func GroupGen() (*big.Int, *big.Int, *big.Int) {
	p, ok := new(big.Int).SetString(P, 16)
	if !ok {
		panic("failed to parse hex number p")
	}

	g, ok := new(big.Int).SetString(G, 16)
	if !ok {
		panic("failed to parse hex number g")
	}

	h, ok := new(big.Int).SetString(H, 16)
	if !ok {
		panic("failed to parse hex number h")
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
// X is the private key in (1, q-1)
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
	b := make([]byte, p.BitLen()/8)
	_, err := io.ReadFull(randReader, b)
	if err != nil {
		return nil, err
	}

	k := new(big.Int).SetBytes(b)
	// todo: make sure k below a number N?

	twistedELGKey := &TwistedELG{}
	twistedELGKey.X = new(big.Int).Set(k)

	// compute public key.
	twistedELGKey.Y = new(big.Int).Exp(g, k, p)

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
	// todo: change?
	pLen := (pub.P.BitLen() + 7) / 8
	if len(msg) > pLen-11 {
		err = errors.New("elgamal: message too long")
		return
	}

	m := new(big.Int).SetBytes(msg)

	r, err := rand.Int(random, pub.P)
	if err != nil {
		return
	}

	c1 = new(big.Int).Exp(pub.Y, r, pub.P)
	s1 := new(big.Int).Exp(pub.G, r, pub.P)
	s2 := new(big.Int).Exp(pub.H, m, pub.P)
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

	s := new(big.Int).Set(priv.X)
	fmt.Printf("priv x is %x\n", s)
	s.ModInverse(s, priv.P)
	fmt.Printf("sk-1 is %x\n", s)
	s.Exp(c1, s, priv.P)
	fmt.Printf("c1 %x ^^ sk-1 is %x\n", c1, s)
	s.DivMod(c2, s, priv.P)
	fmt.Printf("c2 %x / (c1 ^^ sk-1) is %x\n", c2, s)

	// at this point, s = h ^^ m(m is the actual msg, h is the pub parmeter).
	// todo: recove m from h ^^ m.
	fmt.Printf("h ^^ m is %x\n", s)

	return s.Bytes(), nil
}

// nonZeroRandomBytes fills the given slice with non-zero random octets.
func nonZeroRandomBytes(s []byte, rand io.Reader) (err error) {
	_, err = io.ReadFull(rand, s)
	if err != nil {
		return
	}

	for i := 0; i < len(s); i++ {
		for s[i] == 0 {
			_, err = io.ReadFull(rand, s[i:i+1])
			if err != nil {
				return
			}
		}
	}

	return
}
