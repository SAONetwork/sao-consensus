package keeper

import (
	"context"
	"strings"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePaymentAddress(goCtx context.Context, msg *types.MsgUpdatePaymentAddress) (*types.MsgUpdatePaymentAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	accId := msg.GetAccountId()
	proof, found := k.GetDidBindingProofs(ctx, accId)
	if !found {
		return nil, types.ErrBindingNotFound
	}

	accIdSplits := strings.Split(accId, ":")
	if len(accIdSplits) == 3 && accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		paymentAddress := types.PaymentAddress{
			Did:     proof.Proof.Did,
			Address: accIdSplits[2],
		}
		k.SetPaymentAddress(ctx, paymentAddress)
	} else {
		return nil, types.ErrInvalidAccountId
	}

	return &types.MsgUpdatePaymentAddressResponse{}, nil
}
