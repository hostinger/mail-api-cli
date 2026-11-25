## hostinger-mail calendars list

List calendars

### Synopsis

List the calendars of a mailbox: its own calendars and the calendars shared with it, in display order. Call it first to get the `calendarUuid` every calendar event operation needs: the `uuid` of each calendar. The calendar with `default` true is where new events go unless another one is named. `visible` false means the calendar is hidden in the calendar view; its events still exist.

```
hostinger-mail calendars list <mailbox-resource-id> [flags]
```

### Examples

```
  hostinger-mail calendars list my-mailbox-resource-id
```

### Options

```
      --all[=true|false]    Fetch every page and print all items
  -h, --help[=true|false]   help for list
      --page int            Page number (1-based). (default 1)
      --perpage int         Items per page (max 200). (default 25)
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

