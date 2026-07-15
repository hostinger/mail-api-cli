## hostinger-mail webhooks update

Update webhook

### Synopsis

Partially update a webhook.

```
hostinger-mail webhooks update <mailbox-resource-id> <webhook> [flags]
```

### Options

```
      --description string   
      --events strings       (one of: message.received)
  -h, --help                 help for update
      --name string          
      --status string        (one of: active, paused, disabled)
      --url string           
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

