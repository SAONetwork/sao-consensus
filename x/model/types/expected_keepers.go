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
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// OrderKeeper
type OrderKeeper interface {
	GetOrder(ctx sdk.Context, orderId uint64) (ordertypes.Order, bool)
}

// SaoKeeper
type NodeKeeper interface {
	GetNode(ctx sdk.Context, creator string) (val nodetypes.Node, found bool)
}
