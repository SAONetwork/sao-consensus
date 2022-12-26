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

	if proposal.Owner != "all" {
		_, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
		if err != nil {
			return &types.MsgUpdataPermissionResponse{}, err
		}
	} else {
		return &types.MsgUpdataPermissionResponse{}, sdkerrors.Wrap(types.ErrorInvalidOwner, "cannot update an open data model")
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
