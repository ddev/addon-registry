---
title: massiws/ddev-bearer
github_url: https://github.com/massiws/ddev-bearer
description: "Integrate Bearer into DDEV to scan your projects against top security and privacy risks"
user: massiws
repo: ddev-bearer
repo_id: 1156317873
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-02-12
updated_at: 2026-02-14
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/massiws/ddev-bearer/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/massiws/ddev-bearer/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/massiws/ddev-bearer)](https://github.com/massiws/ddev-bearer/commits)
[![release](https://img.shields.io/github/v/release/massiws/ddev-bearer)](https://github.com/massiws/ddev-bearer/releases/latest)

# DDEV Bearer

## Overview

[Bearer](https://github.com/bearer/bearer) CLI is a static application security testing (SAST) tool.

It scans your source code and analyzes your data flows to discover, filter and prioritize security and privacy risks, with built in rules to cover the [OWASP Top 10](https://owasp.org/www-project-top-ten/) and [CWE Top 25](https://cwe.mitre.org/top25/archive/2023/2023_top25_list.html).

## Installation

```bash
# Install the add-on
ddev add-on get massiws/ddev-bearer

# Restart DDEV
ddev restart

# Generates a default config to `bearer.yml`
ddev bearer init
```

After installation, you may want to customize the default configuration settings by editing the `bearer.yml` file:
see [docs](https://docs.bearer.com/reference/config/) for more information.

**Important**: restart DDEV after making changes to the `bearer.yml` file.

Make sure to commit the `.ddev` directory and the `bearer.yml` file to version control.


## Usage

Display available commands and usage information:
  ```bash
  ddev bearer
  ```

Scan project using default configuration in `bearer.yml`:
  ```bash
  ddev bearer scan .
  ```

Scan project only for specified [Severity Levels](https://docs.bearer.com/guides/configure-scan/#limit-severity-levels):
  ```bash
  ddev bearer scan . --severity critical,high
  ```

Scan all project files searching for hardcoded credentials (see [Scanner Types](https://docs.bearer.com/explanations/scanners/):
  ```bash
  ddev bearer scan . --scanner=secrets
  ```

Scan a specific file or folder, also adding _context_ (see [Bearer Flags](https://docs.bearer.com/reference/commands/#bearer%20scan-flags))
  ```bash
  ddev bearer scan <file/path> --context=health
  ```

Ignore a reported risk adding the fingerprint to your ignore file:
  ```bash
  ddev bearer ignore add <fingerprint> --author Mish --comment "Possible false positive"
  ```
**TIP**: to avoid specify _author_ each time, you may want to configure your [git username in DDEV globals](https://docs.ddev.com/en/stable/users/extend/in-container-configuration/#git-configuration):
  ```bash
  ln -s $HOME/.gitconfig $HOME/.ddev/homeadditions/.gitconfig
  ```

See Bearer documentation for a complete list of [commands](https://docs.bearer.com/reference/commands) and [flags](https://docs.bearer.com/reference/commands/#bearer%20init-flags).

## Credits

**Contributed and maintained by [@massiws](https://github.com/massiws)**
