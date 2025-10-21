---
title: d34dman/ddev-notification-server
github_url: https://github.com/d34dman/ddev-notification-server
description: ""
user: d34dman
repo: ddev-notification-server
repo_id: 1003743764
ddev_version_constraint: ">= v1.24.3"
dependencies: ["redis"]
type: contrib
created_at: 2025-06-17
updated_at: 2025-08-26
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/d34dman/ddev-notification-server/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/d34dman/ddev-notification-server/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/d34dman/ddev-notification-server)](https://github.com/d34dman/ddev-notification-server/commits)
[![release](https://img.shields.io/github/v/release/d34dman/ddev-notification-server)](https://github.com/d34dman/ddev-notification-server/releases/latest)

# DDEV Notification Server

## Overview

This add-on integrates Notification Server into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get d34dman/ddev-notification-server
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Notification Server |
| `ddev logs -s notification-server` | Check Notification Server logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.notification-server --notification-server-docker-image="ghcr.io/d34dman/notification-server:main"
ddev add-on get d34dman/ddev-notification-server
ddev restart
```

Make sure to commit the `.ddev/.env.notification-server` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `NOTIFICATION_SERVER_DOCKER_IMAGE` | `--notification-server-docker-image` | `ghcr.io/d34dman/notification-server:main` |

## Credits

**Contributed and maintained by [@d34dman](https://github.com/d34dman)**
