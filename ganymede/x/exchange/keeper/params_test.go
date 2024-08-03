package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ExchangeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
