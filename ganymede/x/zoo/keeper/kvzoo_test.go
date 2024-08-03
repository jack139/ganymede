package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/zoo/keeper"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNKvzoo(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Kvzoo {
	items := make([]types.Kvzoo, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)
		items[i].ZooKey = strconv.Itoa(i)

		keeper.SetKvzoo(ctx, items[i])
	}
	return items
}

func TestKvzooGet(t *testing.T) {
	keeper, ctx := keepertest.ZooKeeper(t)
	items := createNKvzoo(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetKvzoo(ctx,
			item.Owner,
			item.ZooKey,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestKvzooRemove(t *testing.T) {
	keeper, ctx := keepertest.ZooKeeper(t)
	items := createNKvzoo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKvzoo(ctx,
			item.Owner,
			item.ZooKey,
		)
		_, found := keeper.GetKvzoo(ctx,
			item.Owner,
			item.ZooKey,
		)
		require.False(t, found)
	}
}

func TestKvzooGetAll(t *testing.T) {
	keeper, ctx := keepertest.ZooKeeper(t)
	items := createNKvzoo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKvzoo(ctx)),
	)
}
