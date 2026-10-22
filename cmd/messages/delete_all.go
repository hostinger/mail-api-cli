package messages

import (
	"context"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteAllCmd = &cobra.Command{
	Use:               "delete-all <mailbox-resource-id> <folder>",
	Short:             "Delete all messages",
	Long:              "Permanently delete every message in a folder (empty the folder).",
	Example:           "  hostinger-mail messages delete-all my-mailbox-resource-id my-folder",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete all messages", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DeleteAllMessagesWithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(DeleteAllCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"))
	DeleteAllCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
