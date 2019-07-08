package pgc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pgc/contracts"
	"math/big"
	"time"

	log "github.com/inconshreveable/log15"
)

// Contracts includes all contract to run pgc system.
type Contracts struct {
	DLESigmaVerifier *contracts.Dlesigmaverifier
	IPVerifier       *contracts.Ipverifier
	PGC              *contracts.Pgc
	PGCAddress       common.Address
	PublicParams     *contracts.Publicparams
	RangeVerifier    *contracts.Rangeproofverifier
	SigmaVerifier    *contracts.Sigmaverifier
	TokenConverter   *contracts.Tokenconverter
}

func waitFor(tx common.Hash, client *ethclient.Client) {
	for {
		status, err := client.TransactionReceipt(context.Background(), tx)
		if err != nil || status.Status != 1 {
			time.Sleep(time.Second * 5)
			continue
		}

		return
	}
}

func waitForBlocks(n uint64, client *ethclient.Client) {
	log.Debug("wait for blocks to be mined", "amount", n)
	var currentNumber uint64
	for currentNumber == 0 {
		if currentBlock, err := client.BlockByNumber(context.Background(), nil); err == nil {
			currentNumber = currentBlock.Number().Uint64()
		}
	}
	log.Debug("start pos", "number", currentNumber)
	for {
		block, err := client.BlockByNumber(context.Background(), nil)
		if err != nil || block.Number().Uint64()-currentNumber < n {
			time.Sleep(time.Second * 10)
			continue
		}
		log.Debug("advance to block", "number", block.Number())

		return
	}
}

// NewPGCContracts returns deployed contracts.
func NewPGCContracts(dle, ip, pgcc, pp, rv, sv common.Address, client *ethclient.Client) *Contracts {
	c := Contracts{}

	c.DLESigmaVerifier, _ = contracts.NewDlesigmaverifier(dle, client)
	c.IPVerifier, _ = contracts.NewIpverifier(ip, client)
	c.PGC, _ = contracts.NewPgc(pgcc, client)
	c.PublicParams, _ = contracts.NewPublicparams(pp, client)
	c.RangeVerifier, _ = contracts.NewRangeproofverifier(rv, client)
	c.SigmaVerifier, _ = contracts.NewSigmaverifier(sv, client)

	return &c
}

// DeployToken deploys a erc20 token contract for test.
func DeployToken(auth *bind.TransactOpts, client *ethclient.Client) (*contracts.Token, common.Address) {
	// get current nonce.
	if auth.Nonce == nil {
		nonce, err := client.NonceAt(context.Background(), auth.From, nil)
		if err != nil {
			panic(err)
		}

		auth.Nonce = new(big.Int).SetUint64(nonce)
	}

	tokenAddr, tokenTx, token, err := contracts.DeployToken(auth, client)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy token contracts", "address", tokenAddr.Hex(), "tx", tokenTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	waitFor(tokenTx.Hash(), client)

	return token, tokenAddr
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
	waitFor(ppTx.Hash(), client)

	// deploy sigma contract.
	sigmaAddress, sigmaTx, sigma, err := contracts.DeploySigmaverifier(auth, client, ppAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy sigma verifier", "address", sigmaAddress.Hex(), "tx", sigmaTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.SigmaVerifier = sigma
	waitFor(sigmaTx.Hash(), client)

	// deploy dle sigma contract.
	dleSigmaAddress, dleSigmaTx, dleSigma, err := contracts.DeployDlesigmaverifier(auth, client)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy dle sigma", "address", dleSigmaAddress.Hex(), "tx", dleSigmaTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.DLESigmaVerifier = dleSigma
	waitFor(dleSigmaTx.Hash(), client)

	// deploy ip contract.
	ipAddress, ipTx, ip, err := contracts.DeployIpverifier(auth, client, ppAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy ip verifier", "address", ipAddress.Hex(), "tx", ipTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.IPVerifier = ip
	waitFor(ipTx.Hash(), client)

	// deploy range proof verifier
	rangeAddress, rangeTx, rangeVerifier, err := contracts.DeployRangeproofverifier(auth, client, ppAddress, ipAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy range verifier", "address", rangeAddress.Hex(), "tx", rangeTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.RangeVerifier = rangeVerifier
	waitFor(rangeTx.Hash(), client)

	// deploy token convert contract.
	tokenConvertAddress, tokenConvertTx, tokenConverter, err := contracts.DeployTokenconverter(auth, client)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy token converter", "address", tokenConvertAddress, "tx", tokenConvertTx.Hash())
	c.TokenConverter = tokenConverter
	auth.Nonce.Add(auth.Nonce, one)
	waitFor(tokenConvertTx.Hash(), client)

	// deploy pgc verifyer.
	pgcVerifierAddr, pgcVerifierTx, _, err := contracts.DeployPgcverifier(auth, client, ppAddress, dleSigmaAddress, rangeAddress, sigmaAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy pgcverifier", "address", pgcVerifierAddr, "tx", pgcVerifierTx.Hash())
	auth.Nonce.Add(auth.Nonce, one)
	waitFor(pgcVerifierTx.Hash(), client)

	// deploy pgc contract
	pgcAddress, pgcTx, pgcC, err := contracts.DeployPgc(auth, client, ppAddress, pgcVerifierAddr, tokenConvertAddress)
	if err != nil {
		panic(err)
	}
	log.Debug("deploy pgc", "address", pgcAddress, "tx", pgcTx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, one)
	c.PGC = pgcC
	c.PGCAddress = pgcAddress
	waitFor(pgcTx.Hash(), client)

	return &c
}
