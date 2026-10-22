## hostinger-mail folders create

Create folder

### Synopsis

Create a new folder in the managed mailbox.

```
hostinger-mail folders create <mailbox-resource-id> [flags]
```

### Examples

```
  hostinger-mail folders create my-mailbox-resource-id --name Projects
```

### Options

```
  -h, --help          help for create
      --name string   Folder name. Length 1-100 after trimming.
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

