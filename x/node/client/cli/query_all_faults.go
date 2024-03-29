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

func CmdAllFaults() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-faults [sp] [shard]",
		Short: "Query AllFaults",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			if len(args) != 2 {
				return fmt.Errorf("sp address and shard id are required")
			}

			params := &types.QueryAllFaultsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllFaults(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
