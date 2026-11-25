## hostinger-mail calendar-events list

List calendar events

### Synopsis

List the events of one calendar that overlap a time range. It covers one calendar only: get the calendars from `listCalendars` and call it for each. The range runs from `start` (inclusive) to `end` (exclusive), at most 32 days; for a longer period call it again with the next range. Recurring events are expanded: each occurrence is a separate item with the same `uuid` and its own `recurrenceId`. Results are not paginated.
Items hold summary fields only. For description, location, attendees and organizer call `getCalendarEvent` with the `calendarUuid`, `uuid` (as `eventUuid`) and `recurrenceId` of the item. Times are in UTC; `timeZone` is the zone the event was created in. `participationStatus` is the answer of the mailbox to an invitation from someone else, null for its own events. `status` CANCELLED means the organizer cancelled the event.

```
hostinger-mail calendar-events list <mailbox-resource-id> <calendar-uuid> [flags]
```

### Examples

```
  hostinger-mail calendar-events list my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --start 2025-01-31T10:00:00Z --end 2025-01-31T10:00:00Z
```

### Options

```
      --end start           End of the range, exclusive. ISO 8601 date-time after start and at most 32 days after it.
  -h, --help[=true|false]   help for list
      --start string        Start of the range, inclusive. ISO 8601 date-time with offset.
      --timezone start      IANA time zone used to place all-day events and events without a time zone on days, also across a daylight saving change. Defaults to the fixed offset of start.
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

