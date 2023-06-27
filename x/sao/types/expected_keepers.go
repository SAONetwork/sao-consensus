package types

import (
	types2 "github.com/SaoNetwork/sao/x/did/types"
	modeltypes "github.com/SaoNetwork/sao/x/model/types"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type StakingKeeper interface {
	BondDenom(ctx sdk.Context) string
}

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

	RandomSP(ctx sdk.Context, count int, ignore []string, size int64) []nodetypes.Node

	ShardPledge(ctx sdk.Context, shard *ordertypes.Shard, unitPrice sdk.DecCoin) error

	ShardRelease(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Shard) error

	BlockRewardPledge(duration uint64, size uint64, rewardPerByte sdk.DecCoin) sdk.Dec

	StoreRewardPledge(duration uint64, size uint64, rewardPerByte sdk.DecCoin) sdk.Dec

	SetPledgeDebt(ctx sdk.Context, pledgeDebt nodetypes.PledgeDebt)

	GetPledgeDebt(ctx sdk.Context, sp string) (nodetypes.PledgeDebt, bool)

	GetPool(ctx sdk.Context) (val nodetypes.Pool, found bool)

	SetPool(ctx sdk.Context, pool nodetypes.Pool)

	GetPledge(ctx sdk.Context, sp string) (nodetypes.Pledge, bool)

	SetPledge(ctx sdk.Context, pledge nodetypes.Pledge)

	GetFaultBySpAndShardId(ctx sdk.Context, provider string, shardId uint64) (fault *nodetypes.Fault, found bool)

	SetFault(ctx sdk.Context, fault *nodetypes.Fault)

	GetFault(ctx sdk.Context, faultId string) (val *nodetypes.Fault, found bool)

	RemoveFault(ctx sdk.Context, fault *nodetypes.Fault)

	FishmenInfo(ctx sdk.Context) (fishmenInfo string)

	GetFishingReward(ctx sdk.Context, fishman string) (reward *sdk.Dec, found bool)

	SetFishingReward(ctx sdk.Context, fishman string, reward *sdk.Dec)
}

// EarnKeeper
type EarnKeeper interface {
	OrderPledge(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error
}

// OrderKeeper interface
type OrderKeeper interface {
	NewOrder(ctx sdk.Context, order *ordertypes.Order, sp []string) (uint64, error)
	RenewOrder(ctx sdk.Context, order *ordertypes.Order) (uint64, error)
	GenerateShards(ctx sdk.Context, order *ordertypes.Order, sps []string)
	MigrateShard(ctx sdk.Context, oldShard *ordertypes.Shard, order *ordertypes.Order, from string, to string) *ordertypes.Shard
	GetOrder(ctx sdk.Context, orderId uint64) (ordertypes.Order, bool)
	SetOrder(ctx sdk.Context, order ordertypes.Order)
	RemoveOrder(ctx sdk.Context, orderId uint64)
	FulfillShard(ctx sdk.Context, shard *ordertypes.Shard, sp string, cid string)
	GetOrderShardBySP(ctx sdk.Context, order *ordertypes.Order, sp string) *ordertypes.Shard
	GetShard(ctx sdk.Context, id uint64) (val ordertypes.Shard, found bool)
	RemoveShard(ctx sdk.Context, id uint64)
	NewShardTask(ctx sdk.Context, order *ordertypes.Order, provider string) *ordertypes.Shard
	SetShard(ctx sdk.Context, shard ordertypes.Shard)
}

// ModelKeeper
type ModelKeeper interface {
	NewMeta(ctx sdk.Context, order ordertypes.Order, metadata modeltypes.Metadata) error

	GetModel(ctx sdk.Context, key string) (val modeltypes.Model, found bool)

	GetMetadata(ctx sdk.Context, dataId string) (val modeltypes.Metadata, found bool)

	UpdateMeta(ctx sdk.Context, order ordertypes.Order) error

	DeleteMeta(ctx sdk.Context, dataId string) error

	UpdatePermission(ctx sdk.Context, owner string, dataId string, readonlyDids []string, readwriteDids []string) error

	UpdateMetaStatusAndCommit(ctx sdk.Context, order ordertypes.Order) error

	TerminateOrder(ctx sdk.Context, order ordertypes.Order) error

	CancelOrder(ctx sdk.Context, orderId uint64) error

	ResetMetaDuration(ctx sdk.Context, meta *modeltypes.Metadata)

	ExtendMetaDuration(ctx sdk.Context, dataId string, expiredAt uint64)
}

// DidKeeper
type DidKeeper interface {
	GetCosmosPaymentAddress(ctx sdk.Context, did string) (sdk.AccAddress, error)
	GetSidDocument(ctx sdk.Context, versionId string) (val types2.SidDocument, found bool)
	ValidDid(ctx sdk.Context, did string) error
	CreatorIsBoundToDid(ctx sdk.Context, creator, did string) error
	GetBuiltinDids(ctx sdk.Context) string
}

// MarketKeeper
type MarketKeeper interface {
	Deposit(ctx sdk.Context, order ordertypes.Order) error
	Withdraw(ctx sdk.Context, order ordertypes.Order) (sdk.Coin, error)
	Migrate(ctx sdk.Context, order ordertypes.Order, from, to ordertypes.Shard) error
	WorkerRelease(ctx sdk.Context, order *ordertypes.Order, shard *ordertypes.Shard) error
	WorkerAppend(ctx sdk.Context, order *ordertypes.Order, shard *ordertypes.Shard) error
}
