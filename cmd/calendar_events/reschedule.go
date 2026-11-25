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

var RescheduleCmd = &cobra.Command{
	Use:               "reschedule <mailbox-resource-id> <calendar-uuid> <event-uuid>",
	Short:             "Reschedule calendar event",
	Long:              "Move an event to another time or change its length, keeping everything else (title, attendees, description and so on) as stored, so the event does not have to be fetched first. To change other fields use `updateCalendarEvent`. `calendarUuid` and `eventUuid` come from `listCalendarEvents`. Only an event with `editable` true can be moved; any other is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`.\nFor a recurring event, `scope` `this` moves only the occurrence at `recurrenceId`, `future` ends the series before that occurrence and starts a new one with a new `uuid`, `all` (the default) moves the whole series: only the time of day and the length apply and the series keeps the date of its first occurrence.\nAttendees receive an email update unless `notifyAttendees` is false. Returns the moved event, or the moved occurrence for `scope` `this`.",
	Example:           "  hostinger-mail calendar-events reschedule my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --allday=false --start 2026-09-28T13:00:00+03:00",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "scope", []string{"this", "future", "all"}); err != nil {
			return err
		}
		params := rescheduleParams(cmd)
		body, err := rescheduleBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().RescheduleCalendarEventWithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]), params, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(RescheduleCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"), utils.UUIDArg("event-uuid"))
	RescheduleCmd.Flags().StringP("if-match", "", "", "Optional `etag` from `getCalendarEvent`. When the event changed since that call, nothing is changed and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.")
	RescheduleCmd.Flags().BoolP("allday", "", false, "`true` for an all-day event (dates only, `timeZone` ignored), `false` for an event at a time of day.")
	RescheduleCmd.Flags().StringP("end", "", "", "Exclusive end, same format as `start`. Required for a timed event and not before `start`. For an all-day event the day after the last day (2026-09-29 for one day on 2026-09-28); omitted means one day.")
	RescheduleCmd.Flags().BoolP("notifyattendees", "", false, "`false` saves the change without emailing attendees, and copies of the event in the calendars of other Hostinger mailboxes are not updated. Omitted or null means `true`.")
	RescheduleCmd.Flags().StringP("recurrenceid", "", "", "`recurrenceId` of one occurrence, copied unchanged from a `listCalendarEvents` item. Required when `scope` is `this` or `future`, ignored for `all`.")
	RescheduleCmd.Flags().StringP("scope", "", "", "For a recurring event: `this` moves the occurrence at `recurrenceId`, `future` moves it and every later one, `all` moves the whole series. Omitted or null means `all`. Required when `recurrenceId` is sent. (one of: this, future, all)")
	RescheduleCmd.Flags().StringP("start", "", "", "ISO 8601 date-time with offset; without an offset it is read as UTC. For an all-day event a date is enough (2026-09-28).")
	RescheduleCmd.Flags().StringP("timezone", "", "", "IANA time zone of the event. Required when `allDay` is false, ignored when `allDay` is true.")
	RescheduleCmd.MarkFlagRequired("allday")
	RescheduleCmd.MarkFlagRequired("start")
	RescheduleCmd.RegisterFlagCompletionFunc("if-match", cobra.NoFileCompletions)
	RescheduleCmd.RegisterFlagCompletionFunc("allday", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	RescheduleCmd.RegisterFlagCompletionFunc("end", cobra.NoFileCompletions)
	RescheduleCmd.RegisterFlagCompletionFunc("notifyattendees", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	RescheduleCmd.RegisterFlagCompletionFunc("recurrenceid", cobra.NoFileCompletions)
	RescheduleCmd.RegisterFlagCompletionFunc("scope", cobra.FixedCompletions([]cobra.Completion{"this", "future", "all"}, cobra.ShellCompDirectiveNoFileComp))
	RescheduleCmd.RegisterFlagCompletionFunc("start", cobra.NoFileCompletions)
	RescheduleCmd.RegisterFlagCompletionFunc("timezone", cobra.NoFileCompletions)
}

func rescheduleParams(cmd *cobra.Command) *client.RescheduleCalendarEventParams {
	params := &client.RescheduleCalendarEventParams{}
	if cmd.Flags().Changed("if-match") {
		v, _ := cmd.Flags().GetString("if-match")
		params.IfMatch = &v
	}
	return params
}

func rescheduleBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	allDayVal, _ := cmd.Flags().GetBool("allday")
	body["allDay"] = allDayVal
	if cmd.Flags().Changed("end") {
		v, _ := cmd.Flags().GetString("end")
		body["end"] = v
	}
	if cmd.Flags().Changed("notifyattendees") {
		v, _ := cmd.Flags().GetBool("notifyattendees")
		body["notifyAttendees"] = v
	}
	if cmd.Flags().Changed("recurrenceid") {
		v, _ := cmd.Flags().GetString("recurrenceid")
		body["recurrenceId"] = v
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
	return body, nil
}
