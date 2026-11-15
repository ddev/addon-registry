---
title: "Performant-Labs/ddev-cli-relay"
github_url: "https://github.com/Performant-Labs/ddev-cli-relay"
description: "Send commands to the host, capture the response."
user: "Performant-Labs"
repo: "ddev-cli-relay"
repo_id: 1208079295
default_branch: "main"
tag_name: "v1.1.0"
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2026-04-11"
updated_at: "2026-11-14"
workflow_status: "unknown"
stars: 0
---

# ddev-cli-relay

[![tests](https://github.com/Performant-Labs/ddev-cli-relay/actions/workflows/tests.yml/badge.svg)](https://github.com/Performant-Labs/ddev-cli-relay/actions/workflows/tests.yml)
[![DDEV add-on](https://img.shields.io/badge/DDEV-add--on-blue)](https://addons.ddev.com)

A DDEV add-on that starts a lightweight HTTP relay server on the **host machine**, allowing code running inside a DDEV container to execute local CLI tools on the host.

## Usage

Once installed, the relay is available inside the container at `http://host.docker.internal:8765`.

Every `/execute` request must carry the relay's token. The token is created on
first start at `.ddev/cli-relay/token` (git-ignored, mode `600`) and is visible
inside the container at `/var/www/html/.ddev/cli-relay/token`.

**Check the relay is running:**
```bash
# From inside the container
curl http://host.docker.internal:8765/health
# {"status":"ok","port":8765,"auth":"token"}
```

**Execute a CLI tool from inside the container:**
```bash
curl -s -X POST http://host.docker.internal:8765/execute \
  -H "Authorization: Bearer $(cat /var/www/html/.ddev/cli-relay/token)" \
  -H 'Content-Type: application/json' \
  -d '{
    "binary":  "claude",
    "args":    ["--print"],
    "prompt":  "User: Hello",
    "timeout": 30
  }'
# {"stdout":"Hello! How can I help?","stderr":"","exit_code":0}
```

The prompt is passed to the process via **stdin** — never as a shell argument.

**Start and stop the relay manually** (without a full `ddev restart`):
```bash
ddev cli-relay:start
ddev cli-relay:stop
```

The relay starts automatically on `ddev start` and stops on `ddev stop`.

## Why

CLI tools like `claude`, `codex`, and `gemini-cli` use browser-based OAuth authentication — they must run on the host where auth already happened. PHP running inside a DDEV container has no direct way to exec host binaries. DDEV's `.ddev/commands/host/` scripts are triggered by the developer in a terminal, not programmatically from application code. This add-on bridges the gap with a minimal HTTP server that runs on the host and proxies execution requests from the container.

## Requirements

- DDEV >= v1.24.0
- Python 3.12 or newer, as `python3` on the host `PATH` (tested up to 3.14). macOS's `/usr/bin/python3` is 3.9, which is too old: install a newer one (`brew install python`) and make sure it comes first on `PATH`. The built-in Python is new enough on Ubuntu 24.04 and Debian 13 or later, but not on Ubuntu 22.04 (3.10) or Debian 12 (3.11)

## Installation

```bash
ddev add-on get Performant-Labs/ddev-cli-relay
ddev restart
```

## Configuration

| Environment variable | Default | Description |
|---|---|---|
| `AI_RELAY_PORT` | `8765` | Port the relay listens on |
| `AI_RELAY_BIND` | macOS: `127.0.0.1`; Linux: the `ddev_default` network gateway | Address the relay listens on |
| `AI_RELAY_ALLOWED_BINARIES` | *(none)* | Extra binaries to allow, comma-separated (names or absolute paths) |

Export these in your shell profile before running `ddev start`.

**Allowed binaries.** The relay only runs these CLI tools: `claude`, `codex`,
`gemini`, `gemini-cli`, `opencode`, `ollama`, `lms`, `llm`. A bare name is
resolved through the host's `PATH`; an absolute path is accepted only if it is
where an allowed name resolves, or if it is listed explicitly. To allow more,
either set `AI_RELAY_ALLOWED_BINARIES` or list them one per line in
`.ddev/cli-relay/allowed-binaries.txt` (comments with `#`), then run
`ddev cli-relay:stop && ddev cli-relay:start`.

**Rotating the token.** Delete `.ddev/cli-relay/token` and run `ddev restart`.

## API Reference

### `GET /health`

Returns `200` when the relay is running. Needs no token.

```json
{
  "status": "ok",
  "version": "1.1.0",
  "port": 8765,
  "auth": "token",
  "token_id": "3f2a9c41d0be",
  "features": ["files"],
  "binaries": ["claude", "codex", "llm"]
}
```

`token_id` is the first 12 hex characters of the token's SHA-256, so a client
can check it reached its own project's relay without the token being revealed.
`features` lists optional request fields the relay understands, and `binaries`
the allowed tools it found on the host.

### `POST /execute`

Executes an allowed binary on the host with the given arguments. The `prompt` value is written to the process's stdin.

Required headers: `Authorization: Bearer <token>` and `Content-Type: application/json`.
Requests carrying an `Origin` header (i.e. from a web browser) are refused.

**Request:**
```json
{
  "binary":  "claude",
  "args":    ["--flag", "--other-flag"],
  "prompt":  "text sent to stdin",
  "timeout": 120,
  "files":   [{ "name": "photo.png", "base64": "iVBORw0KGgo..." }]
}
```

| Field | Type | Required | Description |
|---|---|---|---|
| `binary` | string | yes | Allowed binary name, or its absolute path on the host |
| `args` | array of strings | no | CLI flags passed to the binary |
| `prompt` | string | no | Text written to stdin |
| `timeout` | integer | no | Seconds before the process is killed (default `120`, max `3600`) |
| `files` | array of objects | no | Attachments, each `{"name": "...", "base64": "..."}` (at most 20) |

**Attachments.** The relay writes `files` to a new temp directory readable only
by you, then replaces these placeholders in `args` and `prompt`:

| Placeholder | Replaced with |
|---|---|
| `{{file:0}}`, `{{file:1}}`, … | The path of that file (numbered from 0) |
| `{{files_dir}}` | The directory holding the files |

For example, `"args": ["exec", "-i", "{{file:0}}"]` passes the first file to
Codex. Only each name's base name is kept (other characters become `_`), and
the relay prefixes it with the file's position, so `photo.png` becomes
`1-photo.png`. The directory is deleted when the tool exits, including on a
timeout. A placeholder that refers to a file that was not sent is a `400`.

**Errors** (HTTP status, JSON body `{"error": "..."}`):

| Status | Meaning |
|---|---|
| `400` | Malformed request body, invalid `files`, or a placeholder without its file |
| `401` | Missing or wrong token |
| `403` | Browser request (has `Origin`), or binary not allowed |
| `413` | Body larger than 32 MiB (about 24 MiB of attachments) |
| `415` | Content-Type is not `application/json` |

**Response:**
```json
{
  "stdout":    "output from the process",
  "stderr":    "",
  "exit_code": 0
}
```

| Exit code | Meaning |
|---|---|
| `0` | Success |
| `124` | Process timed out |
| `127` | Binary not found |
| other | Process returned an error — check `stderr` |

## Security

The relay runs programs on your machine as you. It is locked down against the
two outside threats that matter for a development laptop:

- **Not on your network.** On macOS it listens on loopback only, so other
  machines on the same Wi-Fi cannot reach it. On Linux it listens on the
  address containers use to reach the host (a Docker bridge address). Peers on
  your LAN can't normally reach that address, but they can if they route that
  subnet through your machine. The token still stops them; add a firewall rule
  if that worries you.
- **Token required.** Each project gets its own random token. Only code that
  can read the project's files can call the relay.
- **No browsers.** Requests with an `Origin` header, or without
  `Content-Type: application/json`, are refused, so a web page you visit cannot
  call it.
- **Allowlisted binaries only.** The relay only runs the allowed CLI tools, at
  the paths it found when it started. It never runs a path the caller supplies,
  and it ignores relative `PATH` entries.

**What it does not protect against: anything that holds the token.** Treat
the token as equivalent to a shell on your machine. The caller chooses the
arguments and the prompt, and several allowed tools have flags that run
arbitrary commands. Claude Code's `--settings` hooks and `--mcp-config`, and
Codex's `--dangerously-bypass-approvals-and-sandbox`, are examples. Code in the
web container can read the token, and with DDEV's file sync it can also edit
`.ddev/` scripts that run on the host. So the relay is no security boundary
between your containers and your machine. Run only code you trust in a project
that has the relay installed.

**More than one project.** Every project uses port 8765 unless you set
`AI_RELAY_PORT`. If a second project's relay finds the port taken,
`ddev start` reports it and fails rather than letting Drupal reach the wrong
project's relay. Give each project that runs at the same time its own port.

**Upgrading from an earlier release.** Releases before this change listened on
`0.0.0.0` with no token. Upgrade **every** project that has the relay
installed, or stop the old relays; an old relay left running in another
project stays reachable from your network. To find any still running, run
`lsof -nP -iTCP:8765 -sTCP:LISTEN`. A `*:8765` address means an old relay.
Update with `ddev add-on get Performant-Labs/ddev-cli-relay && ddev restart`.

Why it is set up this way: [ADR 0001](https://github.com/Performant-Labs/ddev-cli-relay/blob/main/docs/adr/0001-relay-access-control.md).
**This add-on is intended for local development only.**

## Contributing

Pull requests are welcome. For significant changes please open an issue first.

Run the tests with `python3 -m unittest discover -s tests -v` (Python 3.12+, no dependencies).

## License

MIT

## Related

- [ai_provider_cli](https://www.drupal.org/project/ai_provider_cli) — Drupal AI provider that uses this relay to pipe prompts to local CLI tools from within a DDEV container.
