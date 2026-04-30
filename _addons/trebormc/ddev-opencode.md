---
title: trebormc/ddev-opencode
github_url: https://github.com/trebormc/ddev-opencode
description: "DDEV Add-on: OpenCode AI CLI container"
user: trebormc
repo: ddev-opencode
repo_id: 1191702047
default_branch: main
tag_name: v1.0.45
ddev_version_constraint: ">= v1.24.10"
dependencies: ["trebormc/ddev-ai-ssh", "trebormc/ddev-playwright-mcp", "trebormc/ddev-beads", "trebormc/ddev-agents-sync"]
type: contrib
created_at: 2026-03-25
updated_at: 2026-04-18
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/trebormc/ddev-opencode/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/trebormc/ddev-opencode/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/trebormc/ddev-opencode)](https://github.com/trebormc/ddev-opencode/commits)
[![release](https://img.shields.io/github/v/release/trebormc/ddev-opencode)](https://github.com/trebormc/ddev-opencode/releases/latest)

# ddev-opencode

A DDEV add-on that runs [OpenCode](https://github.com/opencode-ai/opencode) in a dedicated container for AI-powered **Drupal** development.

> **Part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace)** — a modular ecosystem of DDEV add-ons for AI-powered Drupal development. Install the full stack with one command: `ddev add-on get trebormc/ddev-ai-workspace`
>
> Created by [Robert Menetray](https://menetray.com) · Sponsored by [DruScan](https://druscan.com)

Agents, rules, and skills for Drupal development are automatically synced from [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) via [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync). No manual git clone needed.

## Quick Start

The **recommended way** to install this add-on is through the [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), which installs all tools and dependencies with a single command:

```bash
ddev add-on get trebormc/ddev-ai-workspace
ddev restart
ddev opencode  # or: ddev oc
```

### Standalone installation

If you only need OpenCode without the rest of the workspace, you can install it individually. This requires familiarity with the DDEV add-on ecosystem and its dependencies:

```bash
ddev add-on get trebormc/ddev-opencode
ddev restart
ddev opencode  # or: ddev oc
```

This automatically installs the required dependencies:
- [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync): auto-syncs AI agents from git
- [ddev-ai-ssh](https://github.com/trebormc/ddev-ai-ssh): SSH access to web container
- [ddev-beads](https://github.com/trebormc/ddev-beads): task tracking
- [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp): browser automation

## Prerequisites

- [DDEV](https://ddev.readthedocs.io/) >= v1.24.10
- An API key (Anthropic, OpenAI, or a LiteLLM proxy)

## Authentication

Run `ddev opencode` and follow the prompts. OpenCode handles authentication natively. No custom commands or manual file editing needed.

Credentials are stored in a shared file on the host (`~/.ddev/opencode/auth.json` by default), so you only need to authenticate **once**. All your DDEV projects share the same credentials automatically.

## Configuration

After installation, environment variables are in `.ddev/.env.opencode`:

```bash
# Shared OpenCode directory for credentials and config.
# Shared across ALL DDEV projects. Change only if you need a custom location.
# Files: auth.json (credentials), config/ (opencode.json, custom overrides)
HOST_OPENCODE_DIR=${HOME}/.ddev/opencode

# Timezone
TZ=UTC
```

### Agent Configuration

Agents, rules, and skills are automatically synced from [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) via [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync). No manual setup is needed.

The sync process resolves model tokens (like `${MODEL_CHEAP}`) to OpenCode model names (like `opencode/gpt-5-nano`) using the `.env.agents` file from the agent repository. See [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) for details on model tokens and customization.

To customize which repos are synced, edit `.ddev/.env.agents-sync`:

```bash
# Sync from multiple repos (later repos override earlier ones)
AGENTS_REPOS=https://github.com/trebormc/drupal-ai-agents.git,https://github.com/your-org/private-agents.git
```

To manually trigger an update: `ddev agents-update`

### Changing agent models

To change which AI models your agents use, override the `.env.agents` file. Create a private git repo with just an `.env.agents` and add it as a second entry:

```bash
AGENTS_REPOS=https://github.com/trebormc/drupal-ai-agents.git,https://github.com/your-org/my-config.git
```

See [Model Token System](https://github.com/trebormc/ddev-agents-sync#model-token-system) for the full list of tokens and how to customize them.

### Local config override

To use a custom `opencode.json` (e.g., to add a LiteLLM proxy or change global permissions), place it in `~/.ddev/opencode/config/opencode.json`. It takes precedence over the default from the agent repository. Agents, rules, and skills are always loaded from the synced volume regardless.

## Architecture

```
┌─────────────────────────────────────────────────┐
│              DDEV Docker Network                 │
│                                                  │
│  ┌──────────────┐     SSH       ┌────────────┐  │
│  │   OpenCode   │──────────────>│    Web     │  │
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

OpenCode communicates with the web container via SSH (`ssh web`), giving it full CLI access to drush, composer, phpunit, phpstan, and any other tool in the web container. SSH keys are auto-generated per project in `.ddev/.agent-ssh-keys/`. Playwright MCP is accessed over HTTP for browser automation and visual testing.

## Commands

| Command | Description |
|---------|-------------|
| `ddev opencode` | Launch the OpenCode TUI |
| `ddev oc` | Alias for `ddev opencode` |
| `ddev opencode tui` | Launch TUI (same as above) |
| `ddev opencode tui Fix login bug` | Launch TUI with a custom tab title |
| `ddev opencode shell` | Open a bash shell in the container |
| `ddev opencode <command>` | Run any command in the container |

### Tab title for multi-project workflows

When working on multiple DDEV projects at the same time, it can be hard to tell which terminal belongs to which project. The `tui` subcommand sets the terminal tab title to **`project-name - custom text`**, so you can identify each terminal at a glance.

The project name (`DDEV_SITENAME`) is always included automatically. If you add extra text after `tui`, it appears as a label. Useful for describing the task you are working on in that terminal.

```bash
# Tab title: "mysite - OpenCode"
ddev opencode

# Tab title: "mysite - OpenCode"  (explicit tui, same result)
ddev opencode tui

# Tab title: "mysite - Fix login redirect bug"
ddev opencode tui Fix login redirect bug

# Tab title: "mysite - TASK-42 migrate users"
ddev opencode tui TASK-42 migrate users
```

This way, if you have three terminals open (two projects, two tasks), each tab shows exactly where you are and what you are doing.

### Shell Helpers

Inside the container (via `ddev opencode shell`), these helper functions are available:

| Helper | Description |
|--------|-------------|
| `drush` | Run drush commands in the web container |
| `composer` | Run composer in the web container |
| `phpunit` | Run PHPUnit tests in the web container |
| `phpstan` | Run PHPStan analysis in the web container |
| `web-exec` | Execute any command in the web container |
| `web-shell` | Open an interactive shell in the web container |
| `bd` | Run Beads task tracking commands |

## Desktop Notifications (optional)

OpenCode can send desktop notifications when tasks complete or need attention. Notifications are pre-configured via `opencode-notifier.json` (ships with [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents)). No setup needed inside the container.

To receive notifications, install the [ai-notify-bridge](https://github.com/trebormc/ai-notify-bridge) on your host (one-time setup):

```bash
curl -fsSL https://raw.githubusercontent.com/trebormc/ai-notify-bridge/main/install.sh | bash
```

If the bridge is not installed or not running, OpenCode works normally. Notification calls fail silently with no impact.

## Autonomous Execution

For autonomous task execution (overnight runs), see [ddev-ralph](https://github.com/trebormc/ddev-ralph).

## Uninstallation

```bash
ddev add-on remove ddev-opencode
ddev restart
```

## Part of DDEV AI Workspace

This add-on is part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), a modular ecosystem of DDEV add-ons for AI-powered Drupal development.

| Repository | Description | Relationship |
|------------|-------------|--------------|
| [ddev-ai-workspace](https://github.com/trebormc/ddev-ai-workspace) | Meta add-on that installs the full AI development stack with one command. | Workspace |
| [ddev-claude-code](https://github.com/trebormc/ddev-claude-code) | [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI container for interactive development. | Alternative AI tool |
| [ddev-ralph](https://github.com/trebormc/ddev-ralph) | Autonomous AI task orchestrator. Delegates work to this container or Claude Code via SSH. | Uses this as backend |
| [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) | Headless Playwright browser for browser automation and visual testing. | Auto-installed dependency |
| [ddev-beads](https://github.com/trebormc/ddev-beads) | [Beads](https://github.com/steveyegge/beads) git-backed task tracker shared by all AI containers. | Auto-installed dependency |
| [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) | Auto-syncs AI agent repositories into a shared Docker volume. Provides agents and config. | Auto-installed dependency |
| [ddev-ai-ssh](https://github.com/trebormc/ddev-ai-ssh) | SSH access to the web container. Generates per-project keys, installs sshd. | Auto-installed dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 10 agents, 12 rules, 24 skills for Drupal development. Synced automatically via ddev-agents-sync. | Agent configuration |

## Disclaimer

This project is an independent initiative by [Robert Menetray](https://menetray.com), sponsored by [DruScan](https://druscan.com). It is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors. Always review changes before deploying to production.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-opencode/blob/main/LICENSE).
