package zoo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/zoo"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KvzooList: []types.Kvzoo{
			{
				Owner:  "0",
				ZooKey: "0",
			},
			{
				Owner:  "1",
				ZooKey: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZooKeeper(t)
	zoo.InitGenesis(ctx, *k, genesisState)
	got := zoo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.KvzooList, got.KvzooList)
	// this line is used by starport scaffolding # genesis/test/assert
}
