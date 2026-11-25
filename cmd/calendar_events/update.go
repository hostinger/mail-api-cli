package calendar_events

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hostinger/mail-api-cli/api"
	"github.com/hostinger/mail-api-cli/client"
	"github.com/hostinger/mail-api-cli/output"
	"github.com/hostinger/mail-api-cli/utils"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:               "update <mailbox-resource-id> <calendar-uuid> <event-uuid>",
	Short:             "Update calendar event",
	Long:              "Change the details of an event: title, times, description, location, attendees, recurrence and so on. Every field replaces the stored value and an omitted optional field is cleared, so call `getCalendarEvent` first and send back the values that should stay. To only move an event in time, use `rescheduleCalendarEvent`. Only an event with `editable` true can be changed; any other is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`.\nFor a recurring event, `scope` `this` changes only the occurrence at `recurrenceId`, `future` ends the series before that occurrence and starts a new one with a new `uuid`, `all` (the default) changes the whole series, keeping the date of its first occurrence.\nAttendees receive an email update, added attendees an invitation and removed ones a cancellation, unless `notifyAttendees` is false. Returns the changed event, or the changed occurrence for `scope` `this`.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --attendees\n    [\n      {\n        \"email\": \"john@example.com\",\n        \"name\": \"John\"\n      }\n    ]\n\n  --recurrencerule\n    {\n      \"byDay\": [\n        {\n          \"day\": \"friday\",\n          \"position\": -1\n        }\n      ],\n      \"byMonth\": [\n        3\n      ],\n      \"byMonthDay\": [\n        15\n      ],\n      \"count\": 10,\n      \"frequency\": \"WEEKLY\",\n      \"interval\": 1,\n      \"until\": \"string\",\n      \"weekStart\": \"monday\"\n    }\n\n  --reminder\n    {\n      \"unit\": \"minutes\",\n      \"value\": 15\n    }",
	Example:           "  hostinger-mail calendar-events update my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --allday=false --start 2026-09-28T13:00:00+03:00 --title Planning",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "scope", []string{"this", "future", "all"}); err != nil {
			return err
		}
		params := updateParams(cmd)
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().UpdateCalendarEventWithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]), params, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"), utils.UUIDArg("event-uuid"))
	UpdateCmd.Flags().StringP("if-match", "", "", "Optional `etag` from `getCalendarEvent`. When the event changed since that call, nothing is changed and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.")
	UpdateCmd.Flags().BoolP("allday", "", false, "`true` for an all-day event (dates only, `timeZone` ignored), `false` for an event at a time of day.")
	UpdateCmd.Flags().StringP("attendees", "", "", "Full list of invited people. Added addresses receive an invitation and removed ones a cancellation. Unique addresses, at most 100 not counting the mailbox itself, across all occurrences. The own address of the mailbox is stored as accepted and not invited. Omitted, null or empty for none. (JSON, @file or -; shape in the help above)")
	UpdateCmd.Flags().StringP("color", "", "", "Color of the event as #RRGGBB or #RRGGBBAA. Omitted or null uses the calendar color.")
	UpdateCmd.Flags().StringP("description", "", "", "Notes of the event. Omitted or null clears it.")
	UpdateCmd.Flags().StringP("end", "", "", "Exclusive end, same format as `start`. Required for a timed event and not before `start`. For an all-day event the day after the last day (2026-09-29 for one day on 2026-09-28); omitted means one day.")
	UpdateCmd.Flags().StringP("location", "", "", "Place of the event. Omitted or null clears it.")
	UpdateCmd.Flags().StringP("meetinglink", "", "", "http or https URL of the online meeting. Omitted or null clears it.")
	UpdateCmd.Flags().BoolP("notifyattendees", "", false, "`false` saves the change without emailing attendees, and copies of the event in the calendars of other Hostinger mailboxes are not updated. Omitted or null means `true`.")
	UpdateCmd.Flags().StringP("recurrenceid", "", "", "`recurrenceId` of one occurrence, copied unchanged from a `listCalendarEvents` item. Required when `scope` is `this` or `future`, ignored for `all`.")
	UpdateCmd.Flags().StringP("recurrencerule", "", "", "Omitted or null makes it a one-off event. A repeating event sends every field of the rule, null when unused: `frequency` DAILY, WEEKLY, MONTHLY or YEARLY, `interval` 1 to 1000, and `count` or `until` to end it, not both. (JSON, @file or -; shape in the help above)")
	UpdateCmd.Flags().StringP("reminder", "", "", "Alarm before the start: `value` 0 to 1000 and `unit`. Omitted or null clears it. (JSON, @file or -; shape in the help above)")
	UpdateCmd.Flags().StringP("scope", "", "", "For a recurring event: `this` changes the occurrence at `recurrenceId`, `future` changes it and every later one, `all` changes the whole series. Omitted or null means `all`. Required when `recurrenceId` is sent. (one of: this, future, all)")
	UpdateCmd.Flags().StringP("start", "", "", "ISO 8601 date-time with offset; without an offset it is read as UTC. For an all-day event a date is enough (2026-09-28).")
	UpdateCmd.Flags().StringP("timezone", "", "", "IANA time zone of the event. Required when `allDay` is false, ignored when `allDay` is true.")
	UpdateCmd.Flags().StringP("title", "", "", "Title of the event.")
	UpdateCmd.MarkFlagRequired("allday")
	UpdateCmd.MarkFlagRequired("start")
	UpdateCmd.MarkFlagRequired("title")
	UpdateCmd.RegisterFlagCompletionFunc("if-match", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("allday", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("attendees", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateCmd.RegisterFlagCompletionFunc("color", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("end", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("location", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("meetinglink", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("notifyattendees", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("recurrenceid", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("recurrencerule", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateCmd.RegisterFlagCompletionFunc("reminder", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateCmd.RegisterFlagCompletionFunc("scope", cobra.FixedCompletions([]cobra.Completion{"this", "future", "all"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("start", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("timezone", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("title", cobra.NoFileCompletions)
}

func updateParams(cmd *cobra.Command) *client.UpdateCalendarEventParams {
	params := &client.UpdateCalendarEventParams{}
	if cmd.Flags().Changed("if-match") {
		v, _ := cmd.Flags().GetString("if-match")
		params.IfMatch = &v
	}
	return params
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
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
	if cmd.Flags().Changed("notifyattendees") {
		v, _ := cmd.Flags().GetBool("notifyattendees")
		body["notifyAttendees"] = v
	}
	if cmd.Flags().Changed("recurrenceid") {
		v, _ := cmd.Flags().GetString("recurrenceid")
		body["recurrenceId"] = v
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
	if cmd.Flags().Changed("scope") {
		v, _ := cmd.Flags().GetString("scope")
		body["scope"] = v
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
