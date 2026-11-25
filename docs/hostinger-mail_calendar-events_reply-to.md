## hostinger-mail calendar-events reply-to

Reply to calendar event

### Synopsis

Answer an invitation from someone else: accept, decline or tentatively accept. It saves the answer on the event and emails it to the organizer, which cannot be recalled. `calendarUuid` and `eventUuid` come from `listCalendarEvents`.
It works only when `getCalendarEvent` returns `replyable` true. An event the mailbox organizes or is not invited to is rejected with `403` and code `ERR_EVENT_NOT_REPLYABLE`, an event the organizer cancelled with `409` and code `ERR_EVENT_CANCELLED`. Send the `recurrenceId` of an item to answer one occurrence of a recurring event, or omit it to answer the whole series. Returns the event with the new `participationStatus`.

```
hostinger-mail calendar-events reply-to <mailbox-resource-id> <calendar-uuid> <event-uuid> [flags]
```

### Examples

```
  hostinger-mail calendar-events reply-to my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --participationstatus ACCEPTED
```

### Options

```
  -h, --help[=true|false]            help for reply-to
      --if-match etag                Optional etag from `getCalendarEvent`. When the event changed since that call, the answer is not saved and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.
      --participationstatus string   The answer to the invitation. (one of: ACCEPTED, DECLINED, TENTATIVE)
      --recurrenceid recurrenceId    recurrenceId of one occurrence, copied unchanged from a `listCalendarEvents` item, to answer only that occurrence. Omitted or null answers the whole event.
```

### Options inherited from parent commands

```
      --config string           Config file (default is $HOME/.hostinger-mail.yaml)
      --dry-run[=true|false]    Print the request as a curl command instead of sending it
  -o, --format string           Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string               Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input[=true|false]   Never prompt; fail when an argument, flag or confirmation is missing
      --verbose[=true|false]    Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger-mail calendar-events](hostinger-mail_calendar-events.md)	 - Calendar Events commands

