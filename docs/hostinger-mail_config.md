## hostinger-mail config

Read and write CLI settings

```
hostinger-mail config [flags]
```

### Options

```
  -h, --help   help for config
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
* [hostinger-mail config get](hostinger-mail_config_get.md)	 - Print the effective value of a setting
* [hostinger-mail config list](hostinger-mail_config_list.md)	 - List the settings and their effective values
* [hostinger-mail config path](hostinger-mail_config_path.md)	 - Print the config file path
* [hostinger-mail config set](hostinger-mail_config_set.md)	 - Set a setting in the config file
* [hostinger-mail config unset](hostinger-mail_config_unset.md)	 - Remove a setting from the config file

