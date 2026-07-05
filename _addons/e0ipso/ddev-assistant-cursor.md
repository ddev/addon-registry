---
title: e0ipso/ddev-assistant-cursor
github_url: https://github.com/e0ipso/ddev-assistant-cursor
description: "ddev add-on for setting up the Cursor assistant."
user: e0ipso
repo: ddev-assistant-cursor
repo_id: 1283844175
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-06-29
updated_at: 2026-06-29
workflow_status: success
stars: 1
---

# ddev-assistant-cursor

## Overview

This DDEV add-on installs the Cursor CLI (`cursor-agent`) into the DDEV web
container and shares your host Cursor config and credentials with the container.

Host config is mounted read-only into seed directories, then mirrored into
writable runtime directories on every `ddev start`.

## Requirements

- DDEV >= 1.24.10
- Cursor CLI installed and authenticated on the host

## Installation

```bash
ddev add-ons install ddev-assistant-cursor
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## What It Does

- Installs `cursor-agent` into the container at `/usr/local/bin/cursor-agent`
- Adds `agent` as an alias for `cursor-agent`
- Makes `cursor-agent` available on `$PATH` for `ddev ssh` and `ddev exec`
- Mirrors host Cursor config from `~/.cursor` into writable container `~/.cursor`
- Mirrors host Cursor credentials from `~/.config/cursor` into writable container `~/.config/cursor`
- Replaces container runtime config from the host seed on every start

## Usage

After installation, verify the setup:

```bash
ddev exec cursor-agent --version
ddev exec which cursor-agent
```

Run Cursor from inside the project web container:

```bash
ddev ssh
cursor-agent
```

Or run it directly:

```bash
ddev exec cursor-agent --version
```

## Configuration

### Cursor Config And Credentials

The add-on mounts your host `~/.cursor` directory into the web container as a
read-only seed at `~/.cred-seed/cursor`. It also mounts your host
`~/.config/cursor` directory read-only at `~/.cred-seed/cursor-auth`.

During every container startup, those seeds are mirrored into writable runtime
directories at `~/.cursor` and `~/.config/cursor`.

Your host config is authoritative. Container-only changes under those runtime
paths are removed on restart, then replaced with a fresh copy of the host seeds.
Configure Cursor on the host; this add-on mirrors the resulting config and
credentials into the DDEV container.

## Troubleshooting

### Cursor CLI Not Found

```bash
ddev restart
ddev exec which cursor-agent
```

## Credits

**Contributed and maintained by [@e0ipso](https://github.com/e0ipso)**
