## hostinger-mail webhooks update

Update webhook

### Synopsis

Partially update a webhook.

```
hostinger-mail webhooks update <mailbox-resource-id> <webhook> [flags]
```

### Examples

```
  hostinger-mail webhooks update my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d
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
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

