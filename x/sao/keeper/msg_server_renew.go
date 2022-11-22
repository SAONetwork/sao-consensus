package keeper

import (
	"context"

	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Renew(goCtx context.Context, msg *types.MsgRenew) (*types.MsgRenewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, dataId := range msg.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			return nil, status.Errorf(codes.NotFound, "dataId %s not found", dataId)
		}

		sps := k.FindSPByDataId(ctx, dataId)

		oldOrder, found := k.order.GetOrder(ctx, metadata.OrderId)

		var order = ordertypes.Order{
			Creator:   msg.Creator,
			Owner:     metadata.Owner,
			Cid:       oldOrder.Cid,
			Expire:    int32(ctx.BlockHeight()) + msg.Timeout,
			Duration:  msg.Duration,
			Status:    types.OrderDataReady,
			Size_:     oldOrder.Size_,
			Replica:   oldOrder.Replica,
			Operation: 2,
		}

		price := sdk.NewInt(1)

		owner_address := k.did.GetCosmosPaymentAddress(ctx, order.Owner)

		amount := sdk.NewCoin(sdk.DefaultBondDenom, price.MulRaw(int64(order.Size_)).MulRaw(int64(order.Replica)))
		balance := k.bank.GetBalance(ctx, owner_address, sdk.DefaultBondDenom)

		logger := k.Logger(ctx)

		logger.Debug("order amount1 ###################", "amount", amount, "owner", owner_address, "balance", balance)

		if balance.IsLT(amount) {
			return nil, sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", amount.Amount.Int64())
		}

		order.Amount = amount

		_, err := k.order.NewOrder(ctx, order, sps)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgRenewResponse{}, nil
}
