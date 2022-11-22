package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/model/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	orderList := make([]*ordertypes.Order, 0)
	for _, dataId := range msg.Data {
		metadata, found := k.GetMetadata(ctx, dataId)
		if !found {
			return nil, status.Errorf(codes.NotFound, "dataId %s not found", dataId)
		}

		order, found := k.order.GetOrder(ctx, metadata.OrderId)
		if !found {
			return nil, status.Errorf(codes.NotFound, "orderId %s not found", metadata.OrderId)
		}

		if _, ok := order.Shards[msg.Creator]; !ok {
			return nil, status.Errorf(codes.NotFound, "shard of %s not found in order %d", msg.Creator, order.Id)
		}

		orderList = append(orderList, &order)
	}

	err := k.order.ShardsPayment(ctx, orderList, msg.Creator)

	if err != nil {
		return nil, err
	}

	return &types.MsgClaimResponse{}, nil
}
