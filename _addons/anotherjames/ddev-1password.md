---
title: anotherjames/ddev-1password
github_url: https://github.com/anotherjames/ddev-1password
description: "1Password-focussed DDEV add-on to use the host SSH agent in the web container"
user: anotherjames
repo: ddev-1password
repo_id: 996810365
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-06-05
updated_at: 2025-06-06
workflow_status: success
stars: 4
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/anotherjames/ddev-1password/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/anotherjames/ddev-1password/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/anotherjames/ddev-1password)](https://github.com/anotherjames/ddev-1password/commits)
[![release](https://img.shields.io/github/v/release/anotherjames/ddev-1password)](https://github.com/anotherjames/ddev-1password/releases/latest)

# DDEV 1password

## Overview

This add-on integrates 1password into your [DDEV](https://ddev.com/) project to
make it available as your SSH agent.

## Installation

```bash
ddev add-on get anotherjames/ddev-1password
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev 1password` | Set up 1password for use on macs as the SSH agent to be used globally for every client (not just DDEV). |

## Credits

- Inspired by [@darthsteven's work](https://github.com/ddev/ddev/issues/3878#issuecomment-2491614576)**
- Contributed and maintained by [@anotherjames](https://github.com/anotherjames)**
