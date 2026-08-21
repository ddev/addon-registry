---
title: "tyler36/ddev-fakecloud"
github_url: "https://github.com/tyler36/ddev-fakecloud"
description: "Local AWS cloud emulator for integration tests."
user: "tyler36"
repo: "ddev-fakecloud"
repo_id: 1327500355
default_branch: "main"
tag_name: "v0.1"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-08-08"
updated_at: "2026-08-08"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/fakecloud/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tyler36/fakecloud/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tyler36/fakecloud)](https://github.com/tyler36/fakecloud/commits)
[![release](https://img.shields.io/github/v/release/tyler36/fakecloud)](https://github.com/tyler36/fakecloud/releases/latest)

# DDEV Fakecloud

## Overview

This add-on integrates [Fakecloud](https://fakecloud.dev/) into your [DDEV](https://ddev.com/) project.

> fakecloud gives you a local AWS environment that behaves like infrastructure, not a mock. Your app uses the regular AWS SDK, CLI, and IaC tools. Unlike today's LocalStack Community setup, you do not need an account, auth token, or paid plan just to keep core development flows local.

## Installation

```bash
ddev add-on get tyler36/fakecloud
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

This addon:

- installs `aws-cli`
- adds [helper commands](#commands)

### Commands

| Command                  | Description                                             |
| ------------------------ | ------------------------------------------------------- |
| `ddev fakecloud`         | Command to run integrated fakecloud, overrides aws ENVs |
|                          | Aliases: `ddev aws-fake`, `ddev fc`                     |
| `ddev aws`               | Command to run integrated AWS-CLI.                      |
| `ddev describe`          | View service status and used ports for Fakecloud        |
| `ddev logs -s fakecloud` | Check Fakecloud logs                                    |

> [!NOTE]
> This addon installs `aws-cli` inside the web container.
> You are responsible for managing its authentication and secrets.

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.fakecloud --fakecloud-docker-image="ghcr.io/faiscadev/fakecloud:latest"
ddev add-on get tyler36/fakecloud
ddev restart
```

Make sure to commit the `.ddev/.env.fakecloud` file to version control.

All customization options (use with caution):

| Variable                 | Flag                       | Default                              |
| ------------------------ | -------------------------- | ------------------------------------ |
| `FAKECLOUD_DOCKER_IMAGE` | `--fakecloud-docker-image` | `ghcr.io/faiscadev/fakecloud:latest` |

## Credits

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
