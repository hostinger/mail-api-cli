package account

import (
	"context"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/spf13/cobra"
)

var CurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the authenticated account",
	Long:  "Returns the authenticated account and the mailboxes it can manage.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().GetCurrentAccountWithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
