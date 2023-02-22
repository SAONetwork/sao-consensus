package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto2 "github.com/ethereum/go-ethereum/crypto"
)

const EXPIRE_DURATION uint64 = 15 * 60
const DEFAULT_NETWORK = "cosmos"

func (k msgServer) Binding(goCtx context.Context, msg *types.MsgBinding) (*types.MsgBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	rootDocId := msg.RootDocId
	accAuth := *msg.AccountAuth
	proof := msg.GetProof()
	did := proof.Did
	accId := msg.GetAccountId()

	// parameter verify
	if "did:sid:"+rootDocId != did {
		logger.Error("rootDocId does not match did in bindingProof", "docId", rootDocId, "did", did)
		return nil, types.ErrInconsistentDid
	}

	now := time.Now().Unix()
	if proof.Timestamp+EXPIRE_DURATION < uint64(now) {
		logger.Error("timestamp is too old", "proof.Timestamp", proof.Timestamp, "now", now)
		return nil, types.ErrOutOfDate
	}

	caip10, err := parseAcccountId(accId)
	if err != nil {
		logger.Error("failed to parse accountId!!", "accountId", accId, "did", did, "err", err)
		return nil, types.ErrInvalidAccountId
	}

	accountList, foundAccList := k.GetAccountList(ctx, did)
	if foundAccList {
		for _, ad := range accountList.AccountDids {
			if ad == accAuth.AccountDid {
				logger.Error("accountDid exists in account list", "did", did, "accountDid", ad)
				return nil, types.ErrAuthExists
			}
		}
	}

	_, found := k.GetAccountAuth(ctx, accAuth.AccountDid)
	if found {
		logger.Error("account Auth exists", "did", did, "accountDid", accAuth.AccountDid)
		return nil, types.ErrAuthExists
	}

	storedAccountId, foundAccId := k.GetAccountId(ctx, accAuth.AccountDid)
	if foundAccId && storedAccountId.AccountId != accId {
		logger.Error("accountId exists but not equal!!", "storedAccountId", storedAccountId.AccountId, "accountId", accId)
		return nil, types.ErrInvalidAccountId
	}

	_, found = k.GetDid(ctx, accId)
	if found {
		logger.Error("account has been bound", "accountId", accId)
		return nil, types.ErrBindingExists
	}

	if err := k.verifyProof(ctx, caip10, proof); err != nil {
		logger.Error("verify proof failed!!", "accountId", accId, "err", err)
		return nil, err
	}

	versions, found := k.GetSidDocumentVersion(ctx, rootDocId)
	if found {
		// if sid exists, check creator is bound to sid
		if !k.CheckCreator(ctx, msg.Creator, did) {
			logger.Error("Creator should bind to did.", "creator", msg.Creator, "did", did)
			return nil, types.ErrInvalidCreator
		}
	} else {
		if msg.Creator != caip10.Address {
			logger.Error("Creator should be the first account binding to sid", "bindingAccount", caip10.Address, "creator", msg.Creator)
			return nil, types.ErrInvalidCreator
		}

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

		_, found = k.GetSidDocument(ctx, newDocId)
		if found {
			logger.Error("docId exists", "doc_id", rootDocId, "did", did)
			return nil, types.ErrDocExists
		}

		k.SetSidDocument(ctx, types.SidDocument{
			VersionId: newDocId,
			Keys:      msg.Keys,
		})

		versions = types.SidDocumentVersion{
			DocId:       newDocId,
			VersionList: []string{newDocId},
		}

		k.SetSidDocumentVersion(ctx, versions)
	}

	// account auth
	if !foundAccList {
		accountList = types.AccountList{
			Did:         did,
			AccountDids: []string{accAuth.AccountDid},
		}
	} else {
		accountList.AccountDids = append(accountList.AccountDids, accAuth.AccountDid)
	}

	k.SetAccountAuth(ctx, accAuth)
	k.SetAccountList(ctx, accountList)
	k.SetDid(ctx, types.Did{
		AccountId: accId,
		Did:       proof.Did,
	})

	// accountId
	if !foundAccId {
		k.SetAccountId(ctx, types.AccountId{AccountDid: accAuth.AccountDid, AccountId: accId})
	}

	// set first binding cosmos address as payment address
	if caip10.Network == DEFAULT_NETWORK && caip10.Chain == ctx.ChainID() {
		_, found := k.GetPaymentAddress(ctx, proof.Did)
		if !found {
			paymentAddress := types.PaymentAddress{
				Did:     proof.Did,
				Address: caip10.Address,
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		}
	}

	return &types.MsgBindingResponse{}, nil
}

func (k *Keeper) verifyProof(ctx sdk.Context, caip10 types.Caip10AccountId, proof *types.BindingProof) error {
	logger := k.Logger(ctx)
	accId := caip10.ToString()
	if caip10.Network == DEFAULT_NETWORK && caip10.Chain == ctx.ChainID() {
		// sao
		signBytes := GetSignData(caip10.Address, proof.Message)

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
		address, err := sdk.Bech32ifyAddressBytes("sao", pubkey.Address())
		if err != nil {
			logger.Error("failed to recover address from given pk", "accountId", accId, "err", err)
			return types.ErrInvalidBindingProof
		}
		if address != caip10.Address {
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
	} else if caip10.Network == "eip155" { // && accIdSplits[1] == "???"
		// eth
		hash := sha256.Sum256([]byte(proof.Message))
		recoveredPublicKey, err := crypto2.SigToPub(hash[:], []byte(proof.Signature))
		if err != nil {
			logger.Error("failed to recover pk!!", "accountId", accId, "err", err)
			return types.ErrInvalidBindingProof
		}

		addr := crypto2.PubkeyToAddress(*recoveredPublicKey)
		if addr.Hex() != caip10.Address {
			logger.Error("inconsistent addre!!", "recovered", addr, "accountId", accId)
			return types.ErrInvalidBindingProof
		}
		return nil
	}
	logger.Error("unsupported accountId!!", "accountId", accId)
	return types.ErrUnsupportedAccountId
}
