# Hostinger Mail API CLI

`hostinger-mail` is the command line interface for the [Hostinger Mail API](https://api.mail.hostinger.com).

> **This repository is generated.** Do not open pull requests here — every file
> is overwritten by the next generation run. The CLI is generated from the
> OpenAPI specification by
> [hostinger/public-api-generator](https://github.com/hostinger/public-api-generator)
> (see the `cli/` directory there). Report issues or contribute in that repository.

## Installation

### Homebrew (macOS & Linux)

```sh
brew install hostinger/tap/hostinger-mail
```

Upgrade with `brew upgrade hostinger-mail`. Shell tab-completion (bash/zsh/fish) is installed
automatically.

### Binary download

Download the binary for your platform from the
[releases page](https://github.com/hostinger/mail-api-cli/releases).

## Configuration

Create `$HOME/.hostinger-mail.yaml`:

```yaml
api_token: <your API token>
```

or set the `HOSTINGER_MAIL_API_TOKEN` environment variable. Generate a token at
[hPanel → API](https://hpanel.hostinger.com/emails).

## Usage

```
hostinger-mail <group> <verb> [args] [flags]
```

Examples:

```
hostinger-mail account current
hostinger-mail feedback submit <mailbox-resource-id>
hostinger-mail folders list <mailbox-resource-id>
hostinger-mail messages list <mailbox-resource-id> <folder>
```

Output format defaults to a table; use `--format json|table|tree` to change it.
Command documentation lives in [docs/](docs/), and `manifest.json` maps every
API operation to its command. Shell autocompletion: see
[AUTOCOMPLETE.md](AUTOCOMPLETE.md).
