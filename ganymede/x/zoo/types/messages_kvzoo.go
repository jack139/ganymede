package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateKvzoo = "create_kvzoo"
	TypeMsgUpdateKvzoo = "update_kvzoo"
	TypeMsgDeleteKvzoo = "delete_kvzoo"
)

var _ sdk.Msg = &MsgCreateKvzoo{}

func NewMsgCreateKvzoo(
	creator string,
	owner string,
	zooKey string,
	zooValue string,
	lastDate string,
	linkOwner string,

) *MsgCreateKvzoo {
	return &MsgCreateKvzoo{
		Creator:   creator,
		Owner:     owner,
		ZooKey:    zooKey,
		ZooValue:  zooValue,
		LastDate:  lastDate,
		LinkOwner: linkOwner,
	}
}

func (msg *MsgCreateKvzoo) Route() string {
	return RouterKey
}

func (msg *MsgCreateKvzoo) Type() string {
	return TypeMsgCreateKvzoo
}

func (msg *MsgCreateKvzoo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateKvzoo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateKvzoo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateKvzoo{}

func NewMsgUpdateKvzoo(
	creator string,
	owner string,
	zooKey string,
	zooValue string,
	lastDate string,
	linkOwner string,

) *MsgUpdateKvzoo {
	return &MsgUpdateKvzoo{
		Creator:   creator,
		Owner:     owner,
		ZooKey:    zooKey,
		ZooValue:  zooValue,
		LastDate:  lastDate,
		LinkOwner: linkOwner,
	}
}

func (msg *MsgUpdateKvzoo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateKvzoo) Type() string {
	return TypeMsgUpdateKvzoo
}

func (msg *MsgUpdateKvzoo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateKvzoo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateKvzoo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteKvzoo{}

func NewMsgDeleteKvzoo(
	creator string,
	owner string,
	zooKey string,

) *MsgDeleteKvzoo {
	return &MsgDeleteKvzoo{
		Creator: creator,
		Owner:   owner,
		ZooKey:  zooKey,
	}
}
func (msg *MsgDeleteKvzoo) Route() string {
	return RouterKey
}

func (msg *MsgDeleteKvzoo) Type() string {
	return TypeMsgDeleteKvzoo
}

func (msg *MsgDeleteKvzoo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteKvzoo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteKvzoo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
