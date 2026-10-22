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

var MoveCmd = &cobra.Command{
	Use:               "move <mailbox-resource-id> <folder> <uid>",
	Short:             "Move message",
	Long:              "Move a single message to a target folder.",
	Example:           "  hostinger-mail messages move my-mailbox-resource-id my-folder 1234567 --targetfolder INBOX.Sent",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := moveBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MoveMessageWithBodyWithResponse(ctx, args[0], args[1], utils.StringToInt(args[2]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(MoveCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"), utils.IntArg("uid"))
	MoveCmd.Flags().StringP("targetfolder", "", "", "Destination folder path.")
	MoveCmd.MarkFlagRequired("targetfolder")
	MoveCmd.RegisterFlagCompletionFunc("targetfolder", cobra.NoFileCompletions)
}

func moveBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	targetFolderVal, _ := cmd.Flags().GetString("targetfolder")
	body["targetFolder"] = targetFolderVal
	return body, nil
}
