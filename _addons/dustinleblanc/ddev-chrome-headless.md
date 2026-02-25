---
title: dustinleblanc/ddev-chrome-headless
github_url: https://github.com/dustinleblanc/ddev-chrome-headless
description: "An add-on for DDEV to easily add a Headless Chrome container for a variety of tasks like running functional Javascript tests"
user: dustinleblanc
repo: ddev-chrome-headless
repo_id: 1164859764
default_branch: main
tag_name: v0.1.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-23
updated_at: 2026-02-23
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/dustinleblanc/ddev-chrome-headless/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dustinleblanc/ddev-chrome-headless/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/dustinleblanc/ddev-chrome-headless)](https://github.com/dustinleblanc/ddev-chrome-headless/commits)
[![release](https://img.shields.io/github/v/release/dustinleblanc/ddev-chrome-headless)](https://github.com/dustinleblanc/ddev-chrome-headless/releases/latest)

# DDEV Chrome Headless

## Overview

This add-on integrates Chrome Headless into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get dustinleblanc/ddev-chrome-headless
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Chrome Headless |
| `ddev logs -s chrome-headless` | Check Chrome Headless logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.chrome-headless --chrome-headless-docker-image="ddev/ddev-utilities:latest"
ddev add-on get dustinleblanc/ddev-chrome-headless
ddev restart
```

Make sure to commit the `.ddev/.env.chrome-headless` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `CHROME_HEADLESS_DOCKER_IMAGE` | `--chrome-headless-docker-image` | `ddev/ddev-utilities:latest` |

## Credits

**Contributed and maintained by [@dustinleblanc](https://github.com/dustinleblanc)**
