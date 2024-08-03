package keeper

import (
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

var _ types.QueryServer = Keeper{}
