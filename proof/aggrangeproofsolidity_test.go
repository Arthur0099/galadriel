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

func TestAggRangeProofContractVerify(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH()
	paramsAddr, _ := deployer.DeployParams(auth, ethclient)
	_, aggCon := deployer.DeployAggRangeProof(auth, ethclient, paramsAddr)

	cases := []struct {
		v      []*big.Int
		expect bool
	}{
		{
			[]*big.Int{big.NewInt(1), big.NewInt(2)},
			true,
		},
		{
			[]*big.Int{big.NewInt(0), big.NewInt(0)},
			true,
		},
		{
			[]*big.Int{big.NewInt(1), big.NewInt(-1)},
			false,
		},
		{
			[]*big.Int{big.NewInt(-1), big.NewInt(1)},
			false,
		},
		{
			[]*big.Int{big.NewInt(-1), big.NewInt(-1)},
			false,
		},
	}

	for _, c := range cases {
		testAggRangeProofContractVerify(t, c.v, c.expect, aggCon)
	}

}

func testAggRangeProofContractVerify(t *testing.T, v []*big.Int, expect bool, aggCon *contracts.Aggrangeproofverifier) {
	params := DAggRangeProofParams()
	p, r := newRandomCommitmentsAggRangeProof(params, v)

	proof, err := GenerateAggRangeProof(params, v, r)
	require.Nil(t, err, "generate agg range proof failed")

	input := proof.ToSolidityInput()
	input.points[8] = p[0].X
	input.points[9] = p[0].Y
	input.points[10] = p[1].X
	input.points[11] = p[1].Y
	actual, err := aggCon.AggVerifyRangeProof(utils.CallOpt(), input.points, input.scalar, input.l, input.r)
	require.Nil(t, err, "call agg range verifier failed", err)
	assert.Equal(t, expect, actual)
}
