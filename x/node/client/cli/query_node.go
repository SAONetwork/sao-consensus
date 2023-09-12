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

			resNodes := res.GetNode()

			resTextNodes := make([]types.NodeInText, 0)

			for _, node := range resNodes {

				textNode := types.NodeInText{
					Creator:         node.Creator,
					Peer:            node.Peer,
					Reputation:      node.Reputation,
					Status:          NodeStatusToText(node.Status),
					LastAliveHeight: node.LastAliveHeight,
					TxAddresses:     node.TxAddresses,
					Role:            NodeRoleToText(node.Role),
					Description:     node.Description,
					Validator:       node.Validator,
				}
				resTextNodes = append(resTextNodes, textNode)
			}

			resText := types.QueryAllNodeInTextResponse{
				Node:       resTextNodes,
				Pagination: res.Pagination,
			}

			return clientCtx.PrintProto(&resText)
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

			node := res.GetNode()

			textNode := types.NodeInText{
				Creator:         node.Creator,
				Peer:            node.Peer,
				Reputation:      node.Reputation,
				Status:          NodeStatusToText(node.Status),
				LastAliveHeight: node.LastAliveHeight,
				TxAddresses:     node.TxAddresses,
				Role:            NodeRoleToText(node.Role),
				Description:     node.Description,
				Validator:       node.Validator,
			}

			textRes := types.QueryGetNodeInTextResponse{
				NodeInText: textNode,
			}

			return clientCtx.PrintProto(&textRes)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
