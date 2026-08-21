---
title: "e0ipso/ddev-assistant-copilot"
github_url: "https://github.com/e0ipso/ddev-assistant-copilot"
description: "ddev add-on for setting up the Copilot assistant."
user: "e0ipso"
repo: "ddev-assistant-copilot"
repo_id: 1283680554
default_branch: "main"
tag_name: "v1.1.1"
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2026-06-29"
updated_at: "2026-07-28"
workflow_status: "success"
stars: 2
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/e0ipso/ddev-assistant-copilot/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/e0ipso/ddev-assistant-copilot/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/e0ipso/ddev-assistant-copilot)](https://github.com/e0ipso/ddev-assistant-copilot/commits)
[![release](https://img.shields.io/github/v/release/e0ipso/ddev-assistant-copilot)](https://github.com/e0ipso/ddev-assistant-copilot/releases/latest)

# DDEV GitHub Copilot CLI

## Overview

This DDEV add-on installs [GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/set-up/install-copilot-cli) inside the DDEV web container and automatically shares your host GitHub CLI and Copilot configuration — including authentication and Copilot settings — with no additional setup required.

Once installed, running `copilot` inside `ddev ssh` or `ddev exec` uses a writable copy of your host configuration and a token derived from your host `gh` authentication.

## Requirements

- DDEV >= v1.24.0
- GitHub CLI (`gh`) installed and authenticated **on the host** — the host is where the token is resolved, so `gh auth status` must succeed there

## Installation

```bash
ddev add-on get e0ipso/ddev-assistant-copilot
ddev restart
```

After installation, commit the `.ddev` directory to version control.

## What it does

- **Installs GitHub CLI** into the container image via the official apt repository, on `$PATH` for every shell
- **Installs GitHub Copilot CLI** on every start via `npm install -g @github/copilot` into `~/.local/bin` (warns and continues if npm install fails)
- **Seeds host configuration** on start: your host `~/.config/gh/` and `~/.copilot/` trees are mounted read-only under `~/.cred-seed/`, then mirrored into the writable container directories on every restart:
  - `~/.config/gh/` — GitHub CLI configuration and authentication (e.g. `hosts.yml`)
  - `~/.copilot/` — Copilot CLI configuration (e.g. `config.json`, hooks)
- **Forwards authentication from the host keychain** — the token is resolved on the host, where the OS keychain lives, and handed to the container on a read-only mount. See [How authentication forwarding works](#how-authentication-forwarding-works).
- **Exports `COPILOT_GITHUB_TOKEN` and `GH_TOKEN`** on shell startup, while preserving any token explicitly injected into the container environment. Interactive shells (`.bashrc`, `.profile`) and non-interactive shells (`/etc/bash.env` via `BASH_ENV`) resolve the token dynamically instead of storing the literal value.
- **Available everywhere** — `gh` and `copilot` are on `$PATH` for both interactive shells (`ddev ssh`) and non-interactive commands (`ddev exec`)

## How authentication forwarding works

GitHub CLI stores your token in the **OS keychain** by default — macOS Keychain,
GNOME Keyring (Secret Service) on Linux, Windows Credential Manager. It writes a
plaintext `oauth_token:` into `~/.config/gh/hosts.yml` only as a silent fallback
when no keychain is reachable, such as on a headless server or WSL.

A container cannot reach the host keychain. So copying `~/.config/gh/` into the
container and running `gh auth token` there only ever works on hosts that took
the plaintext fallback. That is why forwarding used to succeed on headless Linux
and fail on macOS and Linux desktops, and why selecting an account with
`DDEV_ASSISTANT_COPILOT_GH_USER` could not help: the token was not in the file at
all.

This add-on resolves the token **on the host** instead:

1. A `pre-start` hook runs `gh auth token` on the host, honouring the configured
   hostname and user, and writes the result to
   `~/.ddev-assistant-copilot/<project>/token` with mode `600`.
2. That directory is bind-mounted read-only into the container at
   `~/.cred-seed/auth/`.
3. Shell startup resolves a token in this order:
   1. `COPILOT_GITHUB_TOKEN` already present in the container environment
   2. the host-resolved token from the seed mount
   3. in-container `gh auth token` — still correct for hosts that legitimately
      use plaintext storage

The token is refreshed on every `ddev start`, so a rotated or revoked token is
picked up by a restart.

### Where the token does and does not go

The token is deliberately kept out of DDEV's environment plumbing:

- **Not** in `web_environment`, so it never appears in `docker inspect`
- **Not** in `.ddev/.env.web` or any file inside the project directory, so it
  cannot be committed
- **Not** written literally into `/etc/bash.env`, `.bashrc`, or `.profile`
- Stored only in a mode-`600` file under `$HOME`, outside the project

Note that inside the web container, PHP runs as the same user as `copilot`, so
anything reachable by Copilot CLI is reachable by your application code. The
practical benefit over an environment variable is that a file is not dumped by
`phpinfo()`, Xdebug stack traces, or a Symfony/Whoops error page. If that
residual exposure matters for your project, use ephemeral mode below.

### Ephemeral tokens

To keep no token at rest at all, set this in
`.ddev/config.assistant-copilot.yaml`:

```yaml
web_environment:
  - DDEV_ASSISTANT_COPILOT_EPHEMERAL_TOKEN=true
```

`ddev start` then stages no token. Instead, run Copilot through the host command:

```bash
ddev copilot
```

This resolves a fresh token from the host keychain for that invocation and
removes it when the session exits. The trade-off is that plain `copilot` inside
`ddev ssh` is left unauthenticated and will prompt you to log in.

`ddev copilot` works in the default mode too, where it simply refreshes the token
before starting the session.

## Multi-account configuration

The add-on ships with this non-secret default:

```yaml
web_environment:
  - DDEV_ASSISTANT_COPILOT_GH_HOSTNAME=github.com
  # - DDEV_ASSISTANT_COPILOT_GH_USER=your-github-username
```

Use this when you have multiple GitHub accounts, such as a personal account and
a work account, authenticated for the same host. Uncomment and set
`DDEV_ASSISTANT_COPILOT_GH_USER` in your installed
`.ddev/config.assistant-copilot.yaml` to tell the add-on which account should
provide `COPILOT_GITHUB_TOKEN`.

If either variable is empty or removed, the add-on does not pass that option to
`gh auth token`, allowing GitHub CLI's default account/host selection to apply.

## Usage

```bash
# Run Copilot CLI with a token resolved fresh from the host keychain
ddev copilot

# Open a shell with Copilot CLI available
ddev ssh
copilot

# Run Copilot CLI non-interactively
ddev exec copilot --version

# GitHub CLI is also available, and authenticated
ddev exec gh auth status
```

## Why not manual setup?

You can install Copilot CLI and `gh` inside a DDEV container yourself. This add-on automates the parts that are easy to get wrong or forget:

| | This add-on | Manual setup |
|---|---|---|
| **GitHub CLI** | Installed in the image layer via official apt repo; on `$PATH` for every shell | Must install and re-install after image rebuilds |
| **Copilot CLI** | `npm install -g @github/copilot` on every start into `~/.local/bin` | Must run npm install manually; easy to lose on restart |
| **Config approach** | Seeds writable container `~/.config/gh/` and `~/.copilot/` from your host config on restart — zero setup if you already use `gh` and Copilot on the host | Must copy or symlink config by hand; stale container files persist |
| **Authentication** | Token resolved on the host, so it works whether or not `gh` used the OS keychain; exported as `COPILOT_GITHUB_TOKEN` and `GH_TOKEN` without writing literal tokens into shell startup files, `web_environment`, or the project directory | Copying `~/.config/gh` only works on hosts without a keychain; otherwise you must export a token manually in every shell type |
| **Non-interactive shells** | `BASH_ENV=/etc/bash.env` ensures `ddev exec` gets PATH and token | `ddev exec` often misses PATH and env vars |
| **Mount safety** | Pre-start hook ensures host config directories exist before Docker bind-mounts them | Bind-mount fails silently or blocks start if dirs are missing |
| **Tests / CI** | BATS integration tests, GitHub Actions CI matrix (DDEV stable + HEAD), daily scheduled runs | No automated verification |

This add-on does one thing: install GitHub Copilot CLI into your DDEV container and share your existing host configuration. Nothing else.

## Credits

**Contributed and maintained by [@e0ipso](https://github.com/e0ipso)**
