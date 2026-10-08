---
title: "weitzman/ddev-laravelcloud"
github_url: "https://github.com/weitzman/ddev-laravelcloud"
description: "Laravel Cloud provider: copies a Prod database into DDEV"
user: "weitzman"
repo: "ddev-laravelcloud"
repo_id: 1404952163
default_branch: "main"
tag_name: "1.0.0-rc1"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-10-05"
updated_at: "2026-10-05"
workflow_status: "failure"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/weitzman/ddev-laravelcloud/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/weitzman/ddev-laravelcloud/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/weitzman/ddev-laravelcloud)](https://github.com/weitzman/ddev-laravelcloud/commits)
[![release](https://img.shields.io/github/v/release/weitzman/ddev-laravelcloud)](https://github.com/weitzman/ddev-laravelcloud/releases/latest)

# DDEV Laravel Cloud

A [DDEV](https://ddev.com/) pull provider for Drupal sites hosted on [Laravel Cloud](https://cloud.laravel.com/). `ddev pull laravelcloud` downloads the database backup that the [laravelcloud](https://www.drupal.org/project/laravelcloud) module uploads on each deploy and imports it.

## Requirements

- The `drupal/laravelcloud` module, installed with Composer, with its database backup configured (see "Database backup" in the module's README).
- `DB_BACKUP_URL` in the web container, in the form `https://ACCESS_KEY_ID:SECRET_ACCESS_KEY@ENDPOINT_HOST/BUCKET_ID`. Use the bucket's read-only key. Set it in `.ddev/.env` (then `ddev restart`) or in the project's `.env`. Do not commit it.

## Installation

```bash
ddev add-on get weitzman/ddev-laravelcloud
```

Commit the `.ddev` directory afterwards. The add-on installs `.ddev/providers/laravelcloud.yaml` and `.ddev/laravelcloud/db-backup-download`.

## Usage

```bash
ddev pull laravelcloud
```

## Notes

- The backup is not sanitized. Its cache, session and log tables are empty.
- Files are not pulled. Use [stage_file_proxy](https://www.drupal.org/project/stage_file_proxy) instead.

## Credits

**Contributed and maintained by [@weitzman](https://github.com/weitzman)**
