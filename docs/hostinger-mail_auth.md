## hostinger-mail auth

Sign in and check which credentials are used

```
hostinger-mail auth [flags]
```

### Options

```
  -h, --help   help for auth
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

* [hostinger-mail](hostinger-mail.md)	 - Hostinger Mail API Command Line Interface
* [hostinger-mail auth status](hostinger-mail_auth_status.md)	 - Show which credentials the CLI uses

