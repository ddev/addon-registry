---
title: trebormc/ddev-ralph
github_url: https://github.com/trebormc/ddev-ralph
description: "DDEV Add-on: Autonomous task orchestrator"
user: trebormc
repo: ddev-ralph
repo_id: 1191701553
default_branch: main
tag_name: v1.0.37
ddev_version_constraint: ">= v1.23.5"
dependencies: ["trebormc/ddev-playwright-mcp", "trebormc/ddev-beads"]
type: contrib
created_at: 2026-03-25
updated_at: 2026-04-03
workflow_status: disabled
stars: 0
---

[![tests](https://github.com/trebormc/ddev-ralph/actions/workflows/tests.yml/badge.svg)](https://github.com/trebormc/ddev-ralph/actions/workflows/tests.yml)

# ddev-ralph

A DDEV add-on that provides **Ralph Loop**, an autonomous AI task orchestrator for Drupal development. Ralph delegates work to [OpenCode](https://opencode.ai) or [Claude Code](https://docs.anthropic.com/en/docs/claude-code) containers via `docker exec`, running AI agents in a loop with [Beads](https://github.com/steveyegge/beads) task tracking -- perfect for overnight unattended execution.

**Important:** Ralph is a lightweight orchestrator. It does **not** include AI tools itself. You must have at least one backend installed first:
- [ddev-opencode](https://github.com/trebormc/ddev-opencode) for the OpenCode backend
- [ddev-claude-code](https://github.com/trebormc/ddev-claude-code) for the Claude Code backend

## Quick Start

```bash
# 1. Install a backend (pick one or both)
ddev add-on get trebormc/ddev-opencode
# or
ddev add-on get trebormc/ddev-claude-code

# 2. Install Ralph
ddev add-on get trebormc/ddev-ralph
ddev restart

# 3. Create a requirements file
cat > requirements.md << 'EOF'
# My Feature

## Objective
Create a custom Drupal module that does X.

## Requirements
- Feature A with admin form
- Feature B with tests

## Success Criteria
- PHPUnit tests passing
- PHPStan level 8 clean
- PHPCS no errors
EOF

# 4. Run Ralph
ddev ralph --backend opencode
```

## Installation

```bash
ddev add-on get trebormc/ddev-ralph
ddev restart
```

This automatically installs [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) (browser automation) and [ddev-beads](https://github.com/trebormc/ddev-beads) (task tracking) as dependencies.

## Usage

```bash
# Run with OpenCode backend
ddev ralph --backend opencode

# Run with Claude Code backend
ddev ralph --backend claude

# Custom requirements file
ddev ralph --backend opencode --prompt my-tasks.md

# Resume existing tasks (skip the planning phase)
ddev ralph --backend claude --no-replan

# Override model
ddev ralph --backend opencode -m anthropic/claude-sonnet-4-5

# Maximum speed overnight run
ddev ralph --backend claude -i 500 -d 0

# Open shell in Ralph container
ddev ralph shell
```

## How It Works

```
┌──────────────────────────────────────────────────────┐
│  PLANNING PHASE (Iteration 1)                        │
│    Read requirements.md -> Create tasks with bd      │
│    Signal: PLANNING_COMPLETE                         │
│                                                      │
│  EXECUTION PHASE (Iterations 2+)                     │
│    bd ready -> Work on task -> bd close              │
│    Repeat until all tasks done                       │
│                                                      │
│  EXIT: When bd ready returns []                      │
└──────────────────────────────────────────────────────┘
```

1. **Planning:** Ralph reads your `requirements.md` and creates a set of Beads tasks.
2. **Execution:** Each iteration picks the next ready task, delegates it to the AI backend, and closes it when done.
3. **Completion:** When no more tasks remain, Ralph exits.

## Options

| Option | Default | Description |
|--------|---------|-------------|
| `-b, --backend` | `opencode` | AI backend: `opencode` or `claude` |
| `-p, --prompt` | `./requirements.md` | Requirements file path |
| `-m, --model` | *(backend default)* | Model override |
| `-i, --iterations` | `200` | Maximum iterations |
| `-d, --delay` | `1` | Seconds between iterations |
| `--no-replan` | | Resume existing tasks (skip planning) |
| `--no-beads` | | Run without Beads task tracking |

## Requirements File Format

The `requirements.md` file tells Ralph what to build. Structure it with clear objectives, requirements, and success criteria:

```markdown
# Project Title

## Objective
One paragraph describing what needs to be built.

## Requirements
- Specific feature or task (be concrete)
- Another feature with details
- Include technical constraints

## Technical Constraints
- Drupal 10/11 compatible
- PHP 8.1+ with strict types
- Drupal coding standards (PHPCS)
- PHPStan level 8

## Success Criteria
- All PHPUnit tests pass
- PHPCS reports no errors
- PHPStan level 8 clean
- Module enables without errors
```

See `.ddev/ralph-loop/requirements-example.md` for a complete example after installation.

## Audit Fixers

Ralph ships with pre-built requirements files for common code quality fixes. These are ideal for automated cleanup runs:

```bash
# Fix PHPCS coding standard violations
ddev ralph --backend opencode --prompt .ddev/ralph-loop/audit-fixers/fix-phpcs.md

# Fix PHPStan static analysis issues
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/fix-phpstan.md

# Fix Twig template issues
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/fix-twig.md

# Fix code complexity issues
ddev ralph --backend opencode --prompt .ddev/ralph-loop/audit-fixers/fix-complexity.md
```

## Test Generation

Ralph includes prompts for automated test generation across all Drupal test types. Two modes of operation:

### Orchestrator Mode

Analyzes the entire project, detects which test types are missing for each module, and generates them all in a single run:

```bash
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/generate-tests.md
```

The orchestrator creates Beads tasks ordered by priority: Kernel tests (P0), Unit and Functional tests (P1), FunctionalJavascript (P2), Behat/Playwright (P3). It skips test types that do not apply (e.g., no Behat if the project has no `behat.yml`).

### Direct Mode

Run a specific test type when you know exactly what you need:

```bash
# Unit tests — pure PHP logic, mocked dependencies (uses Audit module)
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/generate-unit-tests.md

# Kernel tests — services, entities, DB, config, plugins, hooks (most valuable)
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/generate-kernel-tests.md

# Functional tests — forms, permissions, HTML output
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/generate-functional-tests.md

# FunctionalJavascript tests — AJAX, modals, autocompletes
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/generate-functionaljs-tests.md

# Behat tests — BDD acceptance testing in Gherkin (requires behat.yml)
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/generate-behat-tests.md

# Playwright tests — visual regression, cross-browser, accessibility, E2E
ddev ralph --backend claude --prompt .ddev/ralph-loop/audit-fixers/generate-playwright-tests.md
```

All test generation prompts reference the specialized skills from [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) (`drupal-unit-test`, `drupal-kernel-test`, `drupal-functional-test`, etc.) and follow the `drupal-testing` decision rule for choosing the correct test type per class.

## Architecture

```
┌──────────────────────────────────────────────────────────┐
│                    DDEV Docker Network                    │
│                                                          │
│  ┌─────────────┐                                         │
│  │   Ralph     │  docker exec     ┌──────────────┐      │
│  │ (orchestr.) │─────────────────>│  OpenCode    │      │
│  │             │       OR         │  Container   │      │
│  │ bash + jq   │─────────────────>│──────────────│      │
│  │ docker CLI  │                  │ Claude Code  │      │
│  │             │                  │  Container   │      │
│  └─────────────┘                  └──────────────┘      │
│        │                                │                │
│        │ docker exec                    │ docker exec    │
│        v                                v                │
│   ┌──────────┐                   ┌──────────────┐       │
│   │  Beads   │                   │     Web      │       │
│   │ Container│                   │   (PHP)      │       │
│   │ (.beads) │                   └──────────────┘       │
│   └──────────┘                                          │
│                                  ┌──────────────┐       │
│                                  │  Playwright  │       │
│                                  │     MCP      │       │
│                                  └──────────────┘       │
└──────────────────────────────────────────────────────────┘
```

Ralph only contains bash, jq, and docker CLI. All AI execution happens in the dedicated OpenCode or Claude Code containers, and task tracking runs in the Beads container -- all via `docker exec`. This keeps each component independently maintainable.

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | All tasks completed successfully |
| `1` | Error (missing backend, invalid arguments, runtime failure) |
| `2` | Maximum iterations reached before completion |

## Uninstallation

```bash
ddev add-on remove ddev-ralph
ddev restart
```

## Part of DDEV AI Workspace

This add-on is part of [DDEV AI Workspace](https://github.com/trebormc/ddev-ai-workspace), a modular ecosystem of DDEV add-ons for AI-powered Drupal development.

| Repository | Description | Relationship |
|------------|-------------|--------------|
| [ddev-ai-workspace](https://github.com/trebormc/ddev-ai-workspace) | Meta add-on that installs the full AI development stack with one command. | Workspace |
| [ddev-opencode](https://github.com/trebormc/ddev-opencode) | [OpenCode](https://opencode.ai) AI CLI container for interactive development. | Backend (required, pick one or both) |
| [ddev-claude-code](https://github.com/trebormc/ddev-claude-code) | [Claude Code](https://docs.anthropic.com/en/docs/claude-code) CLI container for interactive development. | Backend (required, pick one or both) |
| [ddev-playwright-mcp](https://github.com/trebormc/ddev-playwright-mcp) | Headless Playwright browser for browser automation and visual testing. | Auto-installed dependency |
| [ddev-beads](https://github.com/trebormc/ddev-beads) | [Beads](https://github.com/steveyegge/beads) git-backed task tracker. Ralph uses it for task planning and progress tracking. | Auto-installed dependency |
| [ddev-agents-sync](https://github.com/trebormc/ddev-agents-sync) | Auto-syncs AI agent repositories into a shared Docker volume. | Not a direct dependency |
| [drupal-ai-agents](https://github.com/trebormc/drupal-ai-agents) | 10 agents, 12 rules, 24 skills for Drupal development. Includes `ralph-planner` and `drupal-test-generator` agents. | Agent configuration |

## Disclaimer

This project is not affiliated with Anthropic, OpenCode, Beads, Playwright, Microsoft, or DDEV. AI-generated code may contain errors -- always review changes before deploying to production. See [menetray.com](https://menetray.com) for more information and [DruScan](https://druscan.com) for Drupal auditing tools.

## License

Apache-2.0. See [LICENSE](https://github.com/trebormc/ddev-ralph/blob/main/LICENSE).
