## hostinger-mail feedback submit

Submit feedback

### Synopsis

Send the user's feedback about the Hostinger Email API to the Hostinger mail team.

Only call this when the user explicitly asks to send feedback, report a problem, or request a feature. Send the user's own words; never include tokens, passwords or email contents. The message is capped at 2000 characters.

A `429` (`ERR_FEEDBACK_RATE_LIMIT`) means feedback for this customer was submitted less than ten seconds ago; wait and retry.

```
hostinger-mail feedback submit <mailbox-resource-id> [flags]
```

### Examples

```
  hostinger-mail feedback submit my-mailbox-resource-id --message 'Sending a message with a 10 MB attachment fails with a timeout.' --score 3
```

### Options

```
  -h, --help             help for submit
      --message string   The user's feedback in their own words. Never include tokens, passwords or email contents.
      --score int        The user's rating of the Hostinger Email API: 1 (poor) to 10 (excellent).
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

* [hostinger-mail feedback](hostinger-mail_feedback.md)	 - Feedback commands

