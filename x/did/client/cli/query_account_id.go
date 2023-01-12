package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListAccountId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-account-id",
		Short: "list all accountId",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAccountIdRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AccountIdAll(context.Background(), params)
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

func CmdShowAccountId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-account-id [account-did]",
		Short: "shows a accountId",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAccountDid := args[0]

			params := &types.QueryGetAccountIdRequest{
				AccountDid: argAccountDid,
			}

			res, err := queryClient.AccountId(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
