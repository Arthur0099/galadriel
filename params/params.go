package params

import (
	"github.com/pgc/proof"
)

// DAggRangeProofParams returns default params for pgc/solidity system to genrate
// verify all proofs.
func DAggRangeProofParams() proof.AggRangeParams {
	return proof.DAggRangeProofParams()
}
