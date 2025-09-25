---
title: Cambrico/ddev-ansible
github_url: https://github.com/Cambrico/ddev-ansible
description: "A DDEV addon to seamlessly install and manage Ansible within your DDEV environment. Simplifies the integration of Ansible for local development and automation tasks."
user: Cambrico
repo: ddev-ansible
repo_id: 892347556
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2024-11-22
updated_at: 2025-09-24
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/cambrico/ddev-ansible/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/cambrico/ddev-ansible/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/cambrico/ddev-ansible)](https://github.com/cambrico/ddev-ansible/commits)
[![release](https://img.shields.io/github/v/release/cambrico/ddev-ansible)](https://github.com/cambrico/ddev-ansible/releases/latest)

# DDEV Ansible

## Overview

This add-on integrates Ansible into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get cambrico/ddev-ansible
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

### Dependencies
Make sure you have [DDEV v1.23.3+ installed](https://ddev.readthedocs.io/en/latest/users/install/ddev-installation/)

## Commands

The following commands are available:

| Command              | Usage                               |
|----------------------|-------------------------------------|
| `ansible`            | `ddev ansible [options]`            |
| `ansible-community`  | `ddev ansible-community [options]`  |
| `ansible-config`     | `ddev ansible-config [options]`     |
| `ansible-connection` | `ddev ansible-connection [options]` |
| `ansible-console`    | `ddev ansible-console [options]`    |
| `ansible-doc`        | `ddev ansible-doc [options]`        |
| `ansible-galaxy`     | `ddev ansible-galaxy [options]`     |
| `ansible-inventory`  | `ddev ansible-inventory [options]`  |
| `ansible-playbook`   | `ddev ansible-playbook [options]`   |
| `ansible-pull`       | `ddev ansible-pull [options]`       |
| `ansible-test`       | `ddev ansible-test [options]`       |
| `ansible-vault`      | `ddev ansible-vault [options]`      |

## Credits

**Contributed and maintained by [@cambrico](https://github.com/cambrico)**
