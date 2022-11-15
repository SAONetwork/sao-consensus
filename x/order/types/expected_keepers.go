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
	// Methods imported from bank should be defined here
}

// NodeKeeper
type NodeKeeper interface {
	GetNode(ctx sdk.Context, creator string) (val nodetypes.Node, found bool)

	IncreaseReputation(ctx sdk.Context, nodeId string, value float32) error

	DecreaseReputation(ctx sdk.Context, nodeId string, value float32) error

	RandomSP(ctx sdk.Context, count int) []nodetypes.Node

	OrderPledge(ctx sdk.Context, sp sdk.AccAddress, amount sdk.Coin) error
}
