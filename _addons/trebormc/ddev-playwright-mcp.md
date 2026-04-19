---
title: trebormc/ddev-playwright-mcp
github_url: https://github.com/trebormc/ddev-playwright-mcp
description: "DDEV Add-on: Headless Playwright browser (MCP)"
user: trebormc
repo: ddev-playwright-mcp
repo_id: 1191701790
default_branch: main
tag_name: v1.0.45
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-03-25
updated_at: 2026-04-18
workflow_status: disabled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/trebormc/ddev-playwright-mcp/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/trebormc/ddev-playwright-mcp/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/trebormc/ddev-playwright-mcp)](https://github.com/trebormc/ddev-playwright-mcp/commits)
[![release](https://img.shields.io/github/v/release/trebormc/ddev-playwright-mcp)](https://github.com/trebormc/ddev-playwright-mcp/releases/latest)

# ddev-playwright-mcp

A DDEV add-on that provides a [Playwright MCP](https://github.com/anthropics/playwright-mcp) server for headless browser automation via the Model Context Protocol. It gives AI agents a real Chromium browser to interact with your Drupal site: navigate pages, fill forms, take screenshots, and verify UI state programmatically.

> **Part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace)** — a modular ecosystem of DDEV add-ons for AI-powered Drupal development. Install the full stack with one command: `ddev add-on get trebormc/ddev-ai-workspace`
>
> Created by [Robert Menetray](https://menetray.com) · Sponsored by [DruScan](https://druscan.com)

> **Note:** This add-on is typically installed automatically as a dependency of [ddev-opencode](https://github.com/trebormc/ddev-opencode), [ddev-claude-code](https://github.com/trebormc/ddev-claude-code), or [ddev-ralph](https://github.com/trebormc/ddev-ralph). You rarely need to install it directly.

## Quick Start

The **recommended way** to install this add-on is through the [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), which installs all tools and dependencies with a single command:

```bash
ddev add-on get trebormc/ddev-ai-workspace
ddev restart
```

This add-on is also **automatically installed** as a dependency when you install [ddev-opencode](https://github.com/trebormc/ddev-opencode), [ddev-claude-code](https://github.com/trebormc/ddev-claude-code), or [ddev-ralph](https://github.com/trebormc/ddev-ralph). You rarely need to install it directly.

### Standalone installation

If you need to install it individually (requires familiarity with the DDEV add-on ecosystem):

```bash
ddev add-on get trebormc/ddev-playwright-mcp
ddev restart
```

## What It Does

This add-on runs a **Playwright MCP server** as a DDEV service with headless Chromium. It exposes an MCP endpoint at `http://playwright-mcp:8931/mcp` accessible from other containers in the DDEV network.

Use cases:
- **AI-driven development**: Let AI agents (OpenCode, Claude Code) interact with your site through a real browser
- **Visual testing**: Take screenshots and verify UI state programmatically
- **Browser automation**: Automate form submissions, navigation, and interactions

## Usage

### From Other Containers

Connect to the MCP server at:
```
http://playwright-mcp:8931/mcp
```

### Screenshots

Screenshots are saved to the `screenshots/` directory in your project root, accessible from both the container and the host.

### Custom Image Version

To use a specific Playwright MCP image version:

```bash
ddev dotenv set .ddev/.env.playwright-mcp --playwright-mcp-image=mcr.microsoft.com/playwright/mcp:v1.2.0
ddev restart
```

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PLAYWRIGHT_MCP_IMAGE` | `mcr.microsoft.com/playwright/mcp:latest` | Docker image to use |
| `TZ` | `UTC` | Timezone |

## Tips

- **Use HTTP** (not HTTPS) when navigating to your DDEV site from Playwright to avoid SSL certificate issues in local environments.
- The MCP server runs in **shared browser context** mode. Cookies and sessions persist across requests.
- Screenshots output directory: `/var/www/html/screenshots` (mapped to `./screenshots/` on host).

## Uninstallation

```bash
ddev add-on remove ddev-playwright-mcp
ddev restart
```

## Part of DDEV AI Workspace

This add-on is part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), a modular ecosystem of DDEV add-ons for AI-powered Drupal development.

| Repository | Description | Relationship |
|------------|-------------|--------------|
| [ddev-ai-workspace](https://github.com/trebormc/ddev-ai-workspace) | Meta add-on that installs the full AI development stack with one command. | Workspace |
| [ddev-opencode](https://github.com/trebormc/ddev-opencode) | [OpenCode](https://opencode.ai) AI CLI container for interactive development. | Auto-installs this add-on |
| [ddev-claude-code](https://github.com/trebormc/ddev-claude-code) | [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI container for interactive development. | Auto-installs this add-on |
| [ddev-ralph](https://github.com/trebormc/ddev-ralph) | Autonomous AI task orchestrator. Delegates work to OpenCode or Claude Code. | Auto-installs this add-on |
| [ddev-beads](https://github.com/trebormc/ddev-beads) | [Beads](https://github.com/steveyegge/beads) git-backed task tracker shared by all AI containers. | Sibling dependency |
| [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) | Auto-syncs AI agent repositories into a shared Docker volume. | Sibling dependency |
| [ddev-ai-ssh](https://github.com/trebormc/ddev-ai-ssh) | SSH access to the web container for AI agents. | Sibling dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 10 agents, 12 rules, 24 skills for Drupal development. Includes Playwright browser testing skill. | Uses this for visual testing |

## Disclaimer

This project is an independent initiative by [Robert Menetray](https://menetray.com), sponsored by [DruScan](https://druscan.com). It is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors. Always review changes before deploying to production.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-playwright-mcp/blob/main/LICENSE).
