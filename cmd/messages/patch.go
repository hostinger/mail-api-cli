package messages

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var PatchCmd = &cobra.Command{
	Use:               "patch <mailbox-resource-id> <folder> <uid>",
	Short:             "Update message flags",
	Long:              "Add and/or remove flags on a single message. Returns the updated message.",
	Example:           "  hostinger-mail messages patch my-mailbox-resource-id my-folder 1234567",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := patchBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().PatchMessageWithBodyWithResponse(ctx, args[0], args[1], utils.StringToInt(args[2]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(PatchCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"), utils.IntArg("uid"))
	PatchCmd.Flags().StringSliceP("addflags", "", nil, "IMAP flags to set on the message, e.g. \\Seen, \\Flagged, \\Answered, $forwarded.")
	PatchCmd.Flags().StringSliceP("removeflags", "", nil, "IMAP flags to clear from the message.")
	PatchCmd.RegisterFlagCompletionFunc("addflags", cobra.NoFileCompletions)
	PatchCmd.RegisterFlagCompletionFunc("removeflags", cobra.NoFileCompletions)
}

func patchBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("addflags") {
		v, _ := cmd.Flags().GetStringSlice("addflags")
		body["addFlags"] = v
	}
	if cmd.Flags().Changed("removeflags") {
		v, _ := cmd.Flags().GetStringSlice("removeflags")
		body["removeFlags"] = v
	}
	return body, nil
}
