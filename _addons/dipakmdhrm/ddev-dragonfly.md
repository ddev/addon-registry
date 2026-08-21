---
title: "dipakmdhrm/ddev-dragonfly"
github_url: "https://github.com/dipakmdhrm/ddev-dragonfly"
description: "DDEV add-on for DragonflyDB — a Redis-compatible in-memory data store"
user: "dipakmdhrm"
repo: "ddev-dragonfly"
repo_id: 1251293362
default_branch: "main"
tag_name: "0.1.0"
ddev_version_constraint: ""
dependencies: []
type: "contrib"
created_at: "2026-05-27"
updated_at: "2026-05-27"
workflow_status: "unknown"
stars: 0
---

# ddev-dragonfly

A DDEV add-on that provides [DragonflyDB](https://www.dragonflydb.io/) as a Redis-compatible in-memory data store for local development.

DragonflyDB is a modern replacement for Redis that is fully compatible with the Redis API while offering better performance and lower memory usage for many workloads.

## Installation

```bash
ddev add-on get ddev-dragonfly
ddev restart
```

After installation, the DragonflyDB service is available at `dragonfly:6379` from within the DDEV network.

## Commands

| Command | Description |
|---------|-------------|
| `ddev redis-cli` | Run redis-cli commands against DragonflyDB |
| `ddev dragonfly-flush` | Flush all keys (FLUSHALL ASYNC) |

### Examples

```bash
ddev redis-cli PING          # Returns PONG
ddev redis-cli SET foo bar   # Set a key
ddev redis-cli GET foo       # Get a key
ddev redis-cli INFO          # Server info
ddev redis-cli DBSIZE        # Number of keys
ddev dragonfly-flush          # Flush all keys
```

## Configuration

Override defaults by setting environment variables in `.ddev/config.yaml`:

```yaml
web_environment:
  - DRAGONFLY_DOCKER_IMAGE=docker.dragonflydb.io/dragonflydb/dragonfly:v1.25.5
  - DRAGONFLY_HOSTNAME=dragonfly
```

| Variable | Default | Description |
|----------|---------|-------------|
| `DRAGONFLY_DOCKER_IMAGE` | `docker.dragonflydb.io/dragonflydb/dragonfly:v1.38.1` | Docker image for DragonflyDB |
| `DRAGONFLY_HOSTNAME` | `dragonfly` | Hostname for the DragonflyDB container |

### Default flags

The DragonflyDB container starts with these flags:

- `--port=6379` — listen port
- `--maxmemory=512mb` — memory limit
- `--cache_mode=true` — automatic eviction of least-recently-used keys when memory limit is reached
- `--proactor_threads=2` — number of I/O threads (suitable for local dev)

To customize flags, create a `docker-compose.dragonfly_extra.yaml` override:

```yaml
services:
  dragonfly:
    command: dragonfly --port=6379 --maxmemory=1gb --cache_mode=true --proactor_threads=4
```

## Drupal integration

For Drupal 9+ projects, the add-on automatically:

1. Copies `settings.ddev.dragonfly.php` to `sites/default/`
2. Appends an include line to `settings.php`

This configures Drupal's Redis module to use DragonflyDB as the cache backend. You need the [Redis module](https://www.drupal.org/project/redis) installed and the PhpRedis PHP extension enabled (included in DDEV by default).

## Comparison with ddev-redis

| Feature | ddev-redis | ddev-dragonfly |
|---------|-----------|----------------|
| Backend | Redis | DragonflyDB |
| Redis API compatible | Yes | Yes |
| Memory efficiency | Standard | Better for many workloads |
| Multi-threaded | No (single-threaded) | Yes |
| Cluster mode | Optional | Not needed (multi-threaded) |
| Configuration | Config files | Command-line flags |

Both add-ons use the same Drupal Redis module and PhpRedis extension. You can switch between them by removing one and installing the other.

**Note:** Do not run both ddev-redis and ddev-dragonfly simultaneously — they serve the same purpose and Drupal settings will conflict.

## Removal

```bash
ddev add-on remove dragonfly
ddev restart
```

This removes the DragonflyDB container and cleans up `settings.ddev.dragonfly.php` from your Drupal site if it contains the `#ddev-generated` marker.
