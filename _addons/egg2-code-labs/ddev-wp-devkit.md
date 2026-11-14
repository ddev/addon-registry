---
title: "egg2-code-labs/ddev-wp-devkit"
github_url: "https://github.com/egg2-code-labs/ddev-wp-devkit"
description: "DDEV add-on for local WordPress development with real databases from production, staging or a teammate."
user: "egg2-code-labs"
repo: "ddev-wp-devkit"
repo_id: 1440035315
default_branch: "main"
tag_name: "v0.2.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-10-30"
updated_at: "2026-11-13"
workflow_status: "cancelled"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/egg2-code-labs/ddev-wp-devkit/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/egg2-code-labs/ddev-wp-devkit/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/egg2-code-labs/ddev-wp-devkit)](https://github.com/egg2-code-labs/ddev-wp-devkit/commits)
[![release](https://img.shields.io/github/v/release/egg2-code-labs/ddev-wp-devkit)](https://github.com/egg2-code-labs/ddev-wp-devkit/releases/latest)

# ddev-wp-devkit

DDEV add-on for local WordPress development with real databases from production, staging or a teammate.
Import a database and the site runs locally, without caches, security plugins, cron or outgoing mail getting in the way. Optionally share the database between developers via a remote folder (e.g. Google Drive) with a pointer in git.

## What it does

| When | What |
|---|---|
| `ddev start` | Creates `wp-config.php` if none exists (or only the DDEV-generated one), activates the git hooks. |
| `ddev import-db` | Rewrites URLs to the DDEV URL (incl. multisite tables, JSON-escaped and URL-encoded variants, `http` → `https`). |
| Every request | Disables cache, security/2FA, backup, analytics and SMTP plugins **at runtime**; the database stays unchanged. |
| Every request | `DISABLE_WP_CRON`, Action Scheduler async runner off, all mail to Mailpit (even with SMTP plugins), `WP_DEBUG_LOG` on, `WP_CACHE` off. |
| PHP | `.user.ini` files are ignored (e.g. Wordfence WAF with a production server path). |

The add-on **never writes plugin state to the database**. Plugins stay active in the database and are only filtered out locally. Saving the plugins screen does not remove them from the database either.

## Installation

```bash
ddev add-on get egg2-code-labs/ddev-wp-devkit
ddev restart
ddev wpdev status
```

After installation, make sure to commit the `.ddev` directory to version control.

## Requirements

- DDEV v1.24.10 or newer
- A DDEV project of type `wordpress`
- WordPress core in the docroot or in one subdirectory below it (like DDEV itself supports; the generated `wp-config.php` detects the location of `wp-settings.php`)
- `python3` on the host (used by `ddev dbsync`)
- `rclone` (only for `ddev dbsync` with a remote folder)
- `git` (database sharing relies on git hooks)

## Commands

```bash
ddev wpdev status            # overview: URL, multisite, plugins disabled locally, ...
ddev wpdev prepare           # rewrite URLs (runs automatically after ddev import-db)
ddev wpdev admin             # create local admin "devadmin" / "devadmin" or reset its password
ddev wpdev admin name pass   # with your own name/password
ddev wpdev setup             # create wp-config.php (runs automatically on ddev start)
ddev wpdev hash [--verbose]  # content hash of the database (used by ddev dbsync)
ddev wpdev edits show 10     # last edits recorded for ddev dbsync
```

## Sharing the database (`ddev dbsync`)

The database lives in a shared remote folder (Google Drive via rclone, any other rclone remote, or a locally synced folder). Git only tracks `db/snapshot.json` with file name, SHA-256 and content hash of the dump.

| When | What |
|---|---|
| `git commit` (pre-commit) | Recorded edits, or content hash ≠ pointer → `ddev export-db`, upload, the new pointer is added to the commit |
| `git checkout` / `merge` / `rebase` | Pointer ≠ loaded snapshot → download, SHA check, `ddev import-db`. Unsaved local changes are never overwritten. |
| Merge conflict in `db/snapshot.json` | Both sides changed the database → `ddev dbsync resolve ours\|theirs` |

### How changes are detected

Plugins write to the database all the time (caches, counters, notices, update checks), so comparing database contents cannot tell real edits from noise. By default (`WPDEV_DB_CHANGE_DETECTION=edits`) wp-devkit therefore records **edits** while WordPress runs. A database write counts as an edit when it happens

- in a non-GET request of a logged-in user (saving posts, settings, menus, media, page builders, REST/AJAX saves),
- in an admin URL with a nonce opened by a logged-in user (activating a plugin, trashing a post and other GET actions; any parameter whose name contains "nonce", or the `X-WP-Nonce` header),
- in a WP-CLI command (`ddev wp ...`), except read-only commands (`list`, `get`, `status`, ... also nested like `wp post meta get`) and `cron`, `transient`, `cache`, `rewrite flush`.

Anonymous visitors, page views, list filters and searches, logins, the heartbeat, screen options, transients and statistics hits never count, whichever plugin writes. Each edit is one line in `.ddev/.wpdev/dirty` (`ddev wpdev edits show`, `ddev dbsync status`). Add project-specific noise with `WPDEV_DIRTY_IGNORE_EXTRA`, a regex matched against the signature shown there, e.g. `^POST /wp-admin/admin-ajax\.php action=my_plugin_ping$`.

In addition, the content hash covers the content tables only (posts, postmeta, terms, comments, links, WPML translation links; extend with `WPDEV_DB_HASH_CONTENT_TABLES`). It catches content changes made without WordPress, e.g. with direct SQL. Settings changed that way are invisible: record them with `ddev dbsync mark-dirty "what changed"`.

Edits made while a commit exports the database are kept for the next snapshot; if an export fails, its edits stay recorded. If the edits or the hash cannot be read, the commit is stopped rather than passed without a snapshot.

Edit recording needs the wp-devkit runtime in `wp-config.php` (automatic for the generated one, see [Custom wp-config.php](#custom-wp-configphp)). Without it, `ddev dbsync` warns and compares the full hash (`full-v1`) instead. Set `WPDEV_*` variables in `.ddev/.env` (not just `WPDEV_DB_CHANGE_DETECTION` and `WPDEV_DIRTY_IGNORE_EXTRA`, but also `WPDEV_DB_HASH_CONTENT_TABLES` and the `WPDEV_DB_HASH_SKIP_*` lists read inside the container; `.env.local` never reaches WordPress or the web container) and run `ddev restart` after changing them.

`WPDEV_DB_CHANGE_DETECTION=hash` restores the previous behaviour: no edit recording, the hash covers every table minus known noise (`WPDEV_DB_HASH_SKIP_*`), and every plugin that writes in the background needs its own exclusion.

**Hash policies.** Every snapshot pointer stores the `hash_policy` its hash was computed with (`content-v1` for `edits`, `full-v1` for `hash`). wp-devkit only compares hashes computed with the same policy, so switching the mode never hides a change. A hash it cannot compare counts as changed: after switching from `edits` to `hash` (a `content-v1` hash cannot vouch for unchanged settings) and for snapshots written by wp-devkit versions without hash policies, the next commit exports a new snapshot, and `pull` asks to commit or `--force` first. Upgrade wp-devkit for the whole team at once: older versions do not know about policies and export a snapshot on every commit after a newer version wrote a pointer.

```bash
ddev dbsync status | pull [--force] | push | resolve ours|theirs
ddev dbsync mark-dirty "note"    # record a change wp-devkit cannot see (direct SQL, database GUI)
ddev dbsync init <dump>          # first import from production: rewrite URLs + create devadmin
ddev dbsync golive <url> <id>    # production export: reassign devadmin's content to <id>, delete devadmin, rewrite URLs back
```

Setup: set `WPDEV_DB_REMOTE=<rclone-remote>:` in `.ddev/.env`; every developer creates an rclone remote with that name (`rclone config create <name> drive scope=drive root_folder_id=<folder-id>`). Alternatively use a local folder in `.ddev/.env.local`. The git hooks are activated on `ddev start` via `core.hooksPath`; if it is already set to something else, it is left unchanged.

Skip the hook once: `WPDEV_DB_SKIP=1 git commit ...`

Always commit with a plain `git commit` (or `git commit -a`). `git commit --only <file>` would bypass the pointer staged by the pre-commit hook and silently desynchronize `db/snapshot.json` from the uploaded dump.

## Configuration

In `.ddev/.env` (committed), then `ddev restart`:

| Variable | Default | Meaning |
|---|---|---|
| `WPDEV_MULTISITE` | empty | `subdirectory` or `subdomain` for multisite networks. `prepare` reports when it is missing. |
| `WPDEV_TABLE_PREFIX` | `wp_` | Table prefix |
| `WPDEV_DISABLE_PLUGINS` | see `wp-devkit/wp-config-devkit.php` | Replaces the default list (plugin folder names, separated by spaces/commas) |
| `WPDEV_DISABLE_PLUGINS_EXTRA` | empty | Disable additionally |
| `WPDEV_KEEP_PLUGINS` | empty | Allow again from the default list |
| `WPDEV_DISABLE_CRON` | `true` | WP-Cron and Action Scheduler async runner off |
| `WPDEV_BLOCK_EXTERNAL_HTTP` | `false` | Block outgoing HTTP requests |
| `WPDEV_HTTP_ALLOW` | `*.wordpress.org` | Allowed hosts when blocked |
| `WPDEV_SR_SKIP_TABLES` | statistics/log tables | Tables skipped during URL rewrite (wildcards) |
| `WPDEV_ADMIN_USER` / `_PASSWORD` / `_EMAIL` | `devadmin` / `devadmin` / `<user>@example.test` | Defaults for `ddev wpdev admin` |
| `WPDEV_ADMIN_AUTO` | `false` | Create the admin after every import if missing |
| `WPDEV_DB_REMOTE` | empty | Target for `ddev dbsync`: rclone remote (`name:path`) or absolute local path |
| `WPDEV_DB_CHANGE_DETECTION` | `edits` | `edits`: record edits + content-table hash. `hash`: hash over all tables only (previous behaviour) |
| `WPDEV_DIRTY_IGNORE_EXTRA` | empty | Regex: request signatures that are not edits (see `ddev wpdev edits show`) |
| `WPDEV_DB_HASH_CONTENT_TABLES` | empty | Regex: additional tables for the content hash (mode `edits`) |
| `WPDEV_DB_HASH_SKIP_TABLES` / `_OPTIONS` / `_META` | see `commands/web/wpdev` | Regex: what does not affect the content hash (replaces the default) |
| `WPDEV_DB_HASH_SKIP_TABLES_EXTRA` / `_OPTIONS_EXTRA` / `_META_EXTRA` | empty | Regex: added to the defaults above, e.g. a plugin's view counter meta key |

Project-specific PHP constants (e.g. dev license keys, WPML settings) go into `.ddev/wp-config-project.php`, which is loaded first.

## Custom wp-config.php

If the project has its own committed `wp-config.php`, the add-on leaves it alone. Add this line before `wp-config-ddev.php` is loaded:

```php
if ( getenv( 'IS_DDEV_PROJECT' ) === 'true' ) { require '/var/www/html/.ddev/wp-devkit/wp-config-devkit.php'; }
```

## Limitations

- **Base64-encoded content** (e.g. WPBakery `vc_raw_html`) is not covered by the URL rewrite.
- WP-CLI skips serialized objects of plugin classes (warning "Skipping an uninitialized class"), usually only admin notices.
- Multisite: user names need at least 4 characters, `admin` is not allowed.
- Subdomain multisite also needs `additional_hostnames` in the DDEV config.
- Databases cannot be merged. `ddev dbsync` only makes parallel changes visible (merge conflict); resolving means picking one side and redoing the other side's changes by hand.
- Edit detection is a heuristic. A plugin that writes on every logged-in POST (e.g. a tracking ping) creates a snapshot on the next commit: `ddev wpdev edits show` names the request, exclude it with `WPDEV_DIRTY_IGNORE_EXTRA`. GET actions are only detected in wp-admin: front-end links that change the database via a nonce (confirmations, unsubscribe) are missed, record them with `ddev dbsync mark-dirty`. Because any parameter whose name contains "nonce" counts, a plugin that adds such a parameter to all of its admin links makes every admin visit an edit, exclude it with `WPDEV_DIRTY_IGNORE_EXTRA`. Settings changed outside WordPress (direct SQL, `wp db query`, database GUIs) are not detected: use `ddev dbsync mark-dirty`.
- With `WPDEV_DB_CHANGE_DETECTION=hash`, if a plugin writes to the database on every request, every commit creates a new snapshot. Compare `ddev wpdev hash --verbose` before/after a page request and exclude the option via `WPDEV_DB_HASH_SKIP_OPTIONS` / `_META` / `_TABLES`.
- The content hash aggregates the hashed tables in SQL and runs on every `git commit` (pre-commit hook). On small and medium databases this takes seconds; on very large databases it can take considerably longer. `ddev wpdev hash --verbose` shows which tables dominate the run time.
- rclone's shared Google Drive client ID is being retired in 2026. Create your own OAuth client: https://rclone.org/drive/#making-your-own-client-id

## Credits

- Initial implementation of the plugin code by [Florian Behrend](https://github.com/florianbehrend).

**Contributed and maintained by [@egg2-code-labs](https://github.com/egg2-code-labs)**
