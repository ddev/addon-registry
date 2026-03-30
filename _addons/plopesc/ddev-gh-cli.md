---
title: plopesc/ddev-gh-cli
github_url: https://github.com/plopesc/ddev-gh-cli
description: "ddev add-on for setting up the GH Client."
user: plopesc
repo: ddev-gh-cli
repo_id: 1165690350
default_branch: main
tag_name: v1.1.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-24
updated_at: 2026-02-25
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/plopesc/ddev-gh-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/plopesc/ddev-gh-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/plopesc/ddev-gh-cli)](https://github.com/plopesc/ddev-gh-cli/commits)
[![release](https://img.shields.io/github/v/release/plopesc/ddev-gh-cli)](https://github.com/plopesc/ddev-gh-cli/releases/latest)

# DDEV GH CLI

## Overview

This add-on integrates GH CLI into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get plopesc/ddev-gh-cli
ddev restart
```

If your GitHub token is already stored in `~/.config/gh/hosts.yml`, no additional configuration is required. Otherwise, you must set either the `GH_TOKEN` or `GITHUB_TOKEN` environment variable with a valid token.

To retrieve your token, run the following command on your host machine:
```sh
gh auth token
```

## Usage

| Command                                     | Description                              |
|---------------------------------------------|------------------------------------------|
| `ddev exec gh --version`                    | Check the installed version              |
| `ddev exec gh --help`                       | View available commands                  |

The full documentation about GH CLI can be gound at the [GH CLI documentation page](https://cli.github.com/).

## Credits

**Contributed and maintained by [@plopesc](https://github.com/plopesc)**
