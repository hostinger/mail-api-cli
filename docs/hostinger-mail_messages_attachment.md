## hostinger-mail messages attachment

Download message attachment

### Synopsis

Download a message attachment as `application/octet-stream`.

```
hostinger-mail messages attachment <mailbox-resource-id> <folder> <uid> <attachment-id> [flags]
```

### Examples

```
  hostinger-mail messages attachment my-mailbox-resource-id my-folder 1234567 my-attachment-id
```

### Options

```
  -h, --help   help for attachment
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

