---
title: haase-fabian/ddev-neos
github_url: https://github.com/haase-fabian/ddev-neos
description: "neos environment variables for ddev"
user: haase-fabian
repo: ddev-neos
repo_id: 535660811
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2022-09-12
updated_at: 2024-10-28
stars: 0
---

[![tests](https://github.com/haase-fabian/ddev-neos/actions/workflows/tests.yml/badge.svg)](https://github.com/haase-fabian/ddev-neos/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2022.svg)

## What is ddev-neos?

This repository allows you to quickly install a Neos CMS environment as a Ddev project using just ddev get haase-fabian/ddev-redis.

## Installation

For DDEV v1.23.5 or above run

```bash
ddev add-on get haase-fabian/ddev-neos
```

For earlier versions of DDEV run

```bash
ddev get haase-fabian/ddev-neos
```

Then restart your project

```bash
ddev restart
```

## Explanation

This recipe for ddev installs a .ddev/docker-compose.neos.yaml adding neos specific environment variables to the web container.
And it adds the `ddev flow` command to run ./flow commands.
