## hostinger-mail webhooks update

Update webhook

### Synopsis

Partially update a webhook.

```
hostinger-mail webhooks update <mailbox-resource-id> <webhook> [flags]
```

### Options

```
      --description string                   Free-text note about the webhook purpose. Send null to clear.
      --events strings                       Event types that trigger a delivery. Replaces the current list. (one of: message.received)
  -h, --help                                 help for update
      --name string                          Human-readable webhook name.
      --status string                        Delivery state. Only active webhooks receive events; paused keeps config but stops deliveries. (one of: active, paused, disabled)
      --url Authorization: Bearer <secret>   HTTPS endpoint that receives POST deliveries, authenticated with the webhook secret as Authorization: Bearer <secret>. Must be a public domain name (no IPs or internal hosts).
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

