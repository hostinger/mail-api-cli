## hostinger-mail messages update-flags

Update message flags

### Synopsis

Add and/or remove flags on multiple messages. Returns 200 when all UIDs succeed, 207 with per-UID outcome when some fail.

```
hostinger-mail messages update-flags <mailbox-resource-id> <folder> [flags]
```

### Options

```
      --addflags strings      
  -h, --help                  help for update-flags
      --removeflags strings   
      --uids ints             Message UIDs to update. 1-100 entries, each > 0.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

