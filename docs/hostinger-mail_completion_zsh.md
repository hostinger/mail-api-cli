## hostinger-mail completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(hostinger-mail completion zsh)

To load completions for every new session, execute once:

#### Linux:

	hostinger-mail completion zsh > "${fpath[1]}/_hostinger-mail"

#### macOS:

	hostinger-mail completion zsh > $(brew --prefix)/share/zsh/site-functions/_hostinger-mail

You will need to start a new shell for this setup to take effect.


```
hostinger-mail completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
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

* [hostinger-mail completion](hostinger-mail_completion.md)	 - Generate the autocompletion script for the specified shell

