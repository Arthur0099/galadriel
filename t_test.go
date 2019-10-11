package pgc

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestTT(t *testing.T) {
	s := "owpass$2019$#&#!@&"

	r := Keccak256([]byte(s))
	ss := common.BytesToHash(r).Hex()
	t.Log(ss)
}
