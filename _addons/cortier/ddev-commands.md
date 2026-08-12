---
title: cortier/ddev-commands
github_url: https://github.com/cortier/ddev-commands
description: "Shared DDEV host commands for Cortier projects"
user: cortier
repo: ddev-commands
repo_id: 1323461082
default_branch: main
tag_name: v1.2.0
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-08-04
updated_at: 2026-08-11
workflow_status: success
stars: 0
---

# Overview

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/cortier/ddev-commands/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/cortier/ddev-commands/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/cortier/ddev-commands)](https://github.com/cortier/ddev-commands/commits)

This DDEV add-on installs shared host commands used by Cortier projects.

# Commands

- `ddev api`: Connect to or run a command in an API surface.
- `ddev app`: Connect to or run a command in an app surface.
- `ddev admin`: Connect to or run a command in an admin surface.
- `ddev shop`: Connect to or run a command in a shop surface.
- `ddev launch`: Open the current project's local URL.
- `ddev name`: Print the DDEV project name without its final suffix.
- `ddev path`: Print the configured absolute path for a Cortier repository.
- `ddev surface`: Manage and use API, app, admin, and shop surface connections.
- `ddev url`: Print URLs from `.ddev/config.host.local.yaml`.

# Requirements

The target project must contain `.ddev/config.host.local.yaml`. The `surface`
command also requires `jq` when it reads DDEV metadata.

Configure repository paths in global Git configuration. For example:

```bash
git config --global cortier.paths.rules /absolute/path/to/rules
ddev path rules
```

Add `--json-output` to return the repository name and resolved path as JSON.
Configured paths must identify the root of a standalone Git checkout, not a
nested directory or submodule, and its `origin` must identify the matching
`cortier/<repository>` repository.

API projects can connect to one app, one admin, and one shop. Each frontend can
connect to one API; frontend-to-frontend connections are rejected. Connections
are validated by project-name suffix, DDEV type, and reciprocal connection file.

# Installation

Install the latest release:

```bash
ddev add-on get cortier/ddev-commands
```

Install the add-on from a local checkout:

```bash
ddev add-on get /path/to/ddev-commands
```

Run the installation command again to update the installed commands.
