## hostinger-mail folders list

List folders

### Synopsis

Retrieve a paginated list of folders in the managed mailbox.

```
hostinger-mail folders list <mailbox-resource-id> [flags]
```

### Examples

```
  hostinger-mail folders list my-mailbox-resource-id
```

### Options

```
      --all           Fetch every page and print all items
  -h, --help          help for list
      --page int      Page number (1-based). (default 1)
      --perpage int   Items per page (max 100). (default 25)
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

* [hostinger-mail folders](hostinger-mail_folders.md)	 - Folders commands

