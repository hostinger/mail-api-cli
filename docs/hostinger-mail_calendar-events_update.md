## hostinger-mail calendar-events update

Update calendar event

### Synopsis

Change the details of an event: title, times, description, location, attendees, recurrence and so on. Every field replaces the stored value and an omitted optional field is cleared, so call `getCalendarEvent` first and send back the values that should stay. To only move an event in time, use `rescheduleCalendarEvent`. Only an event with `editable` true can be changed; any other is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`.
For a recurring event, `scope` `this` changes only the occurrence at `recurrenceId`, `future` ends the series before that occurrence and starts a new one with a new `uuid`, `all` (the default) changes the whole series, keeping the date of its first occurrence.
Attendees receive an email update, added attendees an invitation and removed ones a cancellation, unless `notifyAttendees` is false. Returns the changed event, or the changed occurrence for `scope` `this`.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --attendees
    [
      {
        "email": "john@example.com",
        "name": "John"
      }
    ]

  --recurrencerule
    {
      "byDay": [
        {
          "day": "friday",
          "position": -1
        }
      ],
      "byMonth": [
        3
      ],
      "byMonthDay": [
        15
      ],
      "count": 10,
      "frequency": "WEEKLY",
      "interval": 1,
      "until": "string",
      "weekStart": "monday"
    }

  --reminder
    {
      "unit": "minutes",
      "value": 15
    }

```
hostinger-mail calendar-events update <mailbox-resource-id> <calendar-uuid> <event-uuid> [flags]
```

### Examples

```
  hostinger-mail calendar-events update my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --allday=false --start 2026-09-28T13:00:00+03:00 --title Planning
```

### Options

```
      --allday[=true|false]            true for an all-day event (dates only, `timeZone` ignored), `false` for an event at a time of day.
      --attendees string               Full list of invited people. Added addresses receive an invitation and removed ones a cancellation. Unique addresses, at most 100 not counting the mailbox itself, across all occurrences. The own address of the mailbox is stored as accepted and not invited. Omitted, null or empty for none. (JSON, @file or -; shape in the help above)
      --color string                   Color of the event as #RRGGBB or #RRGGBBAA. Omitted or null uses the calendar color.
      --description string             Notes of the event. Omitted or null clears it.
      --end start                      Exclusive end, same format as start. Required for a timed event and not before `start`. For an all-day event the day after the last day (2026-09-29 for one day on 2026-09-28); omitted means one day.
  -h, --help[=true|false]              help for update
      --if-match etag                  Optional etag from `getCalendarEvent`. When the event changed since that call, nothing is changed and the request fails with `412`: fetch the event again and retry. Omit it to apply the request to the current version.
      --location string                Place of the event. Omitted or null clears it.
      --meetinglink string             http or https URL of the online meeting. Omitted or null clears it.
      --notifyattendees[=true|false]   false saves the change without emailing attendees, and copies of the event in the calendars of other Hostinger mailboxes are not updated. Omitted or null means `true`.
      --recurrenceid recurrenceId      recurrenceId of one occurrence, copied unchanged from a `listCalendarEvents` item. Required when `scope` is `this` or `future`, ignored for `all`.
      --recurrencerule frequency       Omitted or null makes it a one-off event. A repeating event sends every field of the rule, null when unused: frequency DAILY, WEEKLY, MONTHLY or YEARLY, `interval` 1 to 1000, and `count` or `until` to end it, not both. (JSON, @file or -; shape in the help above)
      --reminder value                 Alarm before the start: value 0 to 1000 and `unit`. Omitted or null clears it. (JSON, @file or -; shape in the help above)
      --scope this                     For a recurring event: this changes the occurrence at `recurrenceId`, `future` changes it and every later one, `all` changes the whole series. Omitted or null means `all`. Required when `recurrenceId` is sent. (one of: this, future, all)
      --start string                   ISO 8601 date-time with offset; without an offset it is read as UTC. For an all-day event a date is enough (2026-09-28).
      --timezone allDay                IANA time zone of the event. Required when allDay is false, ignored when `allDay` is true.
      --title string                   Title of the event.
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

