package cli

import (
	"strconv"

	"github.com/SaoNetwork/sao/x/order/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"strings"
)

var _ = strconv.Itoa(0)

func CmdOrdersStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "orders-status [order-ids]",
		Short: "Query ordersStatus",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCastOrderIds := strings.Split(args[0], listSeparator)
			reqOrderIds := make([]uint64, len(reqCastOrderIds))
			for i, arg := range reqCastOrderIds {
				value, err := cast.ToUint64E(arg)
				if err != nil {
					return err
				}
				reqOrderIds[i] = value
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryOrdersStatusRequest{

				OrderIds: reqOrderIds,
			}

			res, err := queryClient.OrdersStatus(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
