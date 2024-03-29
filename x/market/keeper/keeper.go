package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/market/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		bank          types.BankKeeper
		order         types.OrderKeeper
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		orderStoreKey storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
	}
)

func NewKeeper(
	bank types.BankKeeper,
	order types.OrderKeeper,
	cdc codec.BinaryCodec,
	storeKey,
	orderStoreKey storetypes.StoreKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		bank:          bank,
		order:         order,
		cdc:           cdc,
		storeKey:      storeKey,
		orderStoreKey: orderStoreKey,
		memKey:        memKey,
		paramstore:    ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
