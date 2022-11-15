package order

import (
	"github.com/SaoNetwork/sao/x/order/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// clean expired order
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.CheckOrder(ctx)
}
