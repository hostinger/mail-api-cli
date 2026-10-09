package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <mailbox-resource-id>",
	Short: "Create webhook",
	Long:  "Create a webhook. The response includes the one-time `secret` — store it securely as it is never returned again.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"active", "paused", "disabled"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().CreateWebhookWithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("description", "", "", "Optional free-text note about the webhook purpose.")
	CreateCmd.Flags().StringSliceP("events", "", nil, "Event types that trigger a delivery. (one of: message.received)")
	CreateCmd.Flags().StringP("name", "", "", "Human-readable webhook name.")
	CreateCmd.Flags().StringP("status", "", "active", "Initial delivery state. Only active webhooks receive events. (one of: active, paused, disabled)")
	CreateCmd.Flags().StringP("url", "", "", "HTTPS endpoint that receives POST deliveries, authenticated with the webhook secret as `Authorization: Bearer <secret>`. Must be a public domain name (no IPs or internal hosts).")
	CreateCmd.MarkFlagRequired("events")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("url")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	eventsVal, _ := cmd.Flags().GetStringSlice("events")
	body["events"] = eventsVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	urlVal, _ := cmd.Flags().GetString("url")
	body["url"] = urlVal
	return body
}
