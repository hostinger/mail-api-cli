## hostinger-mail messages search

Search messages

### Synopsis

Search messages in a folder. Filters in body; pagination and sort via query (`page`, `perPage`, `sort`).

```
hostinger-mail messages search <mailbox-resource-id> <folder> [flags]
```

### Options

```
      --before string    
      --body string      
      --cc string        
      --flags strings    
      --from string      
      --header string    
  -h, --help             help for search
      --larger int       
      --page int          (default 1)
      --perpage int       (default 25)
      --since string     
      --smaller int      
      --sort string       (default "-uid")
      --subject string   
      --text string      
      --to string        
      --uid string       
```

### Options inherited from parent commands

```
      --config string   Config file (default is $HOME/.hostinger-mail.yaml)
      --format string   Output format type (json|table|tree), default: table
```

### SEE ALSO

* [hostinger-mail messages](hostinger-mail_messages.md)	 - Messages commands

