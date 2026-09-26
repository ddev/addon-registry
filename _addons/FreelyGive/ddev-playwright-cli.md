---
title: "FreelyGive/ddev-playwright-cli"
github_url: "https://github.com/FreelyGive/ddev-playwright-cli"
description: "Integrate playwright cli into your ddev setup for AI coding tools in the web container."
user: "FreelyGive"
repo: "ddev-playwright-cli"
repo_id: 1160089818
default_branch: "main"
tag_name: "0.5.0"
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: "contrib"
created_at: "2026-02-17"
updated_at: "2026-09-01"
workflow_status: "failure"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/FreelyGive/ddev-playwright-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/FreelyGive/ddev-playwright-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/FreelyGive/ddev-playwright-cli)](https://github.com/FreelyGive/ddev-playwright-cli/commits)
[![release](https://img.shields.io/github/v/release/FreelyGive/ddev-playwright-cli)](https://github.com/FreelyGive/ddev-playwright-cli/releases/latest)

# DDEV Playwright Cli

## Overview

This add-on integrates [Playwright CLI](https://www.npmjs.com/package/@playwright/cli) into your [DDEV](https://ddev.com/) project, so you can drive a browser from inside the web container — and share a single, cached browser install with other Playwright-based tooling such as Storybook/Vitest browser tests.

## Installation

```bash
ddev add-on get FreelyGive/ddev-playwright-cli
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

Run any `playwright-cli` command through `ddev playwright-cli …`; it executes in the web container.

| Command                                       | Description                                          |
|-----------------------------------------------|------------------------------------------------------|
| `ddev playwright-cli open <url>`              | Open a URL and print a page snapshot.                |
| `ddev playwright-cli config-print`           | Print the resolved Playwright CLI configuration.     |
| `ddev playwright-cli install-browser <name>` | Install an additional browser into the shared cache. |

Because it runs inside the container, it trusts DDEV's mkcert CA, so you can open your project's own `https://<project>.ddev.site/` URLs.

## What the add-on does

On each `ddev start`, the add-on:

- adds `@playwright/cli` to your project's `package.json` (if not already declared) and installs it;
- installs the Chromium browser into DDEV's shared browser cache (`/mnt/ddev-global-cache/playwright-browsers`);
- installs the Playwright CLI Claude skills into `.claude/skills/`;
- adds DDEV's mkcert CA to the browser's trust store so it can open `https://*.ddev.site` URLs.

Commit the resulting changes — the `.ddev/` directory, the `package.json`/lockfile updates, and `.claude/skills/` — to version control.

## Sharing a browser version with other tools (Storybook/Vitest)

The add-on sets `PLAYWRIGHT_BROWSERS_PATH=/mnt/ddev-global-cache/playwright-browsers`, DDEV's persistent global cache, so every Playwright-based tool in the project downloads and reuses browsers from one shared location instead of each fetching its own.

Sharing only works if the tools agree on a browser **revision**, which is tied to the `playwright-core` version. Two tools reuse the same browser only when they resolve to the same `playwright-core`. To guarantee that, pin a single `playwright-core` across `@playwright/cli` and your test tooling — for example with npm `overrides` (or `resolutions` for pnpm/yarn) in `package.json`:

```json
{
  "overrides": {
    "playwright-core": "1.x.y"
  }
}
```

With a single resolved `playwright-core`, the add-on's `install-browser chromium` populates the cache once and your other tools (e.g. `@vitest/browser-playwright` driving Storybook stories) reuse it rather than downloading a second copy.

Because the cache is shared, the add-on also sets `PLAYWRIGHT_SKIP_BROWSER_GC=1`. Playwright normally garbage-collects browser revisions it no longer recognises, which is safe for a per-project cache but destructive for a shared one: an install in one project would delete the revisions other projects pinned to a different `playwright-core`, forcing them to re-download on their next run.

## Credits

**Contributed and maintained by [@FreelyGive](https://github.com/FreelyGive)**
