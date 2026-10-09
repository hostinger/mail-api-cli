## hostinger-mail feedback submit

Submit feedback

### Synopsis

Report a problem or suggestion about this API or the MCP server to the Hostinger mail team.

Report when a call returned 4xx/5xx or unexpected data, was too slow, when documentation was missing or unclear, or when a capability you needed does not exist. Mention the failing operation and the status code you received so the team can find the request. Never include tokens, passwords or email contents: the message is scrubbed of secrets and capped at 2000 characters. Send one report per distinct issue.

A `429` (`ERR_FEEDBACK_RATE_LIMIT`) means feedback for this customer was submitted less than ten seconds ago; wait and retry.

```
hostinger-mail feedback submit <mailbox-resource-id> [flags]
```

### Options

```
  -h, --help             help for submit
      --message string   What happened and what was expected, including the operation and status code involved. Never include tokens, passwords or email contents.
      --score int        How well the API served the task: 1 (poor) to 10 (excellent).
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail feedback](hostinger-mail_feedback.md)	 - Feedback commands

