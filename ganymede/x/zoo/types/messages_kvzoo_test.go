package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/jack139/ganymede/ganymede/testutil/sample"
)

func TestMsgCreateKvzoo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateKvzoo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateKvzoo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateKvzoo{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateKvzoo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKvzoo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateKvzoo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateKvzoo{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteKvzoo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteKvzoo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteKvzoo{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteKvzoo{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
