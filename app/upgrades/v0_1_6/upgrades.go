package v016

import (
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	staking stakingkeeper.Keeper,
	bank bankkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		validators := staking.GetAllValidators(ctx)

		genesis, _ := sdk.AccAddressFromBech32("sao1vnsq3s98qjtrwqlrthtvdv6nhz72ztada4uwlx")

		MintNewToken(ctx, bank, genesis)

		UnbondDelegations(ctx, staking)

		ChangeBondDenom(ctx, staking)

		NewDelegations(ctx, validators, genesis, staking)

		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func MintNewToken(ctx sdk.Context, bank bankkeeper.Keeper, genesis sdk.AccAddress) {
	coins := sdk.Coins{sdk.NewInt64Coin("usao", 10000000000000)}
	bank.MintCoins(ctx, nodetypes.ModuleName, coins)
	bank.SendCoinsFromModuleToAccount(ctx, nodetypes.ModuleName, genesis, coins)
}

func UnbondDelegations(ctx sdk.Context, staking stakingkeeper.Keeper) {
	logger := ctx.Logger().With("upgrade", UpgradeName)
	staking.IterateAllDelegations(ctx, func(delegation stakingtypes.Delegation) bool {
		addr, _ := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		delegator, _ := sdk.AccAddressFromBech32(delegation.DelegatorAddress)
		time, err := staking.Undelegate(ctx, delegator, addr, delegation.Shares)
		logger.Debug("unbond delegation", "t", time, "err", err, "delegator", delegator, "validator", addr)
		return false
	})
}

func ChangeBondDenom(ctx sdk.Context, staking stakingkeeper.Keeper) {
	params := staking.GetParams(ctx)

	params.BondDenom = "usao"

	staking.SetParams(ctx, params)
}

func NewDelegations(ctx sdk.Context, validators []stakingtypes.Validator, genesis sdk.AccAddress, staking stakingkeeper.Keeper) {
	logger := ctx.Logger().With("upgrade", UpgradeName)

	totalShares := sdk.NewDec(0)
	for _, validator := range validators {
		totalShares = totalShares.Add(validator.DelegatorShares)
	}

	coin := sdk.NewDec(1000000000)
	for _, validator := range validators {
		logger.Debug("new delegation", "validator", "coin", validator.BondedTokens())
		share := validator.DelegatorShares.Quo(totalShares).Mul(coin).TruncateInt()
		logger.Debug("new delegation", "genesis", genesis, "coin", coin)
		staking.Delegate(ctx, genesis, share, stakingtypes.Unbonded, validator, true)
	}

}
