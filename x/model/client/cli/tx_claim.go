package cli

import (
    "strconv"
	
	 "strings"
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/SaoNetwork/sao/x/model/types"
)

var _ = strconv.Itoa(0)

func CmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [data]",
		Short: "Broadcast message claim",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argData := strings.Split(args[0], listSeparator)
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaim(
				clientCtx.GetFromAddress().String(),
				argData,
				
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