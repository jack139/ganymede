package keeper

import (
	"context"
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListSentByReceiver(goCtx context.Context, req *types.QueryListSentByReceiverRequest) (*types.QueryListSentByReceiverResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var sentPosts []*types.SentPost

	store := ctx.KVStore(k.storeKey)
	storeLink := prefix.NewStore(store, types.KeyPrefix(types.SentReceiverLinkPrefix))

	// 解析参数: Receiver|askUuid
	params := strings.Split(req.Receiver, "|")
	if len(params)!=2 { // 参数出错
		return &types.QueryListSentByReceiverResponse{SentPost: sentPosts}, nil
	}

	// 找到 Receiver 的尾部
	linkKey := types.PostofficeLinkKey("sentReceiver", params[0]) // 链表头 key
	kvKey := storeLink.Get(linkKey)
	if kvKey == nil { // 未找到 Receiver 表头
		return &types.QueryListSentByReceiverResponse{SentPost: sentPosts}, nil
	}

	if string(kvKey) == "@@LINK:$" {
		// 尾部就是结束标记，直接返回
		return &types.QueryListSentByReceiverResponse{SentPost: sentPosts}, nil
	}


	// 遍历链表
	var count,skip int
	count = 0
	skip = int( (req.Page - 1) * req.Limit )
	for {
		post, found := k.GetSentPostByKey(ctx, kvKey)
		if !found {
			return nil, errors.New("kvKey fault!")
		}

		var askUuid string
		if params[1]!="" { // 参数带 uuid
			if strings.HasPrefix(post.Title, "@EXCH:ASK:") {
				askUuid = post.Title[10:]
			} else if strings.HasPrefix(post.Title, "@EXCH:RPLY:") {
				params2 := strings.Split(post.Title[11:], "|")
				askUuid = params2[0]
			} else {
				askUuid = post.Title
			}
		}

		if params[1]=="" || params[1]==askUuid { // 如果存在 ask_uuid, 进行过滤
			count++

			if count > skip { // 跳过 skip 个
				sentPosts = append(sentPosts, &post)
				if len(sentPosts) >= int(req.Limit) {
					break // 只返回 limit 个
				}
			}
		}

		if post.LinkReceiver == "@@LINK:$" {
			break
		}
		kvKey = []byte(post.LinkReceiver)
	}

	return &types.QueryListSentByReceiverResponse{SentPost: sentPosts}, nil
}
