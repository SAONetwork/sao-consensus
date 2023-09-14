package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/SaoNetwork/sao/x/order/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group order queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListOrder())
	cmd.AddCommand(CmdShowOrder())
	cmd.AddCommand(CmdListShard())
	cmd.AddCommand(CmdShowShard())
	cmd.AddCommand(CmdShardListBySp())

	// this line is used by starport scaffolding # 1

	return cmd
}

func OrderStatusInText(status int32) string {
	switch status {
	case types.OrderDataReady:
		return types.TextOrderDataReady
	case types.OrderPending:
		return types.TextOrderPending
	case types.OrderInProgress:
		return types.TextOrderInProgress
	case types.OrderUnexpected:
		return types.TextOrderUnexpected
	case types.OrderCompleted:
		return types.TextOrderCompleted
	case types.OrderCanceled:
		return types.TextOrderCanceled
	case types.OrderExpired:
		return types.TextOrderExpired
	default:
		return types.TextOrderTerminated

	}
}

func ShardStatusInText(status int) string {
	switch status {
	case types.ShardWaiting:
		return types.TextShardWaiting
	case types.ShardRejected:
		return types.TextShardRejected
	case types.ShardCompleted:
		return types.TextShardCompleted
	case types.ShardTerminated:
		return types.TextShardTerminated
	case types.ShardMigrating:
		return types.TextShardMigrating
	default:
		return types.TextShardTimeout
	}
}
