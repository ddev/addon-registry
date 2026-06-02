---
title: augustash/ddev-pantheon-db
github_url: https://github.com/augustash/ddev-pantheon-db
description: "Faster database download from Pantheon for Drupal and Wordpress."
user: augustash
repo: ddev-pantheon-db
repo_id: 1193784869
default_branch: master
tag_name: 1.0.4
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-03-27
updated_at: 2026-06-01
workflow_status: unknown
stars: 0
---

# ddev-pantheon-db

A DDEV add-on that provides a fast Pantheon database pull from any environment using Terminus backups.

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

### Pulling from a different environment

The pull targets the environment set in `DDEV_PANTHEON_ENVIRONMENT`. To pull
from a different environment for a single run, use the `-e` flag:

```bash
ddev db -f -e=dev
ddev db -f -e=test
ddev db -f -e=pr-123   # a multidev
```

### No environment configured

There is **no default environment** — this is deliberate, so a misconfigured or
brand-new site never silently pulls production. If `DDEV_PANTHEON_ENVIRONMENT`
is unset (and no `-e` is given), `ddev db` skips the pull with a notice.

## How it works

Creates a fresh database backup of the target environment with
`terminus backup:create` and downloads it with `terminus backup:get`. This is
the same fast, backup-based transfer for every environment (`live`, `dev`,
`test`, or any multidev) — there is no separate live-only code path.
