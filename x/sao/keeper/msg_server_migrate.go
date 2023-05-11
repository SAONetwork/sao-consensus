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

	isProvider := false
	if msg.Provider == msg.Creator {
		isProvider = true
	} else {
		provider, found := k.node.GetNode(ctx, msg.Provider)
		if found {
			for _, address := range provider.TxAddresses {
				if address == msg.Creator {
					isProvider = true
				}
			}
		}
	}

	if !isProvider {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, msg.Provider: %s", msg.Creator, msg.Provider)
	}

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

		oldShard := k.order.GetOrderShardBySP(ctx, &oldOrder, msg.Provider)
		if oldShard == nil {
			kv := &types.KV{
				K: dataId,
				V: status.Errorf(codes.NotFound, "FAILED: %s shard not found", oldOrder.Provider).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		ignoreList := make([]string, 0)
		for _, id := range oldOrder.Shards {
			shard, found := k.order.GetShard(ctx, id)
			if !found {
				continue
			}
			ignoreList = append(ignoreList, shard.Sp)
		}

		sps := k.node.RandomSP(ctx, 1, ignoreList)

		newShard := k.order.MigrateShard(ctx, &oldOrder, msg.Provider, sps[0].Creator)

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
