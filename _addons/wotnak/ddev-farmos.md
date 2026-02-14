---
title: wotnak/ddev-farmos
github_url: https://github.com/wotnak/ddev-farmos
description: ""
user: wotnak
repo: ddev-farmos
repo_id: 714356238
default_branch: main
tag_name: 4
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2023-11-04
updated_at: 2026-02-13
workflow_status: disabled
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com) [![tests](https://github.com/wotnak/ddev-farmos/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/wotnak/ddev-farmos/actions/workflows/tests.yml?query=branch%3Amain) [![last commit](https://img.shields.io/github/last-commit/wotnak/ddev-farmos)](https://github.com/wotnak/ddev-farmos/commits) [![release](https://img.shields.io/github/v/release/wotnak/ddev-farmos)](https://github.com/wotnak/ddev-farmos/releases/latest)

# ddev-farmos

[DDEV](https://ddev.com) addon for [farmOS](https://farmos.org/). It will configure DDEV project for running farmOS.

## How to use

To start farmOS project using DDEV, run the following commands:

```shell
mkdir farmos
cd farmos
ddev config --project-type=php
ddev add-on get wotnak/ddev-farmos
ddev restart
ddev composer create-project farmos/project:4.x-dev
ddev drush site:install farm --account-name=admin --account-pass=admin -y
ddev launch
```
