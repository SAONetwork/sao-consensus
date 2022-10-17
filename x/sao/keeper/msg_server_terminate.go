package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Terminate(goCtx context.Context, msg *types.MsgTerminate) (*types.MsgTerminateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	order, found := k.GetOrder(ctx, msg.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.OrderId)
	}

	if order.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotCreator, "only order creator allowed")
	}

	if order.Status != types.OrderCompleted {
		return nil, sdkerrors.Wrapf(types.ErrOrderUnexpectedStatus, "invalid order stauts, expect complete")
	}

	order.Status = types.OrderTerminated

	k.SetOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.TerminateOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
		),
	)

	for provider, shard := range order.Shards {
		// release provider pledge
		coin := sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(shard.Pledge))
		addr, _ := sdk.AccAddressFromBech32(provider)
		k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.Coins{coin})

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.TerminateShardEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
				sdk.NewAttribute(types.ShardEventProvider, provider),
				sdk.NewAttribute(types.EventCid, shard.Cid),
			),
		)
	}

	return &types.MsgTerminateResponse{}, nil
}
