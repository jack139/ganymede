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

func createNReply(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Reply {
	items := make([]types.Reply, n)
	for i := range items {
		items[i].Id = keeper.AppendReply(ctx, items[i])
	}
	return items
}

func TestReplyGet(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNReply(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetReply(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestReplyRemove(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNReply(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveReply(ctx, item.Id)
		_, found := keeper.GetReply(ctx, item.Id)
		require.False(t, found)
	}
}

func TestReplyGetAll(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNReply(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllReply(ctx)),
	)
}

func TestReplyCount(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	items := createNReply(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetReplyCount(ctx))
}
