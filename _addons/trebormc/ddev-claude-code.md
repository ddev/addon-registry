---
title: trebormc/ddev-claude-code
github_url: https://github.com/trebormc/ddev-claude-code
description: "DDEV Add-on: Claude Code CLI container"
user: trebormc
repo: ddev-claude-code
repo_id: 1191702322
default_branch: main
tag_name: v1.0.8
ddev_version_constraint: ">= v1.23.5"
dependencies: ["trebormc/ddev-playwright-mcp", "trebormc/ddev-beads", "trebormc/ddev-agents-sync"]
type: contrib
created_at: 2026-03-25
updated_at: 2026-03-26
workflow_status: unknown
stars: 0
---

[![tests](https://github.com/trebormc/ddev-claude-code/actions/workflows/tests.yml/badge.svg)](https://github.com/trebormc/ddev-claude-code/actions/workflows/tests.yml)

# ddev-claude-code

A DDEV add-on that runs [Claude Code](https://docs.anthropic.com/en/docs/claude-code) (Anthropic's official CLI) in a dedicated container for AI-powered Drupal development.

## Quick Start

```bash
# 1. Install the add-on
ddev add-on get trebormc/ddev-claude-code

# 2. Restart DDEV
ddev restart

# 3. Authenticate (choose one)
ddev claude-code claude login          # OAuth (recommended)
# OR set API key in .ddev/.env.claude-code

# 4. Launch Claude Code
ddev claude-code
```

## Prerequisites

- [DDEV](https://ddev.readthedocs.io/) >= v1.23.5
- An Anthropic API key or OAuth session

## Installation

```bash
ddev add-on get trebormc/ddev-claude-code
ddev restart
```

This automatically installs all dependencies:
- [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) -- auto-syncs AI agents from git (provides CLAUDE.md)
- [ddev-beads](https://github.com/trebormc/ddev-beads) -- task tracking
- [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) -- browser automation

## Authentication

Credentials are stored in a shared directory on the host (`~/.ddev/claude-code/` by default), so you only need to authenticate **once** -- all your DDEV projects share the same session automatically.

**Option A: OAuth login (recommended for Claude Pro/Team/Enterprise subscribers)**

```bash
ddev claude-code claude login
```

This opens a browser for OAuth authentication. Credentials persist across `ddev restart`, new projects, and machine reboots.

**Option B: API key**

```bash
ddev dotenv set .ddev/.env.claude-code --anthropic-api-key=sk-ant-your-key-here
ddev restart
```

## Configuration

After installation, environment variables are in `.ddev/.env.claude-code`:

```bash
# Shared config directory -- credentials, settings, and session data.
# Shared across ALL DDEV projects. Change only if you need a custom location.
HOST_CLAUDE_CONFIG_DIR=${HOME}/.ddev/claude-code

# API key (alternative to OAuth login)
#ANTHROPIC_API_KEY=sk-ant-...

# Timezone
TZ=UTC
```

> **Note:** `HOST_CLAUDE_CONFIG_DIR` must point to an existing directory. The installer creates `~/.ddev/claude-code/` automatically. If you change this value, make sure the directory exists before running `ddev restart`.

### Permissions

The installer creates a default `settings.json` with `bypassPermissions` mode -- all permission prompts are disabled since Claude Code runs inside an isolated DDEV container.

To change this, edit `~/.ddev/claude-code/settings.json`:

```json
{
  "permissions": {
    "defaultMode": "bypassPermissions"
  }
}
```

Available modes: `bypassPermissions` (default, no prompts), `auto` (smart classifier), `acceptEdits` (auto-approve edits only), `default` (prompt for everything).

Since it lives in the shared config directory, permission changes apply to all DDEV projects.

## Architecture

```
┌─────────────────────────────────────────────────┐
│              DDEV Docker Network                 │
│                                                  │
│  ┌──────────────┐  docker exec  ┌────────────┐  │
│  │  Claude Code │──────────────>│    Web     │  │
│  │  Container   │               │  (Drupal)  │  │
│  └──────┬───────┘               └────────────┘  │
│         │ MCP HTTP                               │
│         v                                        │
│  ┌──────────────┐  ┌──────────────┐              │
│  │  Playwright  │  │    Beads     │              │
│  │     MCP      │  │  (bd tasks)  │              │
│  └──────────────┘  └──────────────┘              │
└─────────────────────────────────────────────────┘
```

Claude Code communicates with the web container via `docker exec` (through the mounted Docker socket), giving it full CLI access to drush, composer, phpunit, phpstan, and any other tool in the web container. Playwright MCP is accessed over HTTP for browser automation and visual testing.

## Commands

| Command | Description |
|---------|-------------|
| `ddev claude-code` | Start Claude Code interactive session |
| `ddev claude-code shell` | Open a bash shell in the container |
| `ddev claude-code <command>` | Run any command in the container |

### Shell Helpers

Inside the container (via `ddev claude-code shell`), these helper functions are available:

| Helper | Description |
|--------|-------------|
| `drush` | Run drush commands in the web container |
| `composer` | Run composer in the web container |
| `web-exec` | Execute any command in the web container |
| `web-shell` | Open an interactive shell in the web container |
| `bd` | Run Beads task tracking commands |

## Agents and CLAUDE.md

When [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) is installed (auto-installed as dependency), Claude Code automatically gets:

- **13 specialized agents** in `.claude/agents/` (drupal-dev, three-judges, etc.)
- **CLAUDE.md** with Drupal development instructions in the project root
- **Rules and skills** for Drupal development workflows

Agent `.md` files use model tokens (like `${MODEL_CHEAP}`) that are resolved to Claude Code aliases (like `haiku`) during sync. See [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) for the full list of agents, tokens, and customization options.

You can place your own `CLAUDE.md` in your Drupal project root -- it won't be overwritten if it already exists.

### Customizing agents and models

To add private agents or change which models agents use, edit `.ddev/.env.agents-sync`:

```bash
# Add a private repo with custom agents or model overrides
AGENTS_REPOS=https://github.com/trebormc/drupal-ai-agents.git,https://github.com/your-org/private-agents.git
```

See [Model Token System](https://github.com/trebormc/ddev-agents-sync#model-token-system) for details on changing agent models globally.

## Desktop Notifications

Claude Code can send desktop notifications when sessions finish. Add a stop hook to `.claude/settings.json`:

```json
{
  "hooks": {
    "stop": [
      {
        "matcher": "",
        "hooks": [
          {
            "type": "command",
            "command": "curl -s -X POST http://host.docker.internal:5454/notify -H 'Content-Type: application/json' -d '{\"title\":\"Claude Code\",\"message\":\"Session finished\"}'"
          }
        ]
      }
    ]
  }
}
```

Then start the notification bridge on your host:

```bash
./scripts/start-notify-bridge.sh
```

See the [DDEV AI workspace](https://github.com/trebormc/ddev-ai-workspace) for full notification setup details.

## Autonomous Execution

For autonomous task execution (overnight runs), see [ddev-ralph](https://github.com/trebormc/ddev-ralph).

## Part of DDEV AI Workspace

This add-on is part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), a modular ecosystem of DDEV add-ons for AI-powered Drupal development.

| Repository | Description | Relationship |
|------------|-------------|--------------|
| [ddev-ai-workspace](https://github.com/trebormc/ddev-ai-workspace) | Meta add-on that installs the full AI development stack with one command. | Workspace |
| [ddev-opencode](https://github.com/trebormc/ddev-opencode) | [OpenCode](https://opencode.ai) AI CLI container for interactive development. | Alternative AI tool |
| [ddev-ralph](https://github.com/trebormc/ddev-ralph) | Autonomous AI task orchestrator. Delegates work to this container or OpenCode via `docker exec`. | Uses this as backend |
| [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) | Headless Playwright browser for browser automation and visual testing. | Auto-installed dependency |
| [ddev-beads](https://github.com/trebormc/ddev-beads) | [Beads](https://github.com/steveyegge/beads) git-backed task tracker shared by all AI containers. | Auto-installed dependency |
| [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) | Auto-syncs AI agent repositories into a shared Docker volume. Provides CLAUDE.md. | Auto-installed dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 13 agents, 4 rules, 14 skills for Drupal development. Synced automatically via ddev-agents-sync. | Agent configuration |

## Disclaimer

This project is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors -- always review changes before deploying to production. See [menetray.com](https://menetray.com) for more information and [DruScan](https://druscan.com) for Drupal auditing tools.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-claude-code/blob/main/LICENSE).
