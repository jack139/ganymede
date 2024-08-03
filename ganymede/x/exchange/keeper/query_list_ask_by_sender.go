package keeper

import (
	"context"
	"errors"
	"strings"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAskBySender(goCtx context.Context, req *types.QueryListAskBySenderRequest) (*types.QueryListAskBySenderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var asks []*types.Ask

	store := ctx.KVStore(k.storeKey)
	storeLink := prefix.NewStore(store, types.KeyPrefix(types.AskSenderLinkPrefix))

	// 解析参数: sender|askId
	params := strings.Split(req.Sender, "|")
	if len(params)!=2 { // 参数出错
		return &types.QueryListAskBySenderResponse{Ask: asks}, nil
	}

	// 找到 sender 的尾部
	linkKey := types.ExchangeLinkKey("sender", params[0]) // 链表头 key
	kvKey := storeLink.Get(linkKey)
	if kvKey == nil { // 未找到 sender 表头
		return &types.QueryListAskBySenderResponse{Ask: asks}, nil
	}

	if string(kvKey) == "@@LINK:$" {
		// 尾部就是结束标记，直接返回
		return &types.QueryListAskBySenderResponse{Ask: asks}, nil
	}


	// 遍历链表
	var count,skip int
	count = 0
	skip = int( (req.Page - 1) * req.Limit )
	for {
		ask, found := k.GetAskByKey(ctx, kvKey)
		if !found {
			return nil, errors.New("kvKey fault!")
		}

		var askUuid string
		if params[1]!="" { // 参数带 uuid
			// 解析 payload json
			var askData map[string]string
			if err := json.Unmarshal([]byte(ask.Payload), &askData); err != nil {  // 解析 json
				return nil, err
			}

			askUuid = askData["uuid"]
		}

		if params[1]=="" || params[1]==askUuid { // 如果存在 ask_uuid, 进行过滤
			count++

			if count > skip { // 跳过 skip 个
				asks = append(asks, &ask)
				if len(asks) >= int(req.Limit) {
					break // 只返回 limit 个
				}
			}

		}
		
		if ask.LinkSender == "@@LINK:$" {
			break
		}
		kvKey = []byte(ask.LinkSender)
	}

	return &types.QueryListAskBySenderResponse{Ask: asks}, nil
}
