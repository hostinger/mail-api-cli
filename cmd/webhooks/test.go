package webhooks

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use:   "test <mailbox-resource-id> <webhook>",
	Short: "Test webhook",
	Long:  "Send a test delivery to the webhook URL and return the upstream response.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().TestWebhookWithResponse(context.TODO(), args[0], uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
