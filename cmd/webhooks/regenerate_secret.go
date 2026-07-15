package webhooks

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/spf13/cobra"
)

var RegenerateSecretCmd = &cobra.Command{
	Use:   "regenerate-secret <mailbox-resource-id> <webhook>",
	Short: "Regenerate webhook secret",
	Long:  "Regenerate the webhook secret. The previous secret is immediately invalidated. The new secret is returned once.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().RegenerateWebhookSecretWithResponse(context.TODO(), args[0], uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
