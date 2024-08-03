package keeper

import (
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
)

var _ types.QueryServer = Keeper{}
