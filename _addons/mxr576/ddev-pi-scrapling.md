---
title: mxr576/ddev-pi-scrapling
github_url: https://github.com/mxr576/ddev-pi-scrapling
description: "Adds dynamic web-fetching capability to PI Coding Agent via Scrapling"
user: mxr576
repo: ddev-pi-scrapling
repo_id: 1271596147
default_branch: main
tag_name: 1.0.0-alpha1
ddev_version_constraint: ">= v1.24.10"
dependencies: ["mxr576/ddev-pi"]
type: contrib
created_at: 2026-06-16
updated_at: 2026-06-17
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/mxr576/ddev-pi-scrapling/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/mxr576/ddev-pi-scrapling/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/mxr576/ddev-pi-scrapling)](https://github.com/mxr576/ddev-pi-scrapling/commits)
[![release](https://img.shields.io/github/v/release/mxr576/ddev-pi-scrapling)](https://github.com/mxr576/ddev-pi-scrapling/releases/latest)

# DDEV Pi Scrapling

## Overview

This add-on enhances the PI Coding Agent inside the [`ddev-pi`](https://github.com/mxr576/ddev-pi) DDEV addon with support for dynamic web-fetch capabilities.

It provisions the Python `scrapling` package alongside headless Playwright browsers and registers the official Scrapling Skill with the PI coding agent.

## Installation

```bash
ddev add-on get mxr576/ddev-pi-scrapling
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

The Scrapling capabilities are automatically available inside the `pi` container. You can interact with the agent or view registered skills:

| Command | Description |
| ------- | ----------- |
| `ddev exec -s pi pi list` | View registered skills, including `scrapling-official` |
| `ddev logs -s pi` | Check PI agent logs |

## Customization

To customize the version or git ref of the Scrapling Skill:

1. Update the `PI_SCRAPLING_VERSION` variable in your `.ddev/.env.pi` file (e.g. `PI_SCRAPLING_VERSION="main"`).
2. Restart the project:
   ```bash
   ddev restart
   ```

All customization options:

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `PI_SCRAPLING_VERSION` | `main` | Version / git ref of the Scrapling Skill to register |

## Credits

**Contributed and maintained by [@mxr576](https://github.com/mxr576)**
