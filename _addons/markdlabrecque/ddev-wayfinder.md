---
title: markdlabrecque/ddev-wayfinder
github_url: https://github.com/markdlabrecque/ddev-wayfinder
description: "A DDEV addon for building a container for a Wayfinder search backend"
user: markdlabrecque
repo: ddev-wayfinder
repo_id: 1324790351
default_branch: main
tag_name: 
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-08-06
updated_at: 2026-08-06
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/markdlabrecque/ddev-wayfinder/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/markdlabrecque/ddev-wayfinder/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/markdlabrecque/ddev-wayfinder)](https://github.com/markdlabrecque/ddev-wayfinder/commits)
[![release](https://img.shields.io/github/v/release/markdlabrecque/ddev-wayfinder)](https://github.com/markdlabrecque/ddev-wayfinder/releases/latest)

# DDEV Wayfinder

## Overview

This add-on integrates Wayfinder into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get markdlabrecque/ddev-wayfinder
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Wayfinder |
| `ddev logs -s wayfinder` | Check Wayfinder logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.wayfinder --wayfinder-docker-image="ddev/ddev-utilities:latest"
ddev add-on get markdlabrecque/ddev-wayfinder
ddev restart
```

Make sure to commit the `.ddev/.env.wayfinder` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `WAYFINDER_DOCKER_IMAGE` | `--wayfinder-docker-image` | `ddev/ddev-utilities:latest` |

## Credits

**Contributed and maintained by [@markdlabrecque](https://github.com/markdlabrecque)**
