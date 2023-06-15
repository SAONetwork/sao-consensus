package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetFishingReward set a specific reward in the store from its index
func (k Keeper) SetFishingReward(ctx sdk.Context, fishman string, reward *sdk.Dec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FishingRewardKey))
	store.Set([]byte(fishman), []byte(reward.String()))
}

// GetFishingReward returns a reward from its index
func (k Keeper) GetFishingReward(
	ctx sdk.Context,
	fishman string,
) (reward *sdk.Dec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FishingRewardKey))

	rewardBytes := store.Get([]byte(fishman))
	if rewardBytes == nil {
		return nil, false
	}

	val, err := sdk.NewDecFromStr(string(rewardBytes))
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, false
	}

	return &val, true
}
