---
title: Cambrico/ddev-mysqlshell
github_url: https://github.com/Cambrico/ddev-mysqlshell
description: "A DDEV addon to seamlessly install and manage MySQL Shell within your DDEV environment. Simplifies the integration of MySQL Shell for database management and querying in local development."
user: Cambrico
repo: ddev-mysqlshell
repo_id: 939360133
ddev_version_constraint: ">= v1.24.2"
dependencies: []
type: contrib
created_at: 2025-02-26
updated_at: 2025-03-19
workflow_status: disabled
stars: 0
---

[![tests](https://github.com/cambrico/ddev-mysqlshell/actions/workflows/tests.yml/badge.svg)](https://github.com/cambrico/ddev-mysqlshell/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# DDEV MySQL Shell

* [What is ddev-mysqlshell?](#what-is-ddev-mysql-shell)
* [Commands](#commands)

## What is DDEV MySQL Shell?

This repository provides an add-on for [DDEV](https://ddev.readthedocs.io)
to integrate MySQL Shell into your DDEV projects.

In DDEV addons can be installed from the command line using the `ddev add-on get` command (or `ddev get` for versions of DDEV prior to v1.23.5), for example, `ddev add-on get cambrico/ddev-mysqlshell`.

### Dependencies
Make sure you have [DDEV v1.24.2+ installed](https://ddev.readthedocs.io/en/latest/users/install/ddev-installation/)

## Commands
The following commands are available:

| Command          | Usage                           |
|------------------|---------------------------------|
| `mysqlsh`        | `ddev mysqlsh [options]`        |
| `mysqlsh-export` | `ddev mysqlsh-export`           |
| `mysqlsh-import` | `ddev mysqlsh-import [options]` |


**Contributed and maintained by [@cambrico](https://github.com/cambrico)**
