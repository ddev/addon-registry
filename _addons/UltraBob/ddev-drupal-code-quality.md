---
title: UltraBob/ddev-drupal-code-quality
github_url: https://github.com/UltraBob/ddev-drupal-code-quality
description: ""
user: UltraBob
repo: ddev-drupal-code-quality
repo_id: 1124009554
default_branch: main
tag_name: v0.9.3-beta
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2025-12-28
updated_at: 2026-02-04
workflow_status: success
stars: 0
---

# ddev-drupal-code-quality

[![tests](https://github.com/UltraBob/ddev-drupal-code-quality/actions/workflows/tests.yml/badge.svg)](https://github.com/UltraBob/ddev-drupal-code-quality/actions/workflows/tests.yml)

DDEV add-on that installs local code quality tooling based on Drupal.org GitLab CI template defaults (PHPStan, PHPCS, ESLint,
Stylelint, Prettier, CSpell) for local CLI/IDE usage.

## Installation

```bash
ddev add-on get UltraBob/ddev-drupal-code-quality
```

See `README_ADDON.md` for complete usage documentation.

## Repository structure

- `commands/`: DDEV web commands copied into the project `.ddev/commands` directory.
- `drupal-code-quality/`: project-root configs and `.ddev` shims copied by the installer.
- `dcq-install.sh`: conflict-aware installer invoked by `install.yaml`.
- `install.yaml`: DDEV add-on install definition.
