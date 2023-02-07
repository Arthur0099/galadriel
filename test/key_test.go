package test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/pgc/proof"
)

func TestGenerateKey(t *testing.T) {
	k := proof.MustGenerateKey(proof.DefaultKeyParams())

	fmt.Println(hex.EncodeToString(k.D.Bytes()))
}
