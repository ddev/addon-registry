---
title: "libredb/ddev-libredb-studio"
github_url: "https://github.com/libredb/ddev-libredb-studio"
description: "LibreDB Studio is a lightweight, high-performance, and secure web-based SQL-NoSQL database editor in DDEV"
user: "libredb"
repo: "ddev-libredb-studio"
repo_id: 1372630375
default_branch: "main"
tag_name: "v0.1.0"
ddev_version_constraint: ">= v1.25.4"
dependencies: []
type: "contrib"
created_at: "2026-09-16"
updated_at: "2026-10-06"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/libredb/ddev-libredb-studio/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/libredb/ddev-libredb-studio/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/libredb/ddev-libredb-studio)](https://github.com/libredb/ddev-libredb-studio/commits)
[![release](https://img.shields.io/github/v/release/libredb/ddev-libredb-studio)](https://github.com/libredb/ddev-libredb-studio/releases/latest)

# DDEV LibreDB Studio

## Overview

[LibreDB Studio](https://github.com/libredb/libredb-studio) is an open-source, web-based SQL IDE with AI query assistance.

This add-on integrates LibreDB Studio into your [DDEV](https://ddev.com/) project.
Studio opens already connected to the project database, whether it is MySQL, MariaDB or PostgreSQL.
MongoDB, Redis and SQL Server from the DDEV add-ons can be connected as well, one setting each.

## Installation

```bash
ddev add-on get libredb/ddev-libredb-studio
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.
The login secrets are not part of it: they live in `.ddev/.env.libredb-studio.local`, which Git must ignore (see [Login](#login)).

DDEV v1.25.4 or later is required.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev libredb-studio` | Print the login and open LibreDB Studio in your browser (`https://<project>.ddev.site:9401`) |
| `ddev describe` | View service status and used ports for LibreDB Studio |
| `ddev logs -s libredb-studio` | Check LibreDB Studio logs |

### Login

Studio has no passwordless mode, so the add-on generates a login for each developer:

* The admin email is `admin@ddev.site`.
* The admin password and the session signing key (`JWT_SECRET`) are random, created on install or on the first `ddev start` without them, and kept from then on.
* Both are stored in `.ddev/.env.libredb-studio.local`, which only you can read and Git ignores, so a teammate who clones the project gets their own on their first `ddev start`.
* `ddev libredb-studio` prints the login before it opens the browser.

LibreDB Studio 0.18.0 and later keep their accounts in the add-on's data volume, `ddev-<project>_libredb-studio`, and take the admin password from the secrets file only on their first start.
A password created later, after the secrets file was deleted or the add-on was removed and added again, does not sign in while that volume holds the earlier account.
To start over with new values, run `ddev stop`, delete `.ddev/.env.libredb-studio.local` and the volume with `docker volume rm ddev-<project>_libredb-studio`, then run `ddev start`.
This also removes Studio's saved queries and settings.

DDEV keeps `.ddev/.env.*.local` files out of Git through `.ddev/.gitignore`, but only while DDEV manages that file.
Before it writes the secrets, and on every start, the add-on asks Git whether the file is ignored, and stops with an error if it is not.
If you manage `.ddev/.gitignore` yourself (it has no `#ddev-generated` line), add this line to it:

```gitignore
/.env.*.local
```

The secrets also leave that file in two places on your machine.
DDEV writes the resolved values into `.ddev/.ddev-docker-compose-full.yaml`, which Git ignores and which the add-on makes readable by you only after every start.
They are also part of the container's environment, so `docker inspect` shows them to anyone who can use Docker on this machine.
This is the same exposure as any other value DDEV passes to a container, and the login guards a local development tool, not a shared server.

### Connections

| Connection | Seeded when | Credentials |
| ---------- | ----------- | ----------- |
| DDEV db | Always, unless the project omits the `db` container | `db` / `db`, database `db` |
| DDEV mongo | `LIBREDB_STUDIO_SEED_MONGO=true` and [ddev/ddev-mongo](https://github.com/ddev/ddev-mongo) | `MONGO_INITDB_ROOT_USERNAME` / `MONGO_INITDB_ROOT_PASSWORD`, default `db` / `db` |
| DDEV redis | `LIBREDB_STUDIO_SEED_REDIS=true` and [ddev/ddev-redis](https://github.com/ddev/ddev-redis) | None. The optimized ddev-redis configuration (ACL users with a password) is not supported yet: the seeded connection then fails with `NOAUTH` |
| DDEV sqlsrv | `LIBREDB_STUDIO_SEED_SQLSRV=true` and [ddev/ddev-sqlsrv](https://github.com/ddev/ddev-sqlsrv) | `SA` / `MSSQL_SA_PASSWORD`, default `Password12!` |

Every connection is managed: its password stays on the server, and the connection cannot be edited in the browser.
Every connection is visible to the admin account only, because each one logs in with a superuser account.

Each connection reaches the service by the container name DDEV gives it in this project, such as `ddev-<project>-db`, and never by the short service name.
All DDEV projects share one Docker network, on which a short name like `db` also answers for another running project's container.

For example, to add Redis:

```bash
ddev add-on get ddev/ddev-redis
ddev dotenv set .ddev/.env.libredb-studio --libredb-studio-seed-redis=true
ddev restart
```

When a connection is switched on and its service is missing, the LibreDB Studio container stops with an error naming the add-on to install, and `ddev restart` fails.
A project without a `db` container starts LibreDB Studio without a db connection, after a 30-second wait for the service; set `LIBREDB_STUDIO_SEED_DB=false` to skip that wait.

### Plain http

When your project is served over plain http (mkcert is not installed), the browser drops Studio's login cookie, and `ddev libredb-studio` stops with a message.
Either install mkcert with `mkcert -install` and run `ddev restart`, or allow the cookie over http:

```bash
ddev dotenv set .ddev/.env.libredb-studio --auth-cookie-secure=false
ddev restart
```

## Advanced Customization

To run a different LibreDB Studio release than the one this add-on version pins:

```bash
ddev dotenv set .ddev/.env.libredb-studio --libredb-studio-docker-image="ghcr.io/libredb/libredb-studio:<version>"
ddev restart
```

The add-on is tested with the pinned release only.

Make sure to commit the `.ddev/.env.libredb-studio` file to version control.
Do not put a password in it: use `.ddev/.env.libredb-studio.local` for that.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `LIBREDB_STUDIO_DOCKER_IMAGE` | `--libredb-studio-docker-image` | `ghcr.io/libredb/libredb-studio:0.18.0` |
| `LIBREDB_STUDIO_ADMIN_EMAIL` | `--libredb-studio-admin-email` | `admin@ddev.site` |
| `LIBREDB_STUDIO_SEED_DB` | `--libredb-studio-seed-db` | `true` |
| `LIBREDB_STUDIO_SEED_MONGO` | `--libredb-studio-seed-mongo` | `false` |
| `LIBREDB_STUDIO_SEED_REDIS` | `--libredb-studio-seed-redis` | `false` |
| `LIBREDB_STUDIO_SEED_SQLSRV` | `--libredb-studio-seed-sqlsrv` | `false` |
| `LIBREDB_STUDIO_WAIT_SECONDS` | `--libredb-studio-wait-seconds` | `30` |
| `AUTH_COOKIE_SECURE` | `--auth-cookie-secure` | Secure on any host but localhost |

Other LibreDB Studio settings, such as an AI provider, can be set in `.ddev/.env.libredb-studio` the same way; see [`.env.example`](https://github.com/libredb/libredb-studio/blob/main/.env.example).
Keep their keys and tokens in `.ddev/.env.libredb-studio.local`.

## Removal

```bash
ddev add-on remove libredb-studio
ddev restart
```

Removal deletes `.ddev/.env.libredb-studio.local`.
Saved queries and settings stay in the Docker volume `ddev-<project>_libredb-studio` until you remove it with `docker volume rm`.
The volume also keeps Studio's admin account, so after the add-on is added again its new password signs in only once that volume is removed (see [Login](#login)).

## Credits

**Contributed and maintained by [LibreDB](https://github.com/libredb)**
