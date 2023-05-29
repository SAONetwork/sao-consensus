package cli

import (
	"fmt"
	"strconv"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdFault() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fault [fault-id]",
		Short: "Query Fault",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			if len(args) != 1 {
				return fmt.Errorf("fault id is required")
			}

			params := &types.QueryFaultRequest{
				FaultId: args[0],
			}

			res, err := queryClient.Fault(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
