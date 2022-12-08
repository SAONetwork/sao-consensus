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
	accId := msg.GetAccountId()
	accIdSplits := strings.Split(accId, ":")
	if len(accIdSplits) == 3 && accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		// err if did is a key did or empty, which means update payment address for sid
		did, err := saodidparser.Parse(msg.Did)
		if err != nil {
			return nil, types.ErrInvalidDid
		}
		switch did.Method {
		case "sid":
			proof, found := k.GetDidBingingProof(ctx, accId)
			if !found {
				return nil, types.ErrBindingNotFound
			}
			if msg.Did != proof.Proof.Did {
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
			return nil, types.ErrUnsupportedDid
		}
	} else {
		return nil, types.ErrInvalidAccountId
	}

	return &types.MsgUpdatePaymentAddressResponse{}, nil
}
