package cli

import (
	"context"
	"strconv"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-node [status]",
		Short: "list all node",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			var status = types.NODE_STATUS_NA
			if len(args) == 1 {
				s, err := strconv.ParseUint(args[0], 10, 32)
				if err != nil {
					return err
				}
				status = uint32(s)
			}

			params := &types.QueryAllNodeRequest{
				Pagination: pageReq,
				Status:     status,
			}

			res, err := queryClient.NodeAll(context.Background(), params)
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

func CmdShowNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-node [creator]",
		Short: "shows a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argCreator := args[0]

			params := &types.QueryGetNodeRequest{
				Creator: argCreator,
			}

			res, err := queryClient.Node(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
