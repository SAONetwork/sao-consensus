package keeper

import (
	"fmt"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NewOrder(ctx sdk.Context, order types.Order, sps []nodetypes.Node) uint64 {

	order.Id = k.AppendOrder(ctx, order)

	k.GenerateShards(ctx, order, sps)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	k.SetOrder(ctx, order)
	return order.Id
}

func (k Keeper) GenerateShards(ctx sdk.Context, order types.Order, sps []nodetypes.Node) {

	if len(sps) > 0 {
		shards := make(map[string]*types.Shard, 0)
		for _, sp := range sps {
			shards[sp.Creator] = &types.Shard{
				OrderId: order.Id,
				Status:  types.ShardWaiting,
				Cid:     order.Cid,
			}
		}

		order.Shards = shards

		order.Status = types.OrderDataReady

		for provider, shard := range order.Shards {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(types.NewShardEventType,
					sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
					sdk.NewAttribute(types.OrderEventProvider, order.Provider),
					sdk.NewAttribute(types.ShardEventProvider, provider),
					sdk.NewAttribute(types.EventCid, shard.Cid),
					sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
					sdk.NewAttribute(types.OrderEventProvider, order.Provider),
				),
			)
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.OrderDataReadyEventType,
				sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
				sdk.NewAttribute(types.OrderEventCreator, order.Creator),
				sdk.NewAttribute(types.OrderEventProvider, order.Provider),
				sdk.NewAttribute(types.EventCid, order.Cid),
			),
		)
	}
	k.SetOrder(ctx, order)
}
