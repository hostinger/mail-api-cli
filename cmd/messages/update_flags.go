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

var UpdateFlagsCmd = &cobra.Command{
	Use:               "update-flags <mailbox-resource-id> <folder>",
	Short:             "Update flags on multiple messages",
	Long:              "Add and/or remove flags on multiple messages. Returns 200 when all UIDs succeed, 207 with per-UID outcome when some fail.",
	Example:           "  hostinger-mail messages update-flags my-mailbox-resource-id my-folder --uids 42,43,44",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateFlagsBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().UpdateMessageFlagsWithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateFlagsCmd, utils.StringArg("mailbox-resource-id"), utils.StringArg("folder"))
	UpdateFlagsCmd.Flags().StringSliceP("addflags", "", nil, "IMAP flags to set on every listed message, e.g. \\Seen, \\Flagged, \\Answered, $forwarded.")
	UpdateFlagsCmd.Flags().StringSliceP("removeflags", "", nil, "IMAP flags to clear from every listed message.")
	UpdateFlagsCmd.Flags().IntSliceP("uids", "", nil, "Message UIDs to update. 1-100 entries, each > 0.")
	UpdateFlagsCmd.MarkFlagRequired("uids")
	UpdateFlagsCmd.RegisterFlagCompletionFunc("addflags", cobra.NoFileCompletions)
	UpdateFlagsCmd.RegisterFlagCompletionFunc("removeflags", cobra.NoFileCompletions)
	UpdateFlagsCmd.RegisterFlagCompletionFunc("uids", cobra.NoFileCompletions)
}

func updateFlagsBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("addflags") {
		v, _ := cmd.Flags().GetStringSlice("addflags")
		body["addFlags"] = v
	}
	if cmd.Flags().Changed("removeflags") {
		v, _ := cmd.Flags().GetStringSlice("removeflags")
		body["removeFlags"] = v
	}
	uidsVal, _ := cmd.Flags().GetIntSlice("uids")
	body["uids"] = uidsVal
	return body, nil
}
