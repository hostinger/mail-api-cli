package messages

import (
	"context"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var TextCmd = &cobra.Command{
	Use:               "text <mailbox-resource-id> <folder> <uid>",
	Short:             "Get message text content",
	Long:              "Retrieve rendered text (plain + HTML) of a message. Marks message as `\\Seen`.",
	Example:           "  hostinger-mail messages text my-mailbox-resource-id my-folder 1234567",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GetMessageTextWithResponse(ctx, args[0], args[1], utils.StringToInt(args[2]))
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
	utils.RegisterArgs(TextCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"), utils.IntArg("uid"))
}
