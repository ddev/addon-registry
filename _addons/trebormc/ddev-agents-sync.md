---
title: trebormc/ddev-agents-sync
github_url: https://github.com/trebormc/ddev-agents-sync
description: "DDEV Add-on: Syncs shared AI agent configurations across containers"
user: trebormc
repo: ddev-agents-sync
repo_id: 1191720964
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

[![tests](https://github.com/trebormc/ddev-agents-sync/actions/workflows/tests.yml/badge.svg)](https://github.com/trebormc/ddev-agents-sync/actions/workflows/tests.yml)

# ddev-agents-sync

A DDEV add-on that automatically syncs AI agent repositories and generates tool-specific configurations for [OpenCode](https://github.com/trebormc/ddev-opencode) and [Claude Code](https://github.com/trebormc/ddev-claude-code).

On every `ddev start`, this container clones or updates the configured repositories, resolves model aliases, and produces two separate agent directories -- one optimized for each AI tool.

## Quick Start

```bash
# Install the add-on (default: syncs trebormc/drupal-ai-agents)
ddev add-on get trebormc/ddev-agents-sync
ddev restart
```

That's it. OpenCode and Claude Code will automatically pick up the synced agents.

## Prerequisites

- [DDEV](https://ddev.readthedocs.io/) >= v1.23.5

## Installation

```bash
ddev add-on get trebormc/ddev-agents-sync
```

This add-on is automatically installed as a dependency of [ddev-opencode](https://github.com/trebormc/ddev-opencode) and [ddev-claude-code](https://github.com/trebormc/ddev-claude-code).

## Configuration

Edit `.ddev/.env.agents-sync`:

```bash
# Comma-separated list of git repositories to sync
# Later repos override earlier ones (useful for private overrides)
AGENTS_REPOS=https://github.com/trebormc/drupal-ai-agents.git

# Set to "false" to disable automatic sync on ddev start
AGENTS_AUTO_UPDATE=true
```

### Multiple repos (public + private override)

```bash
AGENTS_REPOS=https://github.com/trebormc/drupal-ai-agents.git,https://github.com/your-org/private-agents.git
```

Files from later repos override earlier ones. This lets you use the public Drupal agents as a base and add (or replace) specific agents/skills from a private repo.

## How It Works

```
On ddev start:
  ┌──────────────────────────────────────────────────────────────────┐
  │  agents-sync container                                          │
  │                                                                  │
  │  1. Clone/update each repo in AGENTS_REPOS                      │
  │                                                                  │
  │  2. Merge all repos into /tmp/agents-merged (later repos win)   │
  │                                                                  │
  │  3. Read .env.agents (model alias → real model name mapping)     │
  │                                                                  │
  │  4. Generate /agents-opencode/                                   │
  │     - envsubst: ${MODEL_CHEAP} → anthropic/claude-haiku-4-5     │
  │     - Keeps OpenCode frontmatter (mode, tools object, permission)│
  │     - Removes allowed_tools line                                 │
  │     - Copies opencode.json.example, notifier config, etc.       │
  │                                                                  │
  │  5. Generate /agents-claude/                                     │
  │     - envsubst: ${MODEL_CHEAP} → haiku                          │
  │     - Converts frontmatter to Claude Code format                 │
  │     - Renames allowed_tools → tools (CSV)                        │
  │     - Removes mode, temperature, permission blocks               │
  │                                                                  │
  │  6. Sleep (stay alive for depends_on)                            │
  └──────────────────────────────────────────────────────────────────┘

  ┌─────────────────────────┐  ┌─────────────────────────┐
  │       OpenCode          │  │      Claude Code        │
  │  reads /agents-opencode │  │  reads /agents-claude   │
  │      (read-only)        │  │      (read-only)        │
  └─────────────────────────┘  └─────────────────────────┘
```

The output directories are Docker named volumes (`ddev-{sitename}-agents-opencode` and `ddev-{sitename}-agents-claude`). They persist between restarts and are not visible on the host filesystem.

### Merge strategy

Files are copied from each repo in order. The merge is a simple file-level override:

- `agent/*.md` -- merged (later repos can add or override agents)
- `skills/*/SKILL.md` -- merged (later repos can add or override skills)
- `rules/*.md` -- merged (later repos can add or override rules)
- `.env.agents` -- overridden by the last repo that provides it
- `CLAUDE.md`, `opencode.json.example` -- overridden by the last repo

## Model Token System

Agent `.md` files use **model tokens** instead of hardcoded model names. This allows the same agent definition to work with both OpenCode and Claude Code, and makes it easy to change models globally.

### Available tokens

| Token | Default (OpenCode) | Default (Claude Code) | Use for |
|-------|--------------------|-----------------------|---------|
| `${MODEL_SMART}` | `anthropic/claude-opus-4-6` | `opus` | Quality gates, planning, research |
| `${MODEL_NORMAL}` | `anthropic/claude-sonnet-4-5` | `sonnet` | General-purpose tasks |
| `${MODEL_CHEAP}` | `anthropic/claude-haiku-4-5` | `haiku` | Fast, cost-effective agents |
| `${MODEL_APPLIER}` | `anthropic/claude-haiku-4-5` | `haiku` | Mechanical code application |

### How tokens are resolved

The `.env.agents` file in the agent repository defines the mapping:

```bash
# OpenCode models (provider/model-id format)
OC_MODEL_SMART=anthropic/claude-opus-4-6
OC_MODEL_NORMAL=anthropic/claude-sonnet-4-5
OC_MODEL_CHEAP=anthropic/claude-haiku-4-5
OC_MODEL_APPLIER=anthropic/claude-haiku-4-5

# Claude Code models (native aliases)
CC_MODEL_SMART=opus
CC_MODEL_NORMAL=sonnet
CC_MODEL_CHEAP=haiku
CC_MODEL_APPLIER=haiku
```

During sync, `envsubst` replaces the tokens with the appropriate values for each tool.

### Changing models

To change which models your agents use:

1. **For all projects**: Fork [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents), edit `.env.agents`, and point `AGENTS_REPOS` to your fork.

2. **Per project**: Create a private repo with just an `.env.agents` file and add it as a second repo:
   ```bash
   AGENTS_REPOS=https://github.com/trebormc/drupal-ai-agents.git,https://github.com/your-org/my-model-config.git
   ```
   The `.env.agents` from your repo will override the public one.

### Writing agents with tokens

If you create custom agents in your own repository, use the same tokens in the frontmatter:

```yaml
---
description: My custom agent for code review.
model: ${MODEL_SMART}
mode: subagent
tools:
  read: true
  glob: true
  grep: true
  bash: false
permission:
  bash: deny
allowed_tools: Read, Glob, Grep
---

Your agent prompt here...
```

The sync script will automatically substitute the tokens when generating configs for each tool. See [Fat Frontmatter](#fat-frontmatter) for the full format.

### Fat Frontmatter

Agent `.md` files use a "fat frontmatter" format that contains configuration for **both** OpenCode and Claude Code. Each tool reads the fields it understands and ignores the rest:

```yaml
---
description: Short description of what this agent does.
model: ${MODEL_CHEAP}                  # Token, replaced by sync

# OpenCode fields (Claude Code ignores these)
mode: subagent                          # primary or subagent
temperature: 0.1                        # optional
tools:                                  # tool availability (YAML object)
  read: true
  glob: true
  grep: true
  bash: false
  write: false
  edit: false
permission:                             # permission policy
  bash: deny

# Claude Code field (OpenCode ignores this, sync renames to "tools:")
allowed_tools: Read, Glob, Grep         # tool availability (CSV)
---

Agent system prompt content...
```

During sync:
- **For OpenCode**: the `allowed_tools:` line is removed. Everything else stays.
- **For Claude Code**: `mode:`, `temperature:`, `tools:` (object), and `permission:` are removed. `allowed_tools:` is renamed to `tools:`.

## Commands

### `ddev agents-update`

Manually trigger a sync without restarting DDEV:

```bash
ddev agents-update
```

## Part of DDEV AI Workspace

This add-on is part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), a modular ecosystem of DDEV add-ons for AI-powered Drupal development.

| Repository | Description | Relationship |
|------------|-------------|--------------|
| [ddev-ai-workspace](https://github.com/trebormc/ddev-ai-workspace) | Meta add-on that installs the full AI development stack with one command. | Workspace |
| [ddev-opencode](https://github.com/trebormc/ddev-opencode) | [OpenCode](https://opencode.ai) AI CLI container for interactive development. | Auto-installs this add-on |
| [ddev-claude-code](https://github.com/trebormc/ddev-claude-code) | [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI container for interactive development. | Auto-installs this add-on |
| [ddev-ralph](https://github.com/trebormc/ddev-ralph) | Autonomous AI task orchestrator. Delegates work to OpenCode or Claude Code. | Does not require this add-on |
| [ddev-beads](https://github.com/trebormc/ddev-beads) | [Beads](https://github.com/steveyegge/beads) git-backed task tracker shared by all AI containers. | Sibling dependency |
| [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) | Headless Playwright browser for browser automation and visual testing. | Sibling dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 13 agents, 4 rules, 14 skills for Drupal development. Default repo synced by this add-on. | Content synced by this add-on |

## Disclaimer

This project is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors -- always review changes before deploying to production. See [menetray.com](https://menetray.com) for more information and [DruScan](https://druscan.com) for Drupal auditing tools.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-agents-sync/blob/main/LICENSE).
