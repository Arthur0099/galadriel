package pgc

import (
	"math/big"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/inconshreveable/log15"
)

type tmptest struct {
	l, r []*big.Int
}

var tmprecord tmptest

func TestAggEqualNoAgg(t *testing.T) {
	vb := Params().vb
	aggreateSize := Params().aggreateSize
	bitSize := Params().bitSize
	n := Params().Curve().Params().N
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
		t.Log("failed")
	}

	// auth := getdaccount()
	// client := GetLocal()
	// if auth.Nonce == nil {
	// 	nonce, err := client.NonceAt(context.Background(), auth.From, nil)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	auth.Nonce = new(big.Int).SetUint64(nonce)
	// }
	// params, _, _ := MustDeployContract("publicParams", auth, client)
	// ipVerifier, _, _ := MustDeployContract("ipVerifier", auth, client, params)
	// rangeVerifier, _, _ := MustDeployContract("rangeProofVerifier", auth, client, params, ipVerifier)
	// con, _ := contracts.NewRangeproofverifier(rangeVerifier, client)

	// testWithSolidity(auth, commitList[0], proof, con)
}

func getdaccount() *bind.TransactOpts {
	senderKey := "9e1c3a5e84983cfa9e00e2818ecb610158430cedd64b30f04b49fc4a9ff326e3"
	key, err := crypto.HexToECDSA(senderKey)
	if err != nil {
		panic(err)
	}
	sender := bind.NewKeyedTransactor(key)
	sender.GasLimit = testGasLimit
	sender.GasPrice = new(big.Int).Mul(new(big.Int).SetUint64(10), GWEI)

	return sender
}

// func testWithSolidity(auth *bind.TransactOpts, commit *ECPoint, proof *AggreateBulletProof, solidity *contracts.Rangeproofverifier) {
// 	input := proof.ToSolidityInput()
// 	input.points[8] = commit.X
// 	input.points[9] = commit.Y

// 	tx, err := solidity.VerifyRangeProof(auth, input.points, input.scalar, input.l, input.r)
// 	if err != nil {
// 		panic(err)
// 	}
// 	log.Info("verify tx", "hash", tx.Hash().Hex())

// 	filterOpt := &bind.FilterOpts{}
// 	filterOpt.Start = 0
// 	filterOpt.Context = context.Background()
// 	iter, err := solidity.FilterLogParam(filterOpt)
// 	if err != nil {
// 		return
// 	}

// 	for iter.Next() {
// 		e := iter.Event
// 		log.Info("log param", "var", e.Param, "x", e.X, "y", e.Y)
// 	}
// }

func TestAggreateProof(t *testing.T) {
	curve := BN256()
	n := curve.Params().N

	first := new(big.Int).SetUint64(1)
	firstInv := new(big.Int).ModInverse(first, n)
	log.Info("1 inv", "inv", firstInv)

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

	if !optimizedAggVerify(vb, commitList, proof) {
		t.Fatal("optimized agg verify failed")
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
