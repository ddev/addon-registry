---
title: kwasib/ddev-keydb
github_url: https://github.com/kwasib/ddev-keydb
description: "KeyDB (Drop-In Alternative to Redis) service for DDEV"
user: kwasib
repo: ddev-keydb
repo_id: 980876915
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-05-09
updated_at: 2025-05-18
workflow_status: disabled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/kwasib/ddev-keydb/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/kwasib/ddev-keydb/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/kwasib/ddev-keydb)](https://github.com/kwasib/ddev-keydb/commits)
[![release](https://img.shields.io/github/v/release/kwasib/ddev-keydb)](https://github.com/kwasib/ddev-keydb/releases/latest)

# DDEV KeyDB

## Overview

This add-on integrates KeyDB into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get kwasib/ddev-keydb
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev keydb-cli` | Run `keydb-cli` inside the Redis container |
| `ddev keydb` | Alias for `ddev keydb-cli` |
| `ddev keydb-flush` | Flush all cache inside the Redis container |
| `ddev describe` | View service status and used ports for KeyDB |
| `ddev logs -s keydb` | Check KeyDB logs |

KeyDB is available inside Docker containers with `keydb:6379`.

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.keydb --keydb-docker-image="eqalpha/keydb:latest"
ddev add-on get kwasib/ddev-keydb
ddev restart
```

Make sure to commit the `.ddev/.env.keydb` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `KEYDB_DOCKER_IMAGE` | `--keydb-docker-image` | `eqalpha/keydb:latest` |

## Credits

**Contributed and maintained by [@kwasib](https://github.com/kwasib). 
Based on [ddev/ddev-redis](https://github.com/ddev/ddev-redis)**
