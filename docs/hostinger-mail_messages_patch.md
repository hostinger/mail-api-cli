## hostinger-mail messages patch

Update message flags

### Synopsis

Add and/or remove flags on a single message. Returns the updated message.

```
hostinger-mail messages patch <mailbox-resource-id> <folder> <uid> [flags]
```

### Options

```
      --addflags strings      IMAP flags to set on the message, e.g. \Seen, \Flagged, \Answered, $forwarded.
  -h, --help                  help for patch
      --removeflags strings   IMAP flags to clear from the message.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

