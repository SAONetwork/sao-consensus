package cli

import (
	"strconv"

	"encoding/json"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdUpdateAccountAuths() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-account-auths [did] [update] [remove]",
		Short: "Broadcast message UpdateAccountAuths",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argUpdate := make([]*types.AccountAuth, 0)

			err = json.Unmarshal([]byte("["+args[1]+"]"), &argUpdate)
			if err != nil {
				return err
			}
			argRemove := strings.Split(args[2], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAccountAuths(
				clientCtx.GetFromAddress().String(),
				argDid,
				argUpdate,
				argRemove,
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
