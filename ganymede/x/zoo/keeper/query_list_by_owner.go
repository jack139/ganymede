package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListByOwner(goCtx context.Context, req *types.QueryListByOwnerRequest) (*types.QueryListByOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var kvzoos []*types.Kvzoo
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	storeLink := prefix.NewStore(store, types.KeyPrefix(types.KvzooLinkPrefix))

	// 找到 owner 的尾部
	linkKey := types.KvzooOwnerLinkKey(req.Owner) // 链表头 key
	kvKey := storeLink.Get(linkKey)
	if kvKey == nil { // 未找到
		//return nil, errors.New("cannot find owner tail!")
		return &types.QueryListByOwnerResponse{Kvzoo: kvzoos}, nil
	}

	if string(kvKey) == "@@LINK:$" {
		// 尾部就是结束标记，直接返回
		return &types.QueryListByOwnerResponse{Kvzoo: kvzoos}, nil
	}


	// 遍历链表
	var count,skip int
	count = 0
	skip = int( (req.Page - 1) * req.Limit )
	for {
		kvzoo, found := k.GetKvzooByKey(ctx, kvKey)
		if !found {
			return nil, errors.New("kvKey fault!")
		}

		count++

		if count > skip { // 跳过 skip 个
			kvzoos = append(kvzoos, &kvzoo)
			if len(kvzoos) >= int(req.Limit) {
				break // 只返回 limit 个
			}
		}
		
		if kvzoo.LinkOwner == "@@LINK:$" {
			break
		}
		kvKey = []byte(kvzoo.LinkOwner)
	}

	return &types.QueryListByOwnerResponse{Kvzoo: kvzoos}, nil
}
