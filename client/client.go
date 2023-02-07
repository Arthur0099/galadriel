package client

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	eth       = new(big.Int).SetUint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000)
	amount    = new(big.Int).SetUint64(10)
	amountWei = new(big.Int).Mul(amount, eth)
)

// Client represents rpc client.
type Client struct {
	client *rpc.Client
}

// Transaction represents a tx without sign.
type Transaction struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Nonce    string         `json:"nonce,omitempty"`
	Value    string         `json:"value"`
	GasLimit string         `json:"gasLimit"`
	GasPrice string         `json:"gasPrice"`
	Data     string         `json:"data"`
}

// Dial connect to rpc client
func Dial(url string) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	client := Client{}
	client.client = c

	return &client, nil
}

// GetAccounts returns all accounts.
func (c *Client) GetAccounts() ([]common.Address, error) {
	var res []common.Address
	err := c.client.Call(&res, "eth_accounts")

	return res, err
}

// SendTx send tx without sign it using send_transaction.
func (c *Client) SendTx(tx *Transaction) (common.Hash, error) {
	var res common.Hash

	err := c.client.Call(&res, "eth_sendTransaction", tx)

	return res, err
}

// Mine mines n blocks.
func (c *Client) Mine(n int) error {
	for i := 0; i < n; i++ {
		if err := c.EVMMine(); err != nil {
			return err
		}
	}

	return nil
}

// EVMMine mines a block on chain.
// Only available in ganache-cli.
func (c *Client) EVMMine() error {
	return c.client.Call(nil, "evm_mine")
}

// SendFromFirstAccout sends eth to someone from test local account get by rpc.
func (c *Client) SendFromFirstAccout(to common.Address, amount *big.Int) (common.Hash, error) {
	accounts, err := c.GetAccounts()
	if err != nil {
		return common.Hash{}, err
	}

	return c.SendETH(accounts[0], to, amount)
}

func (c *Client) GetPendingNonce(account common.Address) (string, error) {
	var result hexutil.Uint64
	err := c.client.CallContext(context.Background(), &result, "eth_getTransactionCount", account, "pending")

	return result.String(), err
}

// SendETH send ether to one.
func (c *Client) SendETH(from common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	tx := Transaction{}
	tx.From = from
	tx.To = to
	tx.Value = "0x" + amountWei.String()
	tx.GasLimit = "0x100000"
	tx.GasPrice = "0x9184e72a00"
	tx.Data = "0x0"
	nonce, _ := c.GetPendingNonce(from)
	tx.Nonce = nonce

	return c.SendTx(&tx)
}

// GetAccountWithETH send eth to an random account.
func (c *Client) GetAccountWithETH(client *ethclient.Client) *bind.TransactOpts {
	account := GenerateAccount(client)

	if _, err := c.SendFromFirstAccout(account.From, amount); err != nil {
		panic(err)
	}

	return account
}

const (
	ropstenInfura = "https://ropsten.infura.io/v3/10d08c76bb104f6286f774ec21fa7ac9"
	local         = "http://127.0.0.1:8545"
	scrollUrlL2   = "https://prealpha.scroll.io/l2"
)

// GetClient returns ethclient, otherwise panic.
func GetClient(url string) *ethclient.Client {
	return getClient(url)
}

// GetRopstenInfura returns infura client on main net.
func GetRopstenInfura() *ethclient.Client {
	return getClient(ropstenInfura)
}

// GetScrollClientL2 return client for l2 testnet
func GetScrollClientL2() *ethclient.Client {
	return getClient(scrollUrlL2)
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

// GetAccountWithKey returns opts for sending tx.
func GetAccountWithKey(keyhex string, client *ethclient.Client) *bind.TransactOpts {
	key, err := crypto.HexToECDSA(keyhex)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		panic(err)
	}

	return auth
}

// GenerateAccount generates a new random account.
func GenerateAccount(client *ethclient.Client) *bind.TransactOpts {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		panic(err)
	}
	return auth
}

// SetNonce sets nonce to account if nil.
func SetNonce(auth *bind.TransactOpts, client *ethclient.Client) error {
	if auth.Nonce != nil {
		return nil
	}

	nonce, err := client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return err
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)
	return nil
}

// WaitForTx stack until tx being mined.
func WaitForTx(client *ethclient.Client, hash common.Hash) *types.Receipt {
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime) > time.Minute*5 {
			panic(fmt.Sprintf("wait too much time: tx %s", hash.Hex()))
		}
		receipt, err := client.TransactionReceipt(context.Background(), hash)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		if receipt != nil && receipt.Status != 1 {
			panic("tx failed")
		}

		if receipt != nil && receipt.Status == 1 {
			return receipt
		}

		time.Sleep(time.Millisecond * 500)
	}
}
