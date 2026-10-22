package webhooks

import (
	"context"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list <mailbox-resource-id>",
	Short:             "List webhooks",
	Long:              "List webhooks for a mailbox.",
	Example:           "  hostinger-mail webhooks list my-mailbox-resource-id",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"active", "paused", "disabled"}); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ListWebhooksWithResponse(ctx, args[0], params)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		if all, _ := cmd.Flags().GetBool("all"); all {
			call = utils.AllPages(call, func(page int) { params.Page = &page })
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(ListCmd, utils.StringArg("mailbox-resource-id"))
	ListCmd.Flags().StringP("status", "", "", "Return only webhooks with this status. (one of: active, paused, disabled)")
	ListCmd.Flags().IntP("page", "", 1, "Page number (1-based).")
	ListCmd.Flags().IntP("perpage", "", 15, "Items per page (max 1000).")
	ListCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "paused", "disabled"}, cobra.ShellCompDirectiveNoFileComp))
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("perpage", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
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
