---
title: "pabman11/ddev-lazysql"
github_url: "https://github.com/pabman11/ddev-lazysql"
description: "DDEV add-on: open LazySQL connected to the project database with 'ddev lazysql'"
user: "pabman11"
repo: "ddev-lazysql"
repo_id: 1355648230
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ""
dependencies: []
type: "contrib"
created_at: "2026-09-03"
updated_at: "2026-09-03"
workflow_status: "unknown"
stars: 0
---

# DDEV LazySQL

[![tests](https://github.com/pabman11/ddev-lazysql/actions/workflows/tests.yml/badge.svg)](https://github.com/pabman11/ddev-lazysql/actions/workflows/tests.yml)

A [DDEV](https://ddev.com) add-on that adds a `ddev lazysql` command to open
[LazySQL](https://github.com/jorgerojas26/lazysql) connected straight to the project database.

## Requirements

- DDEV v1.24.0 or higher
- [`lazysql`](https://github.com/jorgerojas26/lazysql) installed on the host
- `jq` installed on the host

## Installation

```bash
ddev add-on get pabman11/ddev-lazysql
ddev restart
```

## Usage

With the project running:

```bash
ddev lazysql
```

The command reads the project configuration (`ddev describe -j`), gets the published port of the
database container and launches LazySQL on the host with the right connection string.
MySQL/MariaDB and PostgreSQL are supported.

## Removal

```bash
ddev add-on remove lazysql
```

## Credits

Maintained by [@pabman11](https://github.com/pabman11).
