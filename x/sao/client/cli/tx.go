package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/SaoNetwork/sao/x/sao/types"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

/*
const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)
*/

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdStore())
	cmd.AddCommand(CmdCancel())
	cmd.AddCommand(CmdComplete())
	cmd.AddCommand(CmdReject())
	cmd.AddCommand(CmdTerminate())
	cmd.AddCommand(CmdReady())
	cmd.AddCommand(CmdRenew())
	cmd.AddCommand(CmdUpdataPermission())
	cmd.AddCommand(CmdMigrate())
	cmd.AddCommand(CmdCompleteShard())
	// this line is used by starport scaffolding # 1

	return cmd
}
