package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateUsers = "create_users"
	TypeMsgUpdateUsers = "update_users"
	TypeMsgDeleteUsers = "delete_users"
)

var _ sdk.Msg = &MsgCreateUsers{}

func NewMsgCreateUsers(
	creator string,
	chainAddr string,
	keyName string,
	userType string,
	name string,
	address string,
	phone string,
	accountNo string,
	ref string,
	regDate string,
	status string,
	lastDate string,
	linkStatus string,

) *MsgCreateUsers {
	return &MsgCreateUsers{
		Creator:    creator,
		ChainAddr:  chainAddr,
		KeyName:    keyName,
		UserType:   userType,
		Name:       name,
		Address:    address,
		Phone:      phone,
		AccountNo:  accountNo,
		Ref:        ref,
		RegDate:    regDate,
		Status:     status,
		LastDate:   lastDate,
		LinkStatus: linkStatus,
	}
}

func (msg *MsgCreateUsers) Route() string {
	return RouterKey
}

func (msg *MsgCreateUsers) Type() string {
	return TypeMsgCreateUsers
}

func (msg *MsgCreateUsers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUsers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUsers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUsers{}

func NewMsgUpdateUsers(
	creator string,
	chainAddr string,
	keyName string,
	userType string,
	name string,
	address string,
	phone string,
	accountNo string,
	ref string,
	regDate string,
	status string,
	lastDate string,
	linkStatus string,

) *MsgUpdateUsers {
	return &MsgUpdateUsers{
		Creator:    creator,
		ChainAddr:  chainAddr,
		KeyName:    keyName,
		UserType:   userType,
		Name:       name,
		Address:    address,
		Phone:      phone,
		AccountNo:  accountNo,
		Ref:        ref,
		RegDate:    regDate,
		Status:     status,
		LastDate:   lastDate,
		LinkStatus: linkStatus,
	}
}

func (msg *MsgUpdateUsers) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUsers) Type() string {
	return TypeMsgUpdateUsers
}

func (msg *MsgUpdateUsers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUsers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUsers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUsers{}

func NewMsgDeleteUsers(
	creator string,
	chainAddr string,

) *MsgDeleteUsers {
	return &MsgDeleteUsers{
		Creator:   creator,
		ChainAddr: chainAddr,
	}
}
func (msg *MsgDeleteUsers) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUsers) Type() string {
	return TypeMsgDeleteUsers
}

func (msg *MsgDeleteUsers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUsers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUsers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
