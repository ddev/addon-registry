---
title: qwerty-re/ddev-elasticvue
github_url: https://github.com/qwerty-re/ddev-elasticvue
description: "Elasticvue service for DDEV"
user: qwerty-re
repo: ddev-elasticvue
repo_id: 1331773490
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ""
dependencies: ["qwerty-re/ddev-elasticsearch"]
type: contrib
created_at: 2026-08-12
updated_at: 2026-08-12
workflow_status: unknown
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![last commit](https://img.shields.io/github/last-commit/qwerty-re/ddev-elasticvue)](https://github.com/qwerty-re/ddev-elasticvue/commits)
[![release](https://img.shields.io/github/v/release/qwerty-re/ddev-elasticvue)](https://github.com/qwerty-re/ddev-elasticvue/releases/latest)


# DDEV Elasticvue

## What is ddev-elasticvue?

This repository allows you to quickly install Elasticvue into a [DDEV](https://ddev.readthedocs.io) project using just `ddev get qwerty-re/ddev-elasticvue`.

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get qwerty-re/ddev-elasticvue && ddev restart
```

For earlier versions of DDEV run

```sh
ddev get qwerty-re/ddev-elasticvue && ddev restart
```

You can then visit Elasticvue by running `ddev elasticvue` or visiting the URL shown in `ddev describe`.

## Explanation

Elasticvue is a free and open-source gui for elasticsearch that you can use to manage the data in your cluster.

## Additional Resources

- [Elasticvue official website](https://elasticvue.com/).
- [Elasticvue GitHub repository](https://github.com/cars10/elasticvue).

**Elasticvue is maintained by [@cars10](https://github.com/cars10)**  
**DDEV Elasticvue is maintained by [@qwerty-re](https://github.com/qwerty-re)**
