## hostinger-mail messages list

List messages

### Synopsis

List messages in a folder. Use POST /search for filtering. Sort fields: uid, date, size (prefix with `-` for descending). Default `-uid`.

```
hostinger-mail messages list <mailbox-resource-id> <folder> [flags]
```

### Options

```
  -h, --help          help for list
      --page int      Page number (1-based). (default 1)
      --perpage int   Items per page (max 100). (default 25)
      --sort -        Sort field with optional - prefix for descending. Allowed: uid, date, size. (default "-uid")
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

