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

var MoveBulkCmd = &cobra.Command{
	Use:               "move-bulk <mailbox-resource-id> <folder>",
	Short:             "Move messages",
	Long:              "Move multiple messages from a source folder to a target folder.",
	Example:           "  hostinger-mail messages move-bulk my-mailbox-resource-id my-folder --targetfolder INBOX.Sent --uids 42,43,44",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := moveBulkBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MoveMessagesWithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(MoveBulkCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"))
	MoveBulkCmd.Flags().StringP("targetfolder", "", "", "Destination folder path.")
	MoveBulkCmd.Flags().IntSliceP("uids", "", nil, "Message UIDs to move. 1-100 entries, each > 0.")
	MoveBulkCmd.MarkFlagRequired("targetfolder")
	MoveBulkCmd.MarkFlagRequired("uids")
	MoveBulkCmd.RegisterFlagCompletionFunc("targetfolder", cobra.NoFileCompletions)
	MoveBulkCmd.RegisterFlagCompletionFunc("uids", cobra.NoFileCompletions)
}

func moveBulkBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	targetFolderVal, _ := cmd.Flags().GetString("targetfolder")
	body["targetFolder"] = targetFolderVal
	uidsVal, _ := cmd.Flags().GetIntSlice("uids")
	body["uids"] = uidsVal
	return body, nil
}
