package cli

import (
	"context"
	"strconv"

	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListShard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-shard",
		Short: "list all shard",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllShardRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ShardAll(context.Background(), params)
			if err != nil {
				return err
			}
			/*

				Id         uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
				OrderId    uint64      `protobuf:"varint,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
				Status     int32       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
				Size_      uint64      `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
				Cid        string      `protobuf:"bytes,5,opt,name=cid,proto3" json:"cid,omitempty"`
				Pledge     types.Coin  `protobuf:"bytes,6,opt,name=pledge,proto3" json:"pledge"`
				From       string      `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
				Sp         string      `protobuf:"bytes,8,opt,name=sp,proto3" json:"sp,omitempty"`
				Duration   uint64      `protobuf:"varint,9,opt,name=duration,proto3" json:"duration,omitempty"`
				CreatedAt  uint64      `protobuf:"varint,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
				RenewInfos []RenewInfo `protobuf:"bytes,11,rep,name=renewInfos,proto3" json:"renewInfos"`
			*/
			shards := res.GetShard()

			textShards := make([]types.ShardInText, 0)
			for _, shard := range shards {
				textShard := types.ShardInText{
					Id:         shard.Id,
					OrderId:    shard.OrderId,
					Status:     ShardStatusInText(int(shard.Status)),
					Size_:      shard.Size_,
					Cid:        shard.Cid,
					Pledge:     shard.Pledge,
					From:       shard.From,
					Sp:         shard.Sp,
					Duration:   shard.Duration,
					CreatedAt:  shard.CreatedAt,
					RenewInfos: shard.RenewInfos,
				}

				textShards = append(textShards, textShard)
			}

			resText := types.QueryAllShardInTextResponse{
				Shard:      textShards,
				Pagination: res.Pagination,
			}

			return clientCtx.PrintProto(&resText)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowShard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-shard [id]",
		Short: "shows a shard",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetShardRequest{
				Id: id,
			}

			res, err := queryClient.Shard(context.Background(), params)
			if err != nil {
				return err
			}

			shard := res.GetShard()
			textShard := types.ShardInText{
				Id:         shard.Id,
				OrderId:    shard.OrderId,
				Status:     ShardStatusInText(int(shard.Status)),
				Size_:      shard.Size_,
				Cid:        shard.Cid,
				Pledge:     shard.Pledge,
				From:       shard.From,
				Sp:         shard.Sp,
				Duration:   shard.Duration,
				CreatedAt:  shard.CreatedAt,
				RenewInfos: shard.RenewInfos,
			}

			resText := types.QueryGetShardInTextResponse{
				Shard: textShard,
			}

			return clientCtx.PrintProto(&resText)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
