---
title: Gonzalo2683/ddev-claudecode-native
github_url: https://github.com/Gonzalo2683/ddev-claudecode-native
description: "DDEV addon for Claude Code integration using a self-contained native binary. No npm   dependencies."
user: Gonzalo2683
repo: ddev-claudecode-native
repo_id: 1119117727
default_branch: master
tag_name: v1.0.6
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2025-12-18
updated_at: 2025-12-21
workflow_status: unknown
stars: 1
---

# ddev-claudecode-native

DDEV add-on that installs Claude Code using the **native binary** (no npm/Node.js dependency).

## Credits

Based on [FreelyGive/ddev-claude-code](https://github.com/FreelyGive/ddev-claude-code), for DDEV users. This version replaces npm with the native binary installer for a smaller footprint and faster startup.

## Installation

```bash
ddev add-on get Gonzalo2683/ddev-claudecode-native
ddev restart
```

## Usage

```bash
# Start Claude Code
ddev claude

# Get help
ddev claude --help

# Update Claude Code
ddev claude update
```

On first run, you will be prompted to authenticate. Your configuration is persisted in `.ddev/claude-code/` across restarts.

## Version Control

By default, the `latest` version is installed. To change it, edit `.ddev/web-build/Dockerfile.claude-code-native`:

```dockerfile
# Options: stable, latest, or specific version (e.g., 1.0.58)
ARG CLAUDE_VERSION=stable
```

Then rebuild:

```bash
ddev debug rebuild
```

## GitLab Integration

```bash
ddev glab auth login
```

Configuration is persisted in `.ddev/.glab-cli/`.

## Benefits

- **No Node.js dependency** - Self-contained binary
- **Faster startup** - No npm overhead
- **Smaller footprint** - Single executable
- **Version control** - Pin to stable, latest, or specific version
