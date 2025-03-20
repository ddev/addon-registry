---
title: backdrop-ops/ddev-backdrop-bee
github_url: https://github.com/backdrop-ops/ddev-backdrop-bee
description: "Install bee for Backdrop sites on ddev."
user: backdrop-ops
repo: ddev-backdrop-bee
repo_id: 873926344
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2024-10-17
updated_at: 2025-03-19
stars: 1
---

[![tests](https://github.com/backdrop-ops/ddev-backdrop-bee/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/backdrop-ops/ddev-backdrop-bee/actions/workflows/tests.yml?query=branch%3Amain)
[![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)](https://github.com/backdrop-ops/ddev-backdrop-bee/commits)
[![release](https://img.shields.io/github/v/release/backdrop-ops/ddev-backdrop-bee)](https://github.com/backdrop-ops/ddev-backdrop-bee/releases/latest)

# DDEV Backdrop Bee

## Overview

Bee is a command line utility for Backdrop CMS. It includes commands that allow developers to interact with Backdrop sites, performing actions like:
- Running cron
- Clearing caches
- Downloading and installing Backdrop
- Downloading, enabling and disabling projects
- Viewing information about a site and/or available projects

See the Release notes and the Changelog for details of changes between versions:
- [Backdrop CMS Bee Project](https://backdropcms.org/project/bee)
- [GitHub Project for Bee](https://github.com/backdrop-contrib/bee)

This add-on integrates Bee into your [DDEV](https://ddev.com/) Backdrop CMS project.

## Installation

To install this add-on, run:

```bash
ddev add-on get backdrop-ops/ddev-backdrop-bee
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev bee` | Run command line utility for Backdrop CMS |

Examples:

```bash
ddev bee
ddev bee status
```

## Credits

**Contributed and maintained by [@jenlampton](https://github.com/jenlampton) and [@wylbur](https://github.com/wylbur)**
