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

func CmdAddAccountAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-account-auth [did] [account-auth]",
		Short: "Broadcast message AddAccountAuth",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argAccountAuth := new(types.AccountAuth)
			err = json.Unmarshal([]byte(args[1]), argAccountAuth)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddAccountAuth(
				clientCtx.GetFromAddress().String(),
				argDid,
				argAccountAuth,
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
