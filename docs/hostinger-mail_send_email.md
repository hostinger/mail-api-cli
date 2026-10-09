## hostinger-mail send email

Send email

### Synopsis

Send a message from the managed mailbox. Saves a copy to INBOX.Sent.

```
hostinger-mail send email <mailbox-resource-id> [flags]
```

### Options

```
      --attachments string   Files to attach. Inline images set cid; regular attachments omit it. (JSON)
      --bcc strings          Blind-carbon-copy recipient email addresses. Not visible to other recipients.
      --cc strings           Carbon-copy recipient email addresses.
      --displayname string   Sender display name shown in the From header alongside the mailbox address.
      --forwardof string     Source message this forwards. Copies its Message-Id/References into In-Reply-To/References and flags it $forwarded. Mutually exclusive with inReplyTo. (JSON)
  -h, --help                 help for email
      --html string          HTML body. Optional; if both text and html are omitted the message is sent without a body. Inline images are referenced via cid: URLs matching attachment cid values.
      --inreplyto string     Source message this is a reply to. Copies its Message-Id/References into In-Reply-To/References and flags it \Answered. Mutually exclusive with forwardOf. (JSON)
      --subject string       Message subject line.
      --text string          Plain-text body. Optional; if both text and html are omitted the message is sent without a body.
      --to strings           Primary recipient email addresses.
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail send](hostinger-mail_send.md)	 - Send commands

