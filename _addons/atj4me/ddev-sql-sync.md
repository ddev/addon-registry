---
title: "atj4me/ddev-sql-sync"
github_url: "https://github.com/atj4me/ddev-sql-sync"
description: "DDEV add-on: sync a production MySQL database into your local DDEV environment — Drush-style sql-sync over SSH (app export or mysqldump + rsync + import + sanitize)"
user: "atj4me"
repo: "ddev-sql-sync"
repo_id: 1446485725
default_branch: "main"
tag_name: "v1.0.8"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-11-05"
updated_at: "2026-11-05"
workflow_status: "success"
stars: 0
---

# DDEV sql-sync

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/atj4me/ddev-sql-sync/actions/workflows/tests.yml/badge.svg)](https://github.com/atj4me/ddev-sql-sync/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/atj4me/ddev-sql-sync)](https://github.com/atj4me/ddev-sql-sync/commits)
[![release](https://img.shields.io/github/v/release/atj4me/ddev-sql-sync)](https://github.com/atj4me/ddev-sql-sync/releases/latest)

Sync a production MySQL database into your local [DDEV](https://ddev.com) project — modelled on Drush's `sql-sync`, for any plain SSH-accessible server (a VPS running Laravel, WordPress, Drupal, …), not just hosted platforms.

## What it does

`ddev sql-sync @prod @self`:

1. **Dump** — ssh to the server and create a database dump there, staged in a dedicated private folder (`<tmp>/ddev-dumps/`) with a completion marker carrying its byte size. See *Dump strategies* below.
2. **Download** — rsync the dump into a local folder (default `~/sql-sync`) and verify it against the completion marker before anything is imported. The remote temp files are deleted after a successful import (kept if the import fails, so the same dump can be fetched again with `--no-dump`).
3. **Import** — `ddev import-db` into this project (you are asked to confirm first; the current local database is replaced).
4. **Sanitize** — scrub user emails and passwords in the *local* database (on by default; see *Sanitize*).

The add-on only ever **reads** from the server (a dump). It never writes to a remote database, and `@self` is the only supported target.

## Installation

```bash
ddev add-on get atj4me/ddev-sql-sync
ddev restart
```

## Configuration: `.ddev/sql-sync/aliases.yml`

This file is the whole configuration. Installing the add-on copies `aliases.example.yml` to `aliases.yml` if you don't have one yet; **your `aliases.yml` is never overwritten by add-on updates and never removed** when the add-on is removed.

```yaml
local:
  dir: ~/sql-sync        # where downloaded dumps are kept

prod:
  host: myserver  # an ~/.ssh/config alias or a hostname
  user: deploy           # optional; ssh target becomes user@host
  root: /var/www/example  # the app's path on the server
  tmp: /tmp              # optional; temp folder on the server
```

Define as many aliases as you like (`prod`, `staging`, …) and pick one per run: `ddev sql-sync @staging @self`.

Optional per-alias keys:

| Key | Purpose |
| --- | --- |
| `platform:` | `laravel`, `wordpress`, `drupal` or `php`. Resolved per alias: this key, else the project's DDEV type from `.ddev/config.yaml` (laravel/wordpress/drupal* map to themselves), else `php` — DDEV's generic type, an unrecognised type, and no type at all all land on the generic php platform. Usually unset. |
| `options:` | Extra ssh options, e.g. `options: -i ~/.ssh/prod_key`. |
| `jump:` | Optional jump host (bastion) passed to ssh `-J` (ProxyJump) on every connection, e.g. `jump: bastion.example.com`. Prefer `ProxyJump` in `~/.ssh/config` — that is the durable place, and every tool benefits. This key is only for one-off syncs where you don't want to touch your ssh config. |
| `dump:` | A command on the server whose stdout **is** the dump (e.g. `dump: drush sql-dump`). Replaces both dump strategies below — works for any platform. |
| `database:` | Explicit database credentials (see below). |
| `sanitize:` | Path (relative to the project root) to your own SQL file to sanitize with after import. |

### Database credentials: app config first, explicit block when needed

Normally no credentials are configured at all: at dump time, the add-on reads them from the app's own config **on the server** — `.env` (Laravel), `wp-config.php` (WordPress), or `sites/default/settings.php` (Drupal). SSH authentication stays your own keys/`~/.ssh/config`; nothing is stored by the add-on.

For setups where that can't work (managed database hosts, split credentials), an alias can pin values explicitly:

```yaml
prod:
  # ...
  database:
    name: example_db
    username: example_user
    password: secret
    host: 127.0.0.1   # from the app server's perspective
    port: 3306
```

Any field left out falls back to the app's config, field by field — you can pin just the password, or just host/port for an RDS-style endpoint.

> **Treat a filled-in `password:` as a secret.** Prefer app-config extraction when it works. `aliases.yml` lives under `.ddev/`, which many projects gitignore (check yours), and `chmod 600 .ddev/sql-sync/aliases.yml` is a good idea if you use the block. The password is passed to mysqldump via the `MYSQL_PWD` environment variable only — never as a `-p` argument and never echoed.

### One-time SSH setup (your side)

Make sure you can reach the server non-interactively before the first run:

```bash
ssh your-server-alias 'echo ok'
```

If that prompts for a password, set up a key (`ssh-copy-id`) and, if you use a short alias like `myserver`, a matching entry in your `~/.ssh/config`. This is your own machine's SSH setup — the add-on doesn't manage keys.

## Usage

```bash
ddev sql-sync @prod @self           # source, then target — like cp / rsync / drush sql-sync
ddev sql-sync @prod @self -y        # skip the confirmation prompt
ddev sql-sync @prod @self --dry-run # preflight checks only — nothing is dumped or imported
```

| Flag | Meaning |
| --- | --- |
| `--sanitize` / `--no-sanitize` | Sanitize after import (default on) / skip it. Skipping prints a loud warning: the local DB then holds real production emails and password hashes. |
| `--dump-dir=<path>` | Where the dump is downloaded to (overrides `local.dir`). |
| `--no-dump` | Don't create a new dump; rsync the dump an earlier `--no-sync` run left on the server. Its completion marker must be present and match the downloaded size, or the run refuses to import. |
| `--no-sync` | Only create the dump on the server (with its completion marker); don't download, import, or sanitize. A later `--no-dump` run fetches it. |
| `--dry-run` | Preflight only: check ssh connectivity, platform, dump strategy, credentials (the password is never shown), database size and disk space, then print the plan. Takes no dump and imports nothing. See "Testing safely" for one small footnote. |
| `-y`, `--yes` | Don't ask for confirmation. |
| `-h`, `--help` | Show help. |

Dumps land at `<dump-dir>/<alias>.sql` (or `<alias>.sql.gz` when the dump is gzipped) and are **kept** after import (handy for re-importing). They contain production data — delete them when you no longer need them: `rm ~/sql-sync/*.sql`.

## Platforms

The dump strategy per alias, in order:

1. **`dump:` override** — if the alias sets one, it always wins.
2. **The app's own export tool**, run on the server in the app root, if the platform has one *and* it's installed there (checked at run time; if missing, the next strategy is used instead of failing). A non-interactive ssh session gets a minimal PATH, so WP-CLI is also probed in `~/bin`, `~/.local/bin` and `/usr/local/bin`, and Drush in the app's `vendor/bin`:
   - WordPress → `wp db export` (WP-CLI)
   - Drupal → `drush sql-dump`
   - Laravel → **no native data export exists** (`schema:dump` is structure-only), so Laravel always uses strategy 3.
3. **mysqldump** — with credentials resolved as described above (explicit `database:` block merged over the app's config).

The run tells you which strategy it used: `using the app's own export: wp db export`, `using the app's own export: drush sql-dump`, `using the alias's dump: command`, or `using mysqldump`.

Platform notes:

- **Laravel / any `.env` app** — credentials from `<root>/.env`. Bedrock-style WordPress also fits here (set `platform: laravel`).
- **WordPress (classic)** — credentials from the `DB_*` defines in `<root>/wp-config.php` (or one level up, as WordPress allows).
- **PHP (generic)** — DDEV's generic `php` type, and the terminal default when nothing else matches. No app-native export exists, so it always uses mysqldump; credentials come from `<root>/.env` if the app has one, merged under the alias's `database:` block — for a generic app that block is the expected primary source. No bundled sanitize: a `sanitize:` file is honoured, otherwise the run reports that no sanitize ran.
- **Drupal** — dump via Drush when available; otherwise extraction reads the default `$databases` entry from `settings.php` with a small PHP snippet on the server (best-effort: a settings file that pulls credentials from the environment may not yield them — then use a `dump:` override such as `dump: drush sql-dump`, or pin a `database:` block).

## Sanitize

Runs after import, on the local database only:

- **Laravel** — bundled SQL: `users.email` → `user<id>@example.test`, `users.password` → bcrypt of `password`, `remember_token` → `NULL`.
- **WordPress** — bundled SQL against the site's actual users table: `<prefix>users.user_email` → `user<id>@example.test`, `<prefix>users.user_pass` → `MD5('password')` (WordPress accepts MD5 at login and rehashes transparently). The prefix is read from the server's `wp-config.php` — a custom `$table_prefix` is honoured, not assumed to be `wp_` (on `--no-dump` runs it is detected from the imported tables instead).
- **Drupal** — `ddev drush sql-sanitize` when the local project has Drush (Drupal hashes passwords in PHP, so there is deliberately **no** bundled SQL sanitize for it).
- **Any platform** — an explicit `sanitize:` file in the alias, when set, is used **instead of** the built-in sanitize on every platform — including Drupal, where it wins over `drush sql-sanitize` (explicit per-alias config wins, as with `dump:` and `database:`).

If sanitize is on (the default) but nothing could be applied — no matching table, no Drush, no `sanitize:` file — the run imports and then **stops with a prominent error and a non-zero exit**: it never reports success over an unsanitized database. Pass `--no-sanitize` to accept an unsanitized import explicitly.

## Relationship to Drush `sql-sync`

Mirrored deliberately: the two-position `sql-sync @source @target` form with `@self` as the local site, site aliases in a YAML file, the confirmation prompt, and `--no-dump` / `--no-sync` with their Drush meanings.

Deviations, also deliberate:

- **Sanitize is folded into the sync and ON by default** (Drush makes it a separate `sql-sanitize` command you must remember to run). `--no-sanitize` opts out.
- **Only `@self` as target** in v1 — syncing *to* a remote is out of scope by design; the tool never writes to a remote server.
- Aliases carry connection + app-path + platform in one flat file rather than Drush's per-site alias files, and credential resolution (app config / explicit block) replaces Drush's reliance on the app's own bootstrap, so it works the same for Laravel and WordPress.

## Testing safely

New server, or just cautious? Work up in four steps:

1. **`ddev sql-sync @prod @self --dry-run`** — preflight only. It connects over ssh (BatchMode, short timeout, so it fails fast instead of hanging), detects the platform and dump strategy on the server, resolves the credentials (database name and username are shown; the password never is), estimates the database size from `information_schema`, and checks free space in the server's `tmp` folder and your local dump folder. You get a verdict per check plus the plan — then it stops. No dump is taken and nothing is imported. One footnote: resolving *Drupal* credentials during preflight briefly creates a small temp file (mode 600) in the server's `tmp` folder and deletes it again within the same step — it is the extractor for `settings.php`, not a dump.
2. **`ddev sql-sync @prod @self --no-sync`** — a real dump, created on the server in `<tmp>/ddev-dumps/` with a completion marker. **Nothing is downloaded** — the point of this step is proving the server side works end to end and leaving a known-good dump in place.
3. **`ddev sql-sync @prod @self --no-dump`** — fetch the dump step 2 left behind (its completion marker is checked against the downloaded size before anything is imported), then import and sanitize it — without touching the server's database again.
4. **`ddev sql-sync @prod @self`** — the full run.

What a real run writes on the server: a dump file and a small completion marker in `<tmp>/ddev-dumps/` inside the alias's `tmp` folder (the folder is created mode 700, the dump mode 600), both removed after a successful import. If the import fails, both are kept so a `--no-dump` run can fetch the same dump; if the dump itself fails, its partial file is removed by the remote side. One guard protects small servers: before dumping, a run estimates the database size and aborts if the `tmp` filesystem has less free space than that — whenever the size can be estimated. When it cannot (for example a `dump:` command whose credentials cannot be resolved, or no `mysql` client on the server), the run says so and skips the check rather than pretending it happened — a dump filling the disk is the one realistic way a sync could hurt a small server, so an un-estimable database deserves a second look at the server's free space before a full run.

## Safety notes

- The only operations ever performed on the server are creating the dump and its completion marker in `<tmp>/ddev-dumps/` (removed after a successful import; a failed dump removes its own partial file) — plus whatever a `dump:` command *you* configured chooses to do.
- A downloaded dump is production data. The default sanitize scrubs the local **database**, not the dump file on disk; delete dumps you don't need.
- The add-on stores no credentials of its own anywhere.

## Removal

```bash
ddev add-on remove sql-sync
```

This removes the command and the bundled files. Your `aliases.yml` and any downloaded dumps are left in place — delete them manually if you want them gone.
