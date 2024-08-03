package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

func CmdListKvzoo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-kvzoo",
		Short: "list all kvzoo",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllKvzooRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.KvzooAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowKvzoo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-kvzoo [owner] [zoo-key]",
		Short: "shows a kvzoo",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argOwner := args[0]
			argZooKey := args[1]

			params := &types.QueryGetKvzooRequest{
				Owner:  argOwner,
				ZooKey: argZooKey,
			}

			res, err := queryClient.Kvzoo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
