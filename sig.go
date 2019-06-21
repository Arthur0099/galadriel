package pgc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
)

// Sig is the signature.
type Sig struct {
	R, S *big.Int
}

// ToInputs .
func (s *Sig) ToInputs() [2]*big.Int {
	return [2]*big.Int{s.R, s.S}
}

// Sign sign a data with pgc key.
func Sign(sk *ecdsa.PrivateKey, data []byte) (*Sig, error) {
	n := sk.PublicKey.Params().N
	r, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	params := Params()
	A := new(ECPoint).ScalarMult(params.GetH(), r)

	sig := Sig{}
	sig.R = new(big.Int).Set(A.X)
	sig.R = sig.R.Mod(sig.R, n)
	r.ModInverse(r, n)
	sig.S = new(big.Int).Mul(sig.R, sk.D)
	dataBig := new(big.Int).SetBytes(data)
	dataBig.Mod(dataBig, n)
	sig.S.Add(sig.S, dataBig)
	sig.S.Mul(sig.S, r)
	sig.S.Mod(sig.S, n)

	return &sig, nil
}

// VerifySig checked sig is valid or not.
func VerifySig(pk *ecdsa.PublicKey, data []byte, sig *Sig) bool {
	n := pk.Params().N
	if sig.R.Cmp(n) >= 0 || sig.S.Cmp(n) >= 0 {
		return false
	}

	sig.S.ModInverse(sig.S, n)
	u1 := new(big.Int).Mul(sig.S, new(big.Int).SetBytes(data))
	u2 := new(big.Int).Mul(sig.R, sig.S)
	A := new(ECPoint).ScalarMult(params.GetH(), u1)
	tmp := new(ECPoint).ScalarMult(new(ECPoint).SetFromPublicKey(pk), u2)
	A.Add(A, tmp)

	res := A.X.Mod(A.X, n)
	if res.Cmp(sig.R) != 0 {
		return false
	}
	return true
}
