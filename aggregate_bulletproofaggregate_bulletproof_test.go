package pgc

import (
	"math/big"
	"os"
	"runtime/pprof"
	"testing"
)

func TestAggreateProof(t *testing.T) {
	curve := S256()
	n := curve.Params().N

	bitSize := 32
	aggreateSize := 2
	vb := NewRandomVectorBase(curve, bitSize, aggreateSize)
	g := vb.GetG()
	h := vb.GetH()

	vlist := make([]*big.Int, 0)
	rlist := make([]*big.Int, 0)
	commitList := make([]*ECPoint, 0)
	for i := 0; i < aggreateSize; i++ {
		v := MustGetRandomMsg(bitSize)
		// v := new(big.Int).SetUint64(2)
		r := MustGetRandom(n)
		commit := new(ECPoint).ScalarMult(g, v)
		commit.Add(commit, new(ECPoint).ScalarMult(h, r))

		vlist = append(vlist, v)
		rlist = append(rlist, r)
		commitList = append(commitList, commit)
	}

	// generate proof.
	proof, err := GenerateAggreateBulletProof(vb, vlist, rlist)
	if err != nil {
		t.Fatal(err)
	}

	// validity proof.
	if !VerifyAggreateBulletProof(vb, commitList, proof) {
		t.Fatal("failed")
	}
}

func BenchmarkGenerateAggreateProof(b *testing.B) {
	curve := NoCGOS256()
	n := curve.Params().N

	bitSize := 32
	aggreateSize := 2
	vb := NewRandomVectorBase(curve, bitSize, aggreateSize)
	g := vb.GetG()
	h := vb.GetH()

	vlist := make([]*big.Int, 0)
	rlist := make([]*big.Int, 0)
	commitList := make([]*ECPoint, 0)
	for i := 0; i < aggreateSize; i++ {
		v := MustGetRandomMsg(bitSize)
		// v := new(big.Int).SetUint64(2)
		r := MustGetRandom(n)
		commit := new(ECPoint).ScalarMult(g, v)
		commit.Add(commit, new(ECPoint).ScalarMult(h, r))

		vlist = append(vlist, v)
		rlist = append(rlist, r)
		commitList = append(commitList, commit)
	}
	fd, _ := os.OpenFile("cpu_validator.log", os.O_CREATE|os.O_RDWR, os.ModeSetuid)
	if err := pprof.StartCPUProfile(fd); err != nil {
		b.Fatalf("TestTransferProofValidator %s", err.Error())
	}
	defer func() {
		pprof.StopCPUProfile()
		fd.Close()
	}()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// generate proof.
		proof, err := GenerateAggreateBulletProof(vb, vlist, rlist)
		if err != nil {
			b.Fatal(err)
		}

		if !VerifyAggreateBulletProof(vb, commitList, proof) {
			b.Fatal("failed")
		}
	}
	b.StopTimer()
}
