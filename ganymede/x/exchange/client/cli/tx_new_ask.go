package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

var _ = strconv.Itoa(0)

func CmdNewAsk() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-ask [sender] [replier] [payload] [sent-date]",
		Short: "Broadcast message new-ask",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSender := args[0]
			argReplier := args[1]
			argPayload := args[2]
			argSentDate := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNewAsk(
				clientCtx.GetFromAddress().String(),
				argSender,
				argReplier,
				argPayload,
				argSentDate,
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
