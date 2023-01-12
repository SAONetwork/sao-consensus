package keeper

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	"strings"

	"github.com/SaoNetwork/sao-did/parser"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCosmosPaymentAddress(ctx sdk.Context, did string) (sdk.AccAddress, error) {
	// return sdk.MustAccAddressFromBech32("cosmos1npkx93adc2ml2usfg4hxfpkqzhzxlk2w4hegpe"), nil
	paymentAddress, found := k.GetPaymentAddress(ctx, did)
	if !found {
		return nil, types.ErrPayAddrNotSet
	}
	return sdk.MustAccAddressFromBech32(paymentAddress.Address), nil
}

func (k Keeper) ValidDid(ctx sdk.Context, did string) error {
	if did == "all" {
		return nil
	}

	parsedDid, err := parser.Parse(did)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	switch parsedDid.Method {
	case "key":
		// check payment address
		if _, found := k.GetPaymentAddress(ctx, did); !found {
			return status.Error(codes.NotFound, "payment address not found")
		}
	case "sid":
		// check sid document
		if _, found := k.GetSidDocument(ctx, parsedDid.ID); !found {
			return status.Error(codes.NotFound, "sid document not found")
		}

		// check sid document version
		versionList, found := k.GetSidDocumentVersion(ctx, parsedDid.ID)
		if !found {
			return status.Error(codes.InvalidArgument, "sidId should be a rootDocId")
		}

		// check version
		version := getVersionInfo(parsedDid.Query)
		if version != "" {
			if !inList(version, versionList.VersionList) {
				return status.Error(codes.NotFound, "sid version not found")
			}

			if _, found := k.GetSidDocument(ctx, version); !found {
				return status.Error(codes.NotFound, "versioned sid document not found")
			}
		}

		// check payment address
		if _, found := k.GetPaymentAddress(ctx, did); !found {
			return status.Error(codes.NotFound, "payment address not found")
		}

		// check account auth
		if _, found := k.GetAccountList(ctx, did); !found {
			return status.Error(codes.NotFound, "account list not found")
		}

		// TODO: check pastSeeds, check binding accounts
	}
	return nil
}

func getVersionInfo(query string) string {
	// version-id was changed to versionId in the latest did-core spec
	// https://github.com/w3c/did-core/pull/553
	var versionId string
	for _, q := range strings.Split(query, "&") {
		if strings.Contains(q, "versionId") || strings.Contains(q, "version-id") {
			versionId = strings.Split(q, "=")[1]
			break
		}
	}

	return versionId
}

func inList(obj string, list []string) bool {
	for _, v := range list {
		if v == obj {
			return true
		}
	}
	return false
}

func (k Keeper) CheckCreator(ctx sdk.Context, creator, did string) bool {
	logger := k.Logger(ctx)

	parsedDid, err := parser.Parse(did)
	if err != nil {
		logger.Error("check creator: get invalid did", "did", did, "err", err)
		return false
	}

	if parsedDid.Method == "key" {
		return true
	}

	accountId := "cosmos:sao:" + creator

	bindingProof, found := k.GetDidBindingProof(ctx, accountId)
	if !found {
		logger.Error("check creator: binding proof not found", "account id", accountId)
		return false
	}
	return bindingProof.Proof.Did == did
}

func CalculateDocId(keys []*types.PubKey, timestamp uint64) (string, error) {
	keysmap := make(map[string]string)
	for _, key := range keys {
		keysmap[key.Name] = key.Value
	}

	keysBytes, err := json.Marshal(keysmap)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypto.Sha256([]byte(string(keysBytes) + fmt.Sprint(timestamp)))), nil
}
