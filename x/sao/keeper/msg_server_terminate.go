package keeper

import (
	"context"

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

	// validate the permission for all terminate operations
	meta, isFound := k.Keeper.model.GetMetadata(ctx, msg.Proposal.DataId)
	if !isFound {
		return nil, status.Errorf(codes.NotFound, "dataId:%d not found", msg.Proposal.DataId)
	}

	kid, err := signature.GetKid()
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
	}
	sigDid, err := saodidutil.KidToDid(kid)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
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
