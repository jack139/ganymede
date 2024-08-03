package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UsersAll(goCtx context.Context, req *types.QueryAllUsersRequest) (*types.QueryAllUsersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userss []types.Users
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	usersStore := prefix.NewStore(store, types.KeyPrefix(types.UsersKeyPrefix))

	pageRes, err := query.Paginate(usersStore, req.Pagination, func(key []byte, value []byte) error {
		var users types.Users
		if err := k.cdc.Unmarshal(value, &users); err != nil {
			return err
		}

		userss = append(userss, users)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUsersResponse{Users: userss, Pagination: pageRes}, nil
}

func (k Keeper) Users(goCtx context.Context, req *types.QueryGetUsersRequest) (*types.QueryGetUsersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetUsers(
		ctx,
		req.ChainAddr,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUsersResponse{Users: val}, nil
}
