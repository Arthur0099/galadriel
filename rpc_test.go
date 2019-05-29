package pgc

import (
	"math/big"
	"testing"
)

func TestClient(t *testing.T) {
	url := "http://127.0.0.1:8545"
	c, _ := Dial(url)
	accounts, err := c.GetAccounts()
	if err != nil {
		t.Error(err)
		return
	}

	// test send 1 eth from first account to second account.
	tx := Transaction{}
	tx.From = accounts[0]
	tx.To = accounts[1]
	tx.Value = new(big.Int).SetUint64(1000 * 1000 * 1000 * 1000).String()
	tx.GasLimit = "100000"
	txHash, err := c.SendTx(&tx)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(txHash.Hex())
}
