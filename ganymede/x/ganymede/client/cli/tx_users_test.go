package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/jack139/ganymede/ganymede/testutil/network"
	"github.com/jack139/ganymede/ganymede/x/ganymede/client/cli"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCreateUsers(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz"}
	tests := []struct {
		desc        string
		idChainAddr string

		args []string
		err  error
		code uint32
	}{
		{
			idChainAddr: strconv.Itoa(0),

			desc: "valid",
			args: []string{
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
				fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			args := []string{
				tc.idChainAddr,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateUsers(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}

func TestUpdateUsers(t *testing.T) {
	net := network.New(t)
	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{
		"0",
	}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateUsers(), args)
	require.NoError(t, err)

	tests := []struct {
		desc        string
		idChainAddr string

		args []string
		code uint32
		err  error
	}{
		{
			desc:        "valid",
			idChainAddr: strconv.Itoa(0),

			args: common,
		},
		{
			desc:        "key not found",
			idChainAddr: strconv.Itoa(100000),

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			args := []string{
				tc.idChainAddr,
			}
			args = append(args, fields...)
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdUpdateUsers(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}

func TestDeleteUsers(t *testing.T) {
	net := network.New(t)

	val := net.Validators[0]
	ctx := val.ClientCtx

	fields := []string{"xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz", "xyz"}
	common := []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(net.Config.BondDenom, sdkmath.NewInt(10))).String()),
	}
	args := []string{
		"0",
	}
	args = append(args, fields...)
	args = append(args, common...)
	_, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdCreateUsers(), args)
	require.NoError(t, err)

	tests := []struct {
		desc        string
		idChainAddr string

		args []string
		code uint32
		err  error
	}{
		{
			desc:        "valid",
			idChainAddr: strconv.Itoa(0),

			args: common,
		},
		{
			desc:        "key not found",
			idChainAddr: strconv.Itoa(100000),

			args: common,
			code: sdkerrors.ErrKeyNotFound.ABCICode(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			require.NoError(t, net.WaitForNextBlock())

			args := []string{
				tc.idChainAddr,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdDeleteUsers(), args)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			var resp sdk.TxResponse
			require.NoError(t, ctx.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NoError(t, clitestutil.CheckTxCode(net, ctx, resp.TxHash, tc.code))
		})
	}
}
