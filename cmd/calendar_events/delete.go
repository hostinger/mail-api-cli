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

var DeleteCmd = &cobra.Command{
	Use:               "delete <mailbox-resource-id> <calendar-uuid> <event-uuid>",
	Short:             "Delete calendar event",
	Long:              "Delete an event, one occurrence of a recurring event, or a recurring event from one occurrence on. This cannot be undone. `calendarUuid` and `eventUuid` come from `listCalendarEvents`.\nFor a one-off event omit `scope` and `recurrenceId`. For a recurring event, `scope` `this` removes only the occurrence at `recurrenceId`, `future` removes it and every later one, `all` (the default) deletes the whole series.\nAn event with `editable` true can be deleted with any scope, and its attendees receive a cancellation or the updated series by email. An invitation from someone else on the own calendar of the mailbox can be deleted with `all` or `this`: the organizer receives a decline for the removed occurrences. Any other event is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`. `notifyAttendees` false deletes without sending any email.",
	Example:           "  hostinger-mail calendar-events delete my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "scope", []string{"this", "future", "all"}); err != nil {
			return err
		}
		if err := utils.EnumCheck(cmd, "notifyattendees", []string{"true", "false", "1", "0"}); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "recurrenceid", "date-time"); err != nil {
			return err
		}
		params := deleteParams(cmd)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Delete calendar event", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DeleteCalendarEventWithResponse(ctx, args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]), params)
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
	utils.RegisterArgs(DeleteCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"), utils.UUIDArg("event-uuid"))
	DeleteCmd.Flags().StringP("scope", "", "", "For a recurring event: `this` removes the occurrence at `recurrenceId`, `future` removes it and every later one, `all` deletes the whole series. Omitted means `all`. Required when `recurrenceId` is sent. (one of: this, future, all)")
	DeleteCmd.Flags().StringP("recurrenceid", "", "", "`recurrenceId` of one occurrence, copied unchanged from a `listCalendarEvents` item. Required when `scope` is `this` or `future`. A value that matches no occurrence, or any value on a one-off event, returns `404`.")
	DeleteCmd.Flags().StringP("notifyattendees", "", "true", "`false` deletes without emailing attendees a cancellation or update, or the organizer a decline. Omitted means `true`. (one of: true, false, 1, 0)")
	DeleteCmd.Flags().StringP("if-match", "", "", "Optional `etag` from `getCalendarEvent`. When the event changed since that call, nothing is deleted and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.")
	DeleteCmd.RegisterFlagCompletionFunc("scope", cobra.FixedCompletions([]cobra.Completion{"this", "future", "all"}, cobra.ShellCompDirectiveNoFileComp))
	DeleteCmd.RegisterFlagCompletionFunc("recurrenceid", cobra.NoFileCompletions)
	DeleteCmd.RegisterFlagCompletionFunc("notifyattendees", cobra.FixedCompletions([]cobra.Completion{"true", "false", "1", "0"}, cobra.ShellCompDirectiveNoFileComp))
	DeleteCmd.RegisterFlagCompletionFunc("if-match", cobra.NoFileCompletions)
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deleteParams(cmd *cobra.Command) *client.DeleteCalendarEventParams {
	params := &client.DeleteCalendarEventParams{}
	if cmd.Flags().Changed("scope") {
		v, _ := cmd.Flags().GetString("scope")
		e := client.DeleteCalendarEventParamsScope(v)
		params.Scope = &e
	}
	if cmd.Flags().Changed("recurrenceid") {
		v, _ := cmd.Flags().GetString("recurrenceid")
		t := utils.StringToTime(v)
		params.RecurrenceId = &t
	}
	if cmd.Flags().Changed("notifyattendees") {
		v, _ := cmd.Flags().GetString("notifyattendees")
		e := client.DeleteCalendarEventParamsNotifyAttendees(v)
		params.NotifyAttendees = &e
	}
	if cmd.Flags().Changed("if-match") {
		v, _ := cmd.Flags().GetString("if-match")
		params.IfMatch = &v
	}
	return params
}
