package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Logout(goCtx context.Context, msg *types.MsgLogout) (*types.MsgLogoutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	node, found := k.GetNode(ctx, msg.Creator)

	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

	signers := msg.GetSigners()

	if len(signers) != 1 || signers[0].String() != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrSignerAndCreator, "signer shoud equal to creator")
	}

	if len(signers) != 1 || signers[0].String() != node.Creator {
		return nil, sdkerrors.Wrapf(types.ErrOnlyOwner, "only node owner can execute this action")
	}

	k.RemoveNode(ctx, msg.Creator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.LogoutEventType,
			sdk.NewAttribute(types.NodeEventCreator, node.Creator),
		),
	)
	return &types.MsgLogoutResponse{}, nil
}
