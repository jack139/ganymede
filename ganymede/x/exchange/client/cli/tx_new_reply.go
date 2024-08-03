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

func CmdNewReply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-reply [ask-id] [sender] [replier] [payload] [sent-date]",
		Short: "Broadcast message new-reply",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAskId := args[0]
			argSender := args[1]
			argReplier := args[2]
			argPayload := args[3]
			argSentDate := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNewReply(
				clientCtx.GetFromAddress().String(),
				argAskId,
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
