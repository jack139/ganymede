package postoffice_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/postoffice"
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		PostList: []types.Post{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PostCount: 2,
		SentPostList: []types.SentPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentPostCount: 2,
		TimedoutPostList: []types.TimedoutPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TimedoutPostCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PostofficeKeeper(t)
	postoffice.InitGenesis(ctx, *k, genesisState)
	got := postoffice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	require.Equal(t, genesisState.PostCount, got.PostCount)
	require.ElementsMatch(t, genesisState.SentPostList, got.SentPostList)
	require.Equal(t, genesisState.SentPostCount, got.SentPostCount)
	require.ElementsMatch(t, genesisState.TimedoutPostList, got.TimedoutPostList)
	require.Equal(t, genesisState.TimedoutPostCount, got.TimedoutPostCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
