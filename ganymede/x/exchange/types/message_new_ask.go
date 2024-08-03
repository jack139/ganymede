package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewAsk = "new_ask"

var _ sdk.Msg = &MsgNewAsk{}

func NewMsgNewAsk(creator string, sender string, replier string, payload string, sentDate string) *MsgNewAsk {
	return &MsgNewAsk{
		Creator:  creator,
		Sender:   sender,
		Replier:  replier,
		Payload:  payload,
		SentDate: sentDate,
	}
}

func (msg *MsgNewAsk) Route() string {
	return RouterKey
}

func (msg *MsgNewAsk) Type() string {
	return TypeMsgNewAsk
}

func (msg *MsgNewAsk) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewAsk) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewAsk) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
