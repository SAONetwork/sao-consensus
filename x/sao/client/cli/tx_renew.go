package cli

import (
	"encoding/json"
	"strconv"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRenew() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renew [proposal] [signature] [provider]",
		Short: "Broadcast message renew",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProposal := new(types.RenewProposal)
			err = json.Unmarshal([]byte(args[0]), argProposal)
			if err != nil {
				return err
			}
			argSignature := new(types.JwsSignature)
			err = json.Unmarshal([]byte(args[1]), argSignature)
			if err != nil {
				return err
			}
			argProvider := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRenew(
				clientCtx.GetFromAddress().String(),
				argProposal,
				argSignature,
				argProvider,
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
