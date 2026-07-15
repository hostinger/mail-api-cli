## hostinger-mail messages move-bulk

Move messages

### Synopsis

Move multiple messages from a source folder to a target folder.

```
hostinger-mail messages move-bulk <mailbox-resource-id> <folder> [flags]
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

