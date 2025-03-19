---
title: valicm/ddev-dynamodb
github_url: https://github.com/valicm/ddev-dynamodb
description: "DynamoDB container for DDEV"
user: valicm
repo: ddev-dynamodb
repo_id: 696750542
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-09-26
updated_at: 2025-03-18
stars: 0
---

# DDEV DynamoDB Service

## What is this?

This repository allows you to quickly install DynamoDB database manager into a [Ddev](https://ddev.readthedocs.io) project using just `ddev get valicm/ddev-dynamodb`.


## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get valicm/ddev-dynamodb && ddev restart
```

For earlier versions of DDEV run

```sh
ddev get valicm/ddev-dynamodb && ddev restart
```
