package proof

import (
	"math/big"
	"testing"

	"github.com/pgc/client"
	"github.com/pgc/contracts"
	"github.com/pgc/deployer"
	"github.com/pgc/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRangeProofContractVerify(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH()
	paramsAddr, _ := deployer.DeployParams(auth, ethclient)
	innerAddr, _ := deployer.DeployInnerProduct(auth, ethclient, paramsAddr)
	_, rangeCon := deployer.DeployRangeproofverifier(auth, ethclient, paramsAddr, innerAddr)

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

func testRangeProofContractVerify(t *testing.T, v *big.Int, expect bool, proofCon *contracts.Rangeproofverifier) {
	params := DRangeProofParams()
	p, r := newRandomCommitmentsRangeProof(params, v)

	proof, err := GenerateRangeProof(params, v, r)
	require.Nil(t, err, "generate range proof failed")

	input := proof.ToSolidityInput()
	input.points[8] = p.X
	input.points[9] = p.Y

	actual, err := proofCon.OptimizedVerifyRangeProof(utils.CallOpt(), input.points, input.scalar, input.l, input.r)
	require.Nil(t, err, "call agg range verifier failed", err)
	assert.Equal(t, expect, actual)

	actual, err = proofCon.VerifyRangeProof(utils.CallOpt(), input.points, input.scalar, input.ll, input.rr)
	require.Nil(t, err, "call agg range verifier failed", err)
	assert.Equal(t, expect, actual)
}
