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

var ReplyToCmd = &cobra.Command{
	Use:               "reply-to <mailbox-resource-id> <calendar-uuid> <event-uuid>",
	Short:             "Reply to calendar event",
	Long:              "Answer an invitation from someone else: accept, decline or tentatively accept. It saves the answer on the event and emails it to the organizer, which cannot be recalled. `calendarUuid` and `eventUuid` come from `listCalendarEvents`.\nIt works only when `getCalendarEvent` returns `replyable` true. An event the mailbox organizes or is not invited to is rejected with `403` and code `ERR_EVENT_NOT_REPLYABLE`, an event the organizer cancelled with `409` and code `ERR_EVENT_CANCELLED`. Send the `recurrenceId` of an item to answer one occurrence of a recurring event, or omit it to answer the whole series. Returns the event with the new `participationStatus`.",
	Example:           "  hostinger-mail calendar-events reply-to my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --participationstatus ACCEPTED",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "participationstatus", []string{"ACCEPTED", "DECLINED", "TENTATIVE"}); err != nil {
			return err
		}
		params := replyToParams(cmd)
		body, err := replyToBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReplyToCalendarEventWithBodyWithResponse(ctx, args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]), params, "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ReplyToCmd, utils.StringArg("mailbox-resource-id"), utils.UUIDArg("calendar-uuid"), utils.UUIDArg("event-uuid"))
	ReplyToCmd.Flags().StringP("if-match", "", "", "Optional `etag` from `getCalendarEvent`. When the event changed since that call, the answer is not saved and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.")
	ReplyToCmd.Flags().StringP("participationstatus", "", "", "The answer to the invitation. (one of: ACCEPTED, DECLINED, TENTATIVE)")
	ReplyToCmd.Flags().StringP("recurrenceid", "", "", "`recurrenceId` of one occurrence, copied unchanged from a `listCalendarEvents` item, to answer only that occurrence. Omitted or null answers the whole event.")
	ReplyToCmd.MarkFlagRequired("participationstatus")
	ReplyToCmd.RegisterFlagCompletionFunc("if-match", cobra.NoFileCompletions)
	ReplyToCmd.RegisterFlagCompletionFunc("participationstatus", cobra.FixedCompletions([]cobra.Completion{"ACCEPTED", "DECLINED", "TENTATIVE"}, cobra.ShellCompDirectiveNoFileComp))
	ReplyToCmd.RegisterFlagCompletionFunc("recurrenceid", cobra.NoFileCompletions)
}

func replyToParams(cmd *cobra.Command) *client.ReplyToCalendarEventParams {
	params := &client.ReplyToCalendarEventParams{}
	if cmd.Flags().Changed("if-match") {
		v, _ := cmd.Flags().GetString("if-match")
		params.IfMatch = &v
	}
	return params
}

func replyToBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	participationStatusVal, _ := cmd.Flags().GetString("participationstatus")
	body["participationStatus"] = participationStatusVal
	if cmd.Flags().Changed("recurrenceid") {
		v, _ := cmd.Flags().GetString("recurrenceid")
		body["recurrenceId"] = v
	}
	return body, nil
}
