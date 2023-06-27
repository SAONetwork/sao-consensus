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
	hook.verifySuperStorageNodes(ctx, valAddr)
	return nil
}

func (hook Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) error {
	return nil
}

func (hook Hooks) verifySuperStorageNodes(ctx sdk.Context, valAddr sdk.ValAddress) {
	delegations := hook.k.staking.GetValidatorDelegations(ctx, valAddr)

	unbondingList := make(map[string]sdk.Dec, 0)

	for _, unbonding := range hook.k.staking.GetUnbondingDelegationsFromValidator(ctx, valAddr) {
		if _, ok := unbondingList[unbonding.DelegatorAddress]; !ok {
			unbondingList[unbonding.DelegatorAddress] = sdk.NewDec(0)
		}
		for _, entry := range unbonding.Entries {
			unbondingList[unbonding.DelegatorAddress].AddMut(sdk.NewDecFromInt(entry.Balance))
		}
	}

	verified := make(map[string]bool, 0)
	for _, delegation := range delegations {
		node, found := hook.k.GetNode(ctx, delegation.DelegatorAddress)
		if found {
			// ignore if verified once
			if _, ok := verified[node.Creator]; ok {
				continue
			}
			unbonding := sdk.NewDec(0)
			if _, ok := unbondingList[delegation.DelegatorAddress]; ok {
				unbonding = unbondingList[delegation.DelegatorAddress]
			}
			err := hook.k.CheckDelegationShare(ctx, delegation.DelegatorAddress, valAddr.String(), unbonding)
			if err == nil && node.Role == 0 {
				verified[node.Creator] = true
				hook.k.SetSuperNode(ctx, node.Creator, valAddr.String())
			} else if node.Role == 1 {
				hook.k.SetNormalNode(ctx, node.Creator)
			}
		}
	}
}
