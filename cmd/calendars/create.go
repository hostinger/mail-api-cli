package calendars

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create <mailbox-resource-id>",
	Short:             "Create calendar",
	Long:              "Create a calendar owned by the mailbox. The number of calendars a mailbox can own is limited; when the limit is reached the request fails with `400` and code `ERR_CALENDAR_LIMIT_REACHED`.",
	Example:           "  hostinger-mail calendars create my-mailbox-resource-id --displayname Work",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().CreateCalendarWithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("mailbox-resource-id"))
	CreateCmd.Flags().StringP("color", "", "", "Hex color as #RRGGBB or #RRGGBBAA. Omitted or null uses the default color.")
	CreateCmd.Flags().StringP("description", "", "", "Optional free-text note about the calendar.")
	CreateCmd.Flags().StringP("displayname", "", "", "Name of the calendar.")
	CreateCmd.Flags().BoolP("visible", "", false, "Whether the mailbox shows the calendar in its calendar view. Omitted or null means true.")
	CreateCmd.MarkFlagRequired("displayname")
	CreateCmd.RegisterFlagCompletionFunc("color", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("displayname", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("visible", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("color") {
		v, _ := cmd.Flags().GetString("color")
		body["color"] = v
	}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	displayNameVal, _ := cmd.Flags().GetString("displayname")
	body["displayName"] = displayNameVal
	if cmd.Flags().Changed("visible") {
		v, _ := cmd.Flags().GetBool("visible")
		body["visible"] = v
	}
	return body, nil
}
