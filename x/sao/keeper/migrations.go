package keeper

import (
	"github.com/SaoNetwork/sao/x/sao/migrations/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.SetExpiredShardsBlock, m.keeper.orderStoreKey, m.keeper.cdc)
}
