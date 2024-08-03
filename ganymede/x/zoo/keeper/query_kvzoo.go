package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) KvzooAll(goCtx context.Context, req *types.QueryAllKvzooRequest) (*types.QueryAllKvzooResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var kvzoos []types.Kvzoo
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	kvzooStore := prefix.NewStore(store, types.KeyPrefix(types.KvzooKeyPrefix))

	pageRes, err := query.Paginate(kvzooStore, req.Pagination, func(key []byte, value []byte) error {
		var kvzoo types.Kvzoo
		if err := k.cdc.Unmarshal(value, &kvzoo); err != nil {
			return err
		}

		kvzoos = append(kvzoos, kvzoo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKvzooResponse{Kvzoo: kvzoos, Pagination: pageRes}, nil
}

func (k Keeper) Kvzoo(goCtx context.Context, req *types.QueryGetKvzooRequest) (*types.QueryGetKvzooResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetKvzoo(
		ctx,
		req.Owner,
		req.ZooKey,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKvzooResponse{Kvzoo: val}, nil
}
