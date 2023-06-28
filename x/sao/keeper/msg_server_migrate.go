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
	logger := k.Logger(ctx)

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
		logger.Debug("migrate dataId", "dataId", dataId, "sp", msg.Provider)
		successInfo := "[orderId:new storage provider]: ["
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			kv := &types.KV{
				K: dataId,
				V: status.Errorf(codes.NotFound, "FAILED: dataId %s not found", dataId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}
		commitSet := make(map[string]int)

	orderLoop:
		for i := len(metadata.Orders) - 1; i >= 0; i-- {
			orderId := metadata.Orders[i]
			logger.Debug("migrate order", "orderId", orderId)
			oldOrder, found := k.order.GetOrder(ctx, orderId)
			if !found {
				continue
			}

			logger.Debug("migrate order commit", "commit", oldOrder.Commit)
			if _, ok := commitSet[oldOrder.Commit]; ok {
				continue
			} else {
				commitSet[oldOrder.Commit] = 1
			}

			oldShard := k.order.GetOrderShardBySP(ctx, &oldOrder, msg.Provider)
			if oldShard == nil {
				continue
			}

			logger.Debug("migrate shard", "shardId", oldShard.Id)
			if oldShard.Status != ordertypes.ShardCompleted {
				continue
			}

			ignoreList := make([]string, 0)
			for _, id := range oldOrder.Shards {
				shard, found := k.order.GetShard(ctx, id)
				if !found {
					continue
				}
				// skip if migrating shard is already exist
				if shard.From == msg.Provider {
					continue orderLoop
				}
				ignoreList = append(ignoreList, shard.Sp)
			}

			logger.Debug("migrate ignore list", "list", ignoreList)

			sps := k.node.RandomSP(ctx, 1, ignoreList, int64(oldShard.Size_))

			if len(sps) == 0 {
				continue
			}

			newShard := k.order.MigrateShard(ctx, oldShard, &oldOrder, msg.Provider, sps[0].Creator)

			oldOrder.Shards = append(oldOrder.Shards, newShard.Id)

			k.order.SetOrder(ctx, oldOrder)

			successInfo += fmt.Sprintf("%d:%v  ", orderId, newShard.Sp)
		}
		successInfo += "]"
		kv := &types.KV{
			K: dataId,
			V: fmt.Sprintf("SUCCESS: %s", successInfo),
		}
		resp.Result = append(resp.Result, kv)
	}

	return &resp, nil
}
