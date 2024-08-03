package keeper

import (
	"context"
	"strings"
	"errors"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListReplyBySender(goCtx context.Context, req *types.QueryListReplyBySenderRequest) (*types.QueryListReplyBySenderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var replys []*types.Reply

	store := ctx.KVStore(k.storeKey)
	storeLink := prefix.NewStore(store, types.KeyPrefix(types.ReplySenderLinkPrefix))

	// 解析参数: sender|askId
	params := strings.Split(req.Sender, "|")
	if len(params)!=2 { // 参数出错
		return &types.QueryListReplyBySenderResponse{Reply: replys}, nil
	}

	// 找到 sender 的尾部
	linkKey := types.ExchangeLinkKey("replySender", params[0]) // 链表头 key
	kvKey := storeLink.Get(linkKey)
	if kvKey == nil { // 未找到 sender 表头
		return &types.QueryListReplyBySenderResponse{Reply: replys}, nil
	}

	if string(kvKey) == "@@LINK:$" {
		// 尾部就是结束标记，直接返回
		return &types.QueryListReplyBySenderResponse{Reply: replys}, nil
	}


	// 遍历链表
	var count,skip int
	count = 0
	skip = int( (req.Page - 1) * req.Limit )
	for {
		reply, found := k.GetReplyByKey(ctx, kvKey)
		if !found {
			return nil, errors.New("kvKey fault!")
		}

		var askUuid string
		if params[1]!="" { // 参数带 uuid
			// 解析 payload json
			var replyData map[string]string
			if err := json.Unmarshal([]byte(reply.Payload), &replyData); err != nil {  // 解析 json
				return nil, err
			}

			askUuid = replyData["uuid"]
		}

		if params[1]=="" || params[1]==askUuid { // 如果存在 ask_uuid, 进行过滤
			count++

			if count > skip { // 跳过 skip 个
				replys = append(replys, &reply)
				if len(replys) >= int(req.Limit) {
					break // 只返回 limit 个
				}
			}
		}

		if reply.LinkSender == "@@LINK:$" {
			break
		}
		kvKey = []byte(reply.LinkSender)
	}

	return &types.QueryListReplyBySenderResponse{Reply: replys}, nil
}
