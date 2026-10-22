## hostinger-mail quota get

Get mailbox quota

### Synopsis

Retrieve storage and message quota usage for the managed mailbox.

```
hostinger-mail quota get <mailbox-resource-id> [flags]
```

### Examples

```
  hostinger-mail quota get my-mailbox-resource-id
```

### Options

```
  -h, --help   help for get
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

* [hostinger-mail quota](hostinger-mail_quota.md)	 - Quota commands

