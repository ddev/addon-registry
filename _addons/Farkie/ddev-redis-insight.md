---
title: Farkie/ddev-redis-insight
github_url: https://github.com/Farkie/ddev-redis-insight
description: ""
user: Farkie
repo: ddev-redis-insight
repo_id: 988642656
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-05-22
updated_at: 2025-05-22
workflow_status: disabled
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/Farkie/ddev-redis-insight/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/Farkie/ddev-redis-insight/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/Farkie/ddev-redis-insight)](https://github.com/Farkie/ddev-redis-insight/commits)
[![release](https://img.shields.io/github/v/release/Farkie/ddev-redis-insight)](https://github.com/Farkie/ddev-redis-insight/releases/latest)

# DDEV Redis Insight

## Overview

This add-on integrates Redis Insight into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get Farkie/ddev-redis-insight
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Redis Insight |
| `ddev logs -s redis-insight` | Check Redis Insight logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.redis-insight --redis-insight-docker-image="busybox:stable"
ddev add-on get Farkie/ddev-redis-insight
ddev restart
```

Make sure to commit the `.ddev/.env.redis-insight` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `REDIS_INSIGHT_DOCKER_IMAGE` | `--redis-insight-docker-image` | `busybox:stable` |

## Credits

**Contributed and maintained by [@Farkie](https://github.com/Farkie)**
