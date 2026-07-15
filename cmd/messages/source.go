package messages

import (
	"context"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var SourceCmd = &cobra.Command{
	Use:   "source <mailbox-resource-id> <folder> <uid>",
	Short: "Get message source",
	Long:  "Retrieve raw RFC822 source of a message as `message/rfc822` attachment.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().GetMessageSourceWithResponse(context.TODO(), args[0], args[1], utils.StringToInt(args[2]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
