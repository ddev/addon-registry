---
title: ddev/ddev-memcached
github_url: https://github.com/ddev/ddev-memcached
description: "High-performance Memcached caching service for DDEV"
user: ddev
repo: ddev-memcached
repo_id: 449470423
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: official
created_at: 2022-01-18
updated_at: 2025-04-17
workflow_status: success
stars: 3
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/ddev/ddev-memcached/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/ddev/ddev-memcached/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/ddev/ddev-memcached)](https://github.com/ddev/ddev-memcached/commits)
[![release](https://img.shields.io/github/v/release/ddev/ddev-memcached)](https://github.com/ddev/ddev-memcached/releases/latest)

# DDEV Memcached

## Overview

[Memcached](https://memcached.org/) is a free & open source, high-performance, distributed memory object caching system.

This add-on integrates Memcached into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get ddev/ddev-memcached
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

* The Memcached instance will listen on TCP port 11211 (the Memcached default).
* Configure your application to access Memcached on the host:port `memcached:11211`.
* To reach the Memcached admin interface, run `ddev ssh` to connect to the web container, then use `nc` or `telnet` to connect to the Memcached container on port 11211, i.e. `nc memcached 11211`. You can then run commands such as `stats` to see usage information. See [cheatsheet](https://lzone.de/cheat-sheet/memcached) for more commands.

## Advanced Customization

To change the docker image:

```bash
ddev dotenv set .ddev/.env.memcached --memcached-docker-image=memcached:1.6
ddev add-on get ddev/ddev-memcached
ddev restart
```

Make sure to commit the `.ddev/.env.memcached` file to version control.

## Credits

**Maintained by the [DDEV team](https://ddev.com/support-ddev/)**
