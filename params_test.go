package pgc

import (
	"encoding/json"
	"testing"
)

func TestParams(t *testing.T) {
	p := Params()

	data, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	path := "solidity/params/params"
	WriteToFile(data, path)
}
