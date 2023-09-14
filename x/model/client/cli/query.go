package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/SaoNetwork/sao/x/model/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group model queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListMetadata())
	cmd.AddCommand(CmdShowMetadata())
	cmd.AddCommand(CmdListModel())
	cmd.AddCommand(CmdShowModel())
	cmd.AddCommand(CmdListExpiredData())
	cmd.AddCommand(CmdShowExpiredData())
	// this line is used by starport scaffolding # 1

	return cmd
}

func MetadataStatusInText(status int32) string {
	switch status {
	case types.MetaNew:
		return types.TextMetaNew
	case types.MetaUpdate:
		return types.TextMetaUpdate
	case types.MetaForceUpdate:
		return types.TextMetaForceUpdate
	case types.MetaRenew:
		return types.TextMetaRenew
	default:
		return types.TextMetaComplete
	}
}
