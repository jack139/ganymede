package keeper

import (
	"log"
	"context"

	"github.com/cosmos/cosmos-sdk/codec"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

func (k msgServer) CreateUsers(goCtx context.Context, msg *types.MsgCreateUsers) (*types.MsgCreateUsersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetUsers(
		ctx,
		msg.ChainAddr,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	// msg_add_account 里的操作： 新建 auth/account, 预存点stake
	ak := k.AccountKeeper
	bk := k.BankKeeper

	// 从消息参数中反序列化 Pubkey
	var pk1 secp256k1.PubKey
	aminoCdc := codec.NewLegacyAmino()
	err := aminoCdc.UnmarshalJSON([]byte(msg.Address), &pk1)
	if err!=nil {
		return nil, err
	}

	// 公钥
	pk := cryptotypes.PubKey(&pk1)
	log.Println(pk)

	// 地址
	addr := sdk.AccAddress(pk.Address()) // 等效 pk.Address().Bytes()

	// 添加 acount， 带公钥信息
	baseAccount := authtypes.NewBaseAccount(addr, pk, 0, 0)
	acc1 := ak.NewAccount(ctx, baseAccount)
	ak.SetAccount(ctx, acc1)


	// 挖 1credit, 转给新用户
	feeAmount := sdk.NewCoins(sdk.NewInt64Coin("credit", 1))

	err = bk.MintCoins(ctx, minttypes.ModuleName, feeAmount)
	if err!=nil {
		return nil, err
	}
	err = bk.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, feeAmount)
	if err!=nil {
		return nil, err
	}

	// 存储到 users
	var users = types.Users{
		Creator:    msg.Creator,
		ChainAddr:  msg.ChainAddr,
		KeyName:    msg.KeyName,
		UserType:   msg.UserType,
		Name:       msg.Name,
		Address:    "", // msg.Address 里是 serializedPubkey, 这里不保存
		Phone:      msg.Phone,
		AccountNo:  msg.AccountNo,
		Ref:        msg.Ref,
		RegDate:    msg.RegDate,
		Status:     msg.Status,
		LastDate:   msg.LastDate,
		LinkStatus: msg.LinkStatus,
	}

	k.SetUsers(
		ctx,
		users,
	)
	return &types.MsgCreateUsersResponse{}, nil
}

func (k msgServer) UpdateUsers(goCtx context.Context, msg *types.MsgUpdateUsers) (*types.MsgUpdateUsersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUsers(
		ctx,
		msg.ChainAddr,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var users = types.Users{
		Creator:    msg.Creator,
		ChainAddr:  msg.ChainAddr,
		KeyName:    msg.KeyName,
		UserType:   msg.UserType,
		Name:       msg.Name,
		Address:    msg.Address,
		Phone:      msg.Phone,
		AccountNo:  msg.AccountNo,
		Ref:        msg.Ref,
		RegDate:    msg.RegDate,
		Status:     msg.Status,
		LastDate:   msg.LastDate,
		LinkStatus: msg.LinkStatus,
	}

	k.SetUsersUpdate(ctx, users, valFound.Status, valFound.LinkStatus)

	return &types.MsgUpdateUsersResponse{}, nil
}

func (k msgServer) DeleteUsers(goCtx context.Context, msg *types.MsgDeleteUsers) (*types.MsgDeleteUsersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUsers(
		ctx,
		msg.ChainAddr,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUsers(
		ctx,
		msg.ChainAddr,
	)

	return &types.MsgDeleteUsersResponse{}, nil
}
