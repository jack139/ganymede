package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewReply = "new_reply"

var _ sdk.Msg = &MsgNewReply{}

func NewMsgNewReply(creator string, askId string, sender string, replier string, payload string, sentDate string) *MsgNewReply {
	return &MsgNewReply{
		Creator:  creator,
		AskId:    askId,
		Sender:   sender,
		Replier:  replier,
		Payload:  payload,
		SentDate: sentDate,
	}
}

func (msg *MsgNewReply) Route() string {
	return RouterKey
}

func (msg *MsgNewReply) Type() string {
	return TypeMsgNewReply
}

func (msg *MsgNewReply) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewReply) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewReply) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
