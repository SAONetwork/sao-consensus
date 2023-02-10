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
		Result: make(map[string]string, 0),
	}

	for _, dataId := range msg.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			resp.Result[dataId] = status.Errorf(codes.NotFound, "FAILED: dataId %s not found", dataId).Error()
			continue
		}

		oldOrder, found := k.order.GetOrder(ctx, metadata.OrderId)
		if !found {
			resp.Result[dataId] = sdkerrors.Wrapf(types.ErrOrderNotFound, "FAILED: invalid order id: %d", metadata.OrderId).Error()
			continue
		}

		_, ok := oldOrder.Shards[msg.Creator]
		if !ok {
			resp.Result[dataId] = status.Errorf(codes.NotFound, "FAILED: %s shard not found", msg.Creator).Error()
			continue
		}

		sps := k.node.RandomSP(ctx, 1)

		shard := k.order.MigrateShard(ctx, &oldOrder, msg.Creator, sps[0].Creator)

		oldOrder.Shards[sps[0].Creator] = shard

		oldOrder.Status = ordertypes.OrderInProgress

		k.order.SetOrder(ctx, oldOrder)

		resp.Result[dataId] = fmt.Sprintf("SUCCESS: new storage provider %s", sps[0].Creator)
	}

	return &resp, nil

	return &types.MsgMigrateResponse{}, nil
}
