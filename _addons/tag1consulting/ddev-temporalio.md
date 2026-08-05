---
title: tag1consulting/ddev-temporalio
github_url: https://github.com/tag1consulting/ddev-temporalio
description: "ddev addon for temporal.io"
user: tag1consulting
repo: ddev-temporalio
repo_id: 1301945557
default_branch: main
tag_name: 0.19
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-07-15
updated_at: 2026-08-04
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tag1consulting/ddev-temporalio/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tag1consulting/ddev-temporalio/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tag1consulting/ddev-temporalio)](https://github.com/tag1consulting/ddev-temporalio/commits)
[![release](https://img.shields.io/github/v/release/tag1consulting/ddev-temporalio)](https://github.com/tag1consulting/ddev-temporalio/releases/latest)

# DDEV Temporal.io

## Overview

This add-on integrates [Temporal.io](https://temporal.io) into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get tag1consulting/ddev-temporalio
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Temporalio |
| `ddev logs -s temporalio` | Check Temporalio logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.temporalio --temporalio-docker-image="ddev/ddev-utilities:latest"
ddev add-on get tag1consulting/ddev-temporalio
ddev restart
```

Make sure to commit the `.ddev/.env.temporalio` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `TEMPORALIO_DOCKER_IMAGE` | `--temporalio-docker-image` | `ddev/ddev-utilities:latest` |

## Web UI

The Temporal UI can be found on https://temporalio-ui.sitename.ddev.site/

## Credits

**Contributed and maintained by [@tag1consulting](https://github.com/tag1consulting) with sponsorship from [Tag1 Consulting](https://tag1consulting.com/)**
