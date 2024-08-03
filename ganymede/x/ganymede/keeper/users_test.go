package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/ganymede/keeper"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUsers(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Users {
	items := make([]types.Users, n)
	for i := range items {
		items[i].ChainAddr = strconv.Itoa(i)

		keeper.SetUsers(ctx, items[i])
	}
	return items
}

func TestUsersGet(t *testing.T) {
	keeper, ctx := keepertest.GanymedeKeeper(t)
	items := createNUsers(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUsers(ctx,
			item.ChainAddr,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUsersRemove(t *testing.T) {
	keeper, ctx := keepertest.GanymedeKeeper(t)
	items := createNUsers(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUsers(ctx,
			item.ChainAddr,
		)
		_, found := keeper.GetUsers(ctx,
			item.ChainAddr,
		)
		require.False(t, found)
	}
}

func TestUsersGetAll(t *testing.T) {
	keeper, ctx := keepertest.GanymedeKeeper(t)
	items := createNUsers(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUsers(ctx)),
	)
}
