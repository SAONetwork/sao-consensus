package v4

import (
	v3node "github.com/SaoNetwork/sao/x/node/migrations/v3/types"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {

	store := ctx.KVStore(storeKey)

	pledgeStore := prefix.NewStore(store, types.KeyPrefix(types.PledgeKeyPrefix))
	pledgeIterator := sdk.KVStorePrefixIterator(pledgeStore, []byte{})

	defer pledgeIterator.Close()

	for ; pledgeIterator.Valid(); pledgeIterator.Next() {
		var val v3node.Pledge
		pledgeKey := pledgeIterator.Key()
		cdc.MustUnmarshal(pledgeIterator.Value(), &val)

		pledge := types.Pledge{
			Creator:             val.Creator,
			TotalStoragePledged: val.TotalStoragePledged,
			Reward:              val.Reward,
			RewardDebt:          val.RewardDebt,
			TotalStorage:        val.TotalStorage,
			LoanStrategy:        0,
			LoanPledged:         sdk.NewCoin(val.Reward.Denom, sdk.NewInt(0)),
			InterestDebt:        sdk.NewDecCoin(val.Reward.Denom, sdk.NewInt(0)),
		}

		newVal := cdc.MustMarshal(&pledge)

		pledgeStore.Set(pledgeKey, newVal)

	}
	return nil
}