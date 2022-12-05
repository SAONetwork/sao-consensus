package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/SaoNetwork/sao/x/did/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAddBinding())
	cmd.AddCommand(CmdUnbinding())
	cmd.AddCommand(CmdAddAccountAuth())
	cmd.AddCommand(CmdUpdateAccountAuths())
	cmd.AddCommand(CmdUpdateSidDocument())
	cmd.AddCommand(CmdAddPastSeed())
	cmd.AddCommand(CmdResetStore())
	cmd.AddCommand(CmdUpdatePaymentAddress())
	// this line is used by starport scaffolding # 1

	return cmd
}
