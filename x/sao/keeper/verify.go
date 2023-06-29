package keeper

import (
	saodid "github.com/SaoNetwork/sao-did"
	sid "github.com/SaoNetwork/sao-did/sid"
	saodidtypes "github.com/SaoNetwork/sao-did/types"
	saodidutil "github.com/SaoNetwork/sao-did/util"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dvsekhvalnov/jose2go/base64url"
)

type ProposalApi interface {
	Marshal() (dAtA []byte, err error)
}

func (k Keeper) verifySignature(ctx sdk.Context, owner string, proposal ProposalApi, jwsSignature types.JwsSignature) (string, error) {
	proposalBytes, err := proposal.Marshal()
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrorInvalidProposal, "")
	}

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
	didManager, err := saodid.NewDidManagerWithDid(owner, querySidDocument)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrorInvalidDid, "")
	}

	signature := saodidtypes.JwsSignature{
		Protected: jwsSignature.Protected,
		Signature: jwsSignature.Signature,
	}

	_, err = didManager.VerifyJWS(saodidtypes.GeneralJWS{
		Payload: base64url.Encode(proposalBytes),
		Signatures: []saodidtypes.JwsSignature{
			signature,
		},
	})

	if err != nil {
		return "", sdkerrors.Wrap(types.ErrorInvalidSignature, err.Error())
	}

	kid, err := signature.GetKid()
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrorInvalidSignature, err.Error())
	}
	sigDid, err := saodidutil.KidToDid(kid)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrorInvalidSignature, err.Error())
	}

	return sigDid, nil
}
