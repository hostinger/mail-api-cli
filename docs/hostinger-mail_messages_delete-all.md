## hostinger-mail messages delete-all

Delete all messages

### Synopsis

Permanently delete every message in a folder (empty the folder).

```
hostinger-mail messages delete-all <mailbox-resource-id> <folder> [flags]
```

### Examples

```
  hostinger-mail messages delete-all my-mailbox-resource-id my-folder
```

### Options

```
  -h, --help   help for delete-all
  -y, --yes    Skip the confirmation prompt
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

