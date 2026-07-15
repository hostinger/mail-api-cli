package messages

import (
	"context"
	"log"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var AttachmentCmd = &cobra.Command{
	Use:   "attachment <mailbox-resource-id> <folder> <uid> <attachment-id>",
	Short: "Download message attachment",
	Long:  "Download a message attachment as `application/octet-stream`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(4)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().GetMessageAttachmentWithResponse(context.TODO(), args[0], args[1], utils.StringToInt(args[2]), args[3])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}
