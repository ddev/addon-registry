---
title: trebormc/ddev-claude-code
github_url: https://github.com/trebormc/ddev-claude-code
description: "DDEV Add-on: Claude Code CLI container"
user: trebormc
repo: ddev-claude-code
repo_id: 1191702322
default_branch: main
tag_name: v1.0.37
ddev_version_constraint: ">= v1.23.5"
dependencies: ["trebormc/ddev-playwright-mcp", "trebormc/ddev-beads", "trebormc/ddev-agents-sync"]
type: contrib
created_at: 2026-03-25
updated_at: 2026-04-03
workflow_status: disabled
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

# 3. Launch Claude Code (authenticate on first run)
ddev claude-code  # or: ddev cc
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

Run `ddev claude-code` and follow the prompts. Claude Code handles OAuth and API key authentication natively -- no custom commands or manual file editing needed.

Credentials are stored in a shared directory on the host (`~/.ddev/claude-code/` by default), so you only need to authenticate **once** -- all your DDEV projects share the same session automatically. Credentials persist across `ddev restart`, new projects, and machine reboots.

## Configuration

After installation, environment variables are in `.ddev/.env.claude-code`:

```bash
# Shared config directory -- credentials, settings, and session data.
# Shared across ALL DDEV projects. Change only if you need a custom location.
HOST_CLAUDE_CONFIG_DIR=${HOME}/.ddev/claude-code

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
| `ddev cc` | Alias for `ddev claude-code` |
| `ddev claude-code tui` | Start interactive session (same as above) |
| `ddev claude-code tui Fix login bug` | Start interactive session with a custom tab title |
| `ddev claude-code shell` | Open a bash shell in the container |
| `ddev claude-code <command>` | Run any command in the container |

### Tab title for multi-project workflows

When working on multiple DDEV projects at the same time, it can be hard to tell which terminal belongs to which project. The `tui` subcommand sets the terminal tab title to **`project-name - custom text`**, so you can identify each terminal at a glance.

The project name (`DDEV_SITENAME`) is always included automatically. If you add extra text after `tui`, it appears as a label -- useful for describing the task you are working on in that terminal.

```bash
# Tab title: "mysite - Claude Code"
ddev claude-code

# Tab title: "mysite - Claude Code"  (explicit tui, same result)
ddev claude-code tui

# Tab title: "mysite - Fix login redirect bug"
ddev claude-code tui Fix login redirect bug

# Tab title: "mysite - TASK-42 migrate users"
ddev claude-code tui TASK-42 migrate users
```

This way, if you have three terminals open (two projects, two tasks), each tab shows exactly where you are and what you are doing.

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

## Desktop Notifications (optional)

Claude Code can send desktop notifications when sessions finish. First, install the [ai-notify-bridge](https://github.com/trebormc/ai-notify-bridge) on your host (one-time setup):

```bash
curl -fsSL https://raw.githubusercontent.com/trebormc/ai-notify-bridge/main/install.sh | bash
```

Notification hooks are pre-configured in `settings.json` when you install the add-on. They include the project name and TUI task label automatically. Example notification title: `[mysite] Fix login bug`.

If you already have a `settings.json` from a previous install, add the hooks manually. See the [install.yaml](https://github.com/trebormc/ddev-claude-code/blob/main/install.yaml) for the exact hook configuration.

If the bridge is not installed or not running, the curl call fails silently with no impact on Claude Code.

## Autonomous Execution

For autonomous task execution (overnight runs), see [ddev-ralph](https://github.com/trebormc/ddev-ralph).

## Uninstallation

```bash
ddev add-on remove ddev-claude-code
ddev restart
```

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
