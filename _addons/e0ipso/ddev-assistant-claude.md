---
title: e0ipso/ddev-assistant-claude
github_url: https://github.com/e0ipso/ddev-assistant-claude
description: "ddev add-on for setting up the Claude Code assistant."
user: e0ipso
repo: ddev-assistant-claude
repo_id: 1132377741
default_branch: main
tag_name: v1.1.2
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-01-11
updated_at: 2026-02-18
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/e0ipso/ddev-assistant-claude/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/e0ipso/ddev-assistant-claude/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/e0ipso/ddev-assistant-claude)](https://github.com/e0ipso/ddev-assistant-claude/commits)
[![release](https://img.shields.io/github/v/release/e0ipso/ddev-assistant-claude)](https://github.com/e0ipso/ddev-assistant-claude/releases/latest)

# DDEV Assistant Claude

## Overview

This add-on integrates Assistant Claude into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get e0ipso/ddev-assistant-claude
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Assistant Claude |
| `ddev logs -s assistant-claude` | Check Assistant Claude logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.assistant-claude --assistant-claude-docker-image="ddev/ddev-utilities:latest"
ddev add-on get e0ipso/ddev-assistant-claude
ddev restart
```

Make sure to commit the `.ddev/.env.assistant-claude` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `ASSISTANT_CLAUDE_DOCKER_IMAGE` | `--assistant-claude-docker-image` | `ddev/ddev-utilities:latest` |

## Credits

**Contributed and maintained by [@e0ipso](https://github.com/e0ipso)**
