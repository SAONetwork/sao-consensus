package cli

import (
	"encoding/json"
	"strconv"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Update [did] [new-doc-id] [keys] [timestamp] [update-account-auth] [remove-account-did] [past-seed]",
		Short: "Broadcast message Update",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argNewDocId := args[1]

			argKeys := make([]*types.PubKey, 0)
			err = json.Unmarshal([]byte("["+args[2]+"]"), &argKeys)
			if err != nil {
				return err
			}

			argTimestamp, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			argUpdateAccountAuth := make([]*types.AccountAuth, 0)
			err = json.Unmarshal([]byte("["+args[4]+"]"), &argUpdateAccountAuth)
			if err != nil {
				return err
			}

			argRemoveAccountDid := strings.Split(args[5], listSeparator)
			argPastSeed := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdate(
				clientCtx.GetFromAddress().String(),
				argDid,
				argNewDocId,
				argKeys,
				argTimestamp,
				argUpdateAccountAuth,
				argRemoveAccountDid,
				argPastSeed,
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
