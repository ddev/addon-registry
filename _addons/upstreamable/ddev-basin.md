---
title: upstreamable/ddev-basin
github_url: https://github.com/upstreamable/ddev-basin
description: "Simplify project creation and ddev add-on management "
user: upstreamable
repo: ddev-basin
repo_id: 1174546726
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ">= v1.25.1"
dependencies: []
type: contrib
created_at: 2026-03-06
updated_at: 2026-04-27
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/upstreamable/ddev-basin/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/upstreamable/ddev-basin/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/upstreamable/ddev-basin)](https://github.com/upstreamable/ddev-basin/commits)
[![release](https://img.shields.io/github/v/release/upstreamable/ddev-basin)](https://github.com/upstreamable/ddev-basin/releases/latest)

# DDEV Basin

## Overview

This add-on aims to simplify add-on management for [DDEV](https://ddev.com/)
projects by maintaining a file with the installed add-ons at
`[YOUR_PROJECT_FOLDER]/.ddev/config.basin.yaml`.

A set of DDEV global commands is provided to get Drupal projects running from
scratch with `drupal/recommended-project` as base, drush, and the site
installed at the end of the process.

It also serves as a base for other addons to provide Symfony commands written
in PHP.
See [DDEV Basin Deploy](https://github.com/upstreamable/ddev-basin-deploy/tree/main/src/Command)

## Installation

```bash
ddev add-on get upstreamable/ddev-basin
ddev restart
```

Upon restart the currently installed add-ons will be included in
`.ddev/config.basin.yaml`.
After installation, make sure to commit the `.ddev/config.basin.yaml` to version control.

## Usage

With the add-ons of the project in configuration there is no need to commit all
the add-on files into your repository. The only two files required are
`.ddev/config.yaml` and `.ddev/config.basin.yaml`.
With those present and the `ddev-basin` add-on installed te rest of add-ons
will be installed when the project is started.
When an add-on is installed, the next time the project starts it will be added to
`.ddev/config.basin.yaml`.

## Credits

**Contributed and maintained by [@upstreamable](https://github.com/upstreamable)**
