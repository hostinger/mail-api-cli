package calendar_events

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

var CreateCmd = &cobra.Command{
	Use:               "create <mailbox-resource-id> <calendar-uuid>",
	Short:             "Create calendar event",
	Long:              "Create an event in one calendar of the mailbox. `calendarUuid` comes from `listCalendars`: use the calendar with `default` true unless another one is named. A calendar the mailbox can only read is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`.\nThe mailbox becomes the organizer and every address in `attendees` immediately receives an email invitation, which cannot be recalled. Without attendees no email is sent.\nA timed event needs `allDay` false, `start`, `end` and `timeZone`. An all-day event needs `allDay` true and a `start` date; `end` is optional. Set `recurrenceRule` for a repeating event and `reminder` for an alarm.\nReturns the created event. Its `uuid` is the `eventUuid` for `getCalendarEvent`, `updateCalendarEvent`, `rescheduleCalendarEvent` and `deleteCalendarEvent`.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --attendees\n    [\n      {\n        \"email\": \"john@example.com\",\n        \"name\": \"John\"\n      }\n    ]\n\n  --recurrencerule\n    {\n      \"byDay\": [\n        {\n          \"day\": \"friday\",\n          \"position\": -1\n        }\n      ],\n      \"byMonth\": [\n        3\n      ],\n      \"byMonthDay\": [\n        15\n      ],\n      \"count\": 10,\n      \"frequency\": \"WEEKLY\",\n      \"interval\": 1,\n      \"until\": \"string\",\n      \"weekStart\": \"monday\"\n    }\n\n  --reminder\n    {\n      \"unit\": \"minutes\",\n      \"value\": 15\n    }",
	Example:           "  hostinger-mail calendar-events create my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --allday=false --start 2026-09-28T13:00:00+03:00 --title Planning",
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
			r, err := api.Request().CreateCalendarEventWithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"))
	CreateCmd.Flags().BoolP("allday", "", false, "`true` for an all-day event (dates only, `timeZone` ignored), `false` for an event at a time of day.")
	CreateCmd.Flags().StringP("attendees", "", "", "People to invite; each receives an email invitation. Unique addresses, at most 100 not counting the mailbox itself, across all occurrences. The own address of the mailbox is stored as accepted and not invited. Omitted, null or empty for none. (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("color", "", "", "Color of the event as #RRGGBB or #RRGGBBAA. Omitted or null uses the calendar color.")
	CreateCmd.Flags().StringP("description", "", "", "Notes of the event. Omitted or null for none.")
	CreateCmd.Flags().StringP("end", "", "", "Exclusive end, same format as `start`. Required for a timed event and not before `start`. For an all-day event the day after the last day (2026-09-29 for one day on 2026-09-28); omitted means one day.")
	CreateCmd.Flags().StringP("location", "", "", "Place of the event. Omitted or null for none.")
	CreateCmd.Flags().StringP("meetinglink", "", "", "http or https URL of the online meeting. Omitted or null for none.")
	CreateCmd.Flags().StringP("recurrencerule", "", "", "Omitted or null for a one-off event. A repeating event sends every field of the rule, null when unused: `frequency` DAILY, WEEKLY, MONTHLY or YEARLY, `interval` 1 to 1000, and `count` or `until` to end it, not both. (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("reminder", "", "", "Alarm before the start: `value` 0 to 1000 and `unit`. Omitted or null for none. (JSON, @file or -; shape in the help above)")
	CreateCmd.Flags().StringP("start", "", "", "ISO 8601 date-time with offset; without an offset it is read as UTC. For an all-day event a date is enough (2026-09-28).")
	CreateCmd.Flags().StringP("timezone", "", "", "IANA time zone of the event. Required when `allDay` is false, ignored when `allDay` is true.")
	CreateCmd.Flags().StringP("title", "", "", "Title of the event.")
	CreateCmd.MarkFlagRequired("allday")
	CreateCmd.MarkFlagRequired("start")
	CreateCmd.MarkFlagRequired("title")
	CreateCmd.RegisterFlagCompletionFunc("allday", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("attendees", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("color", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("end", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("location", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("meetinglink", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("recurrencerule", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("reminder", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateCmd.RegisterFlagCompletionFunc("start", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("timezone", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("title", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	allDayVal, _ := cmd.Flags().GetBool("allday")
	body["allDay"] = allDayVal
	if cmd.Flags().Changed("attendees") {
		attendeesVal, _ := cmd.Flags().GetString("attendees")
		attendeesValJSON, err := utils.JSONValue(attendeesVal, "attendees")
		if err != nil {
			return nil, err
		}
		body["attendees"] = attendeesValJSON
	}
	if cmd.Flags().Changed("color") {
		v, _ := cmd.Flags().GetString("color")
		body["color"] = v
	}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	if cmd.Flags().Changed("end") {
		v, _ := cmd.Flags().GetString("end")
		body["end"] = v
	}
	if cmd.Flags().Changed("location") {
		v, _ := cmd.Flags().GetString("location")
		body["location"] = v
	}
	if cmd.Flags().Changed("meetinglink") {
		v, _ := cmd.Flags().GetString("meetinglink")
		body["meetingLink"] = v
	}
	if cmd.Flags().Changed("recurrencerule") {
		recurrenceRuleVal, _ := cmd.Flags().GetString("recurrencerule")
		recurrenceRuleValJSON, err := utils.JSONValue(recurrenceRuleVal, "recurrencerule")
		if err != nil {
			return nil, err
		}
		body["recurrenceRule"] = recurrenceRuleValJSON
	}
	if cmd.Flags().Changed("reminder") {
		reminderVal, _ := cmd.Flags().GetString("reminder")
		reminderValJSON, err := utils.JSONValue(reminderVal, "reminder")
		if err != nil {
			return nil, err
		}
		body["reminder"] = reminderValJSON
	}
	startVal, _ := cmd.Flags().GetString("start")
	body["start"] = startVal
	if cmd.Flags().Changed("timezone") {
		v, _ := cmd.Flags().GetString("timezone")
		body["timeZone"] = v
	}
	titleVal, _ := cmd.Flags().GetString("title")
	body["title"] = titleVal
	return body, nil
}
