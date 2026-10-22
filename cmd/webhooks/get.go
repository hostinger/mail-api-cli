package webhooks

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:               "get <mailbox-resource-id> <webhook>",
	Short:             "Get webhook",
	Long:              "Retrieve a single webhook by id.",
	Example:           "  hostinger-mail webhooks get my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GetWebhookWithResponse(ctx, args[0], uuid.MustParse(args[1]))
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
	utils.RegisterArgs(GetCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("webhook"))
}
