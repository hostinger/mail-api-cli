## hostinger-mail messages move-bulk

Move messages

### Synopsis

Move multiple messages from a source folder to a target folder.

```
hostinger-mail messages move-bulk <mailbox-resource-id> <folder> [flags]
```

### Examples

```
  hostinger-mail messages move-bulk my-mailbox-resource-id my-folder --targetfolder INBOX.Sent --uids 42,43,44
```

### Options

```
  -h, --help                  help for move-bulk
      --targetfolder string   Destination folder path.
      --uids ints             Message UIDs to move. 1-100 entries, each > 0.
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

