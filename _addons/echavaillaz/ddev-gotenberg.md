---
title: echavaillaz/ddev-gotenberg
github_url: https://github.com/echavaillaz/ddev-gotenberg
description: "Gotenberg service for DDEV."
user: echavaillaz
repo: ddev-gotenberg
repo_id: 633840053
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-04-28
updated_at: 2025-02-01
stars: 2
---

[![tests](https://github.com/echavaillaz/ddev-gotenberg/actions/workflows/tests.yml/badge.svg)](https://github.com/echavaillaz/ddev-gotenberg/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## What is ddev-gotenberg?

This repository allows you to quickly install [Gotenberg](https://gotenberg.dev) into a [DDEV](https://ddev.readthedocs.io) project using the instructions below.

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get echavaillaz/ddev-gotenberg
```

For earlier versions of DDEV run

```sh
ddev get echavaillaz/ddev-gotenberg
```

Then restart the project

```sh
ddev restart
```

## Explanation

This Gotenberg recipe for [DDEV](https://ddev.readthedocs.io) installs a [`.ddev/docker-compose.gotenberg.yaml`](https://github.com/echavaillaz/ddev-gotenberg/blob/main/docker-compose.gotenberg.yaml) using the [`gotenberg/gotenberg`](https://hub.docker.com/r/gotenberg/gotenberg) Docker image.

## Interacting with Gotenberg

* The Gotenberg instance will listen on TCP port 300 (the Gotenberg default).
* Configure your application to access Gotenberg on the host:port `gotenberg:3000`.


## Resources

1. [Official documentation](https://gotenberg.dev)
2. [Official repository](https://github.com/gotenberg/gotenberg)
3. [Awesome Gotenberg](https://github.com/gotenberg/awesome-gotenberg)
