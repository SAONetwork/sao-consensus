package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Hooks struct {
	k Keeper
}

var _ stakingtypes.StakingHooks = Hooks{}

var sharesBeforeModified = sdk.NewDec(0)

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
	hook.verifySuperStorageNodes(ctx, valAddr, nil, false)
	return nil
} // Must be called when a validator is deleted

func (hook Hooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	hook.verifySuperStorageNodes(ctx, valAddr, nil, false)
	return nil
} // Must be called when a validator is bonded

func (hook Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	hook.verifySuperStorageNodes(ctx, valAddr, nil, false)
	return nil
} // Must be called when a validator begins unbonding

func (hook Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
} // Must be called when a delegation is created

func (hook Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	del := hook.k.staking.Delegation(ctx, delAddr, valAddr)
	sharesBeforeModified = del.GetShares()
	return nil
} // Must be called when a delegation's shares are modified

func (hook Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	hook.verifySuperStorageNodes(ctx, valAddr, delAddr, true)
	return nil
} // Must be called when a delegation is removed

func (hook Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	hook.verifySuperStorageNodes(ctx, valAddr, delAddr, false)
	return nil
}

func (hook Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) error {
	return nil
}

func (hook Hooks) verifySuperStorageNodes(ctx sdk.Context, valAddr sdk.ValAddress, accAddr sdk.AccAddress, beforeDeletationRemoved bool) {
	delegations := hook.k.staking.GetValidatorDelegations(ctx, valAddr)

	//Records the shares that the validator shares have not been subtracted at the time of the unbond hook call
	sharesToSub := sdk.NewDec(0)
	if accAddr != nil && !sharesBeforeModified.IsZero() {
		del := hook.k.staking.Delegation(ctx, accAddr, valAddr)
		if sharesBeforeModified.GT(del.GetShares()) {
			sharesToSub = sharesBeforeModified.Sub(del.GetShares())
		} else if beforeDeletationRemoved {
			// all shares will be subtracted after delegation removed
			sharesToSub = del.GetShares()
		}

	}

	verified := make(map[string]bool, 0)
	for _, delegation := range delegations {
		node, found := hook.k.GetNode(ctx, delegation.DelegatorAddress)
		if found && (node.Validator == "" || node.Validator == valAddr.String()) {
			// deal with removed delegation
			if beforeDeletationRemoved && delegation.DelegatorAddress == accAddr.String() {
				if node.Role == types.NODE_SUPER {
					hook.k.SetNormalNode(ctx, node.Creator)
				}
				continue
			}
			// ignore if verified once
			if _, ok := verified[node.Creator]; ok {
				continue
			}

			// check status
			if node.Status&types.NODE_STATUS_SUPER_REQUIREMENT != types.NODE_STATUS_SUPER_REQUIREMENT {
				if node.Role == types.NODE_SUPER {
					hook.k.SetNormalNode(ctx, node.Creator)
				}
				continue
			}

			pledge, found := hook.k.GetPledge(ctx, delegation.DelegatorAddress)
			if !found || pledge.TotalStorage < hook.k.VstorageThreshold(ctx) {
				if node.Role == types.NODE_SUPER {
					hook.k.SetNormalNode(ctx, node.Creator)
				}
				continue
			}

			err := hook.k.CheckDelegationShare(ctx, delegation.DelegatorAddress, valAddr.String(), sharesToSub)
			if err == nil {
				if node.Role == types.NODE_NORMAL {
					verified[node.Creator] = true
					hook.k.SetSuperNode(ctx, node.Creator, valAddr.String())
				}
			} else if node.Role == types.NODE_SUPER {
				hook.k.SetNormalNode(ctx, node.Creator)
			}
		}
	}

	// reset shares before modified
	if !sharesBeforeModified.IsZero() {
		sharesBeforeModified = sdk.NewDec(0)
	}
}
