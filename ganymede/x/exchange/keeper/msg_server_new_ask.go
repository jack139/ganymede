package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

func (k msgServer) NewAsk(goCtx context.Context, msg *types.MsgNewAsk) (*types.MsgNewAskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var ask = types.Ask {
		Creator   : msg.Creator,
		Sender    : msg.Sender,
		Replier   : msg.Replier,
		Payload   : msg.Payload,
		SentDate  : msg.SentDate,
	}

	k.AppendAsk(ctx, ask)

	return &types.MsgNewAskResponse{}, nil
}
