## hostinger-mail webhooks

Webhooks commands

```
hostinger-mail webhooks [flags]
```

### Options

```
  -h, --help   help for webhooks
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

* [hostinger-mail](hostinger-mail.md)	 - Hostinger Mail API Command Line Interface
* [hostinger-mail webhooks create](hostinger-mail_webhooks_create.md)	 - Create webhook
* [hostinger-mail webhooks delete](hostinger-mail_webhooks_delete.md)	 - Delete webhook
* [hostinger-mail webhooks get](hostinger-mail_webhooks_get.md)	 - Get webhook
* [hostinger-mail webhooks list](hostinger-mail_webhooks_list.md)	 - List webhooks
* [hostinger-mail webhooks regenerate-secret](hostinger-mail_webhooks_regenerate-secret.md)	 - Regenerate webhook secret
* [hostinger-mail webhooks test](hostinger-mail_webhooks_test.md)	 - Test webhook
* [hostinger-mail webhooks update](hostinger-mail_webhooks_update.md)	 - Update webhook

