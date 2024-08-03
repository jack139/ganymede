package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jack139/ganymede/ganymede/testutil/network"
	"github.com/jack139/ganymede/ganymede/testutil/nullify"
	"github.com/jack139/ganymede/ganymede/x/zoo/client/cli"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithKvzooObjects(t *testing.T, n int) (*network.Network, []types.Kvzoo) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	for i := 0; i < n; i++ {
		kvzoo := types.Kvzoo{
			Owner:  strconv.Itoa(i),
			ZooKey: strconv.Itoa(i),
		}
		nullify.Fill(&kvzoo)
		state.KvzooList = append(state.KvzooList, kvzoo)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.KvzooList
}

func TestShowKvzoo(t *testing.T) {
	net, objs := networkWithKvzooObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc     string
		idOwner  string
		idZooKey string

		args []string
		err  error
		obj  types.Kvzoo
	}{
		{
			desc:     "found",
			idOwner:  objs[0].Owner,
			idZooKey: objs[0].ZooKey,

			args: common,
			obj:  objs[0],
		},
		{
			desc:     "not found",
			idOwner:  strconv.Itoa(100000),
			idZooKey: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idOwner,
				tc.idZooKey,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowKvzoo(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetKvzooResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Kvzoo)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Kvzoo),
				)
			}
		})
	}
}

func TestListKvzoo(t *testing.T) {
	net, objs := networkWithKvzooObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListKvzoo(), args)
			require.NoError(t, err)
			var resp types.QueryAllKvzooResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Kvzoo), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Kvzoo),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListKvzoo(), args)
			require.NoError(t, err)
			var resp types.QueryAllKvzooResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Kvzoo), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Kvzoo),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListKvzoo(), args)
		require.NoError(t, err)
		var resp types.QueryAllKvzooResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Kvzoo),
		)
	})
}
