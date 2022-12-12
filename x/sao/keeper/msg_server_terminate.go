package keeper

import (
	"context"

	saodid "github.com/SaoNetwork/sao-did"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
	didtypes "github.com/SaoNetwork/sao/x/did/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dvsekhvalnov/jose2go/base64url"
)

func (k msgServer) Terminate(goCtx context.Context, msg *types.MsgTerminate) (*types.MsgTerminateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var querySidDocument = func(versionId string) (*didtypes.SidDocument, error) {
		doc, found := k.did.GetSidDocument(ctx, versionId)
		if found {
			return &doc, nil
		} else {
			return nil, nil
		}
	}
	didManager, err := saodid.NewDidManagerWithDid(msg.Proposal.Owner, querySidDocument)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidDid, "")
	}

	proposalBytes, err := msg.Proposal.Marshal()
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidProposal, "")
	}

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

	order, found := k.order.GetOrder(ctx, msg.Proposal.OrderId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrOrderNotFound, "order %d not found", msg.Proposal.OrderId)
	}

	if order.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrNotCreator, "only order creator allowed")
	}

	if order.Metadata != nil && order.Metadata.DataId != "" {
		err := k.model.DeleteMeta(ctx, order.Metadata.DataId)
		if err != nil {
			return nil, err
		}
	}

	err = k.order.TerminateOrder(ctx, order.Id)
	if err != nil {
		return nil, err
	}

	k.market.Withdraw(ctx, order)

	return &types.MsgTerminateResponse{}, nil
}
