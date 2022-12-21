package proof

import (
	"math/big"
	"testing"

	log "github.com/inconshreveable/log15"
	"github.com/pgc/client"
	"github.com/pgc/contracts/rangeproofverifier"
	"github.com/pgc/deployer"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRangeProofContractVerify(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH(ethclient)
	innerAddr, _ := deployer.DeployInnerProduct(auth, ethclient)
	_, rangeCon := deployer.DeployRangeproofverifier(auth, ethclient, innerAddr)

	cases := []struct {
		v      *big.Int
		expect bool
	}{
		{
			big.NewInt(1),
			true,
		},
		{
			big.NewInt(0),
			true,
		},
		{
			big.NewInt(-1),
			false,
		},
	}

	for _, c := range cases {
		testRangeProofContractVerify(t, c.v, c.expect, rangeCon)
	}
}

func testRangeProofContractVerify(t *testing.T, v *big.Int, expect bool, proofCon *rangeproofverifier.Rangeproofverifier) {
	params := DRangeProofParamsWithBitsize(64)
	nparams := toSolidityParams(params)
	p, r := newRandomCommitmentsRangeProof(params, v)
	commit := toBn128(p)

	proof, err := GenerateRangeProof(params, v, r)
	require.Nil(t, err, "generate range proof failed")
	nproof := toSolidityProof(proof)

	input := proof.ToSolidityInput()
	input.points[8] = p.X
	input.points[9] = p.Y

	actual, err := proofCon.OptimizedVerify(utils.CallOpt(), nparams, commit, nproof)
	require.Nil(t, err, "call optimized range verifier failed", err)
	assert.Equal(t, expect, actual)
	log.Info("optimized range verifier success")

	actual, err = proofCon.Verify(utils.CallOpt(), nparams, commit, nproof)
	require.Nil(t, err, "call normal range verifier failed", err)
	assert.Equal(t, expect, actual)
	log.Info("normal range verifier success")
}

func toSolidityProof(proof *RangeProof) rangeproofverifier.RangeProofVerifierRangeProof {
	nproof := rangeproofverifier.RangeProofVerifierRangeProof{}
	nproof.A = toBn128(proof.A)
	nproof.S = toBn128(proof.S)
	nproof.T1 = toBn128(proof.T1)
	nproof.T2 = toBn128(proof.T2)

	nproof.T = new(big.Int).Set(proof.t)
	nproof.Txx = new(big.Int).Set(proof.tx)
	nproof.U = new(big.Int).Set(proof.u)

	nproof.IpProof = rangeproofverifier.RangeProofVerifierIPProof{}
	nproof.IpProof.A = new(big.Int).Set(proof.ipProof.a)
	nproof.IpProof.B = new(big.Int).Set(proof.ipProof.b)
	nproof.IpProof.L = toBn128Array(proof.ipProof.l)
	nproof.IpProof.R = toBn128Array(proof.ipProof.r)

	return nproof
}

func toSolidityParams(params RangeParams) rangeproofverifier.RangeProofVerifierParams {
	np := rangeproofverifier.RangeProofVerifierParams{}
	np.G = toBn128(params.G())
	np.H = toBn128(params.H())
	np.U = toBn128(params.U())
	np.Gv = toBn128Array(params.GV().GetVector())
	np.Hv = toBn128Array(params.HV().GetVector())

	return np
}

func toBn128Array(ori []*utils.ECPoint) []rangeproofverifier.BN128G1Point {
	rs := make([]rangeproofverifier.BN128G1Point, 0)
	for _, o := range ori {
		rs = append(rs, toBn128(o))
	}

	return rs
}

func toBn128(ori *utils.ECPoint) rangeproofverifier.BN128G1Point {
	p := rangeproofverifier.BN128G1Point{}
	p.X = new(big.Int).Set(ori.X)
	p.Y = new(big.Int).Set(ori.Y)

	return p
}
