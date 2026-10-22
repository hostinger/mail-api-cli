## hostinger-mail folders delete

Delete folder

### Synopsis

Delete a folder and all of its subfolders.

```
hostinger-mail folders delete <mailbox-resource-id> <folder> [flags]
```

### Examples

```
  hostinger-mail folders delete my-mailbox-resource-id my-folder
```

### Options

```
  -h, --help   help for delete
  -y, --yes    Skip the confirmation prompt
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

