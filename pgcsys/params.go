package pgcsys

import (
	"math/big"
	"os"

	"github.com/pgc/curve"
	"github.com/pgc/proof"
	"github.com/pgc/utils"
)

// DefaultPgcSysParams only for test
func DefaultPgcSysParams() PgcSys {
	curve := curve.BN256()
	bitsize := 32
	aggsize := 2
	g := "g generator of twisted elg"
	gpoint := utils.NewECPointByString(g, curve)

	h := "h generator of twisted elg"
	hpoint := utils.NewECPointByString(h, curve)

	gv := utils.NewDefaultGV(curve, bitsize*aggsize)
	hv := utils.NewDefaultHV(curve, bitsize*aggsize)

	u := "u generator of innerproduct"
	upoint := utils.NewECPointByString(u, curve)

	arp := pgcSysParams{}
	arp.aggsize = aggsize
	arp.bitsize = bitsize
	arp.g = gpoint
	arp.h = hpoint
	arp.u = upoint
	arp.gv = gv
	arp.hv = hv

	// only for test
	priv, err := proof.HexToKey(os.Getenv("GlobalKey"), &arp)
	if err != nil || priv.D.Cmp(big.NewInt(0)) == 0 {
		panic("Empty GlobalKey")
	}

	arp.pub = new(utils.ECPoint).SetFromPublicKey(&priv.PublicKey)

	return &arp
}
