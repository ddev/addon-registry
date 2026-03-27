---
title: trebormc/ddev-beads
github_url: https://github.com/trebormc/ddev-beads
description: "DDEV Add-on: Beads git-backed task tracker"
user: trebormc
repo: ddev-beads
repo_id: 1191702569
default_branch: main
tag_name: v1.0.3
ddev_version_constraint: ">= v1.23.5"
dependencies: []
type: contrib
created_at: 2026-03-25
updated_at: 2026-03-26
workflow_status: unknown
stars: 0
---

[![tests](https://github.com/trebormc/ddev-beads/actions/workflows/tests.yml/badge.svg)](https://github.com/trebormc/ddev-beads/actions/workflows/tests.yml)

# ddev-beads

A DDEV add-on that provides [Beads](https://github.com/steveyegge/beads) (bd), a git-backed task tracker for AI agents, in a dedicated container. Other AI containers (OpenCode, Claude Code, Ralph) delegate task tracking to this container via `docker exec`.

## Quick Start

```bash
# Install the add-on
ddev add-on get trebormc/ddev-beads
ddev restart

# Use from host
ddev bd ready
ddev bd create "Implement feature X" -p 1
ddev bd close bd-abc --reason "Done"
```

## Prerequisites

- [DDEV](https://ddev.readthedocs.io/) >= v1.23.5

## Installation

```bash
ddev add-on get trebormc/ddev-beads
```

This add-on is automatically installed as a dependency when you install [ddev-opencode](https://github.com/trebormc/ddev-opencode), [ddev-claude-code](https://github.com/trebormc/ddev-claude-code), or [ddev-ralph](https://github.com/trebormc/ddev-ralph).

## Architecture

```
┌──────────────────────────────────────────────────────────┐
│                    DDEV Docker Network                    │
│                                                          │
│  ┌──────────────┐                                        │
│  │   Beads      │  <-- docker exec from other containers │
│  │  Container   │                                        │
│  │  - bd CLI    │  Shared volume: /var/www/html/.beads/  │
│  │  - node 22   │                                        │
│  └──────────────┘                                        │
│        ^    ^    ^                                        │
│        │    │    │  docker exec $BEADS_CONTAINER bd ...   │
│        │    │    │                                        │
│  ┌─────┘    │    └─────┐                                 │
│  │          │          │                                  │
│  OpenCode  Claude   Ralph                                │
│            Code                                          │
└──────────────────────────────────────────────────────────┘
```

All AI containers access Beads via `docker exec $BEADS_CONTAINER bd <command>`. A wrapper function is installed in each container so that `bd` commands work transparently.

The `.beads/` directory lives in the project root (`/var/www/html/.beads/`), shared across all containers via the project volume.

## Commands

### `ddev bd`

Execute Beads commands:

```bash
ddev bd ready                              # List ready tasks
ddev bd create "Implement login" -p 1      # Create task (P1 priority)
ddev bd update bd-abc --status in_progress # Mark in progress
ddev bd update bd-abc --notes "Working..." # Add progress notes
ddev bd close bd-abc --reason "Done"       # Close task
ddev bd sync                               # Sync state
ddev bd prime                              # Get context
```

### Priority Levels

| Level | Meaning |
|-------|---------|
| P0 | Critical -- blockers, security |
| P1 | High -- important features |
| P2 | Medium -- normal tasks |
| P3 | Low -- nice-to-haves |

## Part of DDEV AI Workspace

This add-on is part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), a modular ecosystem of DDEV add-ons for AI-powered Drupal development.

| Repository | Description | Relationship |
|------------|-------------|--------------|
| [ddev-ai-workspace](https://github.com/trebormc/ddev-ai-workspace) | Meta add-on that installs the full AI development stack with one command. | Workspace |
| [ddev-opencode](https://github.com/trebormc/ddev-opencode) | [OpenCode](https://opencode.ai) AI CLI container for interactive development. | Auto-installs this add-on |
| [ddev-claude-code](https://github.com/trebormc/ddev-claude-code) | [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI container for interactive development. | Auto-installs this add-on |
| [ddev-ralph](https://github.com/trebormc/ddev-ralph) | Autonomous AI task orchestrator. Delegates work to OpenCode or Claude Code. | Auto-installs this add-on |
| [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) | Auto-syncs AI agent repositories into a shared Docker volume. | Sibling dependency |
| [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) | Headless Playwright browser for browser automation and visual testing. | Sibling dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 13 agents, 4 rules, 14 skills for Drupal development. Includes Beads workflow rule. | Uses Beads for task tracking |

## Disclaimer

This project is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors -- always review changes before deploying to production. See [menetray.com](https://menetray.com) for more information and [DruScan](https://druscan.com) for Drupal auditing tools.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-beads/blob/main/LICENSE).
