---
title: FreelyGive/ddev-playwright-cli
github_url: https://github.com/FreelyGive/ddev-playwright-cli
description: "Integrate playwright cli into your ddev setup for AI coding tools in the web container."
user: FreelyGive
repo: ddev-playwright-cli
repo_id: 1160089818
default_branch: main
tag_name: 0.1.2
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-17
updated_at: 2026-02-18
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/FreelyGive/ddev-playwright-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/FreelyGive/ddev-playwright-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/FreelyGive/ddev-playwright-cli)](https://github.com/FreelyGive/ddev-playwright-cli/commits)
[![release](https://img.shields.io/github/v/release/FreelyGive/ddev-playwright-cli)](https://github.com/FreelyGive/ddev-playwright-cli/releases/latest)

# DDEV Playwright Cli

## Overview

This add-on integrates Playwright CLI into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get FreelyGive/ddev-playwright-cli
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command                       | Description                               |
|-------------------------------|-------------------------------------------|
| `ddev playwright-cli`         | Execute the Playwright CLI from the host. |

## Credits

**Contributed and maintained by [@FreelyGive](https://github.com/FreelyGive)**
