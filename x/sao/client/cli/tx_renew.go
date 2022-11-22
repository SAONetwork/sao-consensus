package cli

import (
	"strconv"

	"strings"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRenew() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew [data] [duration] [timeout]",
		Short: "Broadcast message renew",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			listSeparator := ","
			argData := strings.Split(args[0], listSeparator)
			argDuration, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argTimeout, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRenew(
				clientCtx.GetFromAddress().String(),
				argData,
				argDuration,
				argTimeout,
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
