package types

import (
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	// Methods imported from bank should be defined here
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// OrderKeeper
type OrderKeeper interface {
	GetOrder(ctx sdk.Context, orderId uint64) (ordertypes.Order, bool)
	RemoveOrder(ctx sdk.Context, id uint64)
	TerminateOrder(ctx sdk.Context, orderId uint64, refundCoin sdk.Coin) error
	RefundOrder(ctx sdk.Context, orderId uint64) error
	GetShard(ctx sdk.Context, id uint64) (val ordertypes.Shard, found bool)
	RemoveShard(ctx sdk.Context, id uint64)
}

// SaoKeeper
type NodeKeeper interface {
	GetNode(ctx sdk.Context, creator string) (val nodetypes.Node, found bool)
	ShardRelease(ctx sdk.Context, sp sdk.AccAddress, shard *ordertypes.Shard) error
}

// DidKeeper
type DidKeeper interface {
	GetCosmosPaymentAddress(ctx sdk.Context, did string) (sdk.AccAddress, error)
}

// MarketKeeper
type MarketKeeper interface {
	Withdraw(ctx sdk.Context, order ordertypes.Order) (sdk.Coin, error)
}
