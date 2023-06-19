package keeper

import (
	v2 "github.com/SaoNetwork/sao/x/node/migrations/v2"
	v3 "github.com/SaoNetwork/sao/x/node/migrations/v3"
	v4 "github.com/SaoNetwork/sao/x/node/migrations/v4"
	v5 "github.com/SaoNetwork/sao/x/node/migrations/v5"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.bank, m.keeper.storeKey, m.keeper.orderStoreKey, m.keeper.cdc)
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	err := v3.UpdateNodeParams(ctx, &m.keeper.paramstore)
	if err != nil {
		return err
	}
	return v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	err := v4.UpdateNodeParams(ctx, &m.keeper.paramstore)
	if err != nil {
		return err
	}
	return v4.MigrateStore(ctx, m.keeper.storeKey, m.keeper.orderStoreKey, m.keeper.cdc)
}

func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v5.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
