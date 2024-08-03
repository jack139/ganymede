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

func (k Keeper) AskAll(goCtx context.Context, req *types.QueryAllAskRequest) (*types.QueryAllAskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var asks []types.Ask
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	askStore := prefix.NewStore(store, types.KeyPrefix(types.AskKey))

	pageRes, err := query.Paginate(askStore, req.Pagination, func(key []byte, value []byte) error {
		var ask types.Ask
		if err := k.cdc.Unmarshal(value, &ask); err != nil {
			return err
		}

		asks = append(asks, ask)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAskResponse{Ask: asks, Pagination: pageRes}, nil
}

func (k Keeper) Ask(goCtx context.Context, req *types.QueryGetAskRequest) (*types.QueryGetAskResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ask, found := k.GetAsk(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAskResponse{Ask: ask}, nil
}
