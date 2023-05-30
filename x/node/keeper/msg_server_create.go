package keeper

import (
	"context"
	"math/big"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Create(goCtx context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_, found := k.GetNode(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyRegistered, "%s", msg.Creator)
	}

	// role: 0 - Normal, 1 - Super
	var node = types.Node{
		Peer:            "",
		Creator:         msg.Creator,
		Reputation:      10000.0,
		Status:          types.NODE_STATUS_NA,
		LastAliveHeight: ctx.BlockHeight(),
		Role:            0,
	}

	if msg.Validator != "" {
		// if validator provided, check if it satisfy super node requirement.
		accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDelegate, "%v", err)
		}
		valAddress, err := sdk.ValAddressFromBech32(msg.Validator)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDelegate, "%v", err)
		}

		// delegator share
		delegate, found := k.staking.GetDelegation(ctx, accAddress, valAddress)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDelegate, "delegator %v delegate to validator %v not found", msg.Creator, msg.Validator)
		}
		share := delegate.Shares.BigInt()

		// validator share
		validator, found := k.staking.GetValidator(ctx, valAddress)
		if !found {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDelegate, "query validator error %v", found)
		}
		totalShares := validator.DelegatorShares.BigInt()

		if new(big.Int).Div(totalShares, share).Cmp(big.NewInt(types.SHARE_THRESHOLD)) > 0 {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDelegate, "insufficient shares in this validator")
		}
		node.Role = 1
		node.Validator = msg.Validator
	}

	k.SetNode(ctx, node)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.CreateEventType,
			sdk.NewAttribute(types.NodeEventCreator, node.Creator),
			sdk.NewAttribute(types.NodeEventPeer, node.Peer),
		),
	)

	return &types.MsgCreateResponse{}, nil
}
