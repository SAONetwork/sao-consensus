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

func (k msgServer) UpdataPermission(goCtx context.Context, msg *types.MsgUpdataPermission) (*types.MsgUpdataPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	logger := k.Logger(ctx)

	var querySidDocument = func(versionId string) (*didtypes.SidDocument, error) {
		doc, found := k.did.GetSidDocument(ctx, versionId)
		if found {
			logger.Error("order amount1 ###################", "stupid", doc)

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

	logger.Error("###################", "proposal", msg.Proposal)
	logger.Error("###################", "proposalBytes", string(proposalBytes))
	logger.Error("###################", "msg.JwsSignature.Protected", msg.JwsSignature.Protected)

	_, err = didManager.VerifyJWS(saodidtypes.GeneralJWS{
		Payload: base64url.Encode(proposalBytes),
		Signatures: []saodidtypes.JwsSignature{
			signature,
		},
	})

	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidSignature, "")
	}

	err = k.model.UpdatePermission(ctx, msg.Proposal.Owner, msg.Proposal.DataId, msg.Proposal.ReadonlyDids, msg.Proposal.ReadwriteDids)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "")
	}

	return &types.MsgUpdataPermissionResponse{}, nil
}
