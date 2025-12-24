---
title: FreelyGive/ddev-dotnet
github_url: https://github.com/FreelyGive/ddev-dotnet
description: "Update ddev web container for dotnet development"
user: FreelyGive
repo: ddev-dotnet
repo_id: 1049119383
default_branch: main
tag_name: 0.1
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-09-02
updated_at: 2025-09-02
workflow_status: disabled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/FreelyGive/ddev-dotnet/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/FreelyGive/ddev-dotnet/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/FreelyGive/ddev-dotnet)](https://github.com/FreelyGive/ddev-dotnet/commits)
[![release](https://img.shields.io/github/v/release/FreelyGive/ddev-dotnet)](https://github.com/FreelyGive/ddev-dotnet/releases/latest)

# DDEV Dotnet

## Overview

This add-on integrates Dotnet into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get FreelyGive/ddev-dotnet
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command               | Description                                |
|-----------------------|--------------------------------------------|
| `ddev dotnet`         | Run dotnet commands in your web container. |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.web --dotnet-docker-image="mcr.microsoft.com/dotnet/sdk:6.0"
ddev add-on get FreelyGive/ddev-dotnet
ddev restart
```

Make sure to commit the `.ddev/.env.web` file to version control.

All customization options (use with caution):

| Variable              | Flag                    | Default                            |
|-----------------------|-------------------------|------------------------------------|
| `DOTNET_DOCKER_IMAGE` | `--dotnet-docker-image` | `mcr.microsoft.com/dotnet/sdk:9.0` |
| `DOTNET_ENVIRONMENT`  | `--dotnet-environment`  | `Development`                      |

## Credits

**Contributed and maintained by [@FreelyGive](https://github.com/FreelyGive)**
