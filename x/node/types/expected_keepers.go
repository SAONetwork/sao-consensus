package types

import (
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
	NewAccount(sdk.Context, types.AccountI) types.AccountI
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) types.AccountI

	GetAllAccounts(ctx sdk.Context) []types.AccountI
	HasAccount(ctx sdk.Context, addr sdk.AccAddress) bool
	SetAccount(ctx sdk.Context, acc types.AccountI)

	IterateAccounts(ctx sdk.Context, process func(types.AccountI) bool)

	ValidatePermissions(macc types.ModuleAccountI) error

	GetModuleAddress(moduleName string) sdk.AccAddress
	GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string)
	GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (types.ModuleAccountI, []string)
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI
	SetModuleAccount(ctx sdk.Context, macc types.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
}

type StakingKeeper interface {
	BondDenom(ctx sdk.Context) string
}

// OrderKeeper interface
type OrderKeeper interface {
	NewOrder(ctx sdk.Context, order *ordertypes.Order, sp []string) (uint64, error)
	GenerateShards(ctx sdk.Context, order *ordertypes.Order, sps []string)
	MigrateShard(ctx sdk.Context, order *ordertypes.Order, from string, to string) *ordertypes.Shard
	GetOrder(ctx sdk.Context, orderId uint64) (ordertypes.Order, bool)
	SetOrder(ctx sdk.Context, order ordertypes.Order)
	TerminateOrder(ctx sdk.Context, orderId uint64, refundCoin sdk.Coin) error
	FulfillShard(ctx sdk.Context, order *ordertypes.Order, sp string, cid string, size uint64) error
	TerminateShard(ctx sdk.Context, shard *ordertypes.Shard, sp string, owner string, orderId uint64) error
	GetOrderShardBySP(ctx sdk.Context, order *ordertypes.Order, sp string) *ordertypes.Shard
	GetShard(ctx sdk.Context, id uint64) (val ordertypes.Shard, found bool)
	RemoveShard(ctx sdk.Context, id uint64)
	SetShard(ctx sdk.Context, shard ordertypes.Shard)
}

type MarketKeeper interface {
	Claim(ctx sdk.Context, denom string, sp string) (sdk.Coin, error)
}
