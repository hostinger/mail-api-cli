package messages

import (
	"context"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <mailbox-resource-id> <folder>",
	Short: "List messages",
	Long:  "List messages in a folder. Use POST /search for filtering. Sort fields: uid, date, size (prefix with `-` for descending). Default `-uid`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ListMessagesWithResponse(context.TODO(), args[0], args[1], listParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListCmd.Flags().IntP("page", "", 1, "Page number (1-based).")
	ListCmd.Flags().IntP("perpage", "", 25, "Items per page (max 100).")
	ListCmd.Flags().StringP("sort", "", "-uid", "Sort field with optional `-` prefix for descending. Allowed: uid, date, size.")
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
