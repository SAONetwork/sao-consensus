package keeper

import (
	"context"
	"fmt"

	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Migrate(goCtx context.Context, msg *types.MsgMigrate) (*types.MsgMigrateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	resp := types.MsgMigrateResponse{
		Result: make([]*types.KV, 0),
	}

	for _, dataId := range msg.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			kv := &types.KV{
				K: dataId,
				V: status.Errorf(codes.NotFound, "FAILED: dataId %s not found", dataId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		oldOrder, found := k.order.GetOrder(ctx, metadata.OrderId)
		if !found {

			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrOrderNotFound, "FAILED: invalid order id: %d", metadata.OrderId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		oldShard := k.order.GetOrderShardBySP(ctx, &oldOrder, msg.Creator)
		if oldShard == nil {
			kv := &types.KV{
				K: dataId,
				V: status.Errorf(codes.NotFound, "FAILED: %s shard not found", msg.Creator).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		ignoreList := make([]string, 0)
		for sp, _ := range oldOrder.Shards {
			ignoreList = append(ignoreList, sp)
		}

		sps := k.node.RandomSP(ctx, 1, ignoreList)

		newShard := k.order.MigrateShard(ctx, &oldOrder, msg.Creator, sps[0].Creator)

		oldOrder.Shards = append(oldOrder.Shards, newShard.Id)

		oldOrder.Status = ordertypes.OrderMigrating

		k.order.SetOrder(ctx, oldOrder)

		kv := &types.KV{
			K: dataId,
			V: fmt.Sprintf("SUCCESS: new storage provider %s", sps[0].Creator),
		}
		resp.Result = append(resp.Result, kv)
	}

	return &resp, nil
}
