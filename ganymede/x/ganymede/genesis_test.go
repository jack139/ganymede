package ganymede_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/ganymede"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		UsersList: []types.Users{
			{
				ChainAddr: "0",
			},
			{
				ChainAddr: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GanymedeKeeper(t)
	ganymede.InitGenesis(ctx, *k, genesisState)
	got := ganymede.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UsersList, got.UsersList)
	// this line is used by starport scaffolding # genesis/test/assert
}
