---
title: SicseDev/ddev-sicse-toolkit
github_url: https://github.com/SicseDev/ddev-sicse-toolkit
description: ""
user: SicseDev
repo: ddev-sicse-toolkit
repo_id: 1226571373
default_branch: main
tag_name: 0.0.1
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-05-01
updated_at: 2026-05-02
workflow_status: success
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/SicseDev/ddev-sicse-toolkit/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/SicseDev/ddev-sicse-toolkit/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/SicseDev/ddev-sicse-toolkit)](https://github.com/SicseDev/ddev-sicse-toolkit/commits)
[![release](https://img.shields.io/github/v/release/SicseDev/ddev-sicse-toolkit)](https://github.com/SicseDev/ddev-sicse-toolkit/releases/latest)

# DDEV Sicse Toolkit

## Overview

This add-on integrates Sicse Toolkit into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get SicseDev/ddev-sicse-toolkit
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Sicse Toolkit |
| `ddev logs -s sicse-toolkit` | Check Sicse Toolkit logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.sicse-toolkit --sicse-toolkit-docker-image="ddev/ddev-utilities:latest"
ddev add-on get SicseDev/ddev-sicse-toolkit
ddev restart
```

Make sure to commit the `.ddev/.env.sicse-toolkit` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `SICSE_TOOLKIT_DOCKER_IMAGE` | `--sicse-toolkit-docker-image` | `ddev/ddev-utilities:latest` |

## Credits

**Contributed and maintained by [@SicseDev](https://github.com/SicseDev)**
