package keeper

import (
	"fmt"

	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		ak            types.AccountKeeper
		bank          types.BankKeeper
		order         types.OrderKeeper
		staking       types.StakingKeeper
		market        types.MarketKeeper
		gov           govkeeper.Keeper
		cdc           codec.BinaryCodec
		govStoreKey   storetypes.StoreKey
		storeKey      storetypes.StoreKey
		orderStoreKey storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
	}
)

func NewKeeper(
	ak types.AccountKeeper,
	bank types.BankKeeper,
	order types.OrderKeeper,
	staking types.StakingKeeper,
	market types.MarketKeeper,
	cdc codec.BinaryCodec,
	govStoreKey,
	storeKey,
	memKey storetypes.StoreKey,
	orderStoreKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		ak:            ak,
		bank:          bank,
		order:         order,
		staking:       staking,
		market:        market,
		cdc:           cdc,
		govStoreKey:   govStoreKey,
		storeKey:      storeKey,
		memKey:        memKey,
		orderStoreKey: orderStoreKey,
		paramstore:    ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bank.MintCoins(ctx, types.ModuleName, newCoins)
}

func (k Keeper) Stats(ctx sdk.Context) {
	logger := k.Logger(ctx)

	addr := k.ak.GetModuleAddress(types.ModuleName)
	balances := k.bank.GetAllBalances(ctx, addr)

	logger.Debug("node module ", "balances", balances)
}
