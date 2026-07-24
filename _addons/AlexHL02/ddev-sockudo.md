---
title: AlexHL02/ddev-sockudo
github_url: https://github.com/AlexHL02/ddev-sockudo
description: "🔌 A DDEV add-on to seamlessly integrate and configure Sockudo for your local development environment."
user: AlexHL02
repo: ddev-sockudo
repo_id: 1258658836
default_branch: main
tag_name: v1.1.0
ddev_version_constraint: ">= v1.24.10"
dependencies: ["ddev/ddev-redis"]
type: contrib
created_at: 2026-06-03
updated_at: 2026-07-23
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/AlexHL02/ddev-sockudo/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/AlexHL02/ddev-sockudo/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/AlexHL02/ddev-sockudo)](https://github.com/AlexHL02/ddev-sockudo/commits)
[![release](https://img.shields.io/github/v/release/AlexHL02/ddev-sockudo)](https://github.com/AlexHL02/ddev-sockudo/releases/latest)

# DDEV Sockudo

## Overview

This add-on integrates Sockudo into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get AlexHL02/ddev-sockudo
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Sockudo |
| `ddev logs -s sockudo` | Check Sockudo logs |

## Connecting

Use the following URLs to connect your Pusher-compatible client:

| Protocol | URL |
| -------- | --- |
| WebSocket (WS) | `ws://<project>.ddev.site:6001/app/<SOCKUDO_APP_KEY>` |
| WebSocket Secure (WSS) | `wss://<project>.ddev.site:6002/app/<SOCKUDO_APP_KEY>` |

With the default credentials:

```
host:   <project>.ddev.site
port:   6001 (ws) / 6002 (wss)
key:    app-key
secret: app-secret
app_id: app-id
```

## Advanced Customization

### Custom `config.json`

For anything beyond the environment variables below (multiple apps, per-app policies,
custom CORS, channel namespaces, delta compression, etc.), you can supply Sockudo's own
[configuration file](https://sockudo.io/docs/reference/configuration):

```bash
cp .ddev/sockudo/config.json.example .ddev/sockudo/config.json
```

Edit `.ddev/sockudo/config.json` as needed, then run `ddev restart`. Once this file exists,
it fully replaces the app/adapter/cache/queue defaults normally derived from the
`SOCKUDO_*` environment variables below — `SOCKUDO_HOST`/`SOCKUDO_PORT` are still applied on
top so the server keeps listening where ddev's routing expects it, but everything else in
`config.json` is authoritative. Make sure to commit `.ddev/sockudo/config.json` to version
control.

To change the Sockudo version:

```bash
ddev dotenv set .ddev/.env.sockudo --sockudo-version="v4.5.1"
ddev add-on get AlexHL02/ddev-sockudo
ddev restart
```

Make sure to commit the `.ddev/.env.sockudo` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `SOCKUDO_DOCKER_IMAGE` | `--sockudo-docker-image` | `ubuntu:22.04` |
| `SOCKUDO_VERSION` | `--sockudo-version` | `v4.5.1` |
| `SOCKUDO_HOSTNAME` | `--sockudo-hostname` | `sockudo` |
| `SOCKUDO_HOST` | `--sockudo-host` | `0.0.0.0` |
| `SOCKUDO_PORT` | `--sockudo-port` | `6001` |
| `SOCKUDO_HTTPS_PORT` | `--sockudo-https-port` | `6002` |
| `SOCKUDO_APP_ID` | `--sockudo-app-id` | `app-id` |
| `SOCKUDO_APP_KEY` | `--sockudo-app-key` | `app-key` |
| `SOCKUDO_APP_SECRET` | `--sockudo-app-secret` | `app-secret` |
| `SOCKUDO_REDIS_HOST` | `--sockudo-redis-host` | `redis` |
| `SOCKUDO_REDIS_PORT` | `--sockudo-redis-port` | `6379` |
| `SOCKUDO_REDIS_DB` | `--sockudo-redis-db` | `0` |
| `SOCKUDO_REDIS_PASSWORD` | `--sockudo-redis-password` | `` |
| `SOCKUDO_METRICS_ENABLED` | `--sockudo-metrics-enabled` | `false` |

## Credits

**Contributed and maintained by [@AlexHL02](https://github.com/AlexHL02)**
