package webhooks

import (
	"context"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <mailbox-resource-id>",
	Short: "List webhooks",
	Long:  "List webhooks for a mailbox.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"active", "paused", "disabled"})
		r, err := api.Request().ListWebhooksWithResponse(context.TODO(), args[0], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().StringP("status", "", "", "Return only webhooks with this status. (one of: active, paused, disabled)")
	ListCmd.Flags().IntP("page", "", 1, "Page number (1-based).")
	ListCmd.Flags().IntP("perpage", "", 15, "Items per page (max 1000).")
}

func listParams(cmd *cobra.Command) *client.ListWebhooksParams {
	params := &client.ListWebhooksParams{}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		e := client.ListWebhooksParamsStatus(v)
		params.Status = &e
	}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("perpage") {
		v, _ := cmd.Flags().GetInt("perpage")
		params.PerPage = &v
	}
	return params
}
