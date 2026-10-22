package messages

import (
	"context"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list <mailbox-resource-id> <folder>",
	Short:             "List messages",
	Long:              "List messages in a folder. Use POST /search for filtering. Sort fields: uid, date, size (prefix with `-` for descending). Default `-uid`.",
	Example:           "  hostinger-mail messages list my-mailbox-resource-id my-folder",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ListMessagesWithResponse(ctx, args[0], args[1], params)
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
	utils.RegisterArgs(ListCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"))
	ListCmd.Flags().IntP("page", "", 1, "Page number (1-based).")
	ListCmd.Flags().IntP("perpage", "", 25, "Items per page (max 100).")
	ListCmd.Flags().StringP("sort", "", "-uid", "Sort field with optional `-` prefix for descending. Allowed: uid, date, size.")
	ListCmd.RegisterFlagCompletionFunc("page", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("perpage", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("sort", cobra.NoFileCompletions)
	ListCmd.Flags().Bool("all", false, "Fetch every page and print all items")
	ListCmd.MarkFlagsMutuallyExclusive("all", "page")
}

func listParams(cmd *cobra.Command) *client.ListMessagesParams {
	params := &client.ListMessagesParams{}
	if cmd.Flags().Changed("page") {
		v, _ := cmd.Flags().GetInt("page")
		params.Page = &v
	}
	if cmd.Flags().Changed("perpage") {
		v, _ := cmd.Flags().GetInt("perpage")
		params.PerPage = &v
	}
	if cmd.Flags().Changed("sort") {
		v, _ := cmd.Flags().GetString("sort")
		params.Sort = &v
	}
	return params
}
