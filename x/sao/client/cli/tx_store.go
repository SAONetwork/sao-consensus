package cli

import (
	"strconv"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [cid] [provider] [duration] [replica]",
		Short: "Broadcast message store",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCid := args[0]
			argProvider := args[1]
			argDuration, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}
			argReplica, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStore(
				clientCtx.GetFromAddress().String(),
				argCid,
				argProvider,
				argDuration,
				argReplica,
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
