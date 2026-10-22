## hostinger-mail webhooks delete

Delete webhook

### Synopsis

Delete a webhook.

```
hostinger-mail webhooks delete <mailbox-resource-id> <webhook> [flags]
```

### Examples

```
  hostinger-mail webhooks delete my-mailbox-resource-id 9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

