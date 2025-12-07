---
title: Corveda/ddev-redash
github_url: https://github.com/Corveda/ddev-redash
description: "DDEV addon for Redash"
user: Corveda
repo: ddev-redash
repo_id: 1111123801
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-12-06
updated_at: 2025-12-06
workflow_status: unknown
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/Corveda/ddev-redash/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/Corveda/ddev-redash/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/Corveda/ddev-redash)](https://github.com/Corveda/ddev-redash/commits)
[![release](https://img.shields.io/github/v/release/Corveda/ddev-redash)](https://github.com/Corveda/ddev-redash/releases/latest)

# DDEV Redash

## Overview

Redash integration for DDEV.

- Runs Redash server + scheduler + workers
- Uses Postgres 17 (pgautoupgrade) dedicated for Redash
- Reuses the [ddev-redis](https://github.com/ddev/ddev-redis) add-on
- Exposes UI at `https://redash.<project>.ddev.site` (no port)

## Requirements

- DDEV v1.24.0 or newer
- Docker
- Project already set up with DDEV

  ```bash
  ddev get ddev/ddev-redis
  ddev restart
  ```

  After restart, DDEV will merge docker-compose.redash.yaml into your project.

## URL

Redash will be available at: `https://redash.<project>.ddev.site`


You can also see this via:

`ddev describe`


Look for "Redash UI" in the output.

## Database initialization

On a new project, you must initialize the Redash database schema once
and keep it up to date.

One-time initialization
# Run migrations (creates schema on an empty DB)
`ddev exec -s redash-server ./manage.py db upgrade`


If you see errors about missing tables (e.g. "relation 'queries' does not exist"),
make sure you've run this on a fresh Redash DB (the default is redash
inside redash-postgres).

Optional: auto-run migrations on startup

You can add this hook to .ddev/config.yaml in your project so Redash
migrations run automatically every time the project starts:
```yaml
hooks:
  post-start:
    - exec: ./manage.py db upgrade || echo 'Redash migration failed (ignored on startup)'
      service: redash-server
```

This keeps the Redash DB schema up to date with the image.

Create an admin user

After migrations succeed, create the first admin user (if Redash doesn't
show you the web-based "Initial Setup" screen):

```
ddev exec -s redash-server ./manage.py users create \
  --admin \
  --password admin \
  --email admin@example.com \
  --name "Admin"
```


Then log in at: `https://redash.<project>.ddev.site`

Environment and secrets

The add-on embeds default environment values in docker-compose.redash.yaml:

```
REDASH_REDIS_URL: "redis://redis:6379/5"
REDASH_DATABASE_URL: "postgresql://redash:redash@redash-postgres/redash"
REDASH_COOKIE_SECRET: "changeme-redash-cookie-secret"
REDASH_SECRET_KEY: "changeme-redash-secret-key"
```

For any non-throwaway setup, you should:

Change `REDASH_COOKIE_SECRET` and `REDASH_SECRET_KEY` to long random strings.

Optionally change the Postgres credentials (`POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`)
and keep `REDASH_DATABASE_URL` in sync.

## Services

This add-on defines these services:

redash-server – Redash web UI / API, fronted by DDEV router

redash-scheduler – schedules periodic queries

redash-scheduled-worker – runs scheduled queries

redash-adhoc-worker – runs ad-hoc queries

redash-worker – handles periodic, emails, default queues

redash-postgres – dedicated Postgres for Redash

redis – expected from ddev/ddev-redis add-on

All of them are wired into DDEV via labels, so you can use:

```
ddev logs -s redash-server
ddev logs -s redash-worker
ddev exec -s redash-scheduler env
```

## Uninstall

From your project:

`ddev delete -Oy   # if you want to remove containers and data`
or just:
`ddev delete`


Then remove the add-on from your project config:

`ddev get --remove corveda/ddev-redash`


Or manually delete docker-compose.redash.yaml under .ddev/.
