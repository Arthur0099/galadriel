package pgc

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ropstenInfura = "https://ropsten.infura.io/v3/10d08c76bb104f6286f774ec21fa7ac9"
	local         = "http://192.168.1.115:8545"
)

// GetRopstenInfura returns infura client on main net.
func GetRopstenInfura() *ethclient.Client {
	return getClient(ropstenInfura)
}

// GetLocal return client on local net.
func GetLocal() *ethclient.Client {
	return getClient(local)
}

// GetLocalRPC returns rpc client on local net.
func GetLocalRPC() *Client {
	client, err := Dial(local)
	if err != nil {
		panic(err)
	}

	return client
}

func getClient(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	return client
}

// GetRopstenAccount returns default ropsten test account.
func GetRopstenAccount() *bind.TransactOpts {
	keyHex := "B38BB7EF4D69CCB1D5A1735887521BC1717AF203AE8BE7F8928E9ECC54FFD5E3"

	key, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		panic(err)
	}

	return bind.NewKeyedTransactor(key)
}
