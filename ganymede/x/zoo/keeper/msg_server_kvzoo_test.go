package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/x/zoo/keeper"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestKvzooMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.ZooKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateKvzoo{Creator: creator,
			Owner:  strconv.Itoa(i),
			ZooKey: strconv.Itoa(i),
		}
		_, err := srv.CreateKvzoo(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetKvzoo(ctx,
			expected.Owner,
			expected.ZooKey,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestKvzooMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateKvzoo
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateKvzoo{Creator: creator,
				Owner:  strconv.Itoa(0),
				ZooKey: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateKvzoo{Creator: "B",
				Owner:  strconv.Itoa(0),
				ZooKey: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateKvzoo{Creator: creator,
				Owner:  strconv.Itoa(100000),
				ZooKey: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ZooKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateKvzoo{Creator: creator,
				Owner:  strconv.Itoa(0),
				ZooKey: strconv.Itoa(0),
			}
			_, err := srv.CreateKvzoo(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateKvzoo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetKvzoo(ctx,
					expected.Owner,
					expected.ZooKey,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestKvzooMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteKvzoo
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteKvzoo{Creator: creator,
				Owner:  strconv.Itoa(0),
				ZooKey: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteKvzoo{Creator: "B",
				Owner:  strconv.Itoa(0),
				ZooKey: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteKvzoo{Creator: creator,
				Owner:  strconv.Itoa(100000),
				ZooKey: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.ZooKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateKvzoo(wctx, &types.MsgCreateKvzoo{Creator: creator,
				Owner:  strconv.Itoa(0),
				ZooKey: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteKvzoo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetKvzoo(ctx,
					tc.request.Owner,
					tc.request.ZooKey,
				)
				require.False(t, found)
			}
		})
	}
}
