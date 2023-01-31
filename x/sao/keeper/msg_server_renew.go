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
	var sigDid string
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &msg.Proposal
	if proposal.Owner != "all" {
		sigDid, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
		if err != nil {
			return nil, err
		}
	} else {
		sigDid = "all"
	}

	resp := types.MsgRenewResponse{
		Result: make(map[string]string, 0),
	}

	for _, dataId := range proposal.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			resp.Result[dataId] = status.Errorf(codes.NotFound, "dataId %s not found", dataId).Error()
			continue
		}

		if metadata.Owner != sigDid {
			// only the data model owner could renew operations
			resp.Result[dataId] = sdkerrors.Wrapf(types.ErrorNoPermission, "No permission to renew the model %s", dataId).Error()
			continue
		}

		sps := k.FindSPByDataId(ctx, dataId)

		oldOrder, found := k.order.GetOrder(ctx, metadata.OrderId)
		if !found {
			return nil, sdkerrors.Wrap(types.ErrOrderNotFound, "")
		}

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

		price := sdk.NewDecWithPrec(1, 3)

		owner_address, err := k.did.GetCosmosPaymentAddress(ctx, order.Owner)
		if err != nil {
			resp.Result[dataId] = err.Error()
			continue
		}

		amount, _ := sdk.NewDecCoinFromDec(sdk.DefaultBondDenom, price.MulInt64(int64(order.Size_)).MulInt64(int64(order.Replica)).MulInt64(int64(order.Duration))).TruncateDecimal()
		balance := k.bank.GetBalance(ctx, owner_address, sdk.DefaultBondDenom)

		logger := k.Logger(ctx)

		logger.Debug("order amount1 ###################", "amount", amount, "owner", owner_address, "balance", balance)

		if balance.IsLT(amount) {
			resp.Result[dataId] = sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", amount.Amount.Int64()).Error()
			continue
		}

		order.Amount = amount
		sps_addr := make([]string, 0)
		for _, sp := range sps {
			sps_addr = append(sps_addr, sp.String())
		}

		k.order.GenerateShards(ctx, &order, sps_addr)

		newOrderId, err := k.order.NewOrder(ctx, &order, sps_addr)
		if err != nil {
			resp.Result[dataId] = err.Error()
			continue
		}
		resp.Result[dataId] = fmt.Sprintf("New order=%d", newOrderId)
	}

	return &resp, nil
}
