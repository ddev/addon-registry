---
title: lexsoft00/ddev-drupal-claude-code
github_url: https://github.com/lexsoft00/ddev-drupal-claude-code
description: "Enhanced Claude Code AI Assistant for **Drupal 11 / PHP 8.3** development with pre-configured plugins, security controls, and Serena MCP integration."
user: lexsoft00
repo: ddev-drupal-claude-code
repo_id: 1133733985
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-01-13
updated_at: 2026-02-08
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/lexsoft00/ddev-drupal-claude-code/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/lexsoft00/ddev-drupal-claude-code/actions/workflows/tests.yml?query=branch%3Amain)
[![release](https://img.shields.io/github/v/release/lexsoft00/ddev-drupal-claude-code)](https://github.com/lexsoft00/ddev-drupal-claude-code/releases/latest)

# DDEV Drupal Claude Code

A [DDEV](https://ddev.com/) add-on that integrates [Claude Code](https://docs.anthropic.com/en/docs/claude-code) with [Serena MCP](https://github.com/oraios/serena) for Drupal 11 development.

## Features

- **Claude Code** running inside your DDEV web container
- **Serena MCP** for semantic code analysis and intelligent editing
- **Drupal 11 optimized** with pre-configured CLAUDE.md and Serena memories
- **Security rules** preventing dangerous operations on core/vendor files
- **Pre-configured settings** with extended thinking and stable updates

## Requirements

- DDEV v1.24.3+

## Installation

```bash
ddev add-on get lexsoft00/ddev-drupal-claude-code
ddev restart
```

## Usage

```bash
# 1. Enter the DDEV container
ddev ssh

# 2. Start Claude Code (first run - login and configure)
claude

# 3. Exit Claude Code (/exit or Ctrl+C)

# 4. Add Serena MCP
claude mcp add serena -- uvx --from git+https://github.com/oraios/serena serena start-mcp-server --context ide-assistant --project "$(pwd)"

# 5. (Optional) Index your project for faster code navigation
uvx --from git+https://github.com/oraios/serena serena project index

# 6. Start Claude Code with Serena
claude
```

Inside Claude Code, you have access to:
- **Serena tools** for code search, symbol navigation, and refactoring
- **Drupal memories** with coding standards, Drush commands, and best practices
- **Context7** for up-to-date library documentation

## What Gets Installed

### DDEV Configuration

- **`config.drupal-claude-code.yaml`** - Post-start hooks that configure Claude Code
- **`.ddev/.claude/`** - Symlinked to `~/.claude` for persisting Claude Code runtime data

### Claude Code Settings (`.ddev/drupal-claude-code/`)

- **`settings.json`** - Extended thinking, stable updates, and statusline
- **`settings.local.json`** - Permission rules protecting core/vendor from edits
- **`statusline.sh`** - Status bar showing git branch and Drupal info

### Serena Configuration (`.ddev/drupal-claude-code/serena/`)

- **`project.yml`** - Project settings for PHP/TypeScript analysis
- **`memories/`** - Drupal knowledge bases:
  - `project-overview.md` - Project structure and patterns
  - `code-style-conventions.md` - Drupal coding standards
  - `commands-reference.md` - Drush and CLI commands
  - `architecture-patterns.md` - Drupal best practices
  - `security-guidelines.md` - Security checklist
  - `task-completion-checklist.md` - Quality gates

### Project Root Files

- **`CLAUDE.md`** - Project-specific instructions for Claude Code
- **`.claudeignore`** - Excludes vendor, node_modules, and generated files from context

## Recommended Plugins

After starting Claude Code, install these recommended plugins using `/plugin`:

| Plugin | Description | Install Command |
|--------|-------------|-----------------|
| **context7** | Up-to-date library documentation | `/plugin install context7@claude-plugins-official` |
| **security-guidance** | Security best practices | `/plugin install security-guidance@claude-plugins-official` |
| **code-simplifier** | Code refactoring assistance | `/plugin install code-simplifier@claude-plugins-official` |
| **frontend-design** | Frontend/UI development | `/plugin install frontend-design@claude-plugins-official` |

Or browse all available plugins interactively:

```bash
/plugin
# Navigate to "Discover" tab and install with Enter
```
