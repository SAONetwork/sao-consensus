package cli

import (
	"fmt"
	"strings"

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
	cmd.AddCommand(CmdListPledgeDebt())
	cmd.AddCommand(CmdShowPledgeDebt())
	cmd.AddCommand(CmdFault())

	cmd.AddCommand(CmdAllFaults())

	cmd.AddCommand(CmdFishmen())

	// this line is used by starport scaffolding # 1

	return cmd
}

func NodeStatusToText(status uint32) string {

	if status == types.NODE_STATUS_NA {
		return types.TEXT_NODE_STATUS_NA
	}

	nodeStatusMap := make(map[uint32]string, 0)

	nodeStatusMap[types.NODE_STATUS_ONLINE] = types.TEXT_NODE_STATUS_ONLINE
	nodeStatusMap[types.NODE_STATUS_SERVE_GATEWAY] = types.TEXT_NODE_STATUS_SERVE_GATEWAY
	nodeStatusMap[types.NODE_STATUS_SERVE_STORAGE] = types.TEXT_NODE_STATUS_SERVE_STORAGE
	nodeStatusMap[types.NODE_STATUS_ACCEPT_ORDER] = types.TEXT_NODE_STATUS_ACCEPT_ORDER
	nodeStatusMap[types.NODE_STATUS_SERVE_INDEXING] = types.TEXT_NODE_STATUS_SERVE_INDEXING
	nodeStatusMap[types.NODE_STATUS_SERVE_FISHING] = types.TEXT_NODE_STATUS_SERVE_FISHING
	nodeStatusMap[types.NODE_STATUS_SERVE_JAILED] = types.TEXT_NODE_STATUS_SERVE_JAILED

	statusText := make([]string, 0)

	for nodeStatus, textNodeStatus := range nodeStatusMap {
		if nodeStatus&uint32(status) == nodeStatus {
			statusText = append(statusText, textNodeStatus)
		}
	}
	return strings.Join(statusText, " | ")
}

func NodeRoleToText(role uint32) string {
	if role == types.NODE_NORMAL {
		return types.TEXT_NODE_NORMAL
	}
	return types.TEXT_NODE_SUPER
}
