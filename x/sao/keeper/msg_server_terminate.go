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
		return nil, status.Errorf(codes.NotFound, "dataId:%s not found", msg.Proposal.DataId)
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

	shardSet := make(map[uint64]int)
	for _, orderId := range meta.Orders {
		order, found := k.order.GetOrder(ctx, orderId)
		if !found {
			continue
		}

		for _, shardId := range order.Shards {
			shardSet[shardId] = 1
		}

		err = k.model.TerminateOrder(ctx, order)
		if err != nil {
			return nil, err
		}
	}

	for shardId := range shardSet {
		k.order.RemoveShard(ctx, shardId)
	}

	err = k.model.DeleteMeta(ctx, msg.Proposal.DataId)
	if err != nil {
		return nil, err
	}

	return &types.MsgTerminateResponse{}, nil
}
