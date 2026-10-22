## hostinger-mail messages get

Get message

### Synopsis

Retrieve a single message by UID.

```
hostinger-mail messages get <mailbox-resource-id> <folder> <uid> [flags]
```

### Examples

```
  hostinger-mail messages get my-mailbox-resource-id my-folder 1234567
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

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

