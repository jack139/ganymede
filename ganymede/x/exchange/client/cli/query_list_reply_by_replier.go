package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

var _ = strconv.Itoa(0)

func CmdListReplyByReplier() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-reply-by-replier [replier] [page] [limit]",
		Short: "Query list-reply-by-replier",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReplier := args[0]
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

			params := &types.QueryListReplyByReplierRequest{

				Replier: reqReplier,
				Page:    reqPage,
				Limit:   reqLimit,
			}

			res, err := queryClient.ListReplyByReplier(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
