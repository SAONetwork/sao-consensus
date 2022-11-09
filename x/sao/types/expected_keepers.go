package types

import (
	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	// Methods imported from bank should be defined here
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// NodeKeeper
type NodeKeeper interface {
	GetNode(ctx sdk.Context, creator string) (val nodetypes.Node, found bool)

	IncreaseReputation(ctx sdk.Context, nodeId string, value float32) error

	DecreaseReputation(ctx sdk.Context, nodeId string, value float32) error

	RandomSP(ctx sdk.Context, count int) []nodetypes.Node

	OrderPledge(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error
}

// EarnKeeper
type EarnKeeper interface {
	OrderPledge(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error
}
