package ganymede

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/ganymede/keeper"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the users
	for _, elem := range genState.UsersList {
		k.SetUsers(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UsersList = k.GetAllUsers(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
