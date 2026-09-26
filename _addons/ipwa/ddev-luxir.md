---
title: "ipwa/ddev-luxir"
github_url: "https://github.com/ipwa/ddev-luxir"
description: "Unofficial DDEV add-on running a Luxir search server"
user: "ipwa"
repo: "ddev-luxir"
repo_id: 1385962641
default_branch: "main"
tag_name: ""
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: "contrib"
created_at: "2026-09-24"
updated_at: "2026-09-24"
workflow_status: "success"
stars: 0
---

# ddev-luxir

A [DDEV](https://ddev.com) add-on that runs a [Luxir](https://luxir.org) search server as the `luxir` service,
reachable at `http://luxir:9400` inside the project network.

> **Unofficial.** Luxir 0.1.0 only ships Linux x86-64 binaries and has no official container image.
> This add-on builds its own image from the official release binaries (checksums verified).
> Luxir is pre-1.0 and has **no authentication or TLS**: never expose it publicly.

## Install

```sh
ddev add-on get ipwa/ddev-luxir
ddev restart
```

## Usage

| Command | What it does |
|---|---|
| `ddev luxir [METHOD] PATH [JSON\|@file]` | Calls the Luxir HTTP API from the web container. `@file` is relative to the current directory; `.ndjson` files are sent as NDJSON. |
| `ddev luxir-status` | Version and CPU tier, health, collections, stats totals. |

```sh
ddev luxir /collections
ddev luxir POST /collections/demo/_update @tests/testdata/books.ndjson
ddev luxir /collections/demo/_search '{"query":"title_t:dune","get_number":true}'
```

To try a local checkout instead: `ddev add-on get /path/to/ddev-luxir`.

## Prebuilt image for CI

The same Dockerfile is published as `ghcr.io/ipwa/luxir:<luxir-version>` (also `sha-<commit>`; never `latest`).
Use it as a service container, for example in GitLab CI:

```yaml
services:
  - name: ghcr.io/ipwa/luxir:0.1.0
    alias: luxir
variables:
  SEARCH_API_LUXIR_URL: http://luxir:9400
```

The add-on itself always builds locally.

## Settings

Set with `ddev dotenv set .ddev/.env.luxir --luxir-max-ram-mb=1024`, then `ddev restart` (or
`ddev debug rebuild -s luxir` after changing `LUXIR_VERSION`).

| Variable | Default | Meaning |
|---|---|---|
| `LUXIR_VERSION` | `0.1.0` | Release to download at build time. Also change the Dockerfile `ARG` default when bumping. |
| `LUXIR_CPU_TIER` | `auto` | `v2`, `v3`, `v4`; `auto` picks the highest tier that starts. |
| `LUXIR_MAX_RAM_MB` | `512` | `--max-ram-mb` |
| `LUXIR_LOG_LEVEL` | `info` | `--log-level` |
| `LUXIR_EXTRA_ARGS` | empty | Extra server flags, split on spaces. |

Data lives in the `luxir` Docker volume (`--store.backend=fs`).

## Drupal and Search API Luxir

Install `drupal/search_api_luxir`, create a server with host `luxir`, port `9400`, scheme `http`.
Do not use `localhost`.

## Limitations

- No authentication/TLS; the router publishes it on `http://<project>.ddev.site:9400` for host-side tools only.
- Luxir has no graceful shutdown: always commit (`"commit": {}`); uncommitted updates are lost on stop.
- `ddev snapshot` does not include the Luxir volume.
- amd64 only. On Apple Silicon it runs under Rosetta (tier `v3` was selected in testing).
- `ddev luxir` runs in the web container, so `@file` paths are relative to where you run it.

## Development

```sh
docker build --platform linux/amd64 -t luxir-dev luxir/
shellcheck luxir/*.sh commands/*/*
bats tests/test.bats --filter-tags '!release'
```

Publishing the image (`.github/workflows/image.yml`) pushes to `ghcr.io/ipwa/luxir:<version>`; the package
must be made public manually in the GitHub package settings.

## Credits

Luxir is developed by the [Luxir project](https://github.com/luxir-search/luxir) (Apache-2.0). This add-on is unofficial and not affiliated with it. Licensed under Apache-2.0.
