package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		ak         types.AccountKeeper
		bank       types.BankKeeper
		order      types.OrderKeeper
		staking    types.StakingKeeper
		market     types.MarketKeeper
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	ak types.AccountKeeper,
	bank types.BankKeeper,
	order types.OrderKeeper,
	staking types.StakingKeeper,
	market types.MarketKeeper,
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		ak:         ak,
		bank:       bank,
		order:      order,
		staking:    staking,
		market:     market,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
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
