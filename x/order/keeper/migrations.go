package keeper

import (
	markettypes "github.com/SaoNetwork/sao/x/market/types"
	v2 "github.com/SaoNetwork/sao/x/order/migrations/v2"
	v3 "github.com/SaoNetwork/sao/x/order/migrations/v3"
	v4 "github.com/SaoNetwork/sao/x/order/migrations/v4"
	v5 "github.com/SaoNetwork/sao/x/order/migrations/v5"
	"github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.RefundOrder, m.keeper.storeKey, m.keeper.modelStoreKey, m.keeper.cdc)
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}

func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	deposit := func(amount sdk.Coin) error {
		return m.keeper.bank.SendCoinsFromModuleToModule(ctx, types.ModuleName, markettypes.ModuleName, sdk.Coins{amount})
	}
	return v4.MigrateStore(ctx, m.keeper.storeKey, m.keeper.modelStoreKey, m.keeper.marketStoreKey, m.keeper.cdc, deposit)
}

func (m Migrator) Migrate4to5(ctx sdk.Context) error {
	return v5.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
