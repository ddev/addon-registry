---
title: kgaut/ddev-drupal-tools
github_url: https://github.com/kgaut/ddev-drupal-tools
description: "DDEV add-on: global host commands for Drupal projects — db-import/db-export and db-{prod,preprod}-{dump,get,import} / ssh-{prod,preprod}"
user: kgaut
repo: ddev-drupal-tools
repo_id: 1316435920
default_branch: main
tag_name: v0.1.0
ddev_version_constraint: ">= v1.25.0"
dependencies: []
type: contrib
created_at: 2026-07-29
updated_at: 2026-07-29
workflow_status: unknown
stars: 0
---

# ddev-drupal-tools

> 🇫🇷 [Version française](https://github.com/kgaut/ddev-drupal-tools/blob/main/README.fr.md)

**Global** [DDEV](https://ddev.com/) add-on for Drupal projects: manage database dumps locally
and pull them from your production / staging servers.

The commands are installed into the global DDEV directory (`~/.ddev/commands/host/`) and are
therefore available in **every** DDEV project on the machine.

Detailed write-ups (in French) on kgaut.net:

- [db-import / db-export: two global DDEV commands to manage dumps](https://kgaut.net/blog/2026/db-import-db-export-deux-commandes-ddev-globales-pour-gerer-ses-dumps)
- [db-prod-* and ssh-prod: global DDEV commands to pull your production database](https://kgaut.net/blog/2026/db-prod-et-ssh-prod-des-commandes-ddev-globales-pour-rapatrier-sa-base-de-production)

## Installation

```bash
ddev add-on get kgaut/ddev-drupal-tools
```

Or from a local clone:

```bash
ddev add-on get /path/to/ddev-drupal-tools
```

> **Note:** even though the add-on installs its files globally, `ddev add-on get` must run
> in the context of a DDEV project: run it from a project directory, or add
> `--project <name>`. Outside a project it fails with "could not find a project".

DDEV has no real notion of a "global add-on" (open feature request:
[ddev/ddev#6145](https://github.com/ddev/ddev/issues/6145)): `global_files` only controls
where the **files** are copied, while the installation itself (name, version, file list) is
recorded **in the project** used as context, in `.ddev/addon-metadata/drupal-tools/manifest.yaml`.
Don't delete that file — DDEV needs it for updates and removal — but keep it out of version
control by ignoring `.ddev/addon-metadata/` in the project's root `.gitignore`.

### Update

Run the install command again **from the project that was used to install the add-on**
(that's where DDEV recorded the installation):

```bash
ddev add-on get kgaut/ddev-drupal-tools
```

Files carrying the `#ddev-generated` marker are replaced with the new version. Running the
command from another project works too, but leaves a second manifest in that project —
better to always anchor on the same one.

### Removal

From the same project:

```bash
ddev add-on remove drupal-tools
```

Because of the per-project tracking described above, `ddev add-on remove` and
`ddev add-on list --installed` only see the add-on from the project holding the manifest.

## Commands

### Local

| Command | Description |
| --- | --- |
| `ddev db-import [dump]` | Drops the database, imports a dump (most recent one from the dumps directory by default), then runs `drush deploy`, `drush cr`, `drush uli`. Flags: `-l` (list dumps), `-n` (dry-run), `-y` (skip confirmation). |
| `ddev db-export` | Clears caches then exports the database to `<dir>/<date>-<project>-dev.sql.gz`. Flags: `--no-gzip`, `--no-cr`, `-n`. |

Formats supported by `db-import`: `.sql`, `.sql.gz`, `.sql.bz2`, `.sql.xz`, `.mysql`, `.mysql.gz`, `.zip`, `.tgz`, `.tar.gz`.
Dump names are tab-completed (`ddev db-import <tab>`), most recent first.

### Remote server (production / staging)

| Command | Description |
| --- | --- |
| `ddev db-prod-dump` | Runs `drush sql-dump --gzip` on the server, into a timestamped file in `PROD_DB_PATH` (the dump stays on the server). |
| `ddev db-prod-get` | Downloads the most recent remote dump into the local dumps directory. |
| `ddev db-prod-import` | Chains `db-prod-get` + `db-import`. |
| `ddev ssh-prod` | Opens an SSH session on the current project's production server. |

Each command has its `preprod` (staging) twin: `db-preprod-dump`, `db-preprod-get`,
`db-preprod-import`, `ssh-preprod` — same files, `PREPROD_` variable prefix.

## Configuration

Everything is configured in the `.env` file at each project's root (never sourced: variables
are extracted with `grep`). The commands are visible everywhere but only useful in configured
projects — a missing variable produces an explicit error.

```bash
# project .env
PROD_USER=kevin
PROD_HOST=my-server.example.org
PROD_PORT=22                           # optional, defaults to 22
PROD_PATH=/var/www/myproject           # project root on the server
PROD_DRUSH=vendor/bin/drush            # drush binary, relative to PROD_PATH
PROD_DB_PATH=/var/www/myproject/dumps  # dumps directory, on the server
PROD_URL=myproject.example.org         # used to name dump files

# same idea for staging, with the PREPROD_ prefix
```

### Local dumps directory

Shared by all commands. Defaults to `files/dumps` (relative to the project root), overridable
with `DB_DUMP_DIR` (project `.env`, then `.ddev/.env`), as a relative or absolute path. The
`db-{prod,preprod}-get` commands also honor `LOCAL_DB_PATH`, which takes precedence over
`DB_DUMP_DIR`.

## Tests

```bash
bats tests
```

Tests run with an isolated `HOME`: they never touch the machine's real `~/.ddev`.

## License

MIT.
