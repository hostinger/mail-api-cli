## hostinger-mail account current

Get the authenticated account

### Synopsis

Returns the authenticated account and the mailboxes it can manage.

```
hostinger-mail account current [flags]
```

### Examples

```
  hostinger-mail account current
```

### Options

```
  -h, --help   help for current
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger-mail account](hostinger-mail_account.md)	 - Account commands

