## hostinger-mail messages delete-bulk

Delete messages

### Synopsis

Permanently delete multiple messages from a folder.

```
hostinger-mail messages delete-bulk <mailbox-resource-id> <folder> [flags]
```

### Options

```
  -h, --help        help for delete-bulk
      --uids ints   Message UIDs to delete. 1-100 entries, each > 0.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

