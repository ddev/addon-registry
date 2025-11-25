---
title: Cambrico/ddev-mysqlshell
github_url: https://github.com/Cambrico/ddev-mysqlshell
description: "A DDEV addon to seamlessly install and manage MySQL Shell within your DDEV environment. Simplifies the integration of MySQL Shell for database management and querying in local development."
user: Cambrico
repo: ddev-mysqlshell
repo_id: 939360133
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-02-26
updated_at: 2025-09-24
workflow_status: disabled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/cambrico/ddev-mysqlshell/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/cambrico/ddev-mysqlshell/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/cambrico/ddev-mysqlshell)](https://github.com/cambrico/ddev-mysqlshell/commits)
[![release](https://img.shields.io/github/v/release/cambrico/ddev-mysqlshell)](https://github.com/cambrico/ddev-mysqlshell/releases/latest)

# DDEV MySQL Shell

## Overview

This add-on integrates MySQL Shell into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get cambrico/ddev-mysqlshell
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Commands

The following commands are available:

| Command          | Usage                           |
|------------------|---------------------------------|
| `mysqlsh`        | `ddev mysqlsh [options]`        |
| `mysqlsh-export` | `ddev mysqlsh-export`           |
| `mysqlsh-import` | `ddev mysqlsh-import [options]` |

## Credits

**Contributed and maintained by [@cambrico](https://github.com/cambrico)**
