package keeper

import (
	"context"
	"fmt"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ipfs/go-cid"
)

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check provider
	node, found := k.node.GetNode(ctx, msg.Provider)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

	// check cid
	_, err := cid.Decode(msg.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invalid cid: %s", msg.Cid)
	}

	var order = ordertypes.Order{
		Creator:  msg.Creator,
		Owner:    msg.Owner,
		Provider: node.Creator,
		Cid:      msg.Cid,
		Expire:   int32(ctx.BlockHeight()) + 86400,
		Duration: msg.Duration,
		Status:   types.OrderPending,
		Replica:  msg.Replica,
	}

	order.Id = k.AppendOrder(ctx, order)

	if order.Provider == msg.Creator {
		// create shard when msg creator is data provider
		err = k.newRandomShard(ctx, &order)
		if err != nil {
			return nil, err
		}

	}

	k.SetOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	return &types.MsgStoreResponse{OrderId: order.Id}, nil
}
