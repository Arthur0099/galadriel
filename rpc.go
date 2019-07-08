package pgc

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
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
	Data     []byte         `json:"data"`
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

// SendETH send ether to one.
func (c *Client) SendETH(from common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	ether := new(big.Int).SetUint64(1000 * 1000 * 1000 * 1000 * 1000 * 1000)
	tx := Transaction{}
	tx.From = from
	tx.To = to
	tx.Value = new(big.Int).Mul(amount, ether).String()
	tx.GasLimit = "100000"

	return c.SendTx(&tx)
}
