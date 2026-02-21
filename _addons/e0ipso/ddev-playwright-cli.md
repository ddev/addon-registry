---
title: e0ipso/ddev-playwright-cli
github_url: https://github.com/e0ipso/ddev-playwright-cli
description: "DDEV add-on that installs @playwright/cli globally in the web container and automatically sets up Claude Code skills"
user: e0ipso
repo: ddev-playwright-cli
repo_id: 1161565491
default_branch: main
tag_name: v1.1.1
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-19
updated_at: 2026-02-19
workflow_status: failure
stars: 2
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/e0ipso/ddev-playwright-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/e0ipso/ddev-playwright-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/e0ipso/ddev-playwright-cli)](https://github.com/e0ipso/ddev-playwright-cli/commits)
[![release](https://img.shields.io/github/v/release/e0ipso/ddev-playwright-cli)](https://github.com/e0ipso/ddev-playwright-cli/releases/latest)

## Overview

This add-on installs [`@playwright/cli`](https://www.npmjs.com/package/@playwright/cli) globally in the DDEV web container and automatically installs Claude Code skills on container start.

What this add-on does:

- Installs `@playwright/cli` globally in the web container via `web-build/Dockerfile.playwright-cli`
- Runs `playwright-cli install --skills` automatically via a post-start hook to install Claude Code skills

## Installation

```bash
ddev add-on get e0ipso/ddev-playwright-cli
ddev restart
```

After installation, commit the `.ddev` directory to version control.

## Usage

| Command                                     | Description                              |
|---------------------------------------------|------------------------------------------|
| `ddev exec playwright-cli --version`        | Check the installed version              |
| `ddev exec playwright-cli --help`           | View available commands                  |
| `ddev exec playwright-cli install --skills` | Manually (re-)install Claude Code skills |

## Credits

**Contributed and maintained by [@e0ipso](https://github.com/e0ipso)**
