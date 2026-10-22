package webhooks

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:               "delete <mailbox-resource-id> <webhook>",
	Short:             "Delete webhook",
	Long:              "Delete a webhook.",
	Example:           "  hostinger-mail webhooks delete my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Delete webhook", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DeleteWebhookWithResponse(ctx, args[0], uuid.MustParse(args[1]))
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("webhook"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}
