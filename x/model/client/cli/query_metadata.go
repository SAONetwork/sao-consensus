package cli

import (
	"context"

	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-metadata",
		Short: "list all metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllMetadataRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.MetadataAll(context.Background(), params)
			if err != nil {
				return err
			}

			metadatas := res.GetMetadata()

			textMetadatas := make([]types.MetadataInText, 0)

			for _, metadata := range metadatas {

				textMetadata := types.MetadataInText{
					DataId:        metadata.DataId,
					Owner:         metadata.Owner,
					Alias:         metadata.Alias,
					GroupId:       metadata.GroupId,
					OrderId:       metadata.OrderId,
					Tags:          metadata.Tags,
					Cid:           metadata.Cid,
					Commits:       metadata.Commits,
					ExtendInfo:    metadata.ExtendInfo,
					Update:        metadata.Update,
					Commit:        metadata.Commit,
					Rule:          metadata.Rule,
					Duration:      metadata.Duration,
					CreatedAt:     metadata.CreatedAt,
					ReadonlyDids:  metadata.ReadonlyDids,
					ReadwriteDids: metadata.ReadwriteDids,
					Orders:        metadata.Orders,
					Status:        MetadataStatusInText(metadata.Status),
				}

				textMetadatas = append(textMetadatas, textMetadata)
			}

			textRes := types.QueryAllMetadataInTextResponse{
				Metadata:   textMetadatas,
				Pagination: res.Pagination,
			}

			return clientCtx.PrintProto(&textRes)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowMetadata() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-metadata [data-id]",
		Short: "shows a metadata",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDataId := args[0]

			params := &types.QueryGetMetadataRequest{
				DataId: argDataId,
			}

			res, err := queryClient.Metadata(context.Background(), params)
			if err != nil {
				return err
			}

			metadata := res.GetMetadata()

			textMetadata := types.MetadataInText{
				DataId:        metadata.DataId,
				Owner:         metadata.Owner,
				Alias:         metadata.Alias,
				GroupId:       metadata.GroupId,
				OrderId:       metadata.OrderId,
				Tags:          metadata.Tags,
				Cid:           metadata.Cid,
				Commits:       metadata.Commits,
				ExtendInfo:    metadata.ExtendInfo,
				Update:        metadata.Update,
				Commit:        metadata.Commit,
				Rule:          metadata.Rule,
				Duration:      metadata.Duration,
				CreatedAt:     metadata.CreatedAt,
				ReadonlyDids:  metadata.ReadonlyDids,
				ReadwriteDids: metadata.ReadwriteDids,
				Orders:        metadata.Orders,
				Status:        MetadataStatusInText(metadata.Status),
			}

			textRes := types.QueryGetMetadataInTextResponse{
				Metadata: textMetadata,
			}

			return clientCtx.PrintProto(&textRes)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
