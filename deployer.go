package pgc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pgc/contracts"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// Contracts includes all contract to run pgc system.
type Contracts struct {
	DLESigmaVerifier *contracts.Dlesigmaverifier
	IPVerifier       *contracts.Ipverifier
	PGC              *contracts.Pgc
	PublicParams     *contracts.Publicparams
	RangeVerifier    *contracts.Rangeproofverifier
	SigmaVerifier    *contracts.Sigmaverifier
}

// DeployPGCContracts deploys all contract of pgc system.
func DeployPGCContracts(auth *bind.TransactOpts, client *ethclient.Client) *Contracts {
	c := Contracts{}

	// get current nonce.
	if auth.Nonce == nil {
		nonce, err := client.NonceAt(context.Background(), auth.From, nil)
		if err != nil {
			panic(err)
		}

		auth.Nonce = new(big.Int).SetUint64(nonce)
	}

	// deploy public params contract.
	ppAddress, ppTx, pp, err := contracts.DeployPublicparams(auth, client)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy public params", "address", ppAddress.Hex(), "tx", ppTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.PublicParams = pp

	// deploy sigma contract.
	sigmaAddress, sigmaTx, sigma, err := contracts.DeploySigmaverifier(auth, client, ppAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy sigma verifier", "address", sigmaAddress.Hex(), "tx", sigmaTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.SigmaVerifier = sigma

	// deploy dle sigma contract.
	dleSigmaAddress, dleSigmaTx, dleSigma, err := contracts.DeployDlesigmaverifier(auth, client)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy dle sigma", "address", dleSigmaAddress.Hex(), "tx", dleSigmaTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.DLESigmaVerifier = dleSigma

	// deploy ip contract.
	ipAddress, ipTx, ip, err := contracts.DeployIpverifier(auth, client, ppAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy ip verifier", "address", ipAddress.Hex(), "tx", ipTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.IPVerifier = ip

	// deploy range proof verifier
	rangeAddress, rangeTx, rangeVerifier, err := contracts.DeployRangeproofverifier(auth, client, ppAddress, ipAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy range verifier", "address", rangeAddress.Hex(), "tx", rangeTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.RangeVerifier = rangeVerifier

	// deploy pgc contract
	pgcAddress, pgcTx, pgcC, err := contracts.DeployPgc(auth, client, ppAddress, dleSigmaAddress, rangeAddress, sigmaAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy pgc", "address", pgcAddress, "tx", pgcTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.PGC = pgcC

	return &c
}
