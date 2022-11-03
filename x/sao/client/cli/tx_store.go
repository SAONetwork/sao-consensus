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
		Use:   "store [owner] [cid] [provider] [duration] [replica] [metadata]",
		Short: "Broadcast message store",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOwner := args[0]

			argCid := args[1]
			argProvider := args[2]
			argDuration, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argReplica, err := cast.ToInt32E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argMetadata := args[5]

			msg := types.NewMsgStore(
				clientCtx.GetFromAddress().String(),
				argOwner,
				argCid,
				argProvider,
				argDuration,
				argReplica,
				argMetadata,
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
