package keeper

import (
	"context"

	saodidparser "github.com/SaoNetwork/sao-did/parser"
	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePaymentAddress(goCtx context.Context, msg *types.MsgUpdatePaymentAddress) (*types.MsgUpdatePaymentAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	did, err := saodidparser.Parse(msg.Did)
	if err := k.CheckCreator(ctx, msg.Creator, msg.Did); err != nil {
		if _, found := k.GetPaymentAddress(ctx, msg.Did); found || did.Method != "key" {
			logger.Error("invalid Creator", "creator", msg.Creator, "did", msg.Did)
			return nil, err
		}
	}

	accId := msg.GetAccountId()
	caip10, err := parseAcccountId(accId)
	if err != nil {
		logger.Error("failed to parse accountId!!", "accountId", accId, "did", msg.Did, "err", err)
		return nil, types.ErrInvalidAccountId
	}

	OldAddr, found := k.GetPaymentAddress(ctx, msg.Did)
	if found {
		if OldAddr.Address == caip10.Address {
			logger.Error("try to update the same address as the old one", "paymentAddress", OldAddr)
			return nil, types.ErrSamePayAddr
		}
	}

	if caip10.Network == DEFAULT_NETWORK && caip10.Chain == ctx.ChainID() {
		// err if did is empty, which means update payment address for sid
		if err != nil {
			logger.Error("failed to parse did", "did", msg.Did)
			return nil, types.ErrInvalidDid
		}
		switch did.Method {
		case "sid":
			storedDid, found := k.GetDid(ctx, accId)
			if !found {
				logger.Error("account id has not bound to did yet", "accountId", msg.AccountId, "did", msg.Did)
				return nil, types.ErrBindingNotFound
			}
			if msg.Did != storedDid.Did {
				logger.Error("did in binding proof is different to did in message", "Did(stored)", storedDid.Did, "Did(msg)", msg.Did)
				return nil, types.ErrInconsistentDid
			}

			paymentAddress := types.PaymentAddress{
				Did:     storedDid.Did,
				Address: caip10.Address,
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		case "key":
			paymentAddress := types.PaymentAddress{
				Did:     msg.Did,
				Address: caip10.Address,
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		default:
			logger.Error("unsupported did", "did", msg.Did)
			return nil, types.ErrUnsupportedDid
		}
	} else {
		logger.Error("invalid chain address", "expectedChainId", ctx.ChainID(), "accountId", msg.AccountId)
		return nil, types.ErrInvalidAccountId
	}

	return &types.MsgUpdatePaymentAddressResponse{}, nil
}
