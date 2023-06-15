package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Hooks struct {
	k Keeper
}

var _ stakingtypes.StakingHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (hook Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) error {

	return nil
} // Must be called when a validator is created

func (hook Hooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) error {

	return nil
} // Must be called when a validator's state changes

func (hook Hooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	hook.verifySuperStorageNodes(ctx, valAddr)
	return nil
} // Must be called when a validator is deleted

func (hook Hooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	hook.verifySuperStorageNodes(ctx, valAddr)
	return nil
} // Must be called when a validator is bonded

func (hook Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	hook.verifySuperStorageNodes(ctx, valAddr)
	return nil
} // Must be called when a validator begins unbonding

func (hook Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
} // Must be called when a delegation is created

func (hook Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
} // Must be called when a delegation's shares are modified

func (hook Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
} // Must be called when a delegation is removed

func (hook Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	logger := ctx.Logger()
	logger.Debug("after delegation", "deladdr", delAddr, "valaddr", valAddr, "keeper", hook.k, "keeper staking", hook.k.staking)
	delegation := hook.k.staking.Delegation(ctx, delAddr, valAddr)
	logger.Debug("after delegation", "founddelegateion", delegation.GetShares())
	if !delegation.GetShares().IsZero() {
		node, found := hook.k.GetNode(ctx, delAddr.String())
		logger.Debug("after delegation", "foundnode", found)
		if !found {
			return nil
		}
		// check current share
		logger.Debug("after delegation", "validator", node.Validator)
		if node.Validator != "" {
			err := hook.k.CheckDelegationShare(ctx, delAddr.String(), node.Validator)
			logger.Debug("after delegation", "check current share", err)
			if err == nil {
				return nil
			}
		}
		// check new delegation share
		err := hook.k.CheckDelegationShare(ctx, delAddr.String(), valAddr.String())
		logger.Debug("after delegation", "check delegation share", err)
		if err == nil && node.Role == 0 {
			logger.Debug("after delegate set super", "sp", delAddr.String(), "val", valAddr.String())
			hook.k.SetSuperNode(ctx, delAddr.String(), valAddr.String())
		} else if node.Role == 1 {
			hook.k.SetNormalNode(ctx, delAddr.String())
		}
	}
	return nil
}

func (hook Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) error {
	return nil
}

func (hook Hooks) verifySuperStorageNodes(ctx sdk.Context, valAddr sdk.ValAddress) {
	delegations := hook.k.staking.GetValidatorDelegations(ctx, valAddr)
	verified := make(map[string]bool, 0)
	for _, delegation := range delegations {
		node, found := hook.k.GetNode(ctx, delegation.DelegatorAddress)
		if found {
			// ignore if verified once
			if _, ok := verified[node.Creator]; ok {
				continue
			}
			err := hook.k.CheckDelegationShare(ctx, delegation.DelegatorAddress, valAddr.String())
			if err == nil && node.Role == 0 {
				verified[node.Creator] = true
				hook.k.SetSuperNode(ctx, node.Creator, valAddr.String())
			} else if node.Role == 1 {
				hook.k.SetNormalNode(ctx, node.Creator)
			}
		}
	}
}
