---
title: upstreamable/ddev-nextcloud-fpm
github_url: https://github.com/upstreamable/ddev-nextcloud-fpm
description: "Integrate the FPM docker image into the nginx-fpm server from DDEV"
user: upstreamable
repo: ddev-nextcloud-fpm
repo_id: 1190971050
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ">= v1.25.1"
dependencies: ["ddev/ddev-redis"]
type: contrib
created_at: 2026-03-24
updated_at: 2026-04-07
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/upstreamable/ddev-nextcloud-fpm/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/upstreamable/ddev-nextcloud-fpm/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/upstreamable/ddev-nextcloud-fpm)](https://github.com/upstreamable/ddev-nextcloud-fpm/commits)
[![release](https://img.shields.io/github/v/release/upstreamable/ddev-nextcloud-fpm)](https://github.com/upstreamable/ddev-nextcloud-fpm/releases/latest)

# DDEV Nextcloud FPM

## Overview

This add-on integrates [Nextcloud](https://nextcloud.com/)
Official FPM Docker image into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get upstreamable/ddev-nextcloud-fpm
ddev restart
```

Note that the config provided by this addon forces the `webserver_type` to be `nginx-fpm`
since it is the only configuration supported.

The php version configured is ignored since it will run the recommended php version by the nextcloud
community shipped in the image. Due to this other development features such as xdebug will not work
for nextcloud.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Nextcloud FPM |
| `ddev logs -s nextcloud` | Check Nextcloud Fpm logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.nextcloud --nextcloud-docker-image="nextcloud:33.0.2-fpm"
ddev add-on get upstreamable/ddev-nextcloud-fpm
ddev restart
```

Make sure to commit the `.ddev/.env.nextcloud` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `NEXTCLOUD_DOCKER_IMAGE` | `--nextcloud-docker-image` | `nextcloud:33.0.2-fpm"` |

## Credits

**Contributed and maintained by [@upstreamable](https://github.com/upstreamable)**
