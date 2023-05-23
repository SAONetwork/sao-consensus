package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListPledgeDebt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-pledge-debt",
		Short: "list all PledgeDebt",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPledgeDebtRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PledgeDebtAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPledgeDebt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-pledge-debt [sp]",
		Short: "shows a PledgeDebt",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argSp := args[0]

			params := &types.QueryGetPledgeDebtRequest{
				Sp: argSp,
			}

			res, err := queryClient.PledgeDebt(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
