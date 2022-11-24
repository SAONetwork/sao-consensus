package cli

import (
	"strconv"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateSidDocument() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-sid-document [signing-key] [encrypt-key] [root-doc-id]",
		Short: "Broadcast message UpdateSidDocument",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSigningKey := args[0]
			argEncryptKey := args[1]
			argRootDocId := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSidDocument(
				clientCtx.GetFromAddress().String(),
				argSigningKey,
				argEncryptKey,
				argRootDocId,
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