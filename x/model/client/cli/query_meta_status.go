package cli

import (
	"strconv"

	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdMetaStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "meta-status [data-ids]",
		Short: "Query metaStatus",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDataIds := strings.Split(args[0], listSeparator)

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryMetaStatusRequest{

				DataIds: reqDataIds,
			}

			res, err := queryClient.MetaStatus(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
