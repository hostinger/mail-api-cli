## hostinger-mail feedback submit

Submit feedback

### Synopsis

Send the user's feedback about the Hostinger Email API to the Hostinger mail team.

Only call this when the user explicitly asks to send feedback, report a problem, or request a feature. Send the user's own words; never include tokens, passwords or email contents. The message is capped at 2000 characters.

A `429` (`ERR_FEEDBACK_RATE_LIMIT`) means feedback for this customer was submitted less than ten seconds ago; wait and retry.

```
hostinger-mail feedback submit <mailbox-resource-id> [flags]
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
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail feedback](hostinger-mail_feedback.md)	 - Feedback commands

