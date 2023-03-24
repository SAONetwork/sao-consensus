package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListOrderFinish() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-order-finish",
		Short: "list all OrderFinish",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrderFinishRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrderFinishAll(context.Background(), params)
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

func CmdShowOrderFinish() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-order-finish [timestamp]",
		Short: "shows a OrderFinish",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argTimestamp, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetOrderFinishRequest{
				Timestamp: argTimestamp,
			}

			res, err := queryClient.OrderFinish(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
