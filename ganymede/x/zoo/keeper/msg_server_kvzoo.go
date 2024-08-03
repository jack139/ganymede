package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

func (k msgServer) CreateKvzoo(goCtx context.Context, msg *types.MsgCreateKvzoo) (*types.MsgCreateKvzooResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetKvzoo(
		ctx,
		msg.Owner,
		msg.ZooKey,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var kvzoo = types.Kvzoo{
		Creator:   msg.Creator,
		Owner:     msg.Owner,
		ZooKey:    msg.ZooKey,
		ZooValue:  msg.ZooValue,
		LastDate:  msg.LastDate,
		LinkOwner: msg.LinkOwner,
	}

	k.SetKvzoo(
		ctx,
		kvzoo,
	)
	return &types.MsgCreateKvzooResponse{}, nil
}

func (k msgServer) UpdateKvzoo(goCtx context.Context, msg *types.MsgUpdateKvzoo) (*types.MsgUpdateKvzooResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKvzoo(
		ctx,
		msg.Owner,
		msg.ZooKey,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var kvzoo = types.Kvzoo{
		Creator:   msg.Creator,
		Owner:     msg.Owner,
		ZooKey:    msg.ZooKey,
		ZooValue:  msg.ZooValue,
		LastDate:  msg.LastDate,
		LinkOwner: msg.LinkOwner,
	}

	k.SetKvzooUpdate(ctx, kvzoo, valFound.Owner, valFound.LinkOwner)

	return &types.MsgUpdateKvzooResponse{}, nil
}

func (k msgServer) DeleteKvzoo(goCtx context.Context, msg *types.MsgDeleteKvzoo) (*types.MsgDeleteKvzooResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetKvzoo(
		ctx,
		msg.Owner,
		msg.ZooKey,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveKvzoo(
		ctx,
		msg.Owner,
		msg.ZooKey,
		valFound.LinkOwner,
	)

	return &types.MsgDeleteKvzooResponse{}, nil
}
