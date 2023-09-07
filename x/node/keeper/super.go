package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math"
)

func (k Keeper) CheckDelegationShare(ctx sdk.Context, delAddr string, valAddr string, sharesToSub sdk.Dec) error {
	// if validator provided, check if it satisfy super node requirement.
	accAddress, err := sdk.AccAddressFromBech32(delAddr)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidDelegate, "%v", err)
	}
	valAddress, err := sdk.ValAddressFromBech32(valAddr)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidDelegate, "%v", err)
	}

	// delegator share
	delegate, found := k.staking.GetDelegation(ctx, accAddress, valAddress)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidDelegate, "delegator %v delegate to validator %v not found", delAddr, valAddr)
	}

	// validator share
	validator, found := k.staking.GetValidator(ctx, valAddress)
	if !found {
		return sdkerrors.Wrapf(types.ErrInvalidDelegate, "query validator error %v", found)
	}

	totalShares := validator.DelegatorShares
	if validator.DelegatorShares.Equal(sharesToSub) {
		return sdkerrors.Wrapf(types.ErrInvalidDelegate, "validator will be removed")
	} else {
		totalShares = totalShares.Sub(sharesToSub)
	}

	ratio := delegate.Shares.Quo(totalShares)

	if ratio.LT(k.ShareThreshold(ctx)) {
		return sdkerrors.Wrapf(types.ErrInvalidDelegate, "insufficient shares in this validator need %.2f but %.2f", k.ShareThreshold(ctx).MustFloat64(), ratio.MustFloat64())
	}

	return nil
}

func (k Keeper) SetSuperNode(ctx sdk.Context, sp string, val string) {
	node, _ := k.GetNode(ctx, sp)
	node.Role = 1
	node.Validator = val
	k.SetNode(ctx, node)
}

func (k Keeper) SetNormalNode(ctx sdk.Context, sp string) {
	node, _ := k.GetNode(ctx, sp)
	node.Role = 0
	k.SetNode(ctx, node)
}

func (k Keeper) CheckNodeShare(ctx sdk.Context, node *types.Node, acc string) bool {
	if node.Validator != "" {
		err := k.CheckDelegationShare(ctx, acc, node.Validator, sdk.NewDec(0))
		if err == nil {
			node.Role = types.NODE_SUPER
			return true
		}
	} else {
		accAddr := sdk.MustAccAddressFromBech32(acc)
		dels := k.staking.GetDelegatorDelegations(ctx, accAddr, math.MaxUint16)
		for _, del := range dels {
			err := k.CheckDelegationShare(ctx, acc, del.ValidatorAddress, sdk.NewDec(0))
			if err == nil {
				node.Role = types.NODE_SUPER
				node.Validator = del.ValidatorAddress

				return true
			}
		}
	}
	return false
}