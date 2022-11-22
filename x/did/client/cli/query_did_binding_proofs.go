package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListDidBindingProofs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-did-binding-proofs",
		Short: "list all DidBindingProofs",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDidBindingProofsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DidBindingProofsAll(context.Background(), params)
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

func CmdShowDidBindingProofs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-did-binding-proofs [account-id]",
		Short: "shows a DidBindingProofs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAccountId := args[0]

			params := &types.QueryGetDidBindingProofsRequest{
				AccountId: argAccountId,
			}

			res, err := queryClient.DidBindingProofs(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
