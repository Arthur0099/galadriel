package pgc

import (
	"testing"
)

func TestSigmaProtocol(t *testing.T) {
	sys := NewTwistedELGamalSystem()
	// generate key pair for alice and bob.
	aliceKey, err := sys.GenerateKey()
	if err != nil {
		t.Error("generate key pair for alice failed", err)
		return
	}
	bobKey, err := sys.GenerateKey()
	if err != nil {
		t.Error("generate key pair for bob failed", err)
		return
	}

	tests := []struct {
		MsgA   []byte
		MsgB   []byte
		Expect bool
		Flag   bool
	}{
		{
			MsgA:   []byte("hh"),
			MsgB:   []byte("hh"),
			Expect: true,
			Flag:   false,
		},
		{
			MsgA:   []byte("a"),
			MsgB:   []byte("b"),
			Expect: false,
			Flag:   true,
		},
	}

	for i, test := range tests {
		// msg to be encoded.
		ct1, err := sys.Encrypt(&aliceKey.PublicKey, test.MsgA)
		if err != nil {
			t.Error(i, "encrypt ct1 for alice failed", err)
			continue
		}
		ct2, err := sys.Encrypt(&bobKey.PublicKey, test.MsgB)
		if err != nil {
			t.Error(i, "encrypt ct2 for bob failed", err)
			continue
		}

		// generate proof.
		sigmaSys := NewSigmaSys()
		sigmaSys.TestFlag = test.Flag
		proof, err := sigmaSys.GenerateProof(&aliceKey.PublicKey, &bobKey.PublicKey, ct1, ct2)
		if err != nil {
			t.Error(i, "generate proof for sigma protocol failed", err)
			continue
		}

		if sigmaSys.VerifySigmaProof(proof) != test.Expect {
			t.Error(i, "sigma proof verify except %v, actual %c", test.Expect, !test.Expect)
		}
	}

}

func TestDLESigmaProof(t *testing.T) {
	sys := NewTwistedELGamalSystem()
	key, err := sys.GenerateKey()
	if err != nil {
		t.Error("key generation failed")
		return
	}

	tests := []struct {
		MsgA   []byte
		MsgB   []byte
		Expect bool
	}{
		{
			MsgA:   []byte("hh"),
			MsgB:   []byte("hh"),
			Expect: true,
		},
		{
			MsgA:   []byte("aa"),
			MsgB:   []byte("bb"),
			Expect: false,
		},
	}

	for i, test := range tests {
		ct1, err := sys.Encrypt(&key.PublicKey, test.MsgA)
		if err != nil {
			t.Error("encrypt failed")
			return
		}
		ct2, err := sys.Encrypt(&key.PublicKey, test.MsgB)
		if err != nil {
			t.Error("encrypt failed")
			return
		}
		if ct1.X.X.Cmp(ct2.X.X) == 0 {
			t.Error("random number is the same in two randmon")
			return
		}

		prover := DLESigmaProver{}
		prover.params = Params()
		proof, err := prover.GenerateProof(ct1, ct2, key)
		if err != nil {
			t.Error("generate proof failed")
			return
		}
		if test.Expect {
			proof.Persist("solidity/dlesigmaproof")
		}

		if VerifyDLESigmaProof(proof) != test.Expect {
			t.Error(i, "dlesigma proof verify except %v, actual %c", test.Expect, !test.Expect)
		}
	}

}
