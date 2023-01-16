package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func (k msgServer) Binding(goCtx context.Context, msg *types.MsgBinding) (*types.MsgBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	// TODO: move parameter valitation to here

	// sid document
	rootDocId := msg.RootDocId
	proof := msg.GetProof()
	did := proof.Did
	accId := msg.GetAccountId()

	if "did:sid:"+rootDocId != did {
		logger.Error("rootDocId does not match did in bindingProof", "docId", rootDocId, "did", did)
		return nil, types.ErrInconsistentDid
	}

	versions, found := k.GetSidDocumentVersion(ctx, rootDocId)
	if !found {
		newDocId, err := CalculateDocId(msg.Keys, proof.Timestamp)
		if err != nil {
			logger.Error("failed to calculate doc Id", "did", did, "err", err)
			return nil, types.ErrInvalidKeys
		}

		// verify and set sid document if sid is new
		if newDocId != rootDocId || did != "did:sid:"+newDocId {
			logger.Error("inconsistent docId", "doc_id", rootDocId, "did", did)
			return nil, types.ErrInconsistentDocId
		}

		versions = types.SidDocumentVersion{
			DocId:       newDocId,
			VersionList: []string{newDocId},
		}

		_, found = k.GetSidDocument(ctx, newDocId)
		if found {
			logger.Error("docId exists", "doc_id", rootDocId, "did", did)
			return nil, types.ErrDocExists
		}

		k.SetSidDocument(ctx, types.SidDocument{
			VersionId: newDocId,
			Keys:      msg.Keys,
		})

		k.SetSidDocumentVersion(ctx, versions)
	} else {
		//TODO: check creator
		//k.CheckCreator(ctx, msg.Creator, did)
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
				logger.Error("accountDid exists in account list", "did", did, "accountDid", ad)
				return nil, types.ErrAuthExists
			}
		}
		accountList.AccountDids = append(accountList.AccountDids, aa.AccountDid)
	}

	_, found = k.GetAccountAuth(ctx, aa.AccountDid)
	if found {
		logger.Error("account Auth exists", "did", did, "accountDid", aa.AccountDid)
		return nil, types.ErrAuthExists
	}

	k.SetAccountAuth(ctx, aa)
	k.SetAccountList(ctx, accountList)

	// TODO: change to accId - did map
	// binding proof
	_, exist := k.GetDidBindingProof(ctx, accId)
	if exist {
		logger.Error("binding proof exists", "accountId", accId)
		return nil, types.ErrBindingExists
	}

	if err := k.verifyProof(ctx, accId, proof); err != nil {
		logger.Error("verify proof failed!!", "accountId", accId, "err", err)
		return nil, err
	}

	newDidBindingProof := types.DidBindingProof{
		AccountId: accId,
		Proof:     proof,
	}
	k.SetDidBindingProof(ctx, newDidBindingProof)

	// accountId
	storedAccountId, found := k.GetAccountId(ctx, aa.AccountDid)
	if found {
		if storedAccountId.AccountId != accId {
			logger.Error("accountId exists but not equal!!", "storedAccountId", storedAccountId.AccountId, "accountId", accId)
			return nil, types.ErrInvalidAccountId
		}
	} else {
		k.SetAccountId(ctx, types.AccountId{AccountDid: aa.AccountDid, AccountId: accId})
	}

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
	logger := k.Logger(ctx)
	accIdSplits := strings.Split(accId, ":")
	if len(accIdSplits) != 3 {
		logger.Error("failed to parse accountId!!", "accountId", accId)
		return types.ErrInvalidAccountId
	}
	if accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		// cosmos
		signBytes := getSignData(accIdSplits[2], proof.Message)

		splitedSig := strings.Split(proof.Signature, ".")
		if splitedSig[0] != "tendermint/PubKeySecp256k1" {
			logger.Error("Unsupported public key type!!", "accountId", accId, "key_type", splitedSig[0])
			return types.ErrInvalidBindingProof
		}

		pkBytes, err := base64.StdEncoding.DecodeString(splitedSig[1])
		if err != nil {
			logger.Error("failed to decode public key!!", "accountId", accId, "err", err)
			return types.ErrInvalidBindingProof
		}

		pubkey := secp256k1.PubKey{Key: pkBytes}
		address, err := sdk.Bech32ifyAddressBytes("cosmos", pubkey.Address())
		if err != nil {
			logger.Error("failed to recover address from given pk", "accountId", accId, "err", err)
			return types.ErrInvalidBindingProof
		}
		if address != accIdSplits[2] {
			logger.Error("address recovered from pk does not match accountId", "get", address, "accountId", accId)
			return types.ErrInvalidBindingProof
		}

		sigBytes, err := base64.StdEncoding.DecodeString(splitedSig[2])
		if err != nil {
			logger.Error("failed to decode signature!!", "accountId", accId, "err", err)
			return types.ErrInvalidBindingProof
		}

		if !pubkey.VerifySignature(signBytes, sigBytes) {
			logger.Error("Invalid signature!!", "accountId", accId)
			return types.ErrInvalidBindingProof
		}

		return nil
	} else if accIdSplits[0] == "eip155" { // && accIdSplits[1] == "???"
		// eth
		hash := sha256.Sum256([]byte(proof.Message))
		recoveredPublicKey, err := crypto2.SigToPub(hash[:], []byte(proof.Signature))
		if err != nil {
			logger.Error("failed to recover pk!!", "accountId", accId, "err", err)
			return types.ErrInvalidBindingProof
		}

		addr := crypto2.PubkeyToAddress(*recoveredPublicKey)
		if addr.Hex() != accIdSplits[2] {
			logger.Error("inconsistent addre!!", "recovered", addr, "accountId", accId)
			return types.ErrInvalidBindingProof
		}
		return nil
	}
	logger.Error("unsupported accountId!!", "accountId", accId)
	return types.ErrUnsupportedAccountId
}
