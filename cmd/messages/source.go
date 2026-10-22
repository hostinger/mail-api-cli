package messages

import (
	"context"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var SourceCmd = &cobra.Command{
	Use:               "source <mailbox-resource-id> <folder> <uid>",
	Short:             "Get message source",
	Long:              "Retrieve raw RFC822 source of a message as `message/rfc822` attachment.",
	Example:           "  hostinger-mail messages source my-mailbox-resource-id my-folder 1234567",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GetMessageSourceWithResponse(ctx, args[0], args[1], utils.StringToInt(args[2]))
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
	utils.RegisterArgs(SourceCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"), utils.IntArg("uid"))
}
