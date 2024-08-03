package keeper

import (
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

type msgServer struct {
	Keeper
	types.AccountKeeper
	types.BankKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, ak types.AccountKeeper, bk types.BankKeeper) types.MsgServer {
	return &msgServer{Keeper: keeper, AccountKeeper: ak, BankKeeper: bk}
}

var _ types.MsgServer = msgServer{}
