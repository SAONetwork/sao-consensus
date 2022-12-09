package keeper

import (
	"context"
	"fmt"

	saodid "github.com/SaoNetwork/sao-did"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
	saodidutil "github.com/SaoNetwork/sao-did/util"
	didtypes "github.com/SaoNetwork/sao/x/did/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Metadata(goCtx context.Context, req *types.QueryMetadataRequest) (*types.QueryMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &req.Proposal

	var sigDid string
	if proposal.Owner != "all" {
		var querySidDocument = func(versionId string) (*didtypes.SidDocument, error) {
			doc, found := k.did.GetSidDocument(ctx, versionId)
			if found {
				return &doc, nil
			} else {
				return nil, nil
			}
		}
		didManager, err := saodid.NewDidManagerWithDid(proposal.Owner, querySidDocument)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrorInvalidDid, "")
		}

		proposalBytes, err := proposal.Marshal()
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrorInvalidProposal, "")
		}

		signature := saodidtypes.JwsSignature{
			Protected: req.JwsSignature.Protected,
			Signature: req.JwsSignature.Signature,
		}

		// logger.Error("###################", "proposal", proposal)
		// logger.Error("###################", "proposalBytes", string(proposalBytes))
		// logger.Error("###################", "msg.JwsSignature.Protected", msg.JwsSignature.Protected)

		_, err = didManager.VerifyJWS(saodidtypes.GeneralJWS{
			Payload: base64url.Encode(proposalBytes),
			Signatures: []saodidtypes.JwsSignature{
				signature,
			},
		})

		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
		}

		kid, err := signature.GetKid()
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
		}
		sigDid, err = saodidutil.KidToDid(kid)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
		}
	} else {
		sigDid = "all"
	}

	var dataId string
	if proposal.Type_ > 1 {
		model, isFound := k.model.GetModel(ctx, fmt.Sprintf("%s-%s-%s",
			proposal.Owner, proposal.Keyword, proposal.GroupId,
		))
		if !isFound {
			return nil, status.Errorf(codes.NotFound, "dataId not found by Alais: %s", proposal.Keyword)
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
	for p, shard := range order.Shards {
		node, node_found := k.node.GetNode(ctx, p)
		if !node_found {
			continue
		}
		meta := types.ShardMeta{
			ShardId:  shard.Id,
			Peer:     node.Peer,
			Cid:      shard.Cid,
			Provider: order.Provider,
		}
		shards[p] = &meta
	}

	return &types.QueryMetadataResponse{
		Metadata: metadata,
		Shards:   shards,
	}, nil
}
