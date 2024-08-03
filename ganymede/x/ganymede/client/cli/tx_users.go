package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

func CmdCreateUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-users [chain-addr] [key-name] [user-type] [name] [address] [phone] [account-no] [ref] [reg-date] [status] [last-date] [link-status]",
		Short: "Create a new users",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChainAddr := args[0]

			// Get value arguments
			argKeyName := args[1]
			argUserType := args[2]
			argName := args[3]
			argAddress := args[4]
			argPhone := args[5]
			argAccountNo := args[6]
			argRef := args[7]
			argRegDate := args[8]
			argStatus := args[9]
			argLastDate := args[10]
			argLinkStatus := args[11]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUsers(
				clientCtx.GetFromAddress().String(),
				indexChainAddr,
				argKeyName,
				argUserType,
				argName,
				argAddress,
				argPhone,
				argAccountNo,
				argRef,
				argRegDate,
				argStatus,
				argLastDate,
				argLinkStatus,
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

func CmdUpdateUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-users [chain-addr] [key-name] [user-type] [name] [address] [phone] [account-no] [ref] [reg-date] [status] [last-date] [link-status]",
		Short: "Update a users",
		Args:  cobra.ExactArgs(12),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexChainAddr := args[0]

			// Get value arguments
			argKeyName := args[1]
			argUserType := args[2]
			argName := args[3]
			argAddress := args[4]
			argPhone := args[5]
			argAccountNo := args[6]
			argRef := args[7]
			argRegDate := args[8]
			argStatus := args[9]
			argLastDate := args[10]
			argLinkStatus := args[11]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUsers(
				clientCtx.GetFromAddress().String(),
				indexChainAddr,
				argKeyName,
				argUserType,
				argName,
				argAddress,
				argPhone,
				argAccountNo,
				argRef,
				argRegDate,
				argStatus,
				argLastDate,
				argLinkStatus,
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

func CmdDeleteUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-users [chain-addr]",
		Short: "Delete a users",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexChainAddr := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteUsers(
				clientCtx.GetFromAddress().String(),
				indexChainAddr,
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
