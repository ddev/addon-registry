---
title: augustash/ddev-pantheon-db
github_url: https://github.com/augustash/ddev-pantheon-db
description: "Faster database download from Pantheon for Drupal and Wordpress."
user: augustash
repo: ddev-pantheon-db
repo_id: 1193784869
default_branch: master
tag_name: 1.0.3
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-03-27
updated_at: 2026-03-27
workflow_status: unknown
stars: 0
---

# ddev-pantheon-db

A DDEV add-on that provides a fast Pantheon database pull using `terminus ldb` for live environments.

## Installation

```bash
ddev add-on get augustash/ddev-pantheon-db
```

If using [augustash/ddev-drupal](https://github.com/augustash/ddev-drupal) or [augustash/ddev-wordpress](https://github.com/augustash/ddev-wordpress), this add-on is installed automatically on `ddev start`.

## Requirements

- `TERMINUS_MACHINE_TOKEN` set in `~/.ddev/global_config.yaml`
- `DDEV_PANTHEON_SITE` and `DDEV_PANTHEON_ENVIRONMENT` set in `.ddev/config.yaml`

## Usage

```bash
ddev db        # Pull database if local db is empty or has only one table
ddev db -f     # Force a fresh database pull
```

For Drupal projects, `ddev db` will also run `composer install`, `drush cr`, `drush cim`, and `drush updb` after the pull.

## How it works

- **Live environments**: Uses `terminus ldb` for a fast local database copy
- **Other environments**: Uses `terminus connection:info` with `mysqldump`
