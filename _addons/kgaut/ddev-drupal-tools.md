---
title: "kgaut/ddev-drupal-tools"
github_url: "https://github.com/kgaut/ddev-drupal-tools"
description: "DDEV add-on: global host commands for Drupal projects — db-import/db-export and db-{prod,preprod}-{dump,get,import} / ssh-{prod,preprod}"
user: "kgaut"
repo: "ddev-drupal-tools"
repo_id: 1317094150
default_branch: "main"
tag_name: "v0.4.1"
ddev_version_constraint: ">= v1.25.0"
dependencies: []
type: "contrib"
created_at: "2026-07-30"
updated_at: "2026-09-17"
workflow_status: "unknown"
stars: 0
---

# ddev-drupal-tools

> 🇫🇷 [Version française](https://github.com/kgaut/ddev-drupal-tools/blob/main/README.fr.md)

**Global** [DDEV](https://ddev.com/) add-on for Drupal projects: manage database dumps locally
and pull them from your production / staging servers. The local commands also fit Symfony and
other project types (see [Project types](#project-types)).

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
control by ignoring `.ddev/addon-metadata/drupal-tools/manifest.yaml` in the project's root
`.gitignore` (ignore just that file, not the whole `addon-metadata/` directory: manifests of
project-level add-ons are meant to be committed).

### Update

Run the install command again **from the project that was used to install the add-on**
(that's where DDEV recorded the installation):

```bash
ddev add-on get kgaut/ddev-drupal-tools
```

Files carrying the `#ddev-generated` marker are replaced with the new version. Running the
command from another project works too, but leaves a second manifest in that project —
better to always anchor on the same one.

The installed files also carry `#ddev-silent-no-warn`. DDEV only recognizes add-on files
through the manifest of the current project; without that marker, `ddev start` would list the
global commands as custom configuration ("Remove unexpected '#ddev-generated' comments…") in
every project but the one holding the manifest. `ddev debug check-custom-config --all` still
shows them.

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
| `ddev db-import [dump]` | Drops the database, imports a dump (most recent one from the dumps directory by default), then runs the steps of the project type (Drupal: `drush deploy`, `drush cr`, `drush uli`). Flags: `-l` (list dumps), `-n` (dry-run), `-y` (skip confirmation). |
| `ddev db-export` | Exports the database to `<dir>/<date>-<project>-dev.sql.gz`, after clearing caches on Drupal projects. Flags: `--no-gzip`, `--no-cr`, `-n`. |

Formats supported by `db-import`: `.sql`, `.sql.gz`, `.sql.bz2`, `.sql.xz`, `.mysql`, `.mysql.gz`, `.zip`, `.tgz`, `.tar.gz`.
Dump names are tab-completed (`ddev db-import <tab>`), most recent first.

### Project types

The steps around an import or an export follow the project's DDEV type (the `type` key of
`.ddev/config.yaml`), the same way DDEV only provides `ddev drush` to Drupal projects and
`ddev console` to Symfony ones:

| DDEV type | `db-import`, after the import | `db-export`, before the export |
| --- | --- | --- |
| `drupal`, `drupal6` … `drupal12` | `drush deploy`, `drush cr`, `drush uli` | `drush cr` (skip with `--no-cr`) |
| `symfony` | `console cache:clear` | — |
| any other type | — | — |

Project-specific steps belong in DDEV's own hooks, which `ddev import-db` (used by
`db-import`) runs anyway:

```yaml
# .ddev/config.yaml
hooks:
  post-import-db:
    - exec: wp cache flush   # e.g. a WordPress project
```

`db-{prod,preprod}-dump` relies on `drush sql-dump` on the server, so it only serves Drupal
projects. Other projects get their dumps with `mysqldump` through `db-{prod,preprod}-get` (see
[Dump without drush](#dump-without-drush)).

### Remote server (production / staging)

| Command | Description |
| --- | --- |
| `ddev db-prod-dump` | Runs `drush sql-dump --gzip` on the server, into a timestamped file in `PROD_DB_PATH` (the dump stays on the server). Drupal only. |
| `ddev db-prod-get` | Puts a production dump into the local dumps directory: the most recent one from `PROD_DB_PATH` on the server (showing its date, with a warning past 24 hours), or a fresh `mysqldump` of `PROD_DB_NAME`, streamed straight to the local directory. |
| `ddev db-prod-import` | Chains `db-prod-get` + `db-import`. |
| `ddev ssh-prod` | Opens an SSH session on the current project's production server. |

Each command has its `preprod` (staging) twin: `db-preprod-dump`, `db-preprod-get`,
`db-preprod-import`, `ssh-preprod` — same files, `PREPROD_` variable prefix.

Every command prints its help with `ddev <command> -h`, without running anything.

## Configuration

Everything is configured in the `.env.local` or `.env` file at each project's root (never
sourced: variables are extracted with `grep`). `.env.local` comes first: Symfony projects commit
their `.env` and keep local values, such as server access, in `.env.local`. The commands are
visible everywhere but only useful in configured projects — a missing variable produces an
explicit error.

```bash
# project .env.local (or .env)
PROD_USER=kevin
PROD_HOST=my-server.example.org
PROD_PORT=22                           # optional, defaults to 22
PROD_PATH=/var/www/myproject           # project root on the server
PROD_DRUSH=vendor/bin/drush            # drush binary, relative to PROD_PATH
PROD_DB_PATH=dumps                     # dumps directory, relative to PROD_PATH
PROD_URL=myproject.example.org         # used to name dump files
PROD_DB_NAME=myproject                 # database name, for a dump without drush (see below)

# same idea for staging, with the PREPROD_ prefix
```

Like `PROD_DRUSH`, a relative `PROD_DB_PATH` is resolved from `PROD_PATH`. Absolute paths
and paths starting with `~` are used as-is.

### Dump without drush

With `PROD_DB_NAME` set and **no** `PROD_DB_PATH`, `db-prod-get` (and therefore
`db-prod-import`) runs `mysqldump` on the server and streams the gzipped output straight to the
local dumps directory: nothing is written on the server. Only `PROD_USER`, `PROD_HOST` and
`PROD_DB_NAME` are required (`PROD_URL` names the file, the DDEV project name otherwise).

- Credentials come from the `~/.my.cnf` of the SSH account (`[client]` section with `user` and
  `password`): no password goes through the add-on.
- Options: `--single-transaction --quick --routines --triggers --no-tablespaces`. Events are left
  out, since they need the `EVENT` privilege and production scheduled jobs have no business
  running locally.
- The stream lands in a `.part` file, ignored by `db-import`. It is kept only if the archive is
  valid and ends with mysqldump's `-- Dump completed` line; otherwise it is deleted, including
  on Ctrl-C.
- When both variables are set, `PROD_DB_PATH` wins: the existing server dumps are fetched.

### Local dumps directory

Shared by all commands. Defaults to `files/dumps` (relative to the project root), overridable
with `DB_DUMP_DIR` (project `.env.local`, then `.env`, then `.ddev/.env`), as a relative or
absolute path. The
`db-{prod,preprod}-get` commands also honor `LOCAL_DB_PATH`, which takes precedence over
`DB_DUMP_DIR`.

## Tests

```bash
bats tests
```

Tests run with an isolated `HOME`: they never touch the machine's real `~/.ddev`.

## License

MIT.
