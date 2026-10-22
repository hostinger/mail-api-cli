## hostinger-mail

Hostinger Mail API Command Line Interface

### Synopsis

Hostinger Mail API Command Line Interface.

Get started:
  1. Sign in:        hostinger-mail config set api_token   (or set HOSTINGER_MAIL_API_TOKEN)

Output is a table on a terminal and JSON when piped; pick one with --format
(table|json|yaml|tree) or filter it with --jq. Missing arguments and required
flags are prompted for in a terminal (disable with --no-input). Commands that
delete or stop things ask for confirmation (skip with --yes), and --dry-run
prints the request as curl without sending it.

Documentation: https://github.com/hostinger/mail-api-cli/tree/main/docs

### Options

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --dry-run         Print the request as a curl command instead of sending it
  -o, --format string   Output format: table|json|yaml|tree (default: table on a terminal, json when piped)
  -h, --help            help for hostinger-mail
      --jq string       Filter the JSON response with a jq expression, e.g. '.data[].id'
      --no-input        Never prompt; fail when an argument, flag or confirmation is missing
      --verbose         Print the config file used and each HTTP request to stderr
```

### SEE ALSO

* [hostinger-mail account](hostinger-mail_account.md)	 - Account commands
* [hostinger-mail auth](hostinger-mail_auth.md)	 - Sign in and check which credentials are used
* [hostinger-mail completion](hostinger-mail_completion.md)	 - Generate the autocompletion script for the specified shell
* [hostinger-mail config](hostinger-mail_config.md)	 - Read and write CLI settings
* [hostinger-mail feedback](hostinger-mail_feedback.md)	 - Feedback commands
* [hostinger-mail folders](hostinger-mail_folders.md)	 - Folders commands
* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands
* [hostinger-mail quota](hostinger-mail_quota.md)	 - Quota commands
* [hostinger-mail send](hostinger-mail_send.md)	 - Send commands
* [hostinger-mail version](hostinger-mail_version.md)	 - Show the current version
* [hostinger-mail webhooks](hostinger-mail_webhooks.md)	 - Webhooks commands

