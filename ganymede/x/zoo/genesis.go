package zoo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/zoo/keeper"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the kvzoo
	for _, elem := range genState.KvzooList {
		k.SetKvzoo(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.KvzooList = k.GetAllKvzoo(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
