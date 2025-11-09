---
title: luho91/e2e-testing
github_url: https://github.com/luho91/e2e-testing
description: ""
user: luho91
repo: e2e-testing
repo_id: 1091727900
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-11-07
updated_at: 2025-11-07
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/luho91/e2e-testing/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/luho91/e2e-testing/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/luho91/e2e-testing)](https://github.com/luho91/e2e-testing/commits)
[![release](https://img.shields.io/github/v/release/luho91/e2e-testing)](https://github.com/luho91/e2e-testing/releases/latest)

# DDEV E2e Testing

## Overview

This add-on integrates E2e Testing into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get luho91/e2e-testing
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for E2e Testing |
| `ddev logs -s e2e-testing` | Check E2e Testing logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.e2e-testing --e2e-testing-docker-image="ddev/ddev-utilities:latest"
ddev add-on get luho91/e2e-testing
ddev restart
```

Make sure to commit the `.ddev/.env.e2e-testing` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `E2E_TESTING_DOCKER_IMAGE` | `--e2e-testing-docker-image` | `ddev/ddev-utilities:latest` |

## Credits

**Contributed and maintained by [@luho91](https://github.com/luho91)**
