package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
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
		keysmap := make(map[string]string)
		for _, key := range msg.Keys {
			keysmap[key.Name] = key.Value
		}

		keysBytes, err := json.Marshal(keysmap)
		if err != nil {
			return nil, types.ErrDocInvalidKeys
		}

		timestamp := proof.Timestamp

		newDocId := hex.EncodeToString(crypto.Sha256([]byte(string(keysBytes) + fmt.Sprint(timestamp))))

		// verify and se sid document if sid is new
		if newDocId != rootDocId || did != "did:sid:"+newDocId {
			fmt.Println(string(keysBytes) + fmt.Sprint(timestamp))
			fmt.Println("new: ", newDocId, "root:", rootDocId, did)
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
	logger := k.Logger(ctx)
	accIdSplits := strings.Split(accId, ":")
	if len(accIdSplits) != 3 {
		logger.Error("failed to parse accountId!! accountId : %v", accId)
		return types.ErrInvalidAccountId
	}
	if accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		// cosmos
		signBytes := getSignData(accIdSplits[2], proof.Message)
		fmt.Println(string(signBytes))

		splitedSig := strings.Split(proof.Signature, ".")
		if splitedSig[0] != "tendermint/PubKeySecp256k1" {
			logger.Error("Unsupported public key type %v!! accountId : %v", splitedSig[0], accId)
			return types.ErrInvalidBindingProof
		}

		pkBytes, err := base64.StdEncoding.DecodeString(splitedSig[1])
		if err != nil {
			logger.Error("failed to decode public key!! accountId : %v", accId)
			return types.ErrInvalidBindingProof
		}

		pubkey := secp256k1.PubKey{Key: pkBytes}
		address, err := sdk.Bech32ifyAddressBytes("cosmos", pubkey.Address())
		if err != nil {
			logger.Error("failed to recover address from given pk, accountId : %v", accId)
			return types.ErrInvalidBindingProof
		}
		if address != accIdSplits[2] {
			logger.Error("address %v recovered from pk is not equal to address in accountId %v", address, accId)
			return types.ErrInvalidBindingProof
		}

		sigBytes, err := base64.StdEncoding.DecodeString(splitedSig[2])
		if err != nil {
			logger.Error("failed to decode signature!! accountId : %v", accId)
			return types.ErrInvalidBindingProof
		}

		if !pubkey.VerifySignature(signBytes, sigBytes) {
			logger.Error("Invalid signature!! accountId : %v", accId)
			return types.ErrInvalidBindingProof
		}

		return nil
	} else if accIdSplits[0] == "eip155" { // && accIdSplits[1] == "???"
		// eth
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

func getSignData(address, message string) []byte {
	// TODO: Amino Sign Doc
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	return []byte(`{"account_number":"0","chain_id":"","fee":{"amount":[],"gas":"0"},"memo":"","msgs":[{"type":"sign/MsgSignData","value":{"data":"` + encodedMessage + `","signer":"` + address + `"}}],"sequence":"0"}`)
}
