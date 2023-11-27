package keeper

import (
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
	builtinDids := k.GetBuiltinDids(ctx)
	if strings.Contains(builtinDids, did) {
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
			return status.Error(codes.Aborted, "sidId should be a rootDocId")
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

		// check past seeds
		pastSeeds, found := k.GetPastSeeds(ctx, did)
		if len(versionList.VersionList) > 1 && len(pastSeeds.Seeds)+1 != len(versionList.VersionList) ||
			found && len(versionList.VersionList) == 1 {
			return status.Error(codes.Aborted, "Invalid pastSeeds length")
		}
	}
	return nil
}

func (k Keeper) CreatorIsBoundToDid(ctx sdk.Context, creator, did string) error {
	logger := k.Logger(ctx)

	accountId := "cosmos:" + ctx.ChainID() + ":" + creator

	storedDid, found := k.GetDid(ctx, accountId)
	if !found {
		logger.Error("check creator: binding did not found", "account id", accountId)
		return types.ErrInvalidCreator
	}
	if storedDid.Did == did {
		return nil
	} else {
		logger.Error("check creator: inconsistent did", "did", did, "storedDid", storedDid.Did)
		return types.ErrInvalidCreator
	}
}

func (k Keeper) SendCoinsFromModuleToDidBalances(ctx sdk.Context, module string, did string, amount sdk.Coin) error {
	if amount.IsZero() {
		return nil
	}

	balances, found := k.GetDidBalances(ctx, did)
	if found {
		balances.Balance.Add(amount)
	} else {
		balances = types.DidBalances{
			Did:     did,
			Balance: amount,
		}
	}

	k.SetDidBalances(ctx, balances)

	err := k.bank.SendCoinsFromModuleToModule(ctx, module, types.ModuleName, sdk.Coins{amount})
	if err != nil {
		return err
	}

	return nil
}
