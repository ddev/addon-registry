---
title: plopesc/ddev-glab-cli
github_url: https://github.com/plopesc/ddev-glab-cli
description: "ddev add-on for setting up the GitLab Client."
user: plopesc
repo: ddev-glab-cli
repo_id: 1192798543
default_branch: main
tag_name: v1.1.1
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-03-26
updated_at: 2026-04-07
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/plopesc/ddev-glab-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/plopesc/ddev-glab-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/plopesc/ddev-glab-cli)](https://github.com/plopesc/ddev-glab-cli/commits)
[![release](https://img.shields.io/github/v/release/plopesc/ddev-glab-cli)](https://github.com/plopesc/ddev-glab-cli/releases/latest)

# DDEV Glab Cli

## Overview

This add-on integrates GitLab Cli (glab) into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get plopesc/ddev-glab-cli
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Authentication

For persistent login, the following environment variables need to be defined:
* GITLAB_HOST
* GITLAB_TOKEN

The full list of environment variables can be found at [CLI documentation](https://docs.gitlab.com/cli/#environment-variables).

For other ways to authenticate, the following command can be executed:
```sh
ddev exec glab auth login
```

## Usage

| Command                                     | Description                              |
|---------------------------------------------|------------------------------------------|
| `ddev exec glab --version`                    | Check the installed version              |
| `ddev exec glab --help`                       | View available commands                  |

The full documentation about GitLab CLI can be gound at the [GitLab CLI documentation page](https://docs.gitlab.com/cli/).

## Customization

### Custom GitLab CLI version

By default, the add-on installs the latest available GitLab CLI version. To pin a specific version, create a `.ddev/docker-compose.glab-cli.yaml` file with the desired version:

```yaml
services:
  web:
    build:
      args:
        GLAB_VERSION: "1.80.0"
```

Then run `ddev restart` to rebuild the container with the specified version.

## Credits

**Contributed and maintained by [@plopesc](https://github.com/plopesc)**
