## hostinger-mail webhooks create

Create webhook

### Synopsis

Create a webhook. The response includes the one-time `secret` — store it securely as it is never returned again.

```
hostinger-mail webhooks create <mailbox-resource-id> [flags]
```

### Options

```
      --description string   
      --events strings       (one of: message.received)
  -h, --help                 help for create
      --name string          
      --status string        (one of: active, paused, disabled) (default "active")
      --url string           
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

