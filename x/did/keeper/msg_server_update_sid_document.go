package keeper

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k msgServer) UpdateSidDocument(goCtx context.Context, msg *types.MsgUpdateSidDocument) (*types.MsgUpdateSidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	rootDocId := msg.RootDocId
	keysBytes, err := json.Marshal(msg.Keys)
	if err != nil {
		return nil, types.ErrInvalidKeys
	}

	// new SidDocument if rootDocId is empty
	timestamp := ctx.BlockTime().String()
	newDocId := hex.EncodeToString(crypto.Sha256([]byte(string(keysBytes) + timestamp)))
	_, found := k.GetSidDocument(ctx, newDocId)
	if found {
		return nil, types.ErrDocExists
	}

	k.SetSidDocument(ctx, types.SidDocument{
		VersionId: newDocId,
		Keys:      msg.Keys,
	})

	if rootDocId == "" {
		k.SetSidDocumentVersion(ctx, types.SidDocumentVersion{
			DocId:       newDocId,
			VersionList: []string{newDocId},
		})
	} else {
		versions, found := k.GetSidDocumentVersion(ctx, rootDocId)
		if !found {
			return nil, types.ErrVersionsNotFound
		}
		versions.VersionList = append(versions.VersionList, newDocId)

		k.SetSidDocumentVersion(ctx, versions)
	}

	return &types.MsgUpdateSidDocumentResponse{newDocId}, nil
}
