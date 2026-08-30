---
title: "e0ipso/ddev-assistant-codex"
github_url: "https://github.com/e0ipso/ddev-assistant-codex"
description: "ddev add-on for setting up the Codex assistant."
user: "e0ipso"
repo: "ddev-assistant-codex"
repo_id: 1132378898
default_branch: "main"
tag_name: "v1.1.0"
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2026-01-11"
updated_at: "2026-06-29"
workflow_status: "disabled"
stars: 0
---

# ddev-assistant-codex

DDEV add-on for running the @openai/codex CLI inside DDEV web containers.

## Installation

```bash
ddev add-ons install ddev-assistant-codex
```

## What This Addon Does

- Installs @openai/codex into the container at `/usr/local/bin/codex`
- Makes `codex` available on `$PATH` for `ddev ssh` and `ddev exec`
- Mirrors host Codex config into a writable container runtime directory on start

## Quick Start

After installation, verify the setup:

```bash
ddev exec codex --version
```

## Configuration

### Codex Credentials And Runtime Config

The add-on mounts your host `~/.codex` directory into the web container as a
read-only seed at `~/.cred-seed/codex`. During every container startup, that
seed is mirrored into the writable runtime config directory at `~/.codex`.

Your host config is authoritative. Container-only changes under `~/.codex` are
removed on restart, then replaced with a fresh copy of the host seed. Configure
Codex features such as MCP servers on the host; this add-on mirrors the resulting
host `~/.codex` config into the DDEV container.

## Troubleshooting

### Codex CLI not found
```bash
ddev restart
ddev exec which codex
```

## Documentation

- [@openai/codex](https://code.codex.com/)
- [DDEV Addons](https://docs.ddev.com/en/stable/users/extend/creating-add-ons/)
