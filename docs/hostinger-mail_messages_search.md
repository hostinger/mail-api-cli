## hostinger-mail messages search

Search messages

### Synopsis

Search messages in a folder. Filters in body; pagination and sort via query (`page`, `perPage`, `sort`).

```
hostinger-mail messages search <mailbox-resource-id> <folder> [flags]
```

### Options

```
      --before string    Only messages received before this date (YYYY-MM-DD).
      --body string      Case-insensitive substring match on the message body only (headers excluded). OR-combined with subject/from/to/cc.
      --cc string        Case-insensitive substring match on the Cc header. OR-combined with subject/from/to/body.
      --flags strings    Only messages carrying all of these IMAP flags, e.g. \Seen, \Flagged, \Answered, $forwarded.
      --from string      Case-insensitive substring match on the From header. OR-combined with subject/to/cc/body.
      --header string    Match a specific header as Name:value, e.g. X-Custom-Header:value. Value match is a substring.
  -h, --help             help for search
      --larger int       Only messages larger than this size in bytes.
      --page int         Page number (1-based). (default 1)
      --perpage int      Items per page (max 100). (default 25)
      --since string     Only messages received on or after this date (YYYY-MM-DD).
      --smaller int      Only messages smaller than this size in bytes.
      --sort -           Sort field with optional - prefix for descending. Allowed: uid, date, size. (default "-uid")
      --subject string   Case-insensitive substring match on the Subject header. OR-combined with from/to/cc/body.
      --text string      Case-insensitive substring match across headers and body.
      --to string        Case-insensitive substring match on the To header. OR-combined with subject/from/cc/body.
      --uid string       IMAP UID set: single UID, range (1:100), open range (100:*), or comma-separated list.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

