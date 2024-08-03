package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestUsersQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GanymedeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUsers(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetUsersRequest
		response *types.QueryGetUsersResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUsersRequest{
				ChainAddr: msgs[0].ChainAddr,
			},
			response: &types.QueryGetUsersResponse{Users: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUsersRequest{
				ChainAddr: msgs[1].ChainAddr,
			},
			response: &types.QueryGetUsersResponse{Users: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUsersRequest{
				ChainAddr: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Users(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestUsersQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GanymedeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNUsers(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUsersRequest {
		return &types.QueryAllUsersRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UsersAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Users), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Users),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.UsersAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Users), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Users),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.UsersAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Users),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.UsersAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
