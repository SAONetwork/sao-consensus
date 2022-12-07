package keeper

import (
	"context"
	saodid "github.com/SaoNetwork/sao-did"
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
		if _, err := saodid.NewDidManagerWithDid(msg.Did, nil); err != nil {
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
		} else {
			paymentAddress := types.PaymentAddress{
				Did:     msg.Did,
				Address: accIdSplits[2],
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		}
	} else {
		return nil, types.ErrInvalidAccountId
	}

	return &types.MsgUpdatePaymentAddressResponse{}, nil
}
