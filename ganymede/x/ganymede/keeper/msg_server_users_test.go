package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/x/ganymede/keeper"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUsersMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.GanymedeKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateUsers{Creator: creator,
			ChainAddr: strconv.Itoa(i),
		}
		_, err := srv.CreateUsers(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetUsers(ctx,
			expected.ChainAddr,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestUsersMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateUsers
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateUsers{Creator: creator,
				ChainAddr: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateUsers{Creator: "B",
				ChainAddr: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateUsers{Creator: creator,
				ChainAddr: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.GanymedeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateUsers{Creator: creator,
				ChainAddr: strconv.Itoa(0),
			}
			_, err := srv.CreateUsers(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateUsers(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetUsers(ctx,
					expected.ChainAddr,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestUsersMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteUsers
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteUsers{Creator: creator,
				ChainAddr: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteUsers{Creator: "B",
				ChainAddr: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteUsers{Creator: creator,
				ChainAddr: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.GanymedeKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateUsers(wctx, &types.MsgCreateUsers{Creator: creator,
				ChainAddr: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteUsers(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetUsers(ctx,
					tc.request.ChainAddr,
				)
				require.False(t, found)
			}
		})
	}
}
