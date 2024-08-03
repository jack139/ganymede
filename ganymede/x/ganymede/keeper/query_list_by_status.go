package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListByStatus(goCtx context.Context, req *types.QueryListByStatusRequest) (*types.QueryListByStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []*types.Users
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	storeLink := prefix.NewStore(store, types.KeyPrefix(types.UsersLinkPrefix))

	// 找到 status 的尾部
	linkKey := types.UsersStatusLinkKey(req.Status) // 链表头 key
	kvKey := storeLink.Get(linkKey)
	if kvKey == nil { // 未找到 status 表头
		//return nil, errors.New("cannot find status tail!")
		return &types.QueryListByStatusResponse{Users: users}, nil
	}

	if string(kvKey) == "@@LINK:$" {
		// 尾部就是结束标记，直接返回
		return &types.QueryListByStatusResponse{Users: users}, nil
	}


	// 遍历链表
	var count,skip int
	count = 0
	skip = int( (req.Page - 1) * req.Limit )
	for {
		kvzoo, found := k.GetUsersByKey(ctx, kvKey)
		if !found {
			return nil, errors.New("kvKey fault!")
		}

		count++

		if count > skip { // 跳过 skip 个
			users = append(users, &kvzoo)
			if len(users) >= int(req.Limit) {
				break // 只返回 limit 个
			}
		}
		
		if kvzoo.LinkStatus == "@@LINK:$" {
			break
		}
		kvKey = []byte(kvzoo.LinkStatus)
	}

	return &types.QueryListByStatusResponse{Users: users}, nil
}
