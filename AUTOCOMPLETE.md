# Enabling shell auto-completion

`hostinger-mail` has auto-complete support. This makes it easier to use the CLI and improves user experience by completing command
names by clicking TAB key. For example if you type `hostinger-mail account cur<TAB>` with auto-completion enabled, shell will automatically append
rest of the command: `hostinger-mail account current`.

Auto-completion can be generated for multiple shells. The currently supported shells are:
- Bash
- Zsh
- fish
- PowerShell

After adding shell auto-completion, remember to refresh your shell profile by logging out from the shell and log back in.

## Bash
### Linux:
```bash
hostinger-mail completion bash > /etc/bash_completion.d/hostinger-mail
````
### macOS:
```bash
hostinger-mail completion bash > /usr/local/etc/bash_completion.d/hostinger-mail
```

## Zsh
If shell completion is not already enabled in your environment, you will need to enable it.  You can execute the following once:
```bash
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

To load completions for each session, execute once:
```bash
hostinger-mail completion zsh > "${fpath[1]}/_hostinger-mail"
```
You will need to start a new shell for this setup to take effect.

## Fish
```bash
hostinger-mail completion fish | source
```
To load completions for each session, execute once:
```bash
hostinger-mail completion fish > ~/.config/fish/completions/hostinger-mail.fish
```

## PowerShell
```powershell
hostinger-mail completion powershell | Out-String | Invoke-Expression
```
To load completions for every new session, run:
```powershell
hostinger-mail completion powershell > hostinger-mail.ps1
```
and source this file from your PowerShell profile.
