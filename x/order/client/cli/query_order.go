package cli

import (
	"context"
	"strconv"

	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-order",
		Short: "list all order",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			did, err := cmd.Flags().GetString("did")
			if err != nil {
				return err
			}
			status, err := cmd.Flags().GetInt32("status")
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrderRequest{
				Pagination: pageReq,
				States:     []int32{status},
				Did:        did,
			}

			res, err := queryClient.OrderAll(context.Background(), params)
			if err != nil {
				return err
			}

			orders := res.GetOrder()

			textOrders := make([]types.OrderInText, 0)

			for _, order := range orders {
				textOrder := types.OrderInText{
					Creator:   order.Creator,
					Owner:     order.Owner,
					Id:        order.Id,
					Provider:  order.Provider,
					Cid:       order.Cid,
					Duration:  order.Duration,
					Status:    OrderStatusInText(order.Status),
					Replica:   order.Replica,
					Shards:    order.Shards,
					Amount:    order.Amount,
					Size_:     order.Size_,
					Operation: order.Operation,
					CreatedAt: order.CreatedAt,
					Timeout:   order.Timeout,
					DataId:    order.DataId,
					Commit:    order.Commit,
					UnitPrice: order.UnitPrice,
				}
				textOrders = append(textOrders, textOrder)
			}

			resText := types.QueryAllOrderInTextResponse{
				Order:      textOrders,
				Pagination: res.Pagination,
			}

			return clientCtx.PrintProto(&resText)
		},
	}

	cmd.Flags().String("did", "", "did")
	cmd.Flags().Int32("status", 0, "status flags")
	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-order [id]",
		Short: "shows a order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetOrderRequest{
				Id: id,
			}

			res, err := queryClient.Order(context.Background(), params)
			if err != nil {
				return err
			}

			order := res.GetOrder()

			textShards := make(map[string]*types.ShardInText, 0)

			for _, shard := range order.Shards {
				textShard := types.ShardInText{
					Id: shard.Id,
				}
				textShards[shard.Sp] = &textShard
			}

			textOrder := types.FullOrderInText{
				Creator:   order.Creator,
				Owner:     order.Owner,
				Id:        order.Id,
				Provider:  order.Provider,
				Cid:       order.Cid,
				Duration:  order.Duration,
				Status:    OrderStatusInText(order.Status),
				Replica:   order.Replica,
				Shards:    textShards,
				Amount:    order.Amount,
				Size_:     order.Size_,
				Operation: order.Operation,
				CreatedAt: order.CreatedAt,
				Timeout:   order.Timeout,
				DataId:    order.DataId,
				Commit:    order.Commit,
				UnitPrice: order.UnitPrice,
			}

			resText := types.QueryGetOrderInTextResponse{
				Order: textOrder,
			}

			return clientCtx.PrintProto(&resText)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
