## hostinger-mail send email

Send email

### Synopsis

Send a message from the managed mailbox. Saves a copy to INBOX.Sent.

JSON flags take inline JSON, @file.json, or - to read stdin. Example values:

  --attachments
    [
      {
        "cid": "logo",
        "content": "JVBERi0xLjQK...",
        "contentType": "application/pdf",
        "encoding": "base64",
        "filename": "invoice.pdf"
      }
    ]

  --forwardof
    {
      "folder": "INBOX",
      "uid": 123
    }

  --inreplyto
    {
      "folder": "INBOX",
      "uid": 123
    }

```
hostinger-mail send email <mailbox-resource-id> [flags]
```

### Examples

```
  hostinger-mail send email my-mailbox-resource-id
```

### Options

```
      --attachments string   Files to attach. Inline images set cid; regular attachments omit it. (JSON, @file or -; shape in the help above)
      --bcc strings          Blind-carbon-copy recipient email addresses. Not visible to other recipients.
      --cc strings           Carbon-copy recipient email addresses.
      --displayname string   Sender display name shown in the From header alongside the mailbox address.
      --forwardof string     Source message this forwards, as an object {uid, folder} — not the Message-ID header string. Copies its Message-Id/References into In-Reply-To/References and flags it $forwarded. Mutually exclusive with inReplyTo. (JSON, @file or -; shape in the help above)
  -h, --help                 help for email
      --html string          HTML body. Optional; if both text and html are omitted the message is sent without a body. Inline images are referenced via cid: URLs matching attachment cid values.
      --inreplyto string     Source message this is a reply to, as an object {uid, folder} — not the Message-ID header string. Copies its Message-Id/References into In-Reply-To/References and flags it \Answered. Mutually exclusive with forwardOf. (JSON, @file or -; shape in the help above)
      --subject string       Message subject line.
      --text string          Plain-text body. Optional; if both text and html are omitted the message is sent without a body.
      --to strings           Primary recipient email addresses.
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

* [hostinger-mail send](hostinger-mail_send.md)	 - Send commands

