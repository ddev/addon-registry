---
title: "BODYA7979/ddev-worktree-kit"
github_url: "https://github.com/BODYA7979/ddev-worktree-kit"
description: "DDEV + Git worktrees + AI agents. Isolated databases, automatic Codex/Claude setup and cleanup, and Drupal multisite."
user: "BODYA7979"
repo: "ddev-worktree-kit"
repo_id: 1444581799
default_branch: "main"
tag_name: "v0.5.0"
ddev_version_constraint: ">= v1.25.4"
dependencies: []
type: "contrib"
created_at: "2026-11-03"
updated_at: "2026-11-03"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/BODYA7979/ddev-worktree-kit/actions/workflows/tests.yml/badge.svg)](https://github.com/BODYA7979/ddev-worktree-kit/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/BODYA7979/ddev-worktree-kit)](https://github.com/BODYA7979/ddev-worktree-kit/commits/main)
[![release](https://img.shields.io/github/v/release/BODYA7979/ddev-worktree-kit)](https://github.com/BODYA7979/ddev-worktree-kit/releases)

# DDEV Worktree Kit

**Isolated DDEV environments for Codex and Claude Code Git worktrees.** Unique URLs, independent databases seeded from your main checkout, Composer/setup commands, Drupal multisite and safe cleanup.

[English documentation](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/index.md) · [Українська](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/README.uk.md) · [Documentation site](https://bodya7979.github.io/ddev-worktree-kit/)

## Requirements

macOS/Linux, **Python 3.10+**, **Git 2.36+**, **DDEV 1.25.4+ stable** and a running local Docker provider with an accessible `docker` CLI. Current database support is MySQL/MariaDB. Host PHP, Composer and Drush are unnecessary; the YAML parser is bundled. Native Windows is unsupported; WSL2 has not been tested end to end. [Full requirements](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/requirements.md).

## Install

Paste this into Codex or Claude Code from the project's repository:

```text
Set up DDEV Worktree Kit here for Codex Desktop and Claude Code.
Read and follow this runbook:
https://raw.githubusercontent.com/BODYA7979/ddev-worktree-kit/main/AGENT_SETUP.md

Install the add-on, merge project/client settings and agent instructions,
and enable automatic setup and cleanup for agent-created worktrees.
Preserve data, unrelated settings and the existing file strategy;
otherwise leave uploads unsynchronized.
Verify setup and report changed files. Do not reset/reseed a database or
create/remove a worktree without my explicit request for that test.
```

Name just one client in the prompt if that is all you use. The agent can install the add-on, read its installed boilerplates and edit the project-scoped configurations itself. The first release must exist for the GitHub install command; before publication, give the agent the local source checkout explicitly.


See [agent-led setup](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/agent-setup.md) for the full workflow.

### Command reference

After the first release is published:

```sh
ddev add-on get BODYA7979/ddev-worktree-kit
```

For local development before publication:

```sh
ddev add-on get /absolute/path/to/ddev-worktree-kit
cp /absolute/path/to/ddev-worktree-kit/examples/worktree-kit.yaml .ddev/worktree-kit.yaml
```

Review and commit `.ddev/worktree-kit/`, `.ddev/commands/host/worktree-kit`, `.ddev/worktree-kit.yaml` and `.ddev/config.yaml`. Project configuration is YAML. Native agent integration is part of installation: [Codex/Claude setup](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/integrations.md).

## Use

`wtk` is DDEV's native alias for `worktree-kit`:

```sh
ddev wtk init
ddev wtk status
ddev wtk files sync          # preview
ddev wtk files sync --apply  # source → copy; preserve extra target files
ddev wtk teardown            # preview DDEV removal; preserve Git
ddev wtk gc                  # preview orphan environments
```

Codex owns Git creation/removal and invokes setup/cleanup callbacks. Claude uses paired native `WorktreeCreate` / `WorktreeRemove` hooks that own their Git checkout and DDEV lifecycle; branches are preserved. No background watcher is needed.

```yaml
setup:
  composer_install: true
  commands:
    - "ddev drush cr"
files:
  strategy: none
  mappings: []
cleanup:
  orphan_detection: true
  auto_delete: false
```

For full Drupal setup, choose `ddev drush deploy -y`. Uploads default to `none`; explicit `shared` and `copy` mappings are available. Automatic DDEV/DB cleanup requires `cleanup.auto_delete: true`. Init never resets an existing database. [Configuration](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/configuration.md) · [Setup commands](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/setup-commands.md) · [Files](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/files.md) · [Cleanup](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/cleanup.md).

## Contribute and publish

```sh
make check      # Ruff, syntax and real-Git contracts
make test-addon # Bats with disposable real DDEV fixtures
make docs       # English + Ukrainian static site; docs dependencies required
make package    # add-on tarball and SHA-256 under dist/
```

[Development](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/development.md) explains separate test/docs environments. [Publishing](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/docs/en/publishing.md) covers GitHub Pages, checked tag-to-release automation, `ddev-get` registry discovery and the separate request for official DDEV maintenance. Workflows do not run until the source is pushed; badges and the site become available after publication.

Runtime: [MIT](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/LICENSE). Template scaffolding and bundled software: [third-party notices](https://github.com/BODYA7979/ddev-worktree-kit/blob/main/THIRD_PARTY_NOTICES.md).
