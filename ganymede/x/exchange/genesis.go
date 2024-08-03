package exchange

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/exchange/keeper"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the ask
	for _, elem := range genState.AskList {
		k.SetAsk(ctx, elem)
	}

	// Set ask count
	k.SetAskCount(ctx, genState.AskCount)
	// Set all the reply
	for _, elem := range genState.ReplyList {
		k.SetReply(ctx, elem)
	}

	// Set reply count
	k.SetReplyCount(ctx, genState.ReplyCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AskList = k.GetAllAsk(ctx)
	genesis.AskCount = k.GetAskCount(ctx)
	genesis.ReplyList = k.GetAllReply(ctx)
	genesis.ReplyCount = k.GetReplyCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
