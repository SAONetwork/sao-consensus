package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/SaoNetwork/sao/x/node/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group node queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListNode())
	cmd.AddCommand(CmdShowNode())
	cmd.AddCommand(CmdShowPool())
	cmd.AddCommand(CmdListPledge())
	cmd.AddCommand(CmdShowPledge())
	cmd.AddCommand(CmdListShard())
	cmd.AddCommand(CmdShowShard())
	// this line is used by starport scaffolding # 1

	return cmd
}
