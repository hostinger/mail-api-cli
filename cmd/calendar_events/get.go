package calendar_events

import (
	"context"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:               "get <mailbox-resource-id> <calendar-uuid> <event-uuid>",
	Short:             "Get calendar event",
	Long:              "Retrieve the full details of one event: title, times, description, location, meeting link, recurrence rule, reminder, organizer, attendees with their answers, and its `etag`. `calendarUuid` and `eventUuid` come from `listCalendarEvents` (the `calendarUuid` and `uuid` of an item). For one occurrence of a recurring event pass the `recurrenceId` of the item; without it the whole series is returned and `recurrenceId` is null. A `recurrenceId` that matches no occurrence returns `404`.\nCall it before `updateCalendarEvent`, which replaces every field, so the values that should stay can be sent back. `replyable` true means the mailbox was invited by someone else and can answer with `replyToCalendarEvent`. `editable` true means the event can be changed or moved.",
	Example:           "  hostinger-mail calendar-events get my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.FormatCheck(cmd, "recurrenceid", "date-time"); err != nil {
			return err
		}
		params := getParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().GetCalendarEventWithResponse(ctx, args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]), params)
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
	utils.RegisterArgs(GetCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"), utils.UUIDArg("event-uuid"))
	GetCmd.Flags().StringP("recurrenceid", "", "", "`recurrenceId` of one occurrence, copied unchanged from a `listCalendarEvents` item. Omit it for a one-off event or for the whole series.")
	GetCmd.RegisterFlagCompletionFunc("recurrenceid", cobra.NoFileCompletions)
}

func getParams(cmd *cobra.Command) *client.GetCalendarEventParams {
	params := &client.GetCalendarEventParams{}
	if cmd.Flags().Changed("recurrenceid") {
		v, _ := cmd.Flags().GetString("recurrenceid")
		t := utils.StringToTime(v)
		params.RecurrenceId = &t
	}
	return params
}
