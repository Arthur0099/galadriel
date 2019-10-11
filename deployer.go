package pgc

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pgc/contracts"

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

func MustDeployContract(name string, auth *bind.TransactOpts, client *ethclient.Client, params ...interface{}) (
	addr common.Address,
	tx *types.Transaction,
	contract interface{},
) {

	var ABI, BIN string

	switch name {
	case "publicParams":
		ABI = contracts.PublicparamsABI
		BIN = contracts.PublicparamsBin
	case "sigmaVerifier":
		ABI = contracts.SigmaverifierABI
		BIN = contracts.SigmaverifierBin
	case "dleSigmaVerifier":
		ABI = contracts.DlesigmaverifierABI
		BIN = contracts.DlesigmaverifierBin
	case "ipVerifier":
		ABI = contracts.IpverifierABI
		BIN = contracts.IpverifierBin
	case "rangeProofVerifier":
		ABI = contracts.RangeproofverifierABI
		BIN = contracts.RangeproofverifierBin
	case "tokenConverter":
		ABI = contracts.TokenconverterABI
		BIN = contracts.TokenconverterBin
	case "pgcVerifier":
		ABI = contracts.PgcverifierABI
		BIN = contracts.PgcverifierBin
	case "pgc":
		ABI = contracts.PgcABI
		BIN = contracts.PgcBin
	default:
		panic("contract not support")
	}

	parsed, err := abi.JSON(strings.NewReader(ABI))
	if err != nil {
		panic(err)
	}

	// log info.
	if len(params) == 0 {
		addr, tx, contract, err = bind.DeployContract(auth, parsed, common.FromHex(BIN), client)
	} else {
		addr, tx, contract, err = bind.DeployContract(auth, parsed, common.FromHex(BIN), client, params...)
	}
	if err != nil {
		panic(err)
	}
	auth.Nonce.Add(auth.Nonce, one)

	log.Debug("deploy contract success", "name", name, "address", addr, "tx", tx.Hash().Hex())
	return
}

func waitFor(tx common.Hash, client *ethclient.Client) {
	for {
		status, err := client.TransactionReceipt(context.Background(), tx)
		if err != nil || status.Status != 1 {
			time.Sleep(time.Second * 1)
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
	params, _, _ := MustDeployContract("publicParams", auth, client)
	c.PublicParams, _ = contracts.NewPublicparams(params, client)

	// deploy sigma contract.
	sigmaVerifier, _, _ := MustDeployContract("sigmaVerifier", auth, client, params)
	c.SigmaVerifier, _ = contracts.NewSigmaverifier(sigmaVerifier, client)

	// deploy dle sigma contract.
	dleSigmaVerifier, _, _ := MustDeployContract("dleSigmaVerifier", auth, client)
	c.DLESigmaVerifier, _ = contracts.NewDlesigmaverifier(dleSigmaVerifier, client)

	// deploy ip contract.
	ipVerifier, _, _ := MustDeployContract("ipVerifier", auth, client, params)
	c.IPVerifier, _ = contracts.NewIpverifier(ipVerifier, client)

	// deploy range proof verifier
	rangeVerifier, _, _ := MustDeployContract("rangeProofVerifier", auth, client, params, ipVerifier)
	c.RangeVerifier, _ = contracts.NewRangeproofverifier(rangeVerifier, client)

	// deploy token convert contract.
	tokenConverter, _, _ := MustDeployContract("tokenConverter", auth, client)
	c.TokenConverter, _ = contracts.NewTokenconverter(tokenConverter, client)

	// deploy pgc verifyer.
	pgcVerifier, _, _ := MustDeployContract("pgcVerifier", auth, client, params, dleSigmaVerifier, rangeVerifier, sigmaVerifier)

	// deploy pgc contract
	pgcCon, _, _ := MustDeployContract("pgc", auth, client, params, pgcVerifier, tokenConverter)
	c.PGC, _ = contracts.NewPgc(pgcCon, client)
	c.PGCAddress = pgcCon

	return &c
}
