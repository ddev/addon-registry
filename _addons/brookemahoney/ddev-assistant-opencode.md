---
title: "brookemahoney/ddev-assistant-opencode"
github_url: "https://github.com/brookemahoney/ddev-assistant-opencode"
description: "ddev add-on for setting up the OpenCode v2 assistant."
user: "brookemahoney"
repo: "ddev-assistant-opencode"
repo_id: 1391324086
default_branch: "main"
tag_name: "1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-27"
updated_at: "2026-09-27"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/brookemahoney/ddev-assistant-opencode/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/brookemahoney/ddev-assistant-opencode/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/brookemahoney/ddev-assistant-opencode)](https://github.com/brookemahoney/ddev-assistant-opencode/commits)
[![release](https://img.shields.io/github/v/release/brookemahoney/ddev-assistant-opencode)](https://github.com/brookemahoney/ddev-assistant-opencode/releases/latest)

# DDEV Assistant Opencode

## About this fork

This is a fork of https://github.com/e0ipso/ddev-assistant-opencode that installs OpenCode V2 instead of V1.

## Overview

This add-on integrates Assistant Opencode into your [DDEV](https://ddev.com/) project's web container.

## Installation

```bash
ddev add-on get brookemahoney/ddev-assistant-opencode
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev exec opencode` | Run OpenCode commands inside the web container |


## Configuration

The add-on mounts your host OpenCode configuration into the web container:

| Host Path | Container Path | Purpose |
| --------- | -------------- | ------- |
| `~/.config/opencode` | `~/.config/opencode` | OpenCode configuration |
| `~/.cache/opencode` | `~/.cache/opencode` | OpenCode cache |
| `~/.local/share/opencode` | `~/.local/share/opencode` | OpenCode data (including auth) |

On first `ddev restart`, the add-on:
1. Installs OpenCode into the web container image at `/usr/local/bin/opencode`
2. Ensures all mounted directories are owned by the web user (not root)

## Credits

**Contributed and maintained by [@brookemahoney](https://github.com/brookemahoney)**
