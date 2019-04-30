package pgc

import (
	"testing"
)

func TestSigmaProtocol(t *testing.T) {
	// generate key pair for alice and bob.
	aliceKey, err := GenerateKey()
	if err != nil {
		t.Error("generate key pair for alice failed", err)
		return
	}
	bobKey, err := GenerateKey()
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
			MsgA:   []byte("test"),
			MsgB:   []byte("test"),
			Expect: true,
			Flag:   false,
		},
		{
			MsgA:   []byte("first"),
			MsgB:   []byte("second"),
			Expect: false,
			Flag:   true,
		},
	}

	for i, test := range tests {
		// msg to be encoded.
		ct1, err := Encrypt(&aliceKey.PublicKey, test.MsgA)
		if err != nil {
			t.Error(i, "encrypt ct1 for alice failed", err)
			continue
		}
		ct2, err := Encrypt(&bobKey.PublicKey, test.MsgB)
		if err != nil {
			t.Error(i, "encrypt ct2 for bob failed", err)
			continue
		}

		// generate proof.
		prove := &SigmaProver{}
		prove.TestFlag = test.Flag
		proof, err := prove.GenerateProof(&aliceKey.PublicKey, &bobKey.PublicKey, ct1, ct2)
		if err != nil {
			t.Error(i, "generate proof for sigma protocol failed", err)
			continue
		}

		if VerifySigmaProof(proof) != test.Expect {
			t.Error(i, "sigma proof verify except %v, actual %c", test.Expect, !test.Expect)
		}
	}

}

func TestDLESigmaProof(t *testing.T) {
	key, err := GenerateKey()
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
			MsgA:   []byte("dlesigmaproof"),
			MsgB:   []byte("dlesigmaproof"),
			Expect: true,
		},
		{
			MsgA:   []byte("sss"),
			MsgB:   []byte("ccc"),
			Expect: false,
		},
	}

	for i, test := range tests {
		ct1, err := Encrypt(&key.PublicKey, test.MsgA)
		if err != nil {
			t.Error("encrypt failed")
			return
		}
		ct2, err := Encrypt(&key.PublicKey, test.MsgB)
		if err != nil {
			t.Error("encrypt failed")
			return
		}
		if ct1.X.X.Cmp(ct2.X.X) == 0 {
			t.Error("random number is the same in two randmon")
			return
		}

		prover := DLESigmaProver{}
		proof, err := prover.GenerateProof(ct1, ct2, key)
		if err != nil {
			t.Error("generate proof failed")
			return
		}

		if VerifyDLESigmaProof(proof) != test.Expect {
			t.Error(i, "dlesigma proof verify except %v, actual %c", test.Expect, !test.Expect)
		}
	}

}
