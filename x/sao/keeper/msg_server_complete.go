package keeper

import (
	"context"
	"fmt"

	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ipfs/go-cid"
)

func (k msgServer) Complete(goCtx context.Context, msg *types.MsgComplete) (*types.MsgCompleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.order.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	if order.Status != types.OrderDataReady && order.Status != types.OrderInProgress {
		return nil, sdkerrors.Wrapf(types.ErrOrderComplete, "order not waiting completed")
	}

	if _, ok := order.Shards[msg.Creator]; !ok {
		return nil, sdkerrors.Wrapf(types.ErrOrderShardProvider, "%s is not the order shard provider")
	}

	shard := order.Shards[msg.Creator]

	if shard.Status == types.ShardCompleted {
		return nil, sdkerrors.Wrapf(types.ErrShardCompleted, "%s already completed the shard task in order %d", msg.Creator, order.Id)
	}

	if shard.Status != types.ShardWaiting {
		return nil, sdkerrors.Wrapf(types.ErrShardUnexpectedStatus, "invalid shard status, expect: wating")
	}

	logger := k.Logger(ctx)

	// check cid
	_, err := cid.Decode(msg.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invali cid: %s", msg.Cid)
	}

	shard.Status = types.ShardCompleted
	shard.Size_ = msg.Size_
	shard.Cid = msg.Cid
	order.Shards[msg.Creator] = shard

	order.Status = types.OrderCompleted

	// set order status
	for _, shard := range order.Shards {
		if shard.Status != types.ShardCompleted {
			order.Status = types.OrderInProgress
		}
	}

	k.order.SetOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, msg.Creator),
		),
	)

	provider := msg.GetSigners()[0]
	balance := k.bank.GetBalance(ctx, provider, sdk.DefaultBondDenom)
	coin := order.Amount

	if balance.IsLT(coin) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", coin.Amount.Int64())
	}

	orderModAddr := k.auth.GetModuleAddress(ordertypes.ModuleName)

	orderModBalance := k.bank.GetBalance(ctx, orderModAddr, sdk.DefaultBondDenom)

	if orderModBalance.IsLT(coin) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", coin.Amount.Int64())
	}

	k.node.IncreaseReputation(ctx, msg.Creator, float32(coin.Amount.Int64()))

	//k.bank.SendCoinsFromAccountToModule(ctx, provider, types.ModuleName, sdk.Coins{coin})

	k.node.OrderPledge(ctx, provider, coin)

	shard.Pledge = coin.Amount.Uint64()

	if order.Status == types.OrderCompleted {

		err := k.bank.SendCoinsFromModuleToModule(ctx, ordertypes.ModuleName, modeltypes.ModuleName, sdk.Coins{order.Amount})

		if err != nil {
			return nil, err
		}

		if order.Metadata != nil {

			if order.Metadata.Update {

			} else {
				err = k.Keeper.model.NewMeta(ctx, order)
				if err != nil {
					logger.Error("failed to store metadata", "err", err.Error())
					return &types.MsgCompleteResponse{}, err
				}
			}
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.OrderCompletedEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			),
		)
	}

	return &types.MsgCompleteResponse{}, nil
}
