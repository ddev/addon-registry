---
title: "mxr576/ddev-pi-claude-code"
github_url: "https://github.com/mxr576/ddev-pi-claude-code"
description: "DDEV add-on to use your Claude subscription in ddev-pi via pi-claude-code."
user: "mxr576"
repo: "ddev-pi-claude-code"
repo_id: 1456396475
default_branch: "main"
tag_name: "1.0.0-alpha1"
ddev_version_constraint: ">= v1.25.2"
dependencies: ["mxr576/ddev-pi"]
type: "contrib"
created_at: "2026-11-13"
updated_at: "2026-11-13"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/mxr576/ddev-pi-claude-code/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/mxr576/ddev-pi-claude-code/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/mxr576/ddev-pi-claude-code)](https://github.com/mxr576/ddev-pi-claude-code/commits)
[![release](https://img.shields.io/github/v/release/mxr576/ddev-pi-claude-code)](https://github.com/mxr576/ddev-pi-claude-code/releases/latest)

# DDEV Pi Claude Code

## Overview

This DDEV add-on integrates [pi-claude-code](https://github.com/earendil-works/pi-claude-code) into your DDEV project's [ddev-pi](https://github.com/mxr576/ddev-pi) environment, allowing you to use your Claude subscription directly within the Pi Coding Agent.

## Dependencies

This add-on extends [mxr576/ddev-pi](https://github.com/mxr576/ddev-pi). `ddev-pi` will be installed automatically if not already present.

## Installation

```bash
ddev add-on get mxr576/ddev-pi-claude-code
ddev restart && ddev start --profiles=pi
```

After installation, make sure to commit the `.ddev` directory to version control.

## Authentication

Authenticate the Claude Code CLI inside the `pi` service:

```bash
ddev pi-claude-login
```

Follow the browser prompt to log in with your Anthropic/Claude account. Credentials are persisted across container restarts in `.ddev/pi/claude`.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev pi-claude-login` | Authenticate Claude Code CLI in the `pi` container |
| `ddev pi` | Launch interactive Pi Coding Agent session with `pi-claude-code` enabled |

## Configuration & Version Pinning

You can specify or pin the `@anthropic-ai/claude-code` version using the `PI_CLAUDE_CODE_VERSION` environment variable in `.ddev/.env.pi`:

```bash
ddev dotenv set .ddev/.env.pi --pi-claude-code-version="2.1.322"
ddev restart && ddev start --profiles=pi
```

All customization options:

| Variable | Flag | Default | Description |
| -------- | ---- | ------- | ----------- |
| `PI_CLAUDE_CODE_VERSION` | `--pi-claude-code-version` | `latest` | Version of `@anthropic-ai/claude-code` npm package to install |

## Credits

**Contributed and maintained by [@mxr576](https://github.com/mxr576)**
