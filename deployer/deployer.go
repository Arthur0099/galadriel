package deployer

import (
	"github.com/pgc/client"
	"github.com/pgc/contracts"
	"github.com/pgc/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DeployParams deploy param contract.
func DeployParams(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *contracts.Publicparams) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployPublicparams(auth, ethclient)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployInnerProduct deploys inner product contract.
func DeployInnerProduct(auth *bind.TransactOpts, ethclient *ethclient.Client, params common.Address) (common.Address, *contracts.Ipverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployIpverifier(auth, ethclient, params)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployAggRangeProof deploys aggrate range proof verifier.
func DeployAggRangeProof(auth *bind.TransactOpts, ethclient *ethclient.Client, params common.Address) (common.Address, *contracts.Aggrangeproofverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployAggrangeproofverifier(auth, ethclient, params)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployRangeproofverifier deploys single range proof verifier.
func DeployRangeproofverifier(auth *bind.TransactOpts, ethclient *ethclient.Client, params, inner common.Address) (common.Address, *contracts.Rangeproofverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployRangeproofverifier(auth, ethclient, params, inner)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployTokenConverter deploys contract to convert token.
func DeployTokenConverter(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *contracts.Tokenconverter) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployTokenconverter(auth, ethclient)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployDleSigma deploys dle sigma contract.
func DeployDleSigma(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *contracts.Dlesigmaverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployDlesigmaverifier(auth, ethclient)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeploySigma deploys sigma contract.
func DeploySigma(auth *bind.TransactOpts, ethclient *ethclient.Client, params common.Address) (common.Address, *contracts.Sigmaverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeploySigmaverifier(auth, ethclient, params)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployPGCVerifier deploys pgc system verifier for verifing proofs.
func DeployPGCVerifier(auth *bind.TransactOpts, ethclient *ethclient.Client,
	params, dleSigma, rangeProof, aggRangeProof, sigma common.Address) (common.Address, *contracts.Pgcverifier) {

	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployPgcverifier(auth, ethclient, params, dleSigma, rangeProof, aggRangeProof, sigma)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployPGCMain deploys pgc system main contract.
func DeployPGCMain(auth *bind.TransactOpts, ethclient *ethclient.Client, params, pgcVerifier, tokenConverter common.Address) (common.Address, *contracts.Pgc) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := contracts.DeployPgc(auth, ethclient, params, pgcVerifier, tokenConverter)
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployPGCSystemAllContract deploys all contract for pgc system.
func DeployPGCSystemAllContract(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *contracts.Pgc) {
	params, _ := DeployParams(auth, ethclient)
	sigma, _ := DeploySigma(auth, ethclient, params)
	dleSigma, _ := DeployDleSigma(auth, ethclient)
	ip, _ := DeployInnerProduct(auth, ethclient, params)
	rangeProof, _ := DeployRangeproofverifier(auth, ethclient, params, ip)
	tokenConverter, _ := DeployTokenConverter(auth, ethclient)
	aggRange, _ := DeployAggRangeProof(auth, ethclient, params)
	pgcVerifier, _ := DeployPGCVerifier(auth, ethclient, params, dleSigma, rangeProof, aggRange, sigma)
	return DeployPGCMain(auth, ethclient, params, pgcVerifier, tokenConverter)
}
