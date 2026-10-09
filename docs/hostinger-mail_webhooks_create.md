## hostinger-mail webhooks create

Create webhook

### Synopsis

Create a webhook. The response includes the one-time `secret` — store it securely as it is never returned again.

```
hostinger-mail webhooks create <mailbox-resource-id> [flags]
```

### Options

```
      --description string                   Optional free-text note about the webhook purpose.
      --events strings                       Event types that trigger a delivery. (one of: message.received)
  -h, --help                                 help for create
      --name string                          Human-readable webhook name.
      --status string                        Initial delivery state. Only active webhooks receive events. (one of: active, paused, disabled) (default "active")
      --url Authorization: Bearer <secret>   HTTPS endpoint that receives POST deliveries, authenticated with the webhook secret as Authorization: Bearer <secret>. Must be a public domain name (no IPs or internal hosts).
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

