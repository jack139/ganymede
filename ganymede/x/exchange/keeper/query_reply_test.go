package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/jack139/ganymede/ganymede/testutil/keeper"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

func TestReplyQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNReply(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetReplyRequest
		response *types.QueryGetReplyResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetReplyRequest{Id: msgs[0].Id},
			response: &types.QueryGetReplyResponse{Reply: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetReplyRequest{Id: msgs[1].Id},
			response: &types.QueryGetReplyResponse{Reply: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetReplyRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Reply(wctx, tc.request)
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

func TestReplyQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ExchangeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNReply(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllReplyRequest {
		return &types.QueryAllReplyRequest{
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
			resp, err := keeper.ReplyAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Reply), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Reply),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ReplyAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Reply), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Reply),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ReplyAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Reply),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ReplyAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
