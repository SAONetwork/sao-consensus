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

func (k msgServer) Renew(goCtx context.Context, msg *types.MsgRenew) (*types.MsgRenewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal := msg.Proposal

	resp := types.MsgRenewResponse{
		Result: make(map[string]string, 0),
	}

	for _, dataId := range proposal.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			resp.Result[dataId] = status.Errorf(codes.NotFound, "dataId %s not found", dataId).Error()
			continue
		}

		sps := k.FindSPByDataId(ctx, dataId)

		oldOrder, found := k.order.GetOrder(ctx, metadata.OrderId)

		var order = ordertypes.Order{
			Creator:   msg.Creator,
			Owner:     metadata.Owner,
			Cid:       oldOrder.Cid,
			Expire:    int32(ctx.BlockHeight()) + proposal.Timeout,
			Duration:  proposal.Duration,
			Status:    types.OrderDataReady,
			Size_:     oldOrder.Size_,
			Replica:   oldOrder.Replica,
			Operation: 3,
		}

		price := sdk.NewInt(1)

		owner_address := k.did.GetCosmosPaymentAddress(ctx, order.Owner)

		amount := sdk.NewCoin(sdk.DefaultBondDenom, price.MulRaw(int64(order.Size_)).MulRaw(int64(order.Replica)))
		balance := k.bank.GetBalance(ctx, owner_address, sdk.DefaultBondDenom)

		logger := k.Logger(ctx)

		logger.Debug("order amount1 ###################", "amount", amount, "owner", owner_address, "balance", balance)

		if balance.IsLT(amount) {
			resp.Result[dataId] = sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", amount.Amount.Int64()).Error()
			continue
		}

		order.Amount = amount

		newOrderId, err := k.order.NewOrder(ctx, order, sps)
		if err != nil {
			resp.Result[dataId] = err.Error()
			continue
		}
		resp.Result[dataId] = fmt.Sprintf("New order=%d", newOrderId)
	}

	return &resp, nil
}
