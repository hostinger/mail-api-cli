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

Upgrade with `brew upgrade hostinger-mail`. Shell tab-completion (bash/zsh/fish) and man pages
(`man hostinger-mail`) are installed automatically.

### Binary download

Download the binary for your platform from the
[releases page](https://github.com/hostinger/mail-api-cli/releases). Then enable tab-completion
with `hostinger-mail completion install` (see [AUTOCOMPLETE.md](AUTOCOMPLETE.md)).

## Configuration

Run `hostinger-mail config set api_token` (it asks for the token without echoing it) or set the
`HOSTINGER_MAIL_API_TOKEN` environment variable. Generate a token at [hPanel → API](https://hpanel.hostinger.com/emails).

`hostinger-mail auth status` shows which credentials are in use. Settings live in
`$HOME/.hostinger-mail.yaml` (`hostinger-mail config list` shows them):

```yaml
api_token: <your API token>
format: json        # default output format
no_input: false     # true: never prompt
```

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

Every command's `--help` lists its required flags separately and shows examples. Flags that take
JSON accept inline JSON, `@file.json`, or `-` for stdin; the help shows the expected shape.

### Output

- Tables on a terminal, JSON when the output is piped or redirected. Choose explicitly with
  `-o/--format table|json|yaml|tree`, or set a default with `hostinger-mail config set format json`.
- `--jq '<expr>'` filters the JSON response (strings print raw), e.g. `--jq '.[].id'`.
- List commands with pages take `--all` to fetch every page.

### Prompts and safety

- In a terminal, missing arguments and required flags are asked for, with choices listed where
  they are known. `--no-input` (or `no_input: true`) turns prompting off.
- Commands that delete, stop, restore or purchase ask for confirmation. Pass `--yes` to skip it;
  non-interactive runs (scripts, CI) must pass `--yes`.
- `--dry-run` prints the request as a `curl` command instead of sending it.
- Commands that start an asynchronous action take `--wait` to block until it finishes.
- `--verbose` prints each HTTP request and its status to stderr.

### Exit codes

| Code | Meaning                                       |
|------|-----------------------------------------------|
| 0    | Success                                       |
| 1    | API, network or other error; aborted prompt   |
| 2    | Invalid arguments, flags or input             |
| 3    | No credentials, or authentication failed (HTTP 401/403) |
| 4    | Not found (HTTP 404)                          |
| 5    | Validation failed (HTTP 422)                  |

Command documentation lives in [docs/](docs/) and in the man pages, and `manifest.json` maps every
API operation to its command. Shell autocompletion: see [AUTOCOMPLETE.md](AUTOCOMPLETE.md).
