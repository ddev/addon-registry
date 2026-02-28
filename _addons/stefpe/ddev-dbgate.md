---
title: stefpe/ddev-dbgate
github_url: https://github.com/stefpe/ddev-dbgate
description: "Powerful and easy to use multi engine database manager as DDEV add-on"
user: stefpe
repo: ddev-dbgate
repo_id: 1168331665
default_branch: main
tag_name: 0.0.1
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-27
updated_at: 2026-02-27
workflow_status: unknown
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/stefpe/ddev-dbgate/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/stefpe/ddev-dbgate/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/stefpe/ddev-dbgate)](https://github.com/stefpe/ddev-dbgate/commits)
[![release](https://img.shields.io/github/v/release/stefpe/ddev-dbgate)](https://github.com/stefpe/ddev-dbgate/releases/latest)

# DDEV Dbgate

## Overview

This add-on integrates Dbgate into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get stefpe/ddev-dbgate
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Dbgate |
| `ddev logs -s dbgate` | Check Dbgate logs |

After starting, Dbgate will be available at:
- `http://dbgate.project.ddev.site:3000`
- `https://dbgate.project.ddev.site:3001`

## Advanced Customization

### Connections

You can manage your database connections in `.ddev/.env.dbgate`. By default, it includes MySQL, Postgres, and Redis with DDEV defaults.

To add or modify connections, edit the `CONNECTIONS` list and add the corresponding `LABEL_`, `SERVER_`, `USER_`, `PASSWORD_`, `PORT_`, and `ENGINE_` variables.

Example for a custom connection:
```bash
CONNECTIONS=mysql,my_custom_db

# ... existing mysql config ...

LABEL_my_custom_db=Custom DB
SERVER_my_custom_db=custom-service
USER_my_custom_db=user
PASSWORD_my_custom_db=pass
PORT_my_custom_db=5432
ENGINE_my_custom_db=postgres@dbgate-plugin-postgres
```

### Docker Image

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.dbgate --dbgate-docker-image="dbgate/dbgate:latest"
ddev add-on get stefpe/ddev-dbgate
ddev restart
```

Make sure to commit the `.ddev/.env.dbgate` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `DBGATE_DOCKER_IMAGE` | `--dbgate-docker-image` | `dbgate/dbgate:latest` |

## Credits

**Contributed and maintained by [@stefpe](https://github.com/stefpe)**
