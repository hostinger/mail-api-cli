## hostinger-mail webhooks list

List webhooks

### Synopsis

List webhooks for a mailbox.

```
hostinger-mail webhooks list <mailbox-resource-id> [flags]
```

### Options

```
  -h, --help            help for list
      --page int        Page number (1-based). (default 1)
      --perpage int     Items per page (max 1000). (default 15)
      --status string   Return only webhooks with this status. (one of: active, paused, disabled)
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

