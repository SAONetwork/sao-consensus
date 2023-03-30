package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListTimeoutOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-timeout-order",
		Short: "list all TimeoutOrder",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllTimeoutOrderRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.TimeoutOrderAll(context.Background(), params)
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

func CmdShowTimeoutOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-timeout-order [height]",
		Short: "shows a TimeoutOrder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argHeight, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetTimeoutOrderRequest{
				Height: argHeight,
			}

			res, err := queryClient.TimeoutOrder(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
