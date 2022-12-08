package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
	"github.com/tendermint/tendermint/crypto"
	"strings"
)

func (k msgServer) Binding(goCtx context.Context, msg *types.MsgBinding) (*types.MsgBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// sid document
	rootDocId := msg.RootDocId
	proof := msg.GetProof()
	did := proof.Did

	if "did:sid:"+rootDocId != did {
		return nil, types.ErrInconsistentDid
	}

	versions, found := k.GetSidDocumentVersion(ctx, rootDocId)
	if !found {
		keysBytes, err := json.Marshal(msg.Keys)
		if err != nil {
			return nil, types.ErrDocInvalidKeys
		}

		timestamp := proof.Timestamp

		newDocId := hex.EncodeToString(crypto.Sha256([]byte(string(keysBytes) + fmt.Sprint(timestamp))))

		// verify and se sid document if sid is new
		if newDocId != rootDocId || did != "did:sid:"+newDocId {
			return nil, types.ErrInconsistentDocId
		}

		versions = types.SidDocumentVersion{
			DocId:       newDocId,
			VersionList: []string{newDocId},
		}

		_, found = k.GetSidDocument(ctx, newDocId)
		if found {
			return nil, types.ErrDocExists
		}

		k.SetSidDocument(ctx, types.SidDocument{
			VersionId: newDocId,
			Keys:      msg.Keys,
		})

		k.SetSidDocumentVersion(ctx, versions)
	}

	// account auth
	aa := *msg.AccountAuth
	accountList, found := k.GetAccountList(ctx, did)
	if !found {
		accountList = types.AccountList{
			Did:         did,
			AccountDids: []string{aa.AccountDid},
		}
	} else {
		for _, ad := range accountList.AccountDids {
			if ad == aa.AccountDid {
				return nil, types.ErrAuthExists
			}
		}
		accountList.AccountDids = append(accountList.AccountDids, aa.AccountDid)
	}

	_, found = k.GetAccountAuth(ctx, aa.AccountDid)
	if found {
		return nil, types.ErrAuthExists
	}

	k.SetAccountAuth(ctx, aa)
	k.SetAccountList(ctx, accountList)

	// binding proof
	accId := msg.GetAccountId()
	_, exist := k.GetDidBindingProof(ctx, accId)
	if exist {
		return nil, types.ErrBindingExists
	}
	if err := k.verifyProof(ctx, accId, proof); err != nil {
		return nil, err
	}

	newDidBindingProof := types.DidBindingProof{
		AccountId: accId,
		Proof:     proof,
	}
	k.SetDidBindingProof(ctx, newDidBindingProof)

	// set first binding cosmos address as payment address
	accIdSplits := strings.Split(accId, ":")
	if len(accIdSplits) == 3 && accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		_, found := k.GetPaymentAddress(ctx, proof.Did)
		if !found {
			paymentAddress := types.PaymentAddress{
				Did:     proof.Did,
				Address: accIdSplits[2],
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		}
	}

	return &types.MsgBindingResponse{}, nil
}

func (k *Keeper) verifyProof(ctx sdk.Context, accId string, proof *types.BindingProof) error {
	accIdSplits := strings.Split(accId, ":")
	if len(accIdSplits) != 3 {
		return types.ErrInvalidAccountId
	}
	if accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		// cosmos
		accAddr, err := sdk.AccAddressFromBech32(accIdSplits[2])
		if err != nil {
			return types.ErrInvalidAccountId
		}
		acc := k.auth.GetAccount(ctx, accAddr)
		if acc.GetPubKey().VerifySignature([]byte(proof.Message), []byte(proof.Signature)) {
			return nil
		} else {
			return types.ErrInvalidBindingProof
		}
	} else if accIdSplits[0] == "eip155" { // && accIdSplits[1] == "???"
		hash := sha256.Sum256([]byte(proof.Message))
		recoverdPublicKey, err := crypto2.SigToPub(hash[:], []byte(proof.Signature))
		if err != nil {
			return types.ErrInvalidBindingProof
		}

		addr := crypto2.PubkeyToAddress(*recoverdPublicKey)
		if addr.Hex() != accIdSplits[2] {
			return types.ErrInvalidBindingProof
		}
		return nil
	}
	return types.ErrUnsupportedAccountId
}
