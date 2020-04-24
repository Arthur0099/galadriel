package deployer

import (
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/pgc/client"
	"github.com/pgc/contracts/ipverifier"
	pgcm "github.com/pgc/contracts/pgc"
	"github.com/pgc/contracts/publicparams"
	"github.com/pgc/contracts/rangeproofverifier"
	"github.com/pgc/contracts/token"
	"github.com/pgc/contracts/tokenconverter"
	"github.com/pgc/contracts/verifier"
	"github.com/pgc/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DeployParams deploy param contract.
func DeployParams(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *publicparams.Publicparams) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := publicparams.DeployPublicparams(auth, ethclient)
	if err != nil {
		panic(err)
	}
	log.Info("send deploy params tx success", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployInnerProduct deploys inner product contract.
func DeployInnerProduct(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *ipverifier.Ipverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := ipverifier.DeployIpverifier(auth, ethclient)
	if err != nil {
		panic(err)
	}
	log.Info("send deploy inner product tx success", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployRangeproofverifier deploys single range proof verifier.
func DeployRangeproofverifier(auth *bind.TransactOpts, ethclient *ethclient.Client, inner common.Address) (common.Address, *rangeproofverifier.Rangeproofverifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := rangeproofverifier.DeployRangeproofverifier(auth, ethclient, inner)
	if err != nil {
		panic(err)
	}
	log.Info("send deploy range proof tx success", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployToken deploys a erc-20 token contract.
func DeployToken(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *token.Token) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := token.DeployToken(auth, ethclient)
	if err != nil {
		panic(err)
	}
	log.Info("send deploy token tx success", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployTokenConverter deploys contract to convert token.
func DeployTokenConverter(auth *bind.TransactOpts, ethclient *ethclient.Client) (common.Address, *tokenconverter.Tokenconverter) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := tokenconverter.DeployTokenconverter(auth, ethclient)
	if err != nil {
		panic(err)
	}
	log.Info("send deploy token converter tx success", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployVerifier deploys pgc proof verifier contract.
func DeployVerifier(auth *bind.TransactOpts, ethclient *ethclient.Client, params common.Address) (common.Address, *verifier.Verifier) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := verifier.DeployVerifier(auth, ethclient, params)
	if err != nil {
		panic(err)
	}
	log.Info("send deploy verifier tx success", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployPGCMain deploys pgc system main contract.
func DeployPGCMain(auth *bind.TransactOpts, ethclient *ethclient.Client, params, pgcVerifier, tokenConverter common.Address) (common.Address, *pgcm.Pgc) {
	if err := client.SetNonce(auth, ethclient); err != nil {
		panic(err)
	}

	addr, tx, con, err := pgcm.DeployPgc(auth, ethclient, params, pgcVerifier, tokenConverter)
	if err != nil {
		panic(err)
	}
	log.Info("send deploy pgc main tx success", "tx", tx.Hash().Hex())
	auth.Nonce.Add(auth.Nonce, utils.One)
	client.WaitForTx(ethclient, tx.Hash())

	return addr, con
}

// DeployPGCSystemAllContract deploys all contract for pgc system.
func DeployPGCSystemAllContract(auth *bind.TransactOpts, ethclient *ethclient.Client) ([]common.Address, *pgcm.Pgc) {
	addrs := make([]common.Address, 0)

	params, _ := DeployParams(auth, ethclient)
	addrs = append(addrs, params)

	tokenConverter, _ := DeployTokenConverter(auth, ethclient)
	addrs = append(addrs, tokenConverter)

	verifierAddr, _ := DeployVerifier(auth, ethclient, params)
	addrs = append(addrs, verifierAddr)

	pgcMain, pgc := DeployPGCMain(auth, ethclient, params, verifierAddr, tokenConverter)
	addrs = append(addrs, pgcMain)

	return addrs, pgc
}

// InitVector init g/h generate vector for agg range proof.
func InitVector(auth *bind.TransactOpts, ethclient *ethclient.Client, addr common.Address, bitsize int) {
	vectorsize := bitsize * 2
	initStep := 32
	step := 45
	if vectorsize <= initStep {
		return
	}
	vectorsize -= initStep
	vr, err := verifier.NewVerifier(addr, ethclient)
	if err != nil {
		panic(err)
	}

	for vectorsize > 0 {
		tx, err := vr.Init(auth, big.NewInt(int64(step)))
		if err != nil {
			panic(err)
		}
		log.Info("send init verifier tx success", "tx", tx.Hash().Hex())
		client.WaitForTx(ethclient, tx.Hash())
		auth.Nonce.Add(auth.Nonce, utils.One)
		vectorsize -= step
	}

}
