package v014

import (
	nodekeeper "github.com/SaoNetwork/sao/x/node/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	node nodekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		setNodeBlockReward(ctx, node)

		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func setNodeBlockReward(ctx sdk.Context, nk nodekeeper.Keeper) {
	nodeParams := nk.GetParams(ctx)
	nodeParams.BlockReward.Amount = nodeParams.BlockReward.Amount.MulRaw(2)
	nk.SetParams(ctx, nodeParams)
}
