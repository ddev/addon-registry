---
title: Cambrico/ddev-ansible
github_url: https://github.com/Cambrico/ddev-ansible
description: "A DDEV addon to seamlessly install and manage Ansible within your DDEV environment. Simplifies the integration of Ansible for local development and automation tasks."
user: Cambrico
repo: ddev-ansible
repo_id: 892347556
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-11-22
updated_at: 2024-11-22
workflow_status: disabled
stars: 0
---

[![tests](https://github.com/cambrico/ddev-ansible/actions/workflows/tests.yml/badge.svg)](https://github.com/cambrico/ddev-ansible/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# DDEV Ansible

* [What is ddev-ansible?](#what-is-ddev-ansible)
* [Commands](#commands)

## What is DDEV Ansible?

This repository provides an add-on for [DDEV](https://ddev.readthedocs.io)
to integrate Ansible into your DDEV projects.

In DDEV addons can be installed from the command line using the `ddev add-on get` command (or `ddev get` for versions of DDEV prior to v1.23.5), for example, `ddev add-on get cambrico/ddev-ansible`.

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

**Contributed and maintained by [@cambrico](https://github.com/cambrico)**
