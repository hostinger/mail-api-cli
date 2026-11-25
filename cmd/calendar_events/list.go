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

var ListCmd = &cobra.Command{
	Use:               "list <mailbox-resource-id> <calendar-uuid>",
	Short:             "List calendar events",
	Long:              "List the events of one calendar that overlap a time range. It covers one calendar only: get the calendars from `listCalendars` and call it for each. The range runs from `start` (inclusive) to `end` (exclusive), at most 32 days; for a longer period call it again with the next range. Recurring events are expanded: each occurrence is a separate item with the same `uuid` and its own `recurrenceId`. Results are not paginated.\nItems hold summary fields only. For description, location, attendees and organizer call `getCalendarEvent` with the `calendarUuid`, `uuid` (as `eventUuid`) and `recurrenceId` of the item. Times are in UTC; `timeZone` is the zone the event was created in. `participationStatus` is the answer of the mailbox to an invitation from someone else, null for its own events. `status` CANCELLED means the organizer cancelled the event.",
	Example:           "  hostinger-mail calendar-events list my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --start 2025-01-31T10:00:00Z --end 2025-01-31T10:00:00Z",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.FormatCheck(cmd, "start", "date-time"); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "end", "date-time"); err != nil {
			return err
		}
		params := listParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ListCalendarEventsWithResponse(ctx, args[0], uuid.MustParse(args[1]), params)
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
	utils.RegisterArgs(ListCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"))
	ListCmd.Flags().StringP("start", "", "", "Start of the range, inclusive. ISO 8601 date-time with offset.")
	ListCmd.Flags().StringP("end", "", "", "End of the range, exclusive. ISO 8601 date-time after `start` and at most 32 days after it.")
	ListCmd.Flags().StringP("timezone", "", "", "IANA time zone used to place all-day events and events without a time zone on days, also across a daylight saving change. Defaults to the fixed offset of `start`.")
	ListCmd.MarkFlagRequired("start")
	ListCmd.MarkFlagRequired("end")
	ListCmd.RegisterFlagCompletionFunc("start", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("end", cobra.NoFileCompletions)
	ListCmd.RegisterFlagCompletionFunc("timezone", cobra.NoFileCompletions)
}

func listParams(cmd *cobra.Command) *client.ListCalendarEventsParams {
	params := &client.ListCalendarEventsParams{}
	startVal, _ := cmd.Flags().GetString("start")
	params.Start = utils.StringToTime(startVal)
	endVal, _ := cmd.Flags().GetString("end")
	params.End = utils.StringToTime(endVal)
	if cmd.Flags().Changed("timezone") {
		v, _ := cmd.Flags().GetString("timezone")
		params.TimeZone = &v
	}
	return params
}
