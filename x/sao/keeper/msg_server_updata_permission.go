package keeper

import (
	"context"
	"fmt"

	saodid "github.com/SaoNetwork/sao-did"
	sid "github.com/SaoNetwork/sao-did/sid"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dvsekhvalnov/jose2go/base64url"
)

func (k msgServer) UpdataPermission(goCtx context.Context, msg *types.MsgUpdataPermission) (*types.MsgUpdataPermissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	logger := k.Logger(ctx)

	var querySidDocument = func(versionId string) (*sid.SidDocument, error) {
		doc, found := k.did.GetSidDocument(ctx, versionId)
		if found {
			var keys = make([]*sid.PubKey, 0)
			for _, pk := range doc.Keys {
				keys = append(keys, &sid.PubKey{
					Name:  pk.Name,
					Value: pk.Value,
				})
			}
			return &sid.SidDocument{
				VersionId: doc.VersionId,
				Keys:      keys,
			}, nil
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
		return nil, err
	}

	err = checkDid(msg.Proposal.ReadwriteDids)
	if err != nil {
		return nil, err
	}

	err = k.model.UpdatePermission(ctx, msg.Proposal.Owner, msg.Proposal.DataId, msg.Proposal.ReadonlyDids, msg.Proposal.ReadwriteDids)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "")
	}

	return &types.MsgUpdataPermissionResponse{}, nil
}
