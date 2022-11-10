package keeper

import (
	"context"
	"fmt"

	saodid "github.com/SaoNetwork/sao-did"
	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
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

	signature := saodid.JwsSignature{
		Protected: msg.JwsSignature.Protected,
		Signature: msg.JwsSignature.Signature,
	}

	_, err = didManager.VerifyJWS(saodid.GeneralJWS{
		Payload: base64url.Encode(proposalBytes),
		Signatures: []saodid.JwsSignature{
			signature,
		},
	})

	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
	}

	var rawMetadata string

	if proposal.DataId != "" {
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
		rawMetadata = metadata.String()
	}

	var order = types.Order{
		Creator:  msg.Creator,
		Owner:    proposal.Owner,
		Provider: node.Creator,
		Cid:      proposal.Cid,
		Expire:   proposal.Timeout,
		Duration: proposal.Duration,
		Status:   types.OrderPending,
		Replica:  proposal.Replica,
		Metadata: rawMetadata,
	}

	order.Id = k.AppendOrder(ctx, order)

	if order.Provider == msg.Creator {
		// create shard when msg creator is data provider
		err = k.newRandomShard(ctx, &order)
		if err != nil {
			return nil, err
		}

	}

	k.SetOrder(ctx, order)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.NewOrderEventType,
			sdk.NewAttribute(types.EventOrderId, fmt.Sprintf("%d", order.Id)),
			sdk.NewAttribute(types.OrderEventCreator, order.Creator),
			sdk.NewAttribute(types.OrderEventProvider, order.Provider),
			sdk.NewAttribute(types.EventCid, order.Cid),
		),
	)

	return &types.MsgStoreResponse{OrderId: order.Id}, nil
}
