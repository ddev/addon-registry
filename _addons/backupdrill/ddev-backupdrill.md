---
title: "backupdrill/ddev-backupdrill"
github_url: "https://github.com/backupdrill/ddev-backupdrill"
description: "DDEV add-on that backs up the project's Postgres database and restores it into a throwaway container to check the backup works"
user: "backupdrill"
repo: "ddev-backupdrill"
repo_id: 1429660515
default_branch: "main"
tag_name: "v0.1.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-10-22"
updated_at: "2026-10-24"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/backupdrill/ddev-backupdrill/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/backupdrill/ddev-backupdrill/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/backupdrill/ddev-backupdrill)](https://github.com/backupdrill/ddev-backupdrill/commits)
[![release](https://img.shields.io/github/v/release/backupdrill/ddev-backupdrill)](https://github.com/backupdrill/ddev-backupdrill/releases/latest)

# ddev-backupdrill

`ddev backupdrill` backs up your project's Postgres database, restores the backup into a throwaway Postgres container, and checks that the tables and rows came back.

It runs the open-source `backupdrill` CLI. That's the engine [BackupDrill](https://backupdrill.com) uses for its scheduled restore drills on Supabase projects.

## Why drill a local database?

Your DDEV database usually has the same schema as production. Running a drill on it shows whether that schema survives a `pg_dump` and `pg_restore` round trip (extensions, views, foreign keys, indexes) before you trust dumps of the real thing. It's also a quick way to see what a restore drill checks and what the report looks like.

It doesn't keep backups of your dev database. The snapshot goes into a temporary bucket that is deleted when the command exits. Use `ddev snapshot` or `ddev export-db` if you want a copy to keep.

## Requirements

- DDEV v1.24.10 or newer
- A project with a Postgres database, for example `ddev config --database=postgres:17`
- On your machine, not in the web container:
  - Docker, which DDEV already needs
  - Node.js 20 or newer, because the CLI is fetched with `npx`
  - `pg_dump` and `pg_restore` with a major version at least as high as the project's Postgres. On macOS, `brew install libpq` and add `$(brew --prefix libpq)/bin` to your `PATH`. On Debian or Ubuntu, `sudo apt install postgresql-client`. If they live somewhere else, point `BACKUPDRILL_PG_DUMP` and `BACKUPDRILL_PG_RESTORE` at them.

The command runs on the host because the drill starts its sandbox with `docker run` and connects to it on `127.0.0.1`. The web container can't do either.

## Installation

```bash
ddev add-on get backupdrill/ddev-backupdrill
```

The add-on refuses to install on a MySQL or MariaDB project. Commit the `.ddev` directory so everyone on the project gets the command.

## Usage

```bash
ddev start
ddev backupdrill
```

What it does:

1. Starts an S3-compatible bucket in a container (versitygw), listening on `127.0.0.1` on a random port, with random keys.
2. Runs `backupdrill backup`, which dumps the project database (the `public` schema) into that bucket along with a manifest: tables, estimated row counts, and the dump's sha256.
3. Runs `backupdrill drill`, which downloads the dump, checks the sha256, restores it into a fresh `postgres:<major>-alpine` container, and compares what came back against the manifest.
4. Removes the temporary bucket, and the sandbox too unless the drill failed and you passed `--keep`. Ctrl+C removes both as well, along with the downloaded dump.

Output from a project with four tables and about 200,000 rows (some progress lines left out):

```text
✓ Postgres 17.11 (Debian 17.11-1.pgdg13+2) — schema(s) public: 4 tables, ~201,254 rows
✓ Database dumped (1.3 MB, sha256 23b232b791f3…)
✓ Backup complete — 4 tables (database only).
→ Drilling snapshot 2026-10-22T11-12-55-623Z

  ✓ archive integrity — sha256 matches (23b232b791f3…)
  ✓ sandbox extensions — installed pg_trgm
  ✓ pg_restore — completed in 0.4s (2 parallel jobs)
  ✓ post-data objects — all post-data objects restored
  ✓ table count — restored 4, manifest 4
  ✓ no missing tables — all manifest tables present
  ✓ populated tables came back — no populated table restored empty

✓ Drill PASSED — snapshot 2026-10-22T11-12-55-623Z, 4 tables / 201,254 rows restored in 0.4s
Removed the temporary bucket. The snapshot was deleted with it; nothing is kept.
```

The exit code is 0 when the drill passes and non-zero when it doesn't. The full drill report also goes to stdout as one line of JSON, if you want to parse it in a script.

### Drill flags

Two flags are passed through to `backupdrill drill`:

| Flag | What it does |
| ---- | ------------ |
| `--check-cmd <command>` | After the built-in checks pass, runs your own command against the restored copy, with its connection string in `BACKUPDRILL_SANDBOX_URL`. Exit 0 passes, anything else fails the drill. |
| `--keep` | If the drill fails, leaves the sandbox container running so you can look inside. The output tells you how to remove it. |

```bash
ddev backupdrill --check-cmd ./scripts/check-restore.sh
```

The check command runs on your machine from the project root, not in the web container. It connects as the sandbox's superuser, so row-level security doesn't apply. Use it for data checks ("every order has a customer", "this table isn't empty"), not to test RLS policies.

Any other flag is refused. Options like `--config`, `--bucket` or `--ping-url` would point the CLI at real storage or report a local result to your production monitoring, which this command is built never to do.

### What it reads, and what it ignores

The command only ever talks to the project database and its own temporary bucket. It doesn't read a `backupdrill.config.json` in the project, and it drops every `BACKUPDRILL_*` variable and `DATABASE_URL` from your shell before it starts the CLI, so a production setting you exported earlier can't redirect it. Three variables are kept:

- `BACKUPDRILL_PG_DUMP` and `BACKUPDRILL_PG_RESTORE`: paths to the Postgres client tools
- `BACKUPDRILL_SCHEMAS`: schemas to dump instead of `public`, comma-separated

## Limitations

- Postgres only. The CLI is built on `pg_dump` and `pg_restore` and has no MySQL or MariaDB support.
- Database only. Uploaded files in your project are not part of the backup.
- Only the `public` schema unless you set `BACKUPDRILL_SCHEMAS`.
- Nothing is kept. Each run makes a new backup and throws it away.
- The CLI comes from npm as `backupdrill@^2.7.1` (the newest 2.x release, 2.7.1 or later), and the first run pulls the versitygw and Postgres images, so it needs network access and takes longer.
- Pressing Ctrl+C a second time skips the cleanup and can leave the sandbox running. List leftover sandboxes with `docker ps --filter name=backupdrill-drill-` and remove one with `docker rm -fv <id>`.
- Tested on macOS (Apple Silicon, OrbStack) and on Ubuntu, x86-64 and arm64, in CI. Windows hasn't been tried.

## Backing up a real database

To back up a production database to your own S3, R2 or B2 bucket, run the `backupdrill` CLI from npm directly. It does the same backup and drill shown here, and for a Supabase project it also copies the Storage files once you give it the project's Storage S3 keys. If you don't want to run it yourself, BackupDrill does it on a schedule for Supabase projects and emails you when a backup or drill fails.

## Removal

```bash
ddev add-on remove backupdrill
```

## Credits

**Contributed and maintained by [@backupdrill](https://github.com/backupdrill)**
