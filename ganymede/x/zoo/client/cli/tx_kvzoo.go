package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

func CmdCreateKvzoo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-kvzoo [owner] [zoo-key] [zoo-value] [last-date] [link-owner]",
		Short: "Create a new kvzoo",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexOwner := args[0]
			indexZooKey := args[1]

			// Get value arguments
			argZooValue := args[2]
			argLastDate := args[3]
			argLinkOwner := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateKvzoo(
				clientCtx.GetFromAddress().String(),
				indexOwner,
				indexZooKey,
				argZooValue,
				argLastDate,
				argLinkOwner,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateKvzoo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-kvzoo [owner] [zoo-key] [zoo-value] [last-date] [link-owner]",
		Short: "Update a kvzoo",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexOwner := args[0]
			indexZooKey := args[1]

			// Get value arguments
			argZooValue := args[2]
			argLastDate := args[3]
			argLinkOwner := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateKvzoo(
				clientCtx.GetFromAddress().String(),
				indexOwner,
				indexZooKey,
				argZooValue,
				argLastDate,
				argLinkOwner,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteKvzoo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-kvzoo [owner] [zoo-key]",
		Short: "Delete a kvzoo",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexOwner := args[0]
			indexZooKey := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteKvzoo(
				clientCtx.GetFromAddress().String(),
				indexOwner,
				indexZooKey,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
