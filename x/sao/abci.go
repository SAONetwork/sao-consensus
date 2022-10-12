package sao

import (
	"fmt"

	"github.com/SaoNetwork/sao/x/sao/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// clean expired order
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	fmt.Println("check order")
	k.CheckOrder(ctx)
}
