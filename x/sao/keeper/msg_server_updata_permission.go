package keeper

import (
	"context"
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) UpdataPermission(goCtx context.Context, msg *types.MsgUpdataPermission) (*types.MsgUpdataPermissionResponse, error) {
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &msg.Proposal
	if proposal == nil {
		return &types.MsgUpdataPermissionResponse{}, status.Errorf(codes.InvalidArgument, "proposal is required")
	}

	isProvider := false
	if msg.Provider == msg.Creator {
		isProvider = true
	} else {
		provider, found := k.node.GetNode(ctx, msg.Provider)
		if found {
			for _, address := range provider.TxAddresses {
				if address == msg.Creator {
					isProvider = true
				}
			}
		}
	}

	if !isProvider {
		return nil, sdkerrors.Wrapf(types.ErrorInvalidProvider, "msg.Creator: %s, msg.Provider: %s", msg.Creator, msg.Provider)
	}

	_, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
	if err != nil {
		return &types.MsgUpdataPermissionResponse{}, err
	}

	checkDid := func(didList []string) error {
		for _, did := range didList {
			err = k.did.ValidDid(ctx, did)
			if err != nil {
				return sdkerrors.Wrap(types.ErrorInvalidDid, fmt.Sprintf("invalid did: %v, err: %v", did, err))
			}
		}
		return nil
	}

	err = checkDid(msg.Proposal.ReadonlyDids)
	if err != nil {
		return &types.MsgUpdataPermissionResponse{}, err
	}

	err = checkDid(msg.Proposal.ReadwriteDids)
	if err != nil {
		return &types.MsgUpdataPermissionResponse{}, err
	}

	err = k.model.UpdatePermission(ctx, msg.Proposal.Owner, msg.Proposal.DataId, msg.Proposal.ReadonlyDids, msg.Proposal.ReadwriteDids)
	if err != nil {
		return &types.MsgUpdataPermissionResponse{}, sdkerrors.Wrap(err, "")
	}

	return &types.MsgUpdataPermissionResponse{}, nil
}
