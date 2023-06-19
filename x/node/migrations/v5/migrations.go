package v5

import (
	v4node "github.com/SaoNetwork/sao/x/node/migrations/v4/types"
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
		var val v4node.Pledge
		pledgeKey := pledgeIterator.Key()
		cdc.MustUnmarshal(pledgeIterator.Value(), &val)

		pledge := types.Pledge{
			Creator:             val.Creator,
			TotalStoragePledged: val.TotalStoragePledged,
			TotalShardPledged:   val.TotalShardPledged,
			Reward:              val.Reward,
			RewardDebt:          val.RewardDebt,
			TotalStorage:        val.TotalStorage,
			UsedStorage:         val.UsedStorage,
			LoanStrategy:        0,
			LoanPledged:         sdk.NewCoin(val.Reward.Denom, sdk.NewInt(0)),
			InterestDebt:        sdk.NewDecCoin(val.Reward.Denom, sdk.NewInt(0)),
		}

		newVal := cdc.MustMarshal(&pledge)

		pledgeStore.Set(pledgeKey, newVal)

	}
	return nil
}
