package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/exchange/keeper"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

func createNAsk(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ask {
	items := make([]types.Ask, n)
	for i := range items {
		items[i].Id = keeper.AppendAsk(ctx, items[i])
	}
	return items
}

func TestAskGet(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNAsk(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAsk(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAskRemove(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNAsk(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAsk(ctx, item.Id)
		_, found := keeper.GetAsk(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAskGetAll(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNAsk(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAsk(ctx)),
	)
}

func TestAskCount(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNAsk(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAskCount(ctx))
}
