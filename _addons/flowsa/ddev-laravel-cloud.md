---
title: "flowsa/ddev-laravel-cloud"
github_url: "https://github.com/flowsa/ddev-laravel-cloud"
description: "DDEV add-on: ddev pull laravel-cloud - pull the production database (and Craft asset files) from Laravel Cloud into your local DDEV project"
user: "flowsa"
repo: "ddev-laravel-cloud"
repo_id: 1318472188
default_branch: "main"
tag_name: "v0.9.1"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-07-31"
updated_at: "2026-07-31"
workflow_status: "success"
stars: 0
---

[![tests](https://github.com/flowsa/ddev-laravel-cloud/actions/workflows/tests.yml/badge.svg)](https://github.com/flowsa/ddev-laravel-cloud/actions/workflows/tests.yml)

# ddev-laravel-cloud

`ddev pull laravel-cloud` for [Laravel](https://laravel.com) and
[Craft CMS](https://craftcms.com) apps hosted on
[Laravel Cloud](https://cloud.laravel.com): replaces your local DDEV database
with a fresh production dump (and, for Craft, syncs the production asset
files), using credentials read live from Laravel Cloud. Production is **only
ever read from** – the recipe defines no push commands.

## What it does

- Reads the production database credentials from your Cloud environment at
  pull time via the [`cloud` CLI](https://cloud.laravel.com/docs/cli) – no
  production secrets are stored in your repository. Craft conventions
  (`CRAFT_DB_*`) are detected first, then standard Laravel ones (`DB_*`).
- Dumps the production database (Postgres or MySQL, from inside the web
  container so client versions match), verifies the dump, and imports it.
- Takes a `ddev snapshot` before importing, so every pull is one
  `ddev snapshot restore` away from undo.
- After import, detects the app type: Craft projects get `craft up`
  (re-applying your committed project config, which resets the base URL to
  the local site) plus cache clears; Laravel projects get
  `artisan migrate --force` plus `cache:clear`.
- **Craft projects only:** downloads asset files straight from the
  environment's public asset base URL (`ASSETS_BASE_URL`, e.g. an R2 bucket)
  into your local asset directory – resumable, parallel, skipping files
  already present. Projects without Craft tables skip this step with a note.

## Requirements

- DDEV v1.24.10+
- The `cloud` CLI installed and authenticated on your host machine:
  `composer global require laravel/cloud-cli && cloud login`
  (that's the only host dependency – JSON parsing and the asset downloader
  run inside the web container, using the `jq` and `python3` DDEV ships there)
- For Postgres apps: the local database major version must match production
  (`database: postgres:17` in `.ddev/config.yaml` for a Postgres 17 cluster).
  `pg_dump` refuses to dump a newer server, and a newer dump will not restore
  into an older local server. The pull checks this and fails early with
  instructions if they differ.

## Install

```bash
ddev add-on get flowsa/ddev-laravel-cloud
ddev dotenv set .ddev/.env.laravel-cloud --laravel-cloud-env-id=env-xxxxxxxx
ddev restart
```

Find the environment ID with `cloud env:list`. Commit
`.ddev/.env.laravel-cloud` – it contains no secrets, only identifiers and
options.

## Usage

```bash
ddev pull laravel-cloud                  # database + asset files
ddev pull laravel-cloud --skip-files     # database only
ddev pull laravel-cloud --skip-db        # asset files only
ddev pull laravel-cloud -y               # skip the confirmation prompt
```

The imported database carries **production user accounts** – sign in with
production credentials afterwards.

## Configuration

Set per-project options in `.ddev/.env.laravel-cloud` (via
`ddev dotenv set`), or override once with
`ddev pull laravel-cloud --environment="KEY=value"`:

| Variable | Default | Purpose |
|---|---|---|
| `LARAVEL_CLOUD_ENV_ID` | – (required) | The Cloud environment to pull from |
| `LARAVEL_CLOUD_TABLE_PREFIX` | `craft_` | Craft table prefix, for the asset file list |
| `LARAVEL_CLOUD_ASSET_ROOT` | `public/uploads/files` | Where Craft asset files land locally |
| `LARAVEL_CLOUD_DOWNLOAD_JOBS` | `6` | Parallel asset downloads |
| `LARAVEL_CLOUD_PGSSLMODE` | `require` | sslmode for the production Postgres connection |

The post-pull steps live in `.ddev/config.laravel-cloud.yaml`; edit it
(removing its `#ddev-generated` line) to customise them.

## Notes

- The Craft asset file list is queried from the **local** database, which on
  a full pull has just been refreshed from production, so list and files
  always match. With `--skip-db` the list reflects whatever your local
  database currently holds.
- Old snapshots accumulate; prune them occasionally with
  `ddev snapshot --cleanup`.
- Asset sync for plain Laravel apps (e.g. a bucket sync) is not implemented
  yet – contributions welcome.

## Credits

Contributed and maintained by [Flow Communications](https://www.flowsa.com).
