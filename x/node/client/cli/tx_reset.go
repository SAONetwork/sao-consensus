package cli

import (
	"strconv"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdReset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset [peer]",
		Short: "Broadcast message reset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPeer := args[0]

			var status = types.NODE_STATUS_ONLINE
			if len(args) == 2 {
				s, err := strconv.ParseUint(args[1], 10, 32)
				if err != nil {
					return err
				}
				status = uint32(s)

				if status&types.NODE_STATUS_ONLINE == 0 {
					return errors.Wrapf(types.ErrInvalidStatus, "%d", status)
				}
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgReset(
				clientCtx.GetFromAddress().String(),
				argPeer,
				status,
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
