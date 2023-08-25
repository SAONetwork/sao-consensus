package v017

import (
	didkeeper "github.com/SaoNetwork/sao/x/did/keeper"
	nodekeeper "github.com/SaoNetwork/sao/x/node/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	node nodekeeper.Keeper,
	did didkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		did.SetBuiltinDids(ctx, "did:key:zQ3shggYEtCZNEiwSeqLdLo97SqS2ERMHB2mgV8hmCGDn4DJ3z")

		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}
