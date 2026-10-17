---
title: "tyler36/ddev-anubis"
github_url: "https://github.com/tyler36/ddev-anubis"
description: "Bot protection via Anubis as a DDEV addon"
user: "tyler36"
repo: "ddev-anubis"
repo_id: 1421573587
default_branch: "main"
tag_name: "v0.1"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-10-16"
updated_at: "2026-10-16"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-anubis/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tyler36/ddev-anubis/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-anubis)](https://github.com/tyler36/ddev-anubis/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-anubis)](https://github.com/tyler36/ddev-anubis/releases/latest)

# DDEV Anubis

## Overview

This add-on integrates [Anubis](https://anubis.techaro.lol/) into your [DDEV](https://ddev.com/) project. Anubis:

> Weigh the soul of incoming HTTP requests to protect your website!

## Installation

```bash
ddev add-on get tyler36/ddev-anubis
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command               | Description                                   |
| --------------------- | --------------------------------------------- |
| `ddev describe`       | View service status and used ports for Anubis |
| `ddev logs -s anubis` | Check Anubis logs                             |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.anubis --anubis-docker-image="ghcr.io/techarohq/anubis:latest"
ddev add-on get tyler36/ddev-anubis
ddev restart
```

Make sure to commit the `.ddev/.env.anubis` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `ANUBIS_DIFFICULTY` | `--anubis-difficulty` | `4` |
| `ANUBIS_DOCKER_IMAGE` | `--anubis-docker-image` | `ghcr.io/techarohq/anubis:latest` |
| `ANUBIS_OG_PASSTHROUGH` | `--anubis-og-passthrough` | `true` |
| `ANUBIS_OG_EXPIRY_TIME` | `--anubis-og-expiry-time` | `24h` |
| `ANUBIS_PORT` | `--anubis-port` | `8080` |
| `ANUBIS_SERVE_ROBOTS_TXT` | `--anubis-serve-robots-txt` | `true` |

## Credits

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
