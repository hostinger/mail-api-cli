## hostinger-mail calendar-events create

Create calendar event

### Synopsis

Create an event in one calendar of the mailbox. `calendarUuid` comes from `listCalendars`: use the calendar with `default` true unless another one is named. A calendar the mailbox can only read is rejected with `403` and code `ERR_EVENT_NOT_EDITABLE`.
The mailbox becomes the organizer and every address in `attendees` immediately receives an email invitation, which cannot be recalled. Without attendees no email is sent.
A timed event needs `allDay` false, `start`, `end` and `timeZone`. An all-day event needs `allDay` true and a `start` date; `end` is optional. Set `recurrenceRule` for a repeating event and `reminder` for an alarm.
Returns the created event. Its `uuid` is the `eventUuid` for `getCalendarEvent`, `updateCalendarEvent`, `rescheduleCalendarEvent` and `deleteCalendarEvent`.

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
hostinger-mail calendar-events create <mailbox-resource-id> <calendar-uuid> [flags]
```

### Examples

```
  hostinger-mail calendar-events create my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --allday=false --start 2026-09-28T13:00:00+03:00 --title Planning
```

### Options

```
      --allday[=true|false]        true for an all-day event (dates only, `timeZone` ignored), `false` for an event at a time of day.
      --attendees string           People to invite; each receives an email invitation. Unique addresses, at most 100 not counting the mailbox itself, across all occurrences. The own address of the mailbox is stored as accepted and not invited. Omitted, null or empty for none. (JSON, @file or -; shape in the help above)
      --color string               Color of the event as #RRGGBB or #RRGGBBAA. Omitted or null uses the calendar color.
      --description string         Notes of the event. Omitted or null for none.
      --end start                  Exclusive end, same format as start. Required for a timed event and not before `start`. For an all-day event the day after the last day (2026-09-29 for one day on 2026-09-28); omitted means one day.
  -h, --help[=true|false]          help for create
      --location string            Place of the event. Omitted or null for none.
      --meetinglink string         http or https URL of the online meeting. Omitted or null for none.
      --recurrencerule frequency   Omitted or null for a one-off event. A repeating event sends every field of the rule, null when unused: frequency DAILY, WEEKLY, MONTHLY or YEARLY, `interval` 1 to 1000, and `count` or `until` to end it, not both. (JSON, @file or -; shape in the help above)
      --reminder value             Alarm before the start: value 0 to 1000 and `unit`. Omitted or null for none. (JSON, @file or -; shape in the help above)
      --start string               ISO 8601 date-time with offset; without an offset it is read as UTC. For an all-day event a date is enough (2026-09-28).
      --timezone allDay            IANA time zone of the event. Required when allDay is false, ignored when `allDay` is true.
      --title string               Title of the event.
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

