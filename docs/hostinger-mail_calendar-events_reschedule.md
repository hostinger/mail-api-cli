## hostinger-mail calendar-events reschedule

Reschedule calendar event

### Synopsis

Move an event to another time or change its length, keeping everything else (title, attendees, description and so on) as stored, so the event does not have to be fetched first. To change other fields use `updateCalendarEvent`. `calendarUuid` and `eventUuid` come from `listCalendarEvents`. Only an event with `editable` true can be moved; any other is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`.
For a recurring event, `scope` `this` moves only the occurrence at `recurrenceId`, `future` ends the series before that occurrence and starts a new one with a new `uuid`, `all` (the default) moves the whole series: only the time of day and the length apply and the series keeps the date of its first occurrence.
Attendees receive an email update unless `notifyAttendees` is false. Returns the moved event, or the moved occurrence for `scope` `this`.

```
hostinger-mail calendar-events reschedule <mailbox-resource-id> <calendar-uuid> <event-uuid> [flags]
```

### Examples

```
  hostinger-mail calendar-events reschedule my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --allday=false --start 2026-09-28T13:00:00+03:00
```

### Options

```
      --allday[=true|false]            true for an all-day event (dates only, `timeZone` ignored), `false` for an event at a time of day.
      --end start                      Exclusive end, same format as start. Required for a timed event and not before `start`. For an all-day event the day after the last day (2026-09-29 for one day on 2026-09-28); omitted means one day.
  -h, --help[=true|false]              help for reschedule
      --if-match etag                  Optional etag from `getCalendarEvent`. When the event changed since that call, nothing is changed and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.
      --notifyattendees[=true|false]   false saves the change without emailing attendees, and copies of the event in the calendars of other Hostinger mailboxes are not updated. Omitted or null means `true`.
      --recurrenceid recurrenceId      recurrenceId of one occurrence, copied unchanged from a `listCalendarEvents` item. Required when `scope` is `this` or `future`, ignored for `all`.
      --scope this                     For a recurring event: this moves the occurrence at `recurrenceId`, `future` moves it and every later one, `all` moves the whole series. Omitted or null means `all`. Required when `recurrenceId` is sent. (one of: this, future, all)
      --start string                   ISO 8601 date-time with offset; without an offset it is read as UTC. For an all-day event a date is enough (2026-09-28).
      --timezone allDay                IANA time zone of the event. Required when allDay is false, ignored when `allDay` is true.
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

