---
title: ddaffie/ddev-pg-search
github_url: https://github.com/ddaffie/ddev-pg-search
description: "DDEV add-on: ParadeDB pg_search (BM25 full-text search) for DDEV's PostgreSQL"
user: ddaffie
repo: ddev-pg-search
repo_id: 1303914361
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-07-17
updated_at: 2026-07-17
workflow_status: success
stars: 0
---

# DDEV pg_search

[![tests](https://github.com/ddaffie/ddev-pg-search/actions/workflows/tests.yml/badge.svg)](https://github.com/ddaffie/ddev-pg-search/actions/workflows/tests.yml)

Adds [ParadeDB's `pg_search`](https://github.com/paradedb/paradedb) — BM25 full-text search
for PostgreSQL — to a DDEV project's database.

This add-on **extends DDEV's stock Postgres image** by installing ParadeDB's prebuilt `.deb`.
It does not replace the image with ParadeDB's own, so DDEV's snapshots, `ddev import-db`,
`ddev export-db`, users and configuration all keep working exactly as before.

## Requirements

- DDEV `v1.24.0`+
- A PostgreSQL project (`postgres:15` … `postgres:18`) — ParadeDB publishes prebuilt
  packages for Postgres 15 and later only. The add-on refuses to install otherwise.

## Install

```bash
ddev add-on get ddaffie/ddev-pg-search
ddev restart
```

That's it. `ddev restart` builds the database image (~30MB download, ~15s) and creates the
extension. Verify:

```bash
ddev psql -c "SELECT extversion FROM pg_extension WHERE extname='pg_search';"
```

## Usage

`pg_search` is created automatically in the `db` database on every start, so just use it:

```sql
CREATE TABLE articles(id int primary key, body text);
INSERT INTO articles VALUES (1, 'the quick brown fox'), (2, 'lazy dog sleeping');

CREATE INDEX articles_bm25 ON articles
  USING bm25 (id, body) WITH (key_field='id');

SELECT id FROM articles WHERE body @@@ 'fox';
```

See the [pg_search documentation](https://docs.paradedb.com/documentation/full-text/overview)
for query syntax, scoring, highlighting and index options.

## Updating pg_search

The add-on installs whatever the **latest** ParadeDB release was at the time the database
image was built. When a newer release appears, `ddev start` tells you:

```
pg_search 0.25.0 is available (installed: 0.24.3)
Update with: ddev utility rebuild -s db
```

Run that command to update:

```bash
ddev utility rebuild -s db
```

It rebuilds the db image without cache (so the new package is fetched) and the add-on then
runs `ALTER EXTENSION pg_search UPDATE` automatically to bring the SQL definitions in line
with the new binary — a step ParadeDB documents as mandatory after any upgrade.

Confirm the two agree:

```bash
ddev psql -c "SELECT extversion FROM pg_extension WHERE extname='pg_search';"
ddev psql -c "SELECT * FROM paradedb.version_info();"
```

> **`ddev restart` does not update pg_search.** Docker caches the image layer, so a plain
> restart keeps serving the installed version indefinitely. Only `ddev utility rebuild -s db`
> picks up a new release. This is deliberate: silently swapping the extension binary under a
> running project would be worse than telling you and letting you choose when.

If BM25 queries misbehave after a large version jump, try rebuilding the index
(`REINDEX INDEX your_bm25_index;`).

### The update check

The check runs on every `ddev start` / `ddev restart` and adds ~300ms (capped at 3s). It is
strictly a notifier and can never break your project: if you are offline, rate-limited, or
GitHub is down, it prints nothing and exits successfully.

It calls the unauthenticated GitHub API, which allows 60 requests/hour per IP shared across
all your DDEV projects. If you exceed that the only symptom is a missing message.

To disable it, remove the second hook from `.ddev/config.pg_search.yaml` and delete the
`#ddev-generated` line at the top of that file so DDEV stops managing it.

## Removal

```bash
ddev add-on remove pg_search
ddev restart
```

Removal **drops the `pg_search` extension**, and with it any `bm25` indexes. Your tables and
rows are not touched, and the database keeps working normally afterwards.

This drop is not optional, and it is worth understanding why. The extension lives in the
*database*, which outlives the image. If it stayed registered after the image lost the
binary, Postgres could no longer load `pg_search` — and then **every** query against a table
carrying a bm25 index fails, not just search queries. Even `SELECT count(*)` and `pg_dump`
fail, leaving the data stranded. Dropping it during removal, while the binary still exists,
is the only moment that can be prevented.

If the project is stopped when you remove the add-on, the drop cannot run. The add-on will
tell you, and you can recover with:

```bash
ddev start && ddev psql -c 'DROP EXTENSION IF EXISTS pg_search CASCADE;'
```

Re-installing the add-on also repairs a database left in that state.

## How it works

| File | Role |
|---|---|
| `db-build/Dockerfile.pg_search` | Appended to DDEV's generated db Dockerfile; installs the `.deb` and sets `shared_preload_libraries` via `/etc/postgresql/conf.d/`. |
| `config.pg_search.yaml` | Post-start hooks: `CREATE EXTENSION` + `ALTER EXTENSION ... UPDATE`, then the update check. |
| `pg_search/check-update.sh` | Compares the installed package against ParadeDB's latest release. |

The Dockerfile picks the right package by reading its own container rather than mapping
versions itself: `PG_MAJOR` (set by the postgres image), `VERSION_CODENAME` (`/etc/os-release`)
and `TARGETARCH` (buildx) are exactly the three fields in ParadeDB's `.deb` filename. So it
follows Debian base changes in the upstream postgres images without any edit here.

`shared_preload_libraries` is set through `conf.d` rather than by editing DDEV's generated
`postgresql.conf`. Note that another add-on writing that same setting into `conf.d` would
override rather than merge with this one.

## Troubleshooting

**Build fails with a 404 fetching the `.deb`** — ParadeDB does not publish a package for that
Postgres major / Debian / architecture combination. Prebuilt packages exist for Postgres 15–18
on amd64 and arm64.

**Build fails calling the GitHub API** — likely the 60/hour unauthenticated rate limit. Wait
and retry. Unlike the update check, the build fails loudly here on purpose: no package means
no extension, and a silently half-built image would be worse.

## Credits

`pg_search` is built by [ParadeDB](https://paradedb.com). This add-on only packages it for DDEV.
