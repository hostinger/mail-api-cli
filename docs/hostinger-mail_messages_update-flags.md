## hostinger-mail messages update-flags

Update flags on multiple messages

### Synopsis

Add and/or remove flags on multiple messages. Returns 200 when all UIDs succeed, 207 with per-UID outcome when some fail.

```
hostinger-mail messages update-flags <mailbox-resource-id> <folder> [flags]
```

### Examples

```
  hostinger-mail messages update-flags my-mailbox-resource-id my-folder --uids 42,43,44
```

### Options

```
      --addflags strings      IMAP flags to set on every listed message, e.g. \Seen, \Flagged, \Answered, $forwarded.
  -h, --help                  help for update-flags
      --removeflags strings   IMAP flags to clear from every listed message.
      --uids ints             Message UIDs to update. 1-100 entries, each > 0.
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

