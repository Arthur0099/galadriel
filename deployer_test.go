package pgc

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

func TestDeployer(t *testing.T) {
	client := GetLocal()
	rpcClient := GetLocalRPC()

	senderKey, _ := crypto.GenerateKey()
	sender := bind.NewKeyedTransactor(senderKey)
	sender.GasLimit = testGasLimit
	sender.GasPrice = new(big.Int).Mul(new(big.Int).SetUint64(10), GWEI)

	amount := new(big.Int).SetUint64(10000)
	_, err := rpcClient.SendFromLocalAccout(sender.From, amount)
	require.Nil(t, err, "send eth failed")

	fakeAddress := common.Address{2}
	params, _, _ := MustDeployContract("publicParams", sender, client)
	MustDeployContract("sigmaVerifier", sender, client, params)
	MustDeployContract("dleSigmaVerifier", sender, client)
	MustDeployContract("ipVerifier", sender, client, params)
	MustDeployContract("rangeProofVerifier", sender, client, params, fakeAddress)
	MustDeployContract("tokenConverter", sender, client)
	MustDeployContract("pgcVerifier", sender, client, params, fakeAddress, fakeAddress, fakeAddress)
	MustDeployContract("pgc", sender, client, params, fakeAddress, fakeAddress)
}
