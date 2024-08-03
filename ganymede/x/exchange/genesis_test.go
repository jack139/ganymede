package exchange_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/exchange"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AskList: []types.Ask{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AskCount: 2,
		ReplyList: []types.Reply{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ReplyCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExchangeKeeper(t)
	exchange.InitGenesis(ctx, *k, genesisState)
	got := exchange.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AskList, got.AskList)
	require.Equal(t, genesisState.AskCount, got.AskCount)
	require.ElementsMatch(t, genesisState.ReplyList, got.ReplyList)
	require.Equal(t, genesisState.ReplyCount, got.ReplyCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
