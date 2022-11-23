package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListSidDocumentVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-sid-document-version",
		Short: "list all SidDocumentVersion",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSidDocumentVersionRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SidDocumentVersionAll(context.Background(), params)
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

func CmdShowSidDocumentVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-sid-document-version [doc-id]",
		Short: "shows a SidDocumentVersion",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDocId := args[0]

			params := &types.QueryGetSidDocumentVersionRequest{
				DocId: argDocId,
			}

			res, err := queryClient.SidDocumentVersion(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
