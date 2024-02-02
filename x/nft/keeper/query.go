package keeper

import (
	"divine/x/nft/types"
)

var _ types.QueryServer = Keeper{}
