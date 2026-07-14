---
title: trebormc/ddev-agents-sync
github_url: https://github.com/trebormc/ddev-agents-sync
description: "DDEV Add-on: Syncs shared AI agent configurations across containers"
user: trebormc
repo: ddev-agents-sync
repo_id: 1191720964
default_branch: main
tag_name: v1.0.58
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-03-25
updated_at: 2026-07-06
workflow_status: disabled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/trebormc/ddev-agents-sync/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/trebormc/ddev-agents-sync/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/trebormc/ddev-agents-sync)](https://github.com/trebormc/ddev-agents-sync/commits)
[![release](https://img.shields.io/github/v/release/trebormc/ddev-agents-sync)](https://github.com/trebormc/ddev-agents-sync/releases/latest)

# ddev-agents-sync

A DDEV add-on that automatically syncs AI agent repositories and generates tool-specific configurations for [OpenCode](https://github.com/trebormc/ddev-opencode) and [Claude Code](https://github.com/trebormc/ddev-claude-code).

> **Part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace)** — a modular ecosystem of DDEV add-ons for AI-powered Drupal development. Install the full stack with one command: `ddev add-on get trebormc/ddev-ai-workspace`
>
> Created by [Robert Menetray](https://menetray.com) · Sponsored by [DruScan](https://druscan.com)

**What problem does this solve?** AI tools like OpenCode and Claude Code each expect agent configurations in a different format. This add-on lets you write agents once (using a shared "fat frontmatter" format) and automatically generates the correct configuration for each tool. It also resolves model tokens, so the same agent definition can use different models depending on the tool.

On every `ddev start`, this container clones or updates the configured repositories, resolves model aliases, and produces two separate agent directories (one optimized for each AI tool).

## Quick Start

The **recommended way** to install this add-on is through the [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), which installs all tools and dependencies with a single command:

```bash
ddev add-on get trebormc/ddev-ai-workspace
ddev restart
```

This add-on is also **automatically installed** as a dependency when you install [ddev-opencode](https://github.com/trebormc/ddev-opencode) or [ddev-claude-code](https://github.com/trebormc/ddev-claude-code). You rarely need to install it directly.

### Standalone installation

If you need to install it individually (requires familiarity with the DDEV add-on ecosystem):

```bash
ddev add-on get trebormc/ddev-agents-sync
ddev restart
```

OpenCode and Claude Code will automatically pick up the synced agents.

## Prerequisites

- [DDEV](https://ddev.readthedocs.io/) >= v1.24.10

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
  │     - envsubst: ${MODEL_CHEAP} → opencode/gpt-5-nano            │
  │     - Keeps OpenCode frontmatter (mode, tools object, permission)│
  │     - Removes allowed_tools line                                 │
  │     - Copies opencode.json.example, notifier config, etc.       │
  │                                                                  │
  │  5. Generate /agents-claude/                                     │
  │     - envsubst: ${MODEL_CHEAP} → haiku                          │
  │     - Converts frontmatter to Claude Code format                 │
  │     - Renames allowed_tools → tools (CSV)                        │
  │     - Removes mode, temperature, maxSteps, tools, permission     │
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

- `.claude/agents/*.md`: merged (later repos can add or override agents)
- `.claude/skills/*/SKILL.md`: merged (later repos can add or override skills)
- `.claude/rules/*.md`: merged (later repos can add or override rules)
- `.env.agents`: overridden by the last repo that provides it
- `CLAUDE.md`, `opencode.json.example`: overridden by the last repo

## Model Token System

Agent `.md` files use **model tokens** instead of hardcoded model names. This allows the same agent definition to work with both OpenCode and Claude Code, and makes it easy to change models globally.

### Available tokens

| Token | Default (OpenCode) | Default (Claude Code) | Use for |
|-------|--------------------|-----------------------|---------|
| `${MODEL_MAIN}` | `opencode/deepseek-v4-flash` | `sonnet` | The orchestrator: default model of the tool's main conversation loop |
| `${MODEL_GENIUS}` | `opencode/glm-5.2` | `opus` | Hardest tasks: important code reviews, architecture |
| `${MODEL_SMART}` | `opencode/glm-5.2` | `opus` | Quality gates, planning, research, delegation advice |
| `${MODEL_NORMAL}` | `opencode/deepseek-v4-pro` | `sonnet` | General-purpose tasks |
| `${MODEL_CHEAP}` | `opencode/deepseek-v4-flash` | `haiku` | Fast, cost-effective agents |
| `${MODEL_APPLIER}` | `opencode/gpt-5-nano` | `haiku` | Mechanical code application |
| `${MODEL_VISION}` | `opencode/minimax-m3` | `sonnet` | Image/screenshot interpretation — MUST accept image input |

When an `.env.agents` file does not define them, `GENIUS` falls back to the `SMART` value, `VISION` falls back to the `NORMAL` value, and `MAIN` falls back to `CHEAP` on OpenCode / `NORMAL` on Claude Code, so older override files keep working unchanged. `GENIUS` defaults to the same model as `SMART` — override it with a stronger model when you have one. `VISION` must be a vision-capable model or image analysis agents will not see anything.

`${MODEL_MAIN}` is special: besides being substituted in configs, the sync injects it as the DEFAULT model of each tool — the top-level `model` in the generated `opencode.json` and the `model` key in the generated Claude Code `settings.generated.json`. The main conversation loop dominates session cost, so it defaults to a cheap tier; expensive models are reserved for the specialist subagents.

### How tokens are resolved

Token values come from `.env.agents` files, loaded as a variable-level cascade — each level only needs to define the variables it wants to override:

| Priority | Location | Scope |
|---|---|---|
| 1 (highest) | `.ddev/.env.agents` in the project | This project only |
| 2 | `~/.ddev/agents-sync/.env.agents` on the host | ALL your DDEV projects |
| 3 (lowest) | `.env.agents` in the agents repo | Managed defaults |

The repo default defines the full mapping:

```bash
# OpenCode models (provider/model-id format)
OC_MODEL_MAIN=opencode/deepseek-v4-flash
OC_MODEL_GENIUS=opencode/glm-5.2
OC_MODEL_SMART=opencode/glm-5.2
OC_MODEL_NORMAL=opencode/deepseek-v4-pro
OC_MODEL_CHEAP=opencode/deepseek-v4-flash
OC_MODEL_APPLIER=opencode/gpt-5-nano
OC_MODEL_VISION=opencode/minimax-m3

# Claude Code models (native aliases)
CC_MODEL_MAIN=sonnet
CC_MODEL_GENIUS=opus
CC_MODEL_SMART=opus
CC_MODEL_NORMAL=sonnet
CC_MODEL_CHEAP=haiku
CC_MODEL_APPLIER=haiku
CC_MODEL_VISION=sonnet
```

During sync, `envsubst` replaces the tokens with the appropriate values for each tool. Changes take effect after `ddev agents-update && ddev restart`.

### Changing models

To change which models your agents use:

1. **For all projects**: Edit `~/.ddev/agents-sync/.env.agents` on the host (created by the installer) and uncomment/set only the variables you want to change — e.g. point everything to your own LiteLLM models.

2. **Per project**: Edit `.ddev/.env.agents` in the project. It overrides both the host file and the repo defaults, again variable by variable.

3. **Managed defaults for a team**: Fork [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents), edit its `.env.agents`, and point `AGENTS_REPOS` to your fork (or add a private repo with just an `.env.agents` as a second repo in `AGENTS_REPOS` — later repos override earlier ones).

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
- **For Claude Code**: `mode:`, `temperature:`, `maxSteps:`, `tools:` (object), and `permission:` are removed. `allowed_tools:` is renamed to `tools:`.

## Git Access

By default AI agents are **not allowed to run git write commands** — they
present a summary of their changes and you commit manually. This is controlled
**per tool** by two flags in the same `.env.agents` cascade as the model tokens,
so they follow the same `repo default < ~/.ddev/agents-sync/.env.agents < .ddev/.env.agents`
priority.

Each flag holds a **comma-separated list of the tools** the capability is
granted to. Valid tool ids: `opencode`, `claude`. An empty value grants it to
no tool (the default).

| Flag | Default | Allows (for the listed tools) |
|------|---------|--------|
| `GIT_ALLOW_COMMIT` | *(empty)* | `git add`, `git commit` |
| `GIT_ALLOW_OPERATIONS` | *(empty)* | `git push` (non-force), `pull`, `fetch`, `merge`, `rebase`, `checkout`/`switch`, `reset`, `restore`, `stash`, `tag`, `cherry-pick` |

List a tool in a flag to enable that capability there — for example, let both
tools commit locally but only Claude Code push:

```bash
# .ddev/.env.agents (this project) or ~/.ddev/agents-sync/.env.agents (all projects)
GIT_ALLOW_COMMIT=opencode,claude
GIT_ALLOW_OPERATIONS=claude
```

Changes take effect after `ddev agents-update && ddev restart`.

**Always blocked, regardless of the flags:** force-push
(`git push --force`/`-f`/`--force-with-lease`) and remote-branch deletion
(`git push --delete`/`-d`) — nothing may rewrite or destroy remote history.

The flags are enforced in **both tools**, not just documented in prompts:
- **OpenCode** — a tool's resolved flags become `allow`/`deny` values in the
  generated `opencode.json` `permission.bash` map.
- **Claude Code** — the sync generates a `settings.generated.json` with a
  `PreToolUse` hook that denies blocked git commands (Claude Code runs in
  `bypassPermissions` mode, so the hook is the actual enforcement).

The **git-workflow** rule shipped to each tool is generated from that tool's
resolved flags, so the prompt text always matches what is technically allowed.

## Commands

### `ddev agents-update`

Manually trigger a sync without restarting DDEV:

```bash
ddev agents-update
```

## Uninstallation

```bash
ddev add-on remove ddev-agents-sync
ddev restart
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
| [ddev-ai-ssh](https://github.com/trebormc/ddev-ai-ssh) | SSH access to the web container for AI agents. | Sibling dependency |
| [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) | Headless Playwright browser for browser automation and visual testing. | Sibling dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 10 agents, 12 rules, 24 skills for Drupal development. Default repo synced by this add-on. | Content synced by this add-on |

## Disclaimer

This project is an independent initiative by [Robert Menetray](https://menetray.com), sponsored by [DruScan](https://druscan.com). It is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors. Always review changes before deploying to production.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-agents-sync/blob/main/LICENSE).
