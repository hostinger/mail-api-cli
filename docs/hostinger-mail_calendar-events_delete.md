## hostinger-mail calendar-events delete

Delete calendar event

### Synopsis

Delete an event, one occurrence of a recurring event, or a recurring event from one occurrence on. This cannot be undone. `calendarUuid` and `eventUuid` come from `listCalendarEvents`.
For a one-off event omit `scope` and `recurrenceId`. For a recurring event, `scope` `this` removes only the occurrence at `recurrenceId`, `future` removes it and every later one, `all` (the default) deletes the whole series.
An event with `editable` true can be deleted with any scope, and its attendees receive a cancellation or the updated series by email. An invitation from someone else on the own calendar of the mailbox can be deleted with `all` or `this`: the organizer receives a decline for the removed occurrences. Any other event is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`. `notifyAttendees` false deletes without sending any email.

```
hostinger-mail calendar-events delete <mailbox-resource-id> <calendar-uuid> <event-uuid> [flags]
```

### Examples

```
  hostinger-mail calendar-events delete my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d
```

### Options

```
  -h, --help[=true|false]           help for delete
      --if-match etag               Optional etag from `getCalendarEvent`. When the event changed since that call, nothing is deleted and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.
      --notifyattendees false       false deletes without emailing attendees a cancellation or update, or the organizer a decline. Omitted means `true`. (one of: true, false, 1, 0) (default "true")
      --recurrenceid recurrenceId   recurrenceId of one occurrence, copied unchanged from a `listCalendarEvents` item. Required when `scope` is `this` or `future`. A value that matches no occurrence, or any value on a one-off event, returns `404`.
      --scope this                  For a recurring event: this removes the occurrence at `recurrenceId`, `future` removes it and every later one, `all` deletes the whole series. Omitted means `all`. Required when `recurrenceId` is sent. (one of: this, future, all)
  -y, --yes[=true|false]            Skip the confirmation prompt
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

