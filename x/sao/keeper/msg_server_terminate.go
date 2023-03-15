package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Terminate(goCtx context.Context, msg *types.MsgTerminate) (*types.MsgTerminateResponse, error) {
	var sigDid string
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &msg.Proposal
	if proposal == nil {
		return nil, status.Errorf(codes.InvalidArgument, "proposal is required")
	}

	if proposal.Owner != "all" {
		sigDid, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, sdkerrors.Wrap(types.ErrorNoPermission, "No permission to delete the open data model")
	}

	// validate the permission for all terminate operations
	meta, isFound := k.Keeper.model.GetMetadata(ctx, msg.Proposal.DataId)
	if !isFound {
		return nil, status.Errorf(codes.NotFound, "dataId:%d not found", msg.Proposal.DataId)
	}

	isValid := meta.Owner == sigDid
	if !isValid {
		for _, readwriteDid := range meta.ReadwriteDids {
			if readwriteDid == sigDid {
				isValid = true
				break
			}
		}

		if !isValid {
			return nil, sdkerrors.Wrap(types.ErrorNoPermission, "No permission to delete the model")
		}
	}

	order, found := k.order.GetOrder(ctx, meta.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", meta.OrderId)
	}

	if order.DataId != "" {
		err := k.model.DeleteMeta(ctx, order.DataId)
		if err != nil {
			return nil, err
		}
	}

	refundCoin, err := k.market.Withdraw(ctx, order)
	if err != nil {
		return nil, err
	}

	err = k.order.TerminateOrder(ctx, order.Id, refundCoin)
	if err != nil {
		return nil, err
	}

	return &types.MsgTerminateResponse{}, nil
}
