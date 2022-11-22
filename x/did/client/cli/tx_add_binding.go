package cli

import (
	"strconv"

	"encoding/json"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-binding [account-id] [proof]",
		Short: "Broadcast message AddBinding",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccountId := args[0]
			argProof := new(types.BindingProof)
			err = json.Unmarshal([]byte(args[1]), argProof)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddBinding(
				clientCtx.GetFromAddress().String(),
				argAccountId,
				argProof,
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
