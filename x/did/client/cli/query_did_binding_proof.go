package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListDidBindingProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-did-binding-proofs",
		Short: "list all DidBindingProof",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDidBindingProofRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DidBindingProofAll(context.Background(), params)
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

func CmdShowDidBindingProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-did-binding-proofs [account-id]",
		Short: "shows a DidBindingProof",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAccountId := args[0]

			params := &types.QueryGetDidBindingProofRequest{
				AccountId: argAccountId,
			}

			res, err := queryClient.DidBindingProof(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
