package keeper

import (
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

var _ types.QueryServer = Keeper{}
