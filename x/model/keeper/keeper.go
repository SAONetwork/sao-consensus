package keeper

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		order         types.OrderKeeper
		account       types.AccountKeeper
		node          types.NodeKeeper
		did           types.DidKeeper
		market        types.MarketKeeper
		bank          types.BankKeeper
		storeKey      storetypes.StoreKey
		orderStoreKey storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
	}
)

func NewKeeper(
	account types.AccountKeeper,
	order types.OrderKeeper,
	did types.DidKeeper,
	bank types.BankKeeper,
	node types.NodeKeeper,
	market types.MarketKeeper,
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
		did:           did,
		account:       account,
		bank:          bank,
		order:         order,
		node:          node,
		market:        market,
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
