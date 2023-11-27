package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/SaoNetwork/sao/x/did/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group did queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListAccountList())
	cmd.AddCommand(CmdShowAccountList())
	cmd.AddCommand(CmdListAccountAuth())
	cmd.AddCommand(CmdShowAccountAuth())
	cmd.AddCommand(CmdGetAllAccountAuths())

	cmd.AddCommand(CmdListSidDocument())
	cmd.AddCommand(CmdShowSidDocument())
	cmd.AddCommand(CmdListSidDocumentVersion())
	cmd.AddCommand(CmdShowSidDocumentVersion())
	cmd.AddCommand(CmdListPastSeeds())
	cmd.AddCommand(CmdShowPastSeeds())
	cmd.AddCommand(CmdListPaymentAddress())
	cmd.AddCommand(CmdShowPaymentAddress())
	cmd.AddCommand(CmdValidateDid())

	cmd.AddCommand(CmdListAccountId())
	cmd.AddCommand(CmdShowAccountId())
	cmd.AddCommand(CmdListDid())
	cmd.AddCommand(CmdShowDid())
	cmd.AddCommand(CmdListKid())
	cmd.AddCommand(CmdShowKid())
	cmd.AddCommand(CmdListDidBalances())
	cmd.AddCommand(CmdShowDidBalances())
	// this line is used by starport scaffolding # 1

	return cmd
}
