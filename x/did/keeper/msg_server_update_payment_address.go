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

	// err if did is empty
	did, err := saodidparser.Parse(msg.Did)
	if err != nil {
		logger.Error("failed to parse did", "did", msg.Did)
		return nil, types.ErrInvalidDid
	}

	accId := msg.GetAccountId()
	caip10, err := parseAcccountId(accId)
	if err != nil {
		logger.Error("failed to parse accountId!!", "accountId", accId, "did", msg.Did, "err", err)
		return nil, types.ErrInvalidAccountId
	}

	OldAddr, found := k.GetPaymentAddress(ctx, msg.Did)
	if found {
		if did.Method == "key" {
			return nil, types.ErrChangePayAddr
		}
		if OldAddr.Address == caip10.Address {
			logger.Error("try to update the same address as the old one", "paymentAddress", OldAddr)
			return nil, types.ErrSamePayAddr
		}
	}

	if err := k.CreatorIsBoundToDid(ctx, msg.Creator, msg.Did); err != nil && did.Method != "key" {
		logger.Error("invalid Creator", "creator", msg.Creator, "did", msg.Did)
		return nil, err
	}
	if caip10.Network == DEFAULT_NETWORK && caip10.Chain == ctx.ChainID() {
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
			if caip10.Address != msg.Creator {
				logger.Error("The address other than the creator cannot be set as the payment address of the key did", "creator", msg.Creator, "accountId", msg.AccountId)
				return nil, types.ErrInvalidAccountId
			}

			if kid, found := k.GetKid(ctx, caip10.Address); found {
				logger.Error("creator has been bound to a kid", "creator", msg.Creator, "kid", kid.Kid)
				return nil, types.ErrKidExist
			}

			paymentAddress := types.PaymentAddress{
				Did:     msg.Did,
				Address: caip10.Address,
			}
			k.SetPaymentAddress(ctx, paymentAddress)

			kid := types.Kid{
				Address: caip10.Address,
				Kid:     msg.Did,
			}
			k.SetKid(ctx, kid)

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
