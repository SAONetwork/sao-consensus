package cli

import (
	"strconv"

	"encoding/json"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdStore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store [proposal] [signature]",
		Short: "Broadcast message store",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProposal := new(types.Proposal)
			err = json.Unmarshal([]byte(args[0]), argProposal)
			if err != nil {
				return err
			}
			argSignature := new(types.JwsSignature)
			err = json.Unmarshal([]byte(args[1]), argSignature)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStore(
				clientCtx.GetFromAddress().String(),
				argProposal,
				argSignature,
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
