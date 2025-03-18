---
title: feuertelegraph/ddev-firebird
github_url: https://github.com/feuertelegraph/ddev-firebird
description: "Firebird database for DDEV"
user: feuertelegraph
repo: ddev-firebird
repo_id: 674353134
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-08-03
updated_at: 2025-03-17
stars: 0
---

## What is this?

This repository allows you to quickly install a Firebird 2.5 SS database into a [DDEV](https://ddev.readthedocs.io) project using just `ddev get feuertelegraph/ddev-firebird`.

## Installation

1. `ddev get feuertelegraph/ddev-firebird`
2. `ddev restart`

## Explanation

This Firebird recipe for [DDEV](https://ddev.readthedocs.io) installs a [`.ddev/docker-compose.firebird.yaml`](https://github.com/feuertelegraph/ddev-firebird/blob/main/docker-compose.firebird.yaml) using the `jacobalberty/firebird:v2.5.9-ss` Docker image.

## Interacting with the Firebird database

* The Firebird database will listen on TCP port 3050 (the Firebird default).
* Configure your application to access the database on the host `firebird` and the path to the database is `/firebird/data/db.fdb`.
