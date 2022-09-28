package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/libp2p/go-libp2p-core/peer"
)

func (k msgServer) Login(goCtx context.Context, msg *types.MsgLogin) (*types.MsgLoginResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_, found := k.GetNode(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyRegistered, "%s", msg.Creator)
	}

	_, err := peer.AddrInfoFromString(msg.Peer)

	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPeer, "%s", msg.Peer)
	}

	var node = types.Node{
		Peer:    msg.Peer,
		Creator: msg.Creator,
	}
	k.SetNode(ctx, node)

	return &types.MsgLoginResponse{}, nil
}
