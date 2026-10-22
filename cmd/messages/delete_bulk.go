package messages

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteBulkCmd = &cobra.Command{
	Use:               "delete-bulk <mailbox-resource-id> <folder>",
	Short:             "Delete messages",
	Long:              "Permanently delete multiple messages from a folder.",
	Example:           "  hostinger-mail messages delete-bulk my-mailbox-resource-id my-folder --uids 42,43,44",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := deleteBulkBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DeleteMessagesWithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(DeleteBulkCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"))
	DeleteBulkCmd.Flags().IntSliceP("uids", "", nil, "Message UIDs to delete. 1-100 entries, each > 0.")
	DeleteBulkCmd.MarkFlagRequired("uids")
	DeleteBulkCmd.RegisterFlagCompletionFunc("uids", cobra.NoFileCompletions)
}

func deleteBulkBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	uidsVal, _ := cmd.Flags().GetIntSlice("uids")
	body["uids"] = uidsVal
	return body, nil
}
