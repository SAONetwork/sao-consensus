package keeper

import (
	"context"
	"strings"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddBinding(goCtx context.Context, msg *types.MsgAddBinding) (*types.MsgAddBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, exist := k.GetDidBingingProof(ctx, msg.GetAccountId())
	if exist {
		return nil, types.ErrBindingExists
	}
	// TODO : add binding proof verify
	accId := msg.GetAccountId()
	proof := msg.GetProof()
	newDidBingingProof := types.DidBingingProof{
		AccountId: accId,
		Proof:     proof,
	}
	k.SetDidBingingProof(ctx, newDidBingingProof)

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

	return &types.MsgAddBindingResponse{}, nil
}
