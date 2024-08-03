package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
)

var _ = strconv.Itoa(0)

func CmdListTimeoutByReceiver() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-timeout-by-receiver [receiver] [page] [limit]",
		Short: "Query list-timeout-by-receiver",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReceiver := args[0]
			reqPage, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}
			reqLimit, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryListTimeoutByReceiverRequest{

				Receiver: reqReceiver,
				Page:     reqPage,
				Limit:    reqLimit,
			}

			res, err := queryClient.ListTimeoutByReceiver(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}