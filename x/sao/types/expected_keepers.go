package types

import (
	types2 "github.com/SaoNetwork/sao/x/did/types"
	modeltypes "github.com/SaoNetwork/sao/x/model/types"
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
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
}

// NodeKeeper
type NodeKeeper interface {
	GetNode(ctx sdk.Context, creator string) (val nodetypes.Node, found bool)

	IncreaseReputation(ctx sdk.Context, nodeId string, value float32) error

	DecreaseReputation(ctx sdk.Context, nodeId string, value float32) error

	RandomSP(ctx sdk.Context, order ordertypes.Order) []nodetypes.Node

	OrderPledge(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error

	OrderRelease(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error
}

// EarnKeeper
type EarnKeeper interface {
	OrderPledge(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error
}

// OrderKeeper interface
type OrderKeeper interface {
	NewOrder(ctx sdk.Context, order *ordertypes.Order, sp []string) (uint64, error)
	GenerateShards(ctx sdk.Context, order *ordertypes.Order, sps []string)
	GetOrder(ctx sdk.Context, orderId uint64) (ordertypes.Order, bool)
	SetOrder(ctx sdk.Context, order ordertypes.Order)
	TerminateOrder(ctx sdk.Context, orderId uint64) error
	FulfillShard(ctx sdk.Context, order *ordertypes.Order, sp string, cid string, size int32) error
	TerminateShard(ctx sdk.Context, shard *ordertypes.Shard, sp string, owner string, orderId uint64) error
}

// ModelKeeper
type ModelKeeper interface {
	NewMeta(ctx sdk.Context, order ordertypes.Order) error

	GetModel(ctx sdk.Context, key string) (val modeltypes.Model, found bool)

	GetMetadata(ctx sdk.Context, dataId string) (val modeltypes.Metadata, found bool)

	UpdateMeta(ctx sdk.Context, order ordertypes.Order) error

	DeleteMeta(ctx sdk.Context, dataId string) error

	UpdatePermission(ctx sdk.Context, owner string, dataId string, readonlyDids []string, readwriteDids []string) error
}

// DidKeeper
type DidKeeper interface {
	GetCosmosPaymentAddress(ctx sdk.Context, did string) (sdk.AccAddress, error)
	GetSidDocument(ctx sdk.Context, versionId string) (val types2.SidDocument, found bool)
	ValidDid(ctx sdk.Context, did string) error
}

// MarketKeeper
type MarketKeeper interface {
	Deposit(ctx sdk.Context, order ordertypes.Order) error
	Withdraw(ctx sdk.Context, order ordertypes.Order) error
}
