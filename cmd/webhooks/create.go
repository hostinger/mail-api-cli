package webhooks

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create <mailbox-resource-id>",
	Short:             "Create webhook",
	Long:              "Create a webhook. The response includes the one-time `secret` — store it securely as it is never returned again.",
	Example:           "  hostinger-mail webhooks create my-mailbox-resource-id --events message.received --name 'New message notifier' --url https://example.com/webhooks/incoming",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"active", "paused", "disabled"}); err != nil {
			return err
		}
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().CreateWebhookWithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("mailbox-resource-id"))
	CreateCmd.Flags().StringP("description", "", "", "Optional free-text note about the webhook purpose.")
	CreateCmd.Flags().StringSliceP("events", "", nil, "Event types that trigger a delivery. (one of: message.received)")
	CreateCmd.Flags().StringP("name", "", "", "Human-readable webhook name.")
	CreateCmd.Flags().StringP("status", "", "active", "Initial delivery state. Only active webhooks receive events. (one of: active, paused, disabled)")
	CreateCmd.Flags().StringP("url", "", "", "HTTPS endpoint that receives POST deliveries, authenticated with the webhook secret as `Authorization: Bearer <secret>`. Must be a public domain name (no IPs or internal hosts).")
	CreateCmd.MarkFlagRequired("events")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("url")
	CreateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("events", cobra.FixedCompletions([]cobra.Completion{"message.received"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "paused", "disabled"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
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
	return body, nil
}
