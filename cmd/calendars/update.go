package calendars

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:     "update <mailbox-resource-id> <calendar-uuid>",
	Short:   "Update calendar",
	Long:    "Update the name, description, color and visibility of a calendar. Name, description and color replace the stored values: an omitted `description` is cleared and an omitted `color` resets to the default color. An omitted `visible` keeps the current value.",
	Example: "  hostinger-mail calendars update my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --displayname Work",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ListCalendarsWithResponse(ctx, args[0], nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().UpdateCalendarWithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"))
	UpdateCmd.Flags().StringP("color", "", "", "Hex color as #RRGGBB or #RRGGBBAA. Omitted or null resets it to the default color.")
	UpdateCmd.Flags().StringP("description", "", "", "Free-text note about the calendar. Omitted or null clears it.")
	UpdateCmd.Flags().StringP("displayname", "", "", "Name of the calendar.")
	UpdateCmd.Flags().BoolP("visible", "", false, "Whether the mailbox shows the calendar in its calendar view. Omitted or null keeps the current value.")
	UpdateCmd.MarkFlagRequired("displayname")
	UpdateCmd.RegisterFlagCompletionFunc("color", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("displayname", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("visible", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
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
