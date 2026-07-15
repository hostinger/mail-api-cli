## hostinger-mail webhooks regenerate-secret

Regenerate webhook secret

### Synopsis

Regenerate the webhook secret. The previous secret is immediately invalidated. The new secret is returned once.

```
hostinger-mail webhooks regenerate-secret <mailbox-resource-id> <webhook> [flags]
```

### Options

```
  -h, --help   help for regenerate-secret
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

