package cli

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func CmdBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binding [account-id] [root-doc-id] [keys] [account-auth] [proof]",
		Short: "Broadcast message Binding",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccountId := args[0]
			argRootDocId := args[1]
			argKeys := strings.Split(args[2], listSeparator)
			keys := make([]*types.PubKey, 0)
			for _, argKey := range argKeys {
				kvpair := strings.Split(argKey, ":")
				if len(kvpair) != 2 {
					return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "invalid public key pairs (%s)", argKeys)
				}
				keys = append(keys, &types.PubKey{
					Name:  kvpair[0],
					Value: kvpair[1],
				})

			}
			argAccountAuth := new(types.AccountAuth)
			err = json.Unmarshal([]byte(args[3]), argAccountAuth)
			if err != nil {
				return err
			}
			argProof := new(types.BindingProof)
			err = json.Unmarshal([]byte(args[4]), argProof)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgBinding(
				clientCtx.GetFromAddress().String(),
				argAccountId,
				argRootDocId,
				keys,
				argAccountAuth,
				argProof,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	//cmd.Flags().StringSlice("public-key", []string{}, "set sid public key map, eg --public-key=<> ")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
