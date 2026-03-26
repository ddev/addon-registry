---
title: trebormc/ddev-opencode
github_url: https://github.com/trebormc/ddev-opencode
description: "DDEV Add-on: OpenCode AI CLI container"
user: trebormc
repo: ddev-opencode
repo_id: 1191702047
default_branch: main
tag_name: v1.0.1
ddev_version_constraint: ">= v1.23.5"
dependencies: ["trebormc/ddev-playwright-mcp", "trebormc/ddev-beads", "trebormc/ddev-agents-sync"]
type: contrib
created_at: 2026-03-25
updated_at: 2026-03-25
workflow_status: unknown
stars: 0
---

[![tests](https://github.com/trebormc/ddev-opencode/actions/workflows/tests.yml/badge.svg)](https://github.com/trebormc/ddev-opencode/actions/workflows/tests.yml)

# ddev-opencode

A DDEV add-on that runs [OpenCode](https://github.com/opencode-ai/opencode) in a dedicated container for AI-powered Drupal development.

Agents, rules, and skills for Drupal development are automatically synced from [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) via [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) -- no manual git clone needed.

## Quick Start

```bash
# 1. Install the add-on
ddev add-on get trebormc/ddev-opencode

# 2. Set up your API credentials
mkdir -p ~/opencode-auth
cp share/auth.json.example ~/opencode-auth/auth.json
# Edit ~/opencode-auth/auth.json with your API tokens

# 3. Point to your auth directory
ddev dotenv set .ddev/.env.opencode \
  --host-opencode-auth-dir="$HOME/opencode-auth/"

# 4. Restart DDEV
ddev restart

# 5. Launch OpenCode
ddev opencode
```

## Prerequisites

- [DDEV](https://ddev.readthedocs.io/) >= v1.23.5
- An API key (Anthropic, OpenAI, or a LiteLLM proxy)

## Installation

```bash
ddev add-on get trebormc/ddev-opencode
ddev restart
```

This automatically installs all dependencies:
- [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) -- auto-syncs AI agents from git
- [ddev-beads](https://github.com/trebormc/ddev-beads) -- task tracking
- [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) -- browser automation

## Authentication

Create an `auth.json` file in the directory pointed to by `HOST_OPENCODE_AUTH_DIR`:

```json
{
  "anthropic": {
    "type": "oauth",
    "access": "sk-ant-oat01-YOUR_ACCESS_TOKEN",
    "refresh": "sk-ant-ort01-YOUR_REFRESH_TOKEN",
    "expires": 0
  }
}
```

Credentials are stored in a shared directory on the host (`~/opencode-auth/` by default), so you only need to configure them **once** -- all your DDEV projects share the same credentials automatically.

## Configuration

After installation, environment variables are in `.ddev/.env.opencode`:

```bash
# Directory containing auth.json with your API keys.
# Shared across ALL DDEV projects. Change only if you need a custom location.
HOST_OPENCODE_AUTH_DIR=${HOME}/opencode-auth/

# Timezone
TZ=UTC
```

### Agent Configuration

Agents, rules, and skills are automatically synced from [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) via [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync). No manual setup is needed.

The sync process resolves model tokens (like `${MODEL_CHEAP}`) to OpenCode model names (like `anthropic/claude-haiku-4-5`) using the `.env.agents` file from the agent repository. See [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) for details on model tokens and customization.

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

If you want to use a custom `opencode.json` (e.g., to add a LiteLLM proxy or change global permissions), set `HOST_OPENCODE_CONFIG_DIR` in `.ddev/.env.opencode`:

```bash
HOST_OPENCODE_CONFIG_DIR=${HOME}/my-opencode-config/
```

Place your `opencode.json` in that directory. It takes precedence over the default from the agent repository. Agents, rules, and skills are always loaded from the synced volume regardless of this setting.

## Architecture

```
┌─────────────────────────────────────────────────┐
│              DDEV Docker Network                 │
│                                                  │
│  ┌──────────────┐  docker exec  ┌────────────┐  │
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

OpenCode communicates with the web container via `docker exec` (through the mounted Docker socket), giving it full CLI access to drush, composer, phpunit, phpstan, and any other tool in the web container. Playwright MCP is accessed over HTTP for browser automation and visual testing.

## Commands

| Command | Description |
|---------|-------------|
| `ddev opencode` | Launch the OpenCode TUI |
| `ddev opencode tui [title]` | Launch TUI with a custom tab title |
| `ddev opencode shell` | Open a bash shell in the container |
| `ddev opencode <command>` | Run any command in the container |

### Shell Helpers

Inside the container (via `ddev opencode shell`), these helper functions are available:

| Helper | Description |
|--------|-------------|
| `drush` | Run drush commands in the web container |
| `composer` | Run composer in the web container |
| `web-exec` | Execute any command in the web container |
| `web-shell` | Open an interactive shell in the web container |
| `bd` | Run Beads task tracking commands |

## Desktop Notifications

OpenCode can send desktop notifications when tasks complete or need attention. If you use [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents), the notification configuration is included by default.

To receive notifications, start the notification bridge on your host:

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
| [ddev-claude-code](https://github.com/trebormc/ddev-claude-code) | [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI container for interactive development. | Alternative AI tool |
| [ddev-ralph](https://github.com/trebormc/ddev-ralph) | Autonomous AI task orchestrator. Delegates work to this container or Claude Code via `docker exec`. | Uses this as backend |
| [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) | Headless Playwright browser for browser automation and visual testing. | Auto-installed dependency |
| [ddev-beads](https://github.com/trebormc/ddev-beads) | [Beads](https://github.com/steveyegge/beads) git-backed task tracker shared by all AI containers. | Auto-installed dependency |
| [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) | Auto-syncs AI agent repositories into a shared Docker volume. Provides agents and config. | Auto-installed dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 13 agents, 4 rules, 14 skills for Drupal development. Synced automatically via ddev-agents-sync. | Agent configuration |

## Disclaimer

This project is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors -- always review changes before deploying to production. See [menetray.com](https://menetray.com) for more information and [DruScan](https://druscan.com) for Drupal auditing tools.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-opencode/blob/main/LICENSE).
