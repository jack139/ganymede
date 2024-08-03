package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ReplyAll(goCtx context.Context, req *types.QueryAllReplyRequest) (*types.QueryAllReplyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var replys []types.Reply
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	replyStore := prefix.NewStore(store, types.KeyPrefix(types.ReplyKey))

	pageRes, err := query.Paginate(replyStore, req.Pagination, func(key []byte, value []byte) error {
		var reply types.Reply
		if err := k.cdc.Unmarshal(value, &reply); err != nil {
			return err
		}

		replys = append(replys, reply)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReplyResponse{Reply: replys, Pagination: pageRes}, nil
}

func (k Keeper) Reply(goCtx context.Context, req *types.QueryGetReplyRequest) (*types.QueryGetReplyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	reply, found := k.GetReply(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetReplyResponse{Reply: reply}, nil
}
