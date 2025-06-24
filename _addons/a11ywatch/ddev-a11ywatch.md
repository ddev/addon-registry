---
title: a11ywatch/ddev-a11ywatch
github_url: https://github.com/a11ywatch/ddev-a11ywatch
description: "A11yWatch ddev addon"
user: a11ywatch
repo: ddev-a11ywatch
repo_id: 596553504
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-02-02
updated_at: 2024-10-24
workflow_status: disabled
stars: 6
---

[![tests](https://github.com/a11ywatch/ddev-a11ywatch/actions/workflows/tests.yml/badge.svg)](https://github.com/a11ywatch/ddev-a11ywatch/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## What is this?

This repository allows you to quickly install the warp speed accessibility and vast coverage tool [A11yWatch Lite](https://github.com/a11ywatch/a11ywatch) into a [Ddev](https://ddev.readthedocs.io) project using the instructions below.

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get a11ywatch/ddev-a11ywatch
```

For earlier versions of DDEV run

```sh
ddev get a11ywatch/ddev-a11ywatch
```

Afterwards, run `ddev restart`

## Explanation

This A11yWatch recipe for [ddev](https://ddev.readthedocs.io) installs a [`.ddev/docker-compose.a11ywatch-standalone.yaml`](https://github.com/a11ywatch/ddev-a11ywatch/blob/main/docker-compose.a11ywatch-standalone.yaml) using the [`A11yWatch`](https://hub.docker.com/r/a11ywatch/a11ywatch/tags) stand-alone docker image.

## Interacting with A11yWatch

* The A11ywatch instance will listen on TCP port 3280 (the A11yWatch default) and port 50050 for gRPC.
* Configure your application to access A11ywatch on the host:port `a11ywatch:3280`.

## Additional Resources

* To get detailed infromation on how to interact or commincate with the [A11yWatch API Info](https://a11ywatch.com/api-info) A11yWatch.
* The [A11yWatch CLI](https://github.com/a11ywatch/a11ywatch) is helpful to perform automated task using the gRPC client.

## Todo

1. Add web panel option start using the `a11ywatch/web` image.

**Contributed by [j-mendez](https://github.com/j-mendez)**
