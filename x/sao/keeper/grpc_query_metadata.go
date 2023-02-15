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

func (k Keeper) Metadata(goCtx context.Context, req *types.QueryMetadataRequest) (*types.QueryMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sigDid string
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &req.Proposal
	if proposal.Owner != "all" {
		sigDid, err = k.verifySignature(ctx, proposal.Owner, proposal, req.JwsSignature)
		if err != nil {
			return nil, err
		}
	} else {
		sigDid = "all"
	}

	var dataId string
	if proposal.KeywordType > 1 {
		model, isFound := k.model.GetModel(ctx, fmt.Sprintf("%s-%s-%s",
			proposal.Owner, proposal.Keyword, proposal.GroupId,
		))
		if !isFound {
			return nil, status.Errorf(codes.NotFound, "dataId not found by Alias: %s", proposal.Keyword)
		}
		dataId = model.Data
	} else {
		dataId = proposal.Keyword
	}

	meta, isFound := k.model.GetMetadata(ctx, dataId)
	if !isFound {
		return nil, status.Errorf(codes.NotFound, "dataId:%s not found", dataId)
	}

	// validate the permission for all query operations
	isValid := meta.Owner == sigDid
	if !isValid {
		for _, readwriteDid := range meta.ReadwriteDids {
			if readwriteDid == sigDid {
				isValid = true
				break
			}
		}

		if !isValid {
			for _, readonlyDid := range meta.ReadonlyDids {
				if readonlyDid == sigDid {
					isValid = true
					break
				}
			}
		}

		if !isValid {
			return nil, sdkerrors.Wrap(types.ErrorNoPermission, "No permission to update the model")
		}
	}

	order, found := k.order.GetOrder(ctx, meta.OrderId)
	if !found {
		return nil, status.Errorf(codes.NotFound, "order:%d not found", meta.OrderId)
	}

	metadata := types.Metadata{
		DataId:     meta.DataId,
		Owner:      meta.Owner,
		Alias:      meta.Alias,
		GroupId:    meta.GroupId,
		OrderId:    meta.OrderId,
		Tags:       meta.Tags,
		Cid:        meta.Cid,
		Commits:    meta.Commits,
		ExtendInfo: meta.ExtendInfo,
		Update:     meta.Update,
		Commit:     meta.Commit,
		Rule:       meta.Rule,
		Duration:   meta.Duration,
		CreatedAt:  meta.CreatedAt,
		Provider:   order.Provider,
		Expire:     order.Expire,
		Status:     order.Status,
		Replica:    order.Replica,
		Amount:     order.Amount,
		Size_:      order.Size_,
		Operation:  order.Operation,
	}

	shards := make(map[string]*types.ShardMeta, 0)

	for i := 0; i < int(meta.Replica); i++ {
		idx := fmt.Sprintf("%s-%d", meta.DataId, i)
		shard, _ := k.node.GetShard(ctx, idx)
		node, node_found := k.node.GetNode(ctx, shard.Node)
		if !node_found {
			continue
		}
		meta := types.ShardMeta{
			Idx:      shard.Idx,
			Peer:     node.Peer,
			Cid:      shard.Cid,
			Provider: order.Provider,
		}
		shards[shard.Node] = &meta
	}

	return &types.QueryMetadataResponse{
		Metadata: metadata,
		Shards:   shards,
	}, nil
}
