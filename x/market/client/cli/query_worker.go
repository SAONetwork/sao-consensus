package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/market/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListWorker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-worker",
		Short: "list all worker",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllWorkerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.WorkerAll(context.Background(), params)
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

func CmdShowWorker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-worker [workername]",
		Short: "shows a worker",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argWorkername := args[0]

			params := &types.QueryGetWorkerRequest{
				Workername: argWorkername,
			}

			res, err := queryClient.Worker(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
