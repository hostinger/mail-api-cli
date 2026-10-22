package messages

import (
	"context"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var AttachmentCmd = &cobra.Command{
	Use:               "attachment <mailbox-resource-id> <folder> <uid> <attachment-id>",
	Short:             "Download message attachment",
	Long:              "Download a message attachment as `application/octet-stream`.",
	Example:           "  hostinger-mail messages attachment my-mailbox-resource-id my-folder 1234567 my-attachment-id",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GetMessageAttachmentWithResponse(ctx, args[0], args[1], utils.StringToInt(args[2]), args[3])
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
	utils.RegisterArgs(AttachmentCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"), utils.IntArg("uid"), utils.StringArg("attachment-id"))
}
