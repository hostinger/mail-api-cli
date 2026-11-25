## hostinger-mail calendars update

Update calendar

### Synopsis

Update the name, description, color and visibility of a calendar. Name, description and color replace the stored values: an omitted `description` is cleared and an omitted `color` resets to the default color. An omitted `visible` keeps the current value.

```
hostinger-mail calendars update <mailbox-resource-id> <calendar-uuid> [flags]
```

### Examples

```
  hostinger-mail calendars update my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d --displayname Work
```

### Options

```
      --color string           Hex color as #RRGGBB or #RRGGBBAA. Omitted or null resets it to the default color.
      --description string     Free-text note about the calendar. Omitted or null clears it.
      --displayname string     Name of the calendar.
  -h, --help[=true|false]      help for update
      --visible[=true|false]   Whether the mailbox shows the calendar in its calendar view. Omitted or null keeps the current value.
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

