---
title: markdlabrecque/ddev-wayfinder
github_url: https://github.com/markdlabrecque/ddev-wayfinder
description: "A DDEV addon for building a container for a Wayfinder search backend"
user: markdlabrecque
repo: ddev-wayfinder
repo_id: 1324790351
default_branch: main
tag_name: v0.1.0
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-08-06
updated_at: 2026-08-11
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/markdlabrecque/ddev-wayfinder/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/markdlabrecque/ddev-wayfinder/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/markdlabrecque/ddev-wayfinder)](https://github.com/markdlabrecque/ddev-wayfinder/commits)
[![release](https://img.shields.io/github/v/release/markdlabrecque/ddev-wayfinder)](https://github.com/markdlabrecque/ddev-wayfinder/releases/latest)

# DDEV Wayfinder

Adds [Wayfinder](https://github.com/markdlabrecque/wayfinder) — a Solr-wire-compatible search
backend — to a [DDEV](https://ddev.com/) project, so Drupal's Search API can talk to it the way
it talks to Solr.

## Installation

```bash
ddev add-on get markdlabrecque/ddev-wayfinder
ddev restart
```

Then commit the `.ddev` directory to version control.

The add-on installs the Wayfinder service and nothing else — it does not touch Drupal, install
modules, or write any Search API configuration. The sections below are instructions for you to
follow, not things that happen on `ddev restart`.

## Connecting Drupal (optional)

### The Drupal module

The Search API backend plugin lives in `search_api_wayfinder`, inside the Wayfinder repository.
It is self-contained — it does **not** depend on `search_api_solr` or Solarium, and it does not
use a Solr connector plugin. It is not on Packagist, so add the repository yourself. With a local
checkout, a `path` repository is the form that works, because the Wayfinder repo has no
`composer.json` at its root for a `vcs` repository to read:

```bash
ddev composer config repositories.wayfinder path /path/to/wayfinder/drupal/search_api_wayfinder
ddev composer require wayfinder/search_api_wayfinder
ddev drush en search_api_wayfinder -y
```

### The Search API server

Add a server at **Configuration → Search and metadata → Search API → Add server** and choose
**Wayfinder** as the backend (not a Solr backend with a connector), then:

| Setting | Value |
| ------- | ----- |
| HTTP protocol | `http` |
| Host | `wayfinder` |
| Port | `8983` |
| Base path | `/wayfinder` |
| Core | `content` |

Two of those are easy to get wrong:

- **Solr path is `/wayfinder`, not `/solr`.** Wayfinder serves its wire API under `/wayfinder/`
  and has no `/solr` compatibility route, so a `/solr` base path produces 404s on every request.
  See the upgrade note in the Wayfinder repo's `drupal/search_api_wayfinder/README.md`.
- **The host is the service name `wayfinder`**, not `localhost`. Drupal runs in the `web`
  container and reaches Wayfinder across the project's Docker network.

You do not need to generate or upload a config set. The schema is baked into the image.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev wayfinder ping` | Core ping — the health check |
| `ddev wayfinder system` | Core-scoped system info |
| `ddev wayfinder info` | Server-level version handshake |
| `ddev wayfinder ui` | Open the admin UI in a browser (`--print` to only print the URL) |
| `ddev wayfinder url` | Print the base URL for the Search API server |
| `ddev describe` | Service status and exposed ports |
| `ddev logs -s wayfinder` | Wayfinder logs |

The admin UI (query tester, schema view, index stats, synonyms) is served at `/ui` on the
router-exposed port — `ddev wayfinder ui` opens it.

Mind the `/ui` path. The URL `ddev describe` lists for the `wayfinder` service is the bare
`https://<project>.ddev.site:8945`, and Wayfinder has no route at `/`: it answers with an
empty `404`, which in a browser is an indistinguishable-from-broken blank page. On DDEV
v1.25.0+ the `ddev describe` row carries a reminder to that effect.

## Versions

The add-on pins an image tag, so each release targets a specific Wayfinder version.

| Add-on version | Wayfinder image |
| -------------- | --------------- |
| `v0.1.0` | `ghcr.io/markdlabrecque/wayfinder:v0.1.0` |

## Advanced customization

Settings live in `.ddev/.env.wayfinder` (commit it):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `WAYFINDER_DOCKER_IMAGE` | `--wayfinder-docker-image` | `ghcr.io/markdlabrecque/wayfinder:v0.1.0` |
| `WAYFINDER_SCHEMA` | `--wayfinder-schema` | `/presets/search-api.toml` |
| `WAYFINDER_CORE` | `--wayfinder-core` | `content` |
| `WAYFINDER_LOG_LEVEL` | `--wayfinder-log-level` | `info` |

```bash
ddev dotenv set .ddev/.env.wayfinder --wayfinder-docker-image="ghcr.io/markdlabrecque/wayfinder:v1.2.3"
ddev restart
```

### Server config

`.ddev/wayfinder/wayfinder.toml` is mounted read-only at `/conf/wayfinder.toml`. Every key has a
working default, so the shipped file is nearly empty. Unknown keys are rejected and the server
will refuse to start — see Wayfinder's `docs/operations.md` for the valid tables (`[indexing]`,
`[query]`, `[resources]`, `[commit]`, `[admin]`, `[extraction]`, `[auth]`).

Basic auth is deliberately not configured. Wayfinder is only reachable on the project network
here, and the credentials in Wayfinder's integration harness are a test fixture — do not copy
them.

### A custom schema

The default schema is the `search-api` preset baked into the image, which already expresses
`search_api_solr`'s field-naming conventions. To use your own, drop it in `.ddev/wayfinder/` and
point the add-on at its in-container path:

```bash
ddev dotenv set .ddev/.env.wayfinder --wayfinder-schema="/conf/schema.toml"
ddev restart
```

### A different core name

`WAYFINDER_CORE` sets the core the `ddev wayfinder` commands talk to. The core name Wayfinder
actually serves comes from the `[core] name` in its schema, and the baked-in preset uses
`content` — so changing `WAYFINDER_CORE` alone will 404. Change it together with a custom schema
whose `[core] name` matches, and set the same value as the Search API server's core.

### Reindexing from scratch

The index lives in a Docker named volume, not in the project directory, so `ddev restart` keeps
it. To throw it away:

```bash
ddev stop
docker volume ls | grep wayfinder-data   # find the exact name
docker volume rm ddev-<project>_wayfinder-data
ddev start
```

The prefix is DDEV's Compose project name, which is the project name lowercased with dots
removed — so `Client.Local` gives `ddev-clientlocal_wayfinder-data`. Look it up rather than
assuming.

## Credits

**Contributed and maintained by [@markdlabrecque](https://github.com/markdlabrecque)**
