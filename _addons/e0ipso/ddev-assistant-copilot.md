---
title: e0ipso/ddev-assistant-copilot
github_url: https://github.com/e0ipso/ddev-assistant-copilot
description: "ddev add-on for setting up the Copilot assistant."
user: e0ipso
repo: ddev-assistant-copilot
repo_id: 1283680554
default_branch: main
tag_name: v1.0.1
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-06-29
updated_at: 2026-06-29
workflow_status: success
stars: 1
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
- GitHub CLI (`gh`) authenticated on the host (for configuration sharing and `COPILOT_GITHUB_TOKEN`)

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
- **Exports `COPILOT_GITHUB_TOKEN`** on start from `gh auth token`, wired into interactive shells (`.bashrc`, `.profile`) and non-interactive shells (`/etc/bash.env` via `BASH_ENV`)
- **Available everywhere** — `gh` and `copilot` are on `$PATH` for both interactive shells (`ddev ssh`) and non-interactive commands (`ddev exec`)

## Usage

```bash
# Open a shell with Copilot CLI available
ddev ssh
copilot

# Run Copilot CLI non-interactively
ddev exec copilot --version

# GitHub CLI is also available
ddev exec gh auth status
```

## Why not manual setup?

You can install Copilot CLI and `gh` inside a DDEV container yourself. This add-on automates the parts that are easy to get wrong or forget:

| | This add-on | Manual setup |
|---|---|---|
| **GitHub CLI** | Installed in the image layer via official apt repo; on `$PATH` for every shell | Must install and re-install after image rebuilds |
| **Copilot CLI** | `npm install -g @github/copilot` on every start into `~/.local/bin` | Must run npm install manually; easy to lose on restart |
| **Config approach** | Seeds writable container `~/.config/gh/` and `~/.copilot/` from your host config on restart — zero setup if you already use `gh` and Copilot on the host | Must copy or symlink config by hand; stale container files persist |
| **Authentication** | `COPILOT_GITHUB_TOKEN` exported automatically from `gh auth token` | Must export token manually in every shell type |
| **Non-interactive shells** | `BASH_ENV=/etc/bash.env` ensures `ddev exec` gets PATH and token | `ddev exec` often misses PATH and env vars |
| **Mount safety** | Pre-start hook ensures host config directories exist before Docker bind-mounts them | Bind-mount fails silently or blocks start if dirs are missing |
| **Tests / CI** | BATS integration tests, GitHub Actions CI matrix (DDEV stable + HEAD), daily scheduled runs | No automated verification |

This add-on does one thing: install GitHub Copilot CLI into your DDEV container and share your existing host configuration. Nothing else.

## Credits

**Contributed and maintained by [@e0ipso](https://github.com/e0ipso)**
