package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) NewShards(ctx sdk.Context, order *ordertypes.Order) []*types.Shard {

	shards := make([]*types.Shard, 0)
	for i := 0; i < int(order.Replica); i++ {
		var sp string
		if order.Operation == 1 {
			sps := k.RandomSP(ctx, 1)
			sp = sps[0].Creator
		} else {
			idx := fmt.Sprintf("%s-%d", order.Metadata.DataId, i)
			shard, found := k.GetShard(ctx, idx)
			if found {
				sp = shard.Node
			}
		}
		if sp != "" {
			shard := k.NewShard(ctx, order, i, sp)
			shards = append(shards, shard)
		}
	}
	return shards
}

func (k Keeper) NewShard(ctx sdk.Context, order *ordertypes.Order, idx int, sp string) *types.Shard {

	shard := types.Shard{
		Idx:     fmt.Sprintf("%s-d", order.Metadata.DataId, idx),
		OrderId: order.Id,
		Status:  types.ShardWaiting,
		Cid:     order.Cid,
		From:    order.Provider,
		Size_:   order.Size_,
		Pledged: sdk.NewInt64Coin(order.Amount.Denom, 0),
		Node:    sp,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewShardEventType,
			sdk.NewAttribute(ordertypes.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(ordertypes.OrderEventProvider, order.Provider),
			sdk.NewAttribute(ordertypes.ShardEventProvider, shard.Node),
			sdk.NewAttribute(ordertypes.EventCid, shard.Cid),
			sdk.NewAttribute(ordertypes.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(ordertypes.OrderEventOperation, fmt.Sprintf("%d", order.Operation)),
		),
	)

	k.SetShard(ctx, shard)

	return &shard
}

func (k Keeper) RenwShard(ctx sdk.Context, shard types.Shard) {

}

func (k Keeper) ActiveShard(ctx sdk.Context, order *ordertypes.Order, shard *types.Shard, cid string, size uint64) error {

	shard.Status = types.ShardCompleted
	shard.Cid = cid
	shard.Size_ = size

	err := k.OrderPledge(ctx, sdk.MustAccAddressFromBech32(shard.Node), order)
	if err != nil {
		err = sdkerrors.Wrap(types.ErrorOrderPledgeFailed, err.Error())
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ShardCompletedEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.ShardEventProvider, shard.Node),
		),
	)

	k.SetShard(ctx, *shard)

	return nil
}

func (k Keeper) GetMetadataShards(ctx sdk.Context, dataId string, count int) map[string]*types.Shard {
	shards := make(map[string]*types.Shard, 0)
	for i := 0; i < count; i++ {
		idx := fmt.Sprintf("%s-%d", dataId, count)
		shard, found := k.GetShard(ctx, idx)
		if found {
			shards[shard.Node] = &shard
		}
	}
	return shards
}
