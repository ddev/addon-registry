---
title: trebormc/ddev-playwright-mcp
github_url: https://github.com/trebormc/ddev-playwright-mcp
description: "DDEV Add-on: Headless Playwright browser (MCP)"
user: trebormc
repo: ddev-playwright-mcp
repo_id: 1191701790
default_branch: main
tag_name: v1.0.1
ddev_version_constraint: ">= v1.23.5"
dependencies: []
type: contrib
created_at: 2026-03-25
updated_at: 2026-03-25
workflow_status: unknown
stars: 0
---

[![tests](https://github.com/trebormc/ddev-playwright-mcp/actions/workflows/tests.yml/badge.svg)](https://github.com/trebormc/ddev-playwright-mcp/actions/workflows/tests.yml)

# ddev-playwright-mcp

A DDEV add-on that provides a [Playwright MCP](https://github.com/anthropics/playwright-mcp) server for headless browser automation via the Model Context Protocol.

> **Note:** This add-on is typically installed automatically as a dependency of [ddev-opencode](https://github.com/trebormc/ddev-opencode), [ddev-claude-code](https://github.com/trebormc/ddev-claude-code), or [ddev-ralph](https://github.com/trebormc/ddev-ralph). You rarely need to install it directly.

## Installation

```bash
ddev add-on get trebormc/ddev-playwright-mcp
ddev restart
```

## What It Does

This add-on runs a **Playwright MCP server** as a DDEV service with headless Chromium. It exposes an MCP endpoint at `http://playwright-mcp:8931/mcp` accessible from other containers in the DDEV network.

Use cases:
- **AI-driven development** -- Let AI agents (OpenCode, Claude Code) interact with your site through a real browser
- **Visual testing** -- Take screenshots and verify UI state programmatically
- **Browser automation** -- Automate form submissions, navigation, and interactions

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
- The MCP server runs in **shared browser context** mode -- cookies and sessions persist across requests.
- Screenshots output directory: `/tmp/playwright-output` (mapped to `./screenshots/` on host).

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
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 13 agents, 4 rules, 14 skills for Drupal development. Includes Playwright browser testing skill. | Uses this for visual testing |

## Disclaimer

This project is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors -- always review changes before deploying to production. See [menetray.com](https://menetray.com) for more information and [DruScan](https://druscan.com) for Drupal auditing tools.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-playwright-mcp/blob/main/LICENSE).
