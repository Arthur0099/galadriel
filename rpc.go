package pgc

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"io/ioutil"
	"math/big"
	"net/http"
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

// TronClient is a http client.
type TronClient struct {
	baseURL string
}

const (
	listAccount    = "/admin/accounts"
	deployContract = "/wallet/deploycontract"
	broadcasttx    = "/wallet/broadcasttransaction"
	getcontract    = "/wallet/getcontract"
	freezebalance  = "/wallet/freezebalance"
)

// NewTronClient create tron http client.
func NewTronClient(url string) *TronClient {
	c := TronClient{}
	c.baseURL = url

	return &c
}

// FreezeBalance .
func (tc *TronClient) FreezeBalance(data interface{}) ([]byte, error) {
	return jsonPost(tc.baseURL+freezebalance, data)
}

// ListAccounts returns accounts of tron private net.
// Warn: only for private net.
func (tc *TronClient) ListAccounts() (res []byte, err error) {
	return get(tc.baseURL + listAccount)
}

// DeployContract deploy contracts.
func (tc *TronClient) DeployContract(data interface{}) ([]byte, error) {
	return jsonPost(tc.baseURL+deployContract, data)
}

// GetContract .
func (tc *TronClient) GetContract(data interface{}) ([]byte, error) {
	return jsonPost(tc.baseURL+getcontract, data)
}

// BroadcastTx .
func (tc *TronClient) BroadcastTx(tx interface{}) ([]byte, error) {
	return jsonPost(tc.baseURL+broadcasttx, tx)
}

// SignTronTx .
func SignTronTx(rawData string, key *ecdsa.PrivateKey) (sig []byte, err error) {
	data, err := hex.DecodeString(rawData)
	hash := Sha256Hash(data)

	// sign for tx.
	sig, err = crypto.Sign(hash, key)
	return
}

func get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}

func jsonPost(url string, data interface{}) (body []byte, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}
