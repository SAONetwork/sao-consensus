package keeper

import (
	"context"
	saodidparser "github.com/SaoNetwork/sao-did/parser"
	"strings"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePaymentAddress(goCtx context.Context, msg *types.MsgUpdatePaymentAddress) (*types.MsgUpdatePaymentAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)
	k.CheckCreator(ctx, msg.Creator, msg.Did)
	accId := msg.GetAccountId()
	accIdSplits := strings.Split(accId, ":")
	if len(accIdSplits) == 3 && accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		// err if did is a key did or empty, which means update payment address for sid
		did, err := saodidparser.Parse(msg.Did)
		if err != nil {
			logger.Error("failed to parse did %v", msg.Did)
			return nil, types.ErrInvalidDid
		}
		switch did.Method {
		case "sid":
			proof, found := k.GetDidBindingProof(ctx, accId)
			if !found {
				logger.Error("account id %v has not bound to did %v yet", msg.AccountId, msg.Did)
				return nil, types.ErrBindingNotFound
			}
			if msg.Did != proof.Proof.Did {
				logger.Error("did %v in binding proof is different to did %v in message", proof.Proof.Did, msg.Did)
				return nil, types.ErrInconsistentDid
			}

			paymentAddress := types.PaymentAddress{
				Did:     proof.Proof.Did,
				Address: accIdSplits[2],
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		case "key":
			paymentAddress := types.PaymentAddress{
				Did:     msg.Did,
				Address: accIdSplits[2],
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		default:
			logger.Error("unsupported did %v", msg.Did)
			return nil, types.ErrUnsupportedDid
		}
	} else {
		logger.Error("you should use a cosmos account in %v chain, but we get %v ", ctx.ChainID(), msg.AccountId)
		return nil, types.ErrInvalidAccountId
	}

	return &types.MsgUpdatePaymentAddressResponse{}, nil
}
