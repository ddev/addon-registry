---
title: "kevinquillen/ddev-arangodb"
github_url: "https://github.com/kevinquillen/ddev-arangodb"
description: "A DDEV service for ArangoDB."
user: "kevinquillen"
repo: "ddev-arangodb"
repo_id: 1253487763
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-05-29"
updated_at: "2026-05-29"
workflow_status: "disabled"
stars: 0
---

# ddev-arangodb

[![tests](https://github.com/kevinquillen/ddev-arangodb/actions/workflows/tests.yml/badge.svg)](https://github.com/kevinquillen/ddev-arangodb/actions/workflows/tests.yml)
![Add-on Registry](https://img.shields.io/badge/DDEV%20add--on-arangodb-blue)
![Project is maintained](https://img.shields.io/maintenance/yes/2026)

A [DDEV](https://ddev.com/) add-on that provisions an
[ArangoDB](https://arangodb.com/) multi-model database service
alongside your DDEV project. Designed to be zero-config for Drupal
sites that consume ArangoDB but useful for any DDEV project that needs
a local graph/document/key-value store.

## What you get

- An ArangoDB 3.12 (community) container reachable from the web
  container at `http://arangodb:8529`.
- The ArangoDB web UI at `https://<project>.ddev.site:8530`, with a
  valid TLS cert provided by DDEV's router.
- Named persistent volumes for database data and Foxx apps that
  survive `ddev stop` and are wiped only by
  `ddev add-on remove arangodb`.
- For Drupal projects: a `settings.ddev.arangodb.php` snippet wired
  into `sites/default/settings.php` that populates
  `$settings['arangodb']` so the Drupal side connects with no extra
  wiring.

## Install

```bash
ddev add-on get kevinquillen/ddev-arangodb
ddev restart
```

First boot takes ~20 seconds while ArangoDB initializes the data
volume and sets the root credential; subsequent boots are fast.

Open the web UI:

```bash
ddev launch :8530
# or visit https://<project>.ddev.site:8530 directly
```

In the ArangoDB web UI's login form:

| Field    | Value                                             |
| -------- | ------------------------------------------------- |
| Username | `root`                                            |
| Password | `ddevpassword` (override via `.ddev/config.arangodb.yaml`) |
| Database | `_system` (or any database you create)            |

## Connect from your project

### From PHP (Drupal-ready)

If your project type is Drupal, the add-on writes
`web/sites/default/settings.ddev.arangodb.php` and includes it from
`settings.php`. It populates:

```php
$settings['arangodb'] = [
  'endpoint' => 'http://arangodb:8529',
  'username' => 'root',
  'password' => getenv('DDEV_ARANGODB_PASSWORD') ?: 'ddevpassword',
  'database' => getenv('DDEV_ARANGODB_DATABASE') ?: '_system',
];
```

Sites that need bespoke credentials can simply not include the file
and define their own block.

### From the CLI

`arangosh` ships in the ArangoDB image:

```bash
docker exec -it ddev-<project>-arangodb arangosh \
  --server.endpoint http+tcp://localhost:8529 \
  --server.username root \
  --server.password ddevpassword
```

Or hit the HTTP API directly from the web container:

```bash
ddev exec "curl -u root:ddevpassword http://arangodb:8529/_api/version"
```

### Connection endpoints

| Caller                                                       | Endpoint                                  |
| ------------------------------------------------------------ | ----------------------------------------- |
| PHP / drush / web container                                  | `http://arangodb:8529` (plaintext, in-network) |
| Host browser (web UI)                                        | `https://<project>.ddev.site:8530`        |
| Host browser (HTTP)                                          | `http://<project>.ddev.site:8529`         |

ArangoDB speaks only HTTP(S), so the DDEV router proxies everything.
No raw-TCP host port binding is required.

## Configuration overrides

The add-on ships `.ddev/config.arangodb.yaml` for per-project tuning.
Edit it and run `ddev restart`:

```yaml
web_environment:
  - DDEV_ARANGODB_PASSWORD=ddevpassword
  - DDEV_ARANGODB_DATABASE=_system
  - DDEV_ARANGODB_MEMORY=
  - ARANGODB_DOCKER_IMAGE=arangodb:3.12
```

To claim ownership of the file (so a future `ddev add-on get` will
not overwrite your edits), delete the `#ddev-generated` marker comment
at the top.

### Changing the root password after first boot

`ARANGO_ROOT_PASSWORD` is only honored on a fresh data volume. To
change it on an existing volume, either drop the volume and reinstall,
or run `arango-secure-installation` inside the container:

```bash
docker exec -it ddev-<project>-arangodb arango-secure-installation
```

### Low-memory hosts

Pin RocksDB's view of total RAM:

```yaml
web_environment:
  - DDEV_ARANGODB_MEMORY=2G
```

This maps to `ARANGODB_OVERRIDE_DETECTED_TOTAL_MEMORY` inside the
container.

### Pin a specific patch version

```yaml
web_environment:
  - ARANGODB_DOCKER_IMAGE=arangodb:3.12.9
```

This add-on targets community edition only. Enterprise tags are not
supported.

## Wipe the database

Three options, in order of severity:

```bash
# Drop a single collection's contents from arangosh.
docker exec -it ddev-<project>-arangodb arangosh \
  --server.username root --server.password ddevpassword \
  --javascript.execute-string 'db._collection("mycol").truncate()'

# Remove the data volume only.
ddev stop
docker volume rm "ddev-${DDEV_PROJECT:-$(basename $PWD)}_arangodb-data"
ddev start

# Remove the add-on entirely (wipes data, Foxx apps,
# Drupal settings include).
ddev add-on remove arangodb
```

## Compatibility

| Layer    | Tested                                       | Notes                                              |
| -------- | -------------------------------------------- | -------------------------------------------------- |
| DDEV CLI | 1.24.10+                                     | Required for `x-ddev` describe extensions          |
| ArangoDB | 3.12 community                               | Pin via `ARANGODB_DOCKER_IMAGE`                    |
| Host OS  | macOS (Apple Silicon + Intel), Linux x86_64  | Multi-arch image                                   |
| Drupal   | 9, 10, 11                                    | Settings injection is conditional on `PROJECT_TYPE` |

## Operational notes

- **Cold-start time.** First boot is ~20s while ArangoDB initializes
  the data volume. The web container waits on ArangoDB's healthcheck,
  so `ddev start` blocks briefly. This is not a hang.
- **Memory.** Defaults size RocksDB from detected RAM. On Docker
  Desktop this can be the whole VM. Set `DDEV_ARANGODB_MEMORY` if
  your VM is small.
- **HTTP only.** ArangoDB speaks HTTP for both UI and API. Inside
  the DDEV network it is unencrypted; the host-facing UI is HTTPS via
  the DDEV router. If you need to expose ArangoDB externally,
  configure TLS at the ArangoDB level.
- **Volume naming.** Volumes are scoped per project as
  `ddev-<project>_arangodb-{data,apps}`, so multiple DDEV sites on
  the same host do not collide.

## Removing the add-on

```bash
ddev add-on remove arangodb
```

This removes:

- `.ddev/docker-compose.arangodb.yaml`
- `.ddev/config.arangodb.yaml`
- `.ddev/settings.ddev.arangodb.php`
- The Drupal `settings.php` include block (left intact if you
  modified it).
- The `ddev-<project>_arangodb-{data,apps}` Docker volumes.

## Testing

```bash
brew install bats-core bats-assert bats-file bats-support
bats ./tests/test.bats
```

CI runs the same suite via
[`ddev/github-action-add-on-test`](https://github.com/ddev/github-action-add-on-test)
on every PR, every push to `main`, and weekly on a cron schedule.

## License

Apache 2.0. See [LICENSE](https://github.com/kevinquillen/ddev-arangodb/blob/main/LICENSE).

ArangoDB community edition is licensed under the Apache 2.0 license
with a 100 GB production dataset limit for non-commercial use; review
the [ArangoDB Community License](https://arangodb.com/community-license/)
before deploying. You are responsible for compliance.
