package webhooks

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:               "update <mailbox-resource-id> <webhook>",
	Short:             "Update webhook",
	Long:              "Partially update a webhook.",
	Example:           "  hostinger-mail webhooks update my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"active", "paused", "disabled"}); err != nil {
			return err
		}
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().UpdateWebhookWithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("webhook"))
	UpdateCmd.Flags().StringP("description", "", "", "Free-text note about the webhook purpose. Send null to clear.")
	UpdateCmd.Flags().StringSliceP("events", "", nil, "Event types that trigger a delivery. Replaces the current list. (one of: message.received)")
	UpdateCmd.Flags().StringP("name", "", "", "Human-readable webhook name.")
	UpdateCmd.Flags().StringP("status", "", "", "Delivery state. Only active webhooks receive events; paused keeps config but stops deliveries. (one of: active, paused, disabled)")
	UpdateCmd.Flags().StringP("url", "", "", "HTTPS endpoint that receives POST deliveries, authenticated with the webhook secret as `Authorization: Bearer <secret>`. Must be a public domain name (no IPs or internal hosts).")
	UpdateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("events", cobra.FixedCompletions([]cobra.Completion{"message.received"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "paused", "disabled"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	if cmd.Flags().Changed("events") {
		v, _ := cmd.Flags().GetStringSlice("events")
		body["events"] = v
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	if cmd.Flags().Changed("url") {
		v, _ := cmd.Flags().GetString("url")
		body["url"] = v
	}
	return body, nil
}
