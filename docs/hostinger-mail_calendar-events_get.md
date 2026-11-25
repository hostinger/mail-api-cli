## hostinger-mail calendar-events get

Get calendar event

### Synopsis

Retrieve the full details of one event: title, times, description, location, meeting link, recurrence rule, reminder, organizer, attendees with their answers, and its `etag`. `calendarUuid` and `eventUuid` come from `listCalendarEvents` (the `calendarUuid` and `uuid` of an item). For one occurrence of a recurring event pass the `recurrenceId` of the item; without it the whole series is returned and `recurrenceId` is null. A `recurrenceId` that matches no occurrence returns `404`.
Call it before `updateCalendarEvent`, which replaces every field, so the values that should stay can be sent back. `replyable` true means the mailbox was invited by someone else and can answer with `replyToCalendarEvent`. `editable` true means the event can be changed or moved.

```
hostinger-mail calendar-events get <mailbox-resource-id> <calendar-uuid> <event-uuid> [flags]
```

### Examples

```
  hostinger-mail calendar-events get my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d
```

### Options

```
  -h, --help[=true|false]           help for get
      --recurrenceid recurrenceId   recurrenceId of one occurrence, copied unchanged from a `listCalendarEvents` item. Omit it for a one-off event or for the whole series.
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

