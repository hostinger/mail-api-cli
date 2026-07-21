## hostinger-mail send email

Send email

### Synopsis

Send a message from the managed mailbox. Saves a copy to INBOX.Sent.

```
hostinger-mail send email <mailbox-resource-id> [flags]
```

### Options

```
      --attachments string    (JSON)
      --bcc strings          
      --cc strings           
      --displayname string   
      --forwardof string     Source message this forwards. Copies its Message-Id/References into In-Reply-To/References and flags it $forwarded. Mutually exclusive with inReplyTo. (JSON)
  -h, --help                 help for email
      --html string          
      --inreplyto string     Source message this is a reply to. Copies its Message-Id/References into In-Reply-To/References and flags it \Answered. Mutually exclusive with forwardOf. (JSON)
      --subject string       
      --text string          
      --to strings           
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail send](hostinger-mail_send.md)	 - Send commands

