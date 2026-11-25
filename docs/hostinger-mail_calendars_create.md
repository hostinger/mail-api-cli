## hostinger-mail calendars create

Create calendar

### Synopsis

Create a calendar owned by the mailbox. The number of calendars a mailbox can own is limited; when the limit is reached the request fails with `400` and code `ERR_CALENDAR_LIMIT_REACHED`.

```
hostinger-mail calendars create <mailbox-resource-id> [flags]
```

### Examples

```
  hostinger-mail calendars create my-mailbox-resource-id --displayname Work
```

### Options

```
      --color string           Hex color as #RRGGBB or #RRGGBBAA. Omitted or null uses the default color.
      --description string     Optional free-text note about the calendar.
      --displayname string     Name of the calendar.
  -h, --help[=true|false]      help for create
      --visible[=true|false]   Whether the mailbox shows the calendar in its calendar view. Omitted or null means true.
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

* [hostinger-mail calendars](hostinger-mail_calendars.md)	 - Calendars commands

