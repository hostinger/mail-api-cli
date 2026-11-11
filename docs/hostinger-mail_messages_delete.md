## hostinger-mail messages delete

Delete message

### Synopsis

Permanently delete a single message.

```
hostinger-mail messages delete <mailbox-resource-id> <folder> <uid> [flags]
```

### Examples

```
  hostinger-mail messages delete my-mailbox-resource-id my-folder 1234567
```

### Options

```
  -h, --help[=true|false]   help for delete
  -y, --yes[=true|false]    Skip the confirmation prompt
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

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

