package proof

import (
	"context"
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pgc/client"
	"github.com/pgc/contracts/ipverifier"
	"github.com/pgc/curve"
	"github.com/pgc/deployer"

	"github.com/stretchr/testify/require"
)

func TestInnerProductContractVerify(t *testing.T) {
	cv := curve.BN256()

	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH()
	_, ipcon := deployer.DeployInnerProduct(auth, ethclient)

	cases := []int{2, 4, 8, 16, 32, 64}
	for _, c := range cases {
		testInnerProductContractVerify(t, cv, c, ipcon)
	}
}

func testInnerProductContractVerify(t *testing.T, cv elliptic.Curve, bitsize int, ipcon *ipverifier.Ipverifier) {
	ipparams := newRandomParams(cv, bitsize)
	p, c, a, b := newRandomcommitments(ipparams, cv.Params().N, bitsize)

	proof, err := GenIPProof(ipparams, p, c, a, b)
	require.Nil(t, err, "generate inner product failed")

	input := proof.ToSolidityInput()

	pinput := [2]*big.Int{p.X, p.Y}
	uinput := [2]*big.Int{ipparams.U().X, ipparams.U().Y}

	opt := bind.CallOpts{}
	opt.From = common.Address{1}
	opt.Pending = false
	opt.Context = context.Background()
	valid, err := ipcon.OptimizedVerifyIPProof(&opt, ipparams.GV().ToSolidityInput(), ipparams.HV().ToSolidityInput(),
		pinput, uinput, c, input.l, input.r, input.a, input.b)

	require.Nil(t, err, "call optimized failed")
	require.True(t, valid, "invalid optimized proof")

	valid, err = ipcon.VerifyIPProof(&opt, ipparams.GV().ToSolidityInput(), ipparams.HV().ToSolidityInput(),
		pinput, uinput, c, input.l, input.r, input.a, input.b)
	require.Nil(t, err, "call failed")
	require.True(t, valid, "invalid proof")
}
