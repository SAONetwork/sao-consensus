package keeper

import (
	"context"
	"encoding/json"

	saodid "github.com/SaoNetwork/sao-did"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/ipfs/go-cid"
)

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	proposal := msg.Proposal

	// check provider
	node, found := k.node.GetNode(ctx, proposal.Provider)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

	// check cid
	_, err := cid.Decode(proposal.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invalid cid: %s", proposal.Cid)
	}

	didManager, err := saodid.NewDidManagerWithDid(proposal.Owner)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidDid, "")
	}

	proposalBytes, _ := proposal.Marshal()

	signature := saodidtypes.JwsSignature{
		Protected: msg.JwsSignature.Protected,
		Signature: msg.JwsSignature.Signature,
	}

	_, err = didManager.VerifyJWS(saodidtypes.GeneralJWS{
		Payload: base64url.Encode(proposalBytes),
		Signatures: []saodidtypes.JwsSignature{
			signature,
		},
	})

	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
	}

	var rawMetadata string

	if proposal != nil {
		metadata := modeltypes.Metadata{
			DataId:     proposal.DataId,
			Owner:      proposal.Owner,
			Alias:      proposal.Alias,
			GroupId:    proposal.GroupId,
			Tags:       proposal.Tags,
			Cid:        proposal.Cid,
			Commit:     proposal.CommitId,
			ExtendInfo: proposal.ExtendInfo,
			Update:     proposal.IsUpdate,
			Rule:       proposal.Rule,
		}
		bytes, err := json.Marshal(metadata)
		if err != nil {
			return nil, err
		}
		rawMetadata = string(bytes)
	}

	var order = ordertypes.Order{
		Creator:  msg.Creator,
		Owner:    proposal.Owner,
		Provider: node.Creator,
		Cid:      proposal.Cid,
		Expire:   proposal.Timeout + int32(ctx.BlockHeight()),
		Duration: proposal.Duration,
		Status:   types.OrderPending,
		Replica:  proposal.Replica,
		Metadata: rawMetadata,
	}

	var sps []nodetypes.Node

	if order.Provider == msg.Creator {
		sps = k.node.RandomSP(ctx, int(order.Replica))
		if order.Replica <= 0 || int(order.Replica) > len(sps) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
		}
	}

	orderId := k.order.NewOrder(ctx, order, sps)

	return &types.MsgStoreResponse{OrderId: orderId}, nil
}
