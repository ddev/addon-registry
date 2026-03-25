---
title: e0ipso/ddev-playwright-cli
github_url: https://github.com/e0ipso/ddev-playwright-cli
description: "DDEV add-on that installs @playwright/cli globally in the web container and automatically sets up Claude Code skills"
user: e0ipso
repo: ddev-playwright-cli
repo_id: 1161565491
default_branch: main
tag_name: v1.2.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-19
updated_at: 2026-03-24
workflow_status: failure
stars: 3
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/e0ipso/ddev-playwright-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/e0ipso/ddev-playwright-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/e0ipso/ddev-playwright-cli)](https://github.com/e0ipso/ddev-playwright-cli/commits)
[![release](https://img.shields.io/github/v/release/e0ipso/ddev-playwright-cli)](https://github.com/e0ipso/ddev-playwright-cli/releases/latest)

## Overview

This add-on installs [`@playwright/cli`](https://www.npmjs.com/package/@playwright/cli) globally in the DDEV web container along with its bundled Chromium browser, and automatically installs Claude Code skills on container start.

[Playwright CLI](https://github.com/microsoft/playwright-cli) is a
token-efficient command-line tool designed for AI coding agents such as Claude
Code, GitHub Copilot, and Cursor. It lets agents interact with Playwright
browsers through simple shell commands instead of writing full test scripts.

What this add-on does:

- Installs `@playwright/cli` and Chromium at Docker build time via `web-build/Dockerfile.playwright-cli` with Docker build cache for fast rebuilds
- Runs `playwright-cli install --skills` automatically via a post-start hook to install Claude Code skills

## Installation

```bash
ddev add-on get e0ipso/ddev-playwright-cli
ddev install-playwright-cli
```

The `ddev install-playwright-cli` command copies the Dockerfile into your build
directory and restarts the web container. This opt-in step keeps the add-on
lightweight for projects that install it but aren't ready to use it yet.

After installation, commit the `.ddev` directory to version control.

## Usage

All Playwright CLI commands are available through `ddev exec playwright-cli`:

```console
# Open a URL in the browser.
ddev exec playwright-cli open https://example.ddev.site

# Take a snapshot of the current page (returns an accessibility tree).
ddev exec playwright-cli snapshot

# Click an element, fill a field, etc.
ddev exec playwright-cli click --selector "text=Log in"
ddev exec playwright-cli fill --selector "#edit-name" --value admin
```

| Command                                     | Description                              |
|---------------------------------------------|------------------------------------------|
| `ddev exec playwright-cli --version`        | Check the installed version              |
| `ddev exec playwright-cli --help`           | View available commands                  |
| `ddev exec playwright-cli install --skills` | Manually (re-)install Claude Code skills |

## Credits

**Contributed and maintained by [@e0ipso](https://github.com/e0ipso)**
