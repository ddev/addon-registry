---
title: dragonwize/ddev-opencode
github_url: https://github.com/dragonwize/ddev-opencode
description: "Use OpenCode with your DDEV project. Being in the container, the agent is isolated and uses the versioned dependencies and tools allowed."
user: dragonwize
repo: ddev-opencode
repo_id: 1146504993
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-01-31
updated_at: 2026-02-23
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/dragonwize/ddev-opencode/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dragonwize/ddev-opencode/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/dragonwize/ddev-opencode)](https://github.com/dragonwize/ddev-opencode/commits)
[![release](https://img.shields.io/github/v/release/dragonwize/ddev-opencode)](https://github.com/dragonwize/ddev-opencode/releases/latest)

# DDEV OpenCode

## Overview

This add-on integrates Opencode into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get dragonwize/ddev-opencode
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command               | Description                                                                                   |
|-----------------------|-----------------------------------------------------------------------------------------------|
| `ddev opencode`       | Start OpenCode in the container.                                                              |
| `ddev opencode-serve` | Start OpenCode server. Then connect to it from outside with `opencode attach <DDEV URL>:4096. |

## Update OpenCode

OpenCode has new versions very frequently. Since the install is cached in the 
Docker container for faster startups, a container rebuild is needed to pull the
latest version.

```bash
ddev add-on get dragonwize/ddev-opencode
ddev restart --no-cache
```

If you want to update quickly, knowing it will not keep on restart:
(Useful to check the new version before saving in your docker container)

```bash
ddev opencode upgrade
```

## Credits

**Contributed and maintained by [@dragonwize](https://github.com/dragonwize)**
