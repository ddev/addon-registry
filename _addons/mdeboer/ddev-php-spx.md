---
title: mdeboer/ddev-php-spx
github_url: https://github.com/mdeboer/ddev-php-spx
description: "DDEV addon to provide php-spx profiler extension for PHP"
user: mdeboer
repo: ddev-php-spx
repo_id: 1199485951
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-04-02
updated_at: 2026-04-02
workflow_status: success
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/mdeboer/ddev-php-spx/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/mdeboer/ddev-php-spx/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/mdeboer/ddev-php-spx)](https://github.com/mdeboer/ddev-php-spx/commits)
[![release](https://img.shields.io/github/v/release/mdeboer/ddev-php-spx)](https://github.com/mdeboer/ddev-php-spx/releases/latest)

# DDEV php-spx (Simple Profiling eXtension)

## Overview

This add-on integrates [php-spx](https://github.com/NoiseByNorthwest/php-spx) (Simple Profiling eXtension) into
your [DDEV](https://ddev.com/)
project.

## Installation

```bash
ddev add-on get mdeboer/ddev-php-spx
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command           | Description                               |
|-------------------|-------------------------------------------|
| `ddev spx status` | See whether SPX extension has been loaded |
| `ddev spx on`     | Enable SPX extension                      |
| `ddev spx off`    | Disable SPX extension                     |

For information about the usage of SPX itself, please refer to the official documentation
at https://github.com/NoiseByNorthwest/php-spx.

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.php-spx --php-spx-docker-image="ddev/ddev-utilities:latest"
ddev add-on get mdeboer/ddev-php-spx
ddev restart
```

Make sure to commit the `.ddev/.env.php-spx` file to version control.

All customization options (use with caution):

| Variable               | Flag                     | Default                      |
|------------------------|--------------------------|------------------------------|
| `PHP_SPX_DOCKER_IMAGE` | `--php-spx-docker-image` | `ddev/ddev-utilities:latest` |

## Credits

**Contributed and maintained by [@mdeboer](https://github.com/mdeboer)**
