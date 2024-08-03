package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

func (k msgServer) NewReply(goCtx context.Context, msg *types.MsgNewReply) (*types.MsgNewReplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var reply = types.Reply {
		Creator   : msg.Creator,
		AskId     : msg.AskId,
		Sender    : msg.Sender,
		Replier   : msg.Replier,
		Payload   : msg.Payload,
		SentDate  : msg.SentDate,
	}

	k.AppendReply(ctx, reply)

	return &types.MsgNewReplyResponse{}, nil
}
