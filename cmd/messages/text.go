package messages

import (
	"context"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var TextCmd = &cobra.Command{
	Use:   "text <mailbox-resource-id> <folder> <uid>",
	Short: "Get message text content",
	Long:  "Retrieve rendered text (plain + HTML) of a message. Marks message as `\\Seen`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().GetMessageTextWithResponse(context.TODO(), args[0], args[1], utils.StringToInt(args[2]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
