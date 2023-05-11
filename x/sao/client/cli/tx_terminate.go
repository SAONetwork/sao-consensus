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

func CmdTerminate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "terminate [proposal] [signature] [provider]",
		Short: "Broadcast message terminate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var proposal types.TerminateProposal
			err = json.Unmarshal([]byte(args[0]), &proposal)
			if err != nil {
				return err
			}
			var signature types.JwsSignature
			err = json.Unmarshal([]byte(args[1]), &signature)
			if err != nil {
				return err
			}
			argProvider := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTerminate(
				clientCtx.GetFromAddress().String(),
				proposal,
				signature,
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
