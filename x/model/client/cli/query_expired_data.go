package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListExpiredData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-expired-data",
		Short: "list all expired_data",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllExpiredDataRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ExpiredDataAll(context.Background(), params)
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

func CmdShowExpiredData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-expired-data [height]",
		Short: "shows a expired_data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argHeight, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetExpiredDataRequest{
				Height: argHeight,
			}

			res, err := queryClient.ExpiredData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
