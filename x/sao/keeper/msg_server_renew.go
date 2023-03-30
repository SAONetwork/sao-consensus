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
		Result: make([]*types.KV, 0),
	}

	ownerAddress, err := k.did.GetCosmosPaymentAddress(ctx, sigDid)
	if err != nil {
		return nil, err
	}
	denom := k.staking.BondDenom(ctx)
	balance := k.bank.GetBalance(ctx, ownerAddress, denom)

	for _, dataId := range proposal.Data {
		metadata, found := k.Keeper.model.GetMetadata(ctx, dataId)
		if !found {
			kv := &types.KV{
				K: dataId,
				V: status.Errorf(codes.NotFound, "FAILED: dataId %s not found", dataId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		if metadata.Owner != sigDid {
			// only the data model owner could renew operations
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrorNoPermission, "FAILED: no permission to renew the model %s", dataId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		sps := k.FindSPByDataId(ctx, dataId)

		oldOrder, found := k.order.GetOrder(ctx, metadata.OrderId)
		if !found {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrOrderNotFound, "FAILED: invalid order id: %d", metadata.OrderId).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		if oldOrder.Status != ordertypes.OrderCompleted {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "FAILED: expected status %d, but get %d", ordertypes.OrderCompleted, oldOrder.Status).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		var order = ordertypes.Order{
			Creator:   msg.Creator,
			Owner:     metadata.Owner,
			Cid:       oldOrder.Cid,
			Timeout:   oldOrder.Timeout,
			Duration:  proposal.Duration,
			Status:    ordertypes.OrderDataReady,
			Size_:     oldOrder.Size_,
			Replica:   oldOrder.Replica,
			Operation: 3,
		}

		price := sdk.NewDecWithPrec(1, 6)

		owner_address, err := k.did.GetCosmosPaymentAddress(ctx, order.Owner)
		if err != nil {
			kv := &types.KV{
				K: dataId,
				V: "FAILED: " + err.Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}

		amount, _ := sdk.NewDecCoinFromDec(denom, price.MulInt64(int64(order.Size_)).MulInt64(int64(order.Replica)).MulInt64(int64(order.Duration))).TruncateDecimal()

		logger := k.Logger(ctx)
		logger.Debug("order amount", "amount", amount, "owner", owner_address, "balance", balance)

		if balance.IsLT(amount) {
			kv := &types.KV{
				K: dataId,
				V: sdkerrors.Wrapf(types.ErrInsufficientCoin, "FAILED: insufficient coin: need %d", amount.Amount.Int64()).Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		} else {
			balance = balance.Sub(amount)
		}

		order.Amount = amount
		sps_addr := make([]string, 0)
		for _, sp := range sps {
			sps_addr = append(sps_addr, sp.String())
		}

		k.order.GenerateShards(ctx, &order, sps_addr)

		k.order.SetOrder(ctx, order)

		newOrderId, err := k.order.NewOrder(ctx, &order, sps_addr)
		if err != nil {
			kv := &types.KV{
				K: dataId,
				V: "FAILED: " + err.Error(),
			}
			resp.Result = append(resp.Result, kv)
			continue
		}
		kv := &types.KV{
			K: dataId,
			V: fmt.Sprintf("SUCCESS: new orderId=%d", newOrderId),
		}
		resp.Result = append(resp.Result, kv)
	}

	return &resp, nil
}
