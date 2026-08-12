---
title: MetaSyntactical/ddev-nub
github_url: https://github.com/MetaSyntactical/ddev-nub
description: "DDEV add-on that installs the Nub CLI into the web container"
user: MetaSyntactical
repo: ddev-nub
repo_id: 1325533045
default_branch: main
tag_name: v0.2.0
ddev_version_constraint: ">=v1.23.0"
dependencies: []
type: contrib
created_at: 2026-08-06
updated_at: 2026-08-10
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![release](https://img.shields.io/github/v/release/MetaSyntactical/ddev-nub)](https://github.com/MetaSyntactical/ddev-nub/releases/latest)
[![license](https://img.shields.io/github/license/MetaSyntactical/ddev-nub)](LICENSE)

# DDEV Nub

## Overview

Nub is a Node.js CLI tool, published as the `@nubjs/nub` npm package.

This add-on installs the Nub CLI into the DDEV web container so it's available for every project without having to install it globally on the host.

## Installation

```bash
ddev add-on get MetaSyntactical/ddev-nub
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

* `ddev nub [command]` runs any Nub CLI command inside the web container.
* `ddev nubx [command]` runs Nub's package-runner (`nubx` is to `nub` what `npx` is to `npm`) inside the web container.

## Configuration

### Pinning the Nub version

By default, the add-on installs the latest Nub release every time the web image is built. To pin a specific version instead:

```bash
ddev dotenv set .ddev/.env.nub --nub-version=<version>
ddev restart
```

For example:

```bash
ddev dotenv set .ddev/.env.nub --nub-version=1.2.3
ddev restart
```

Leave `.ddev/.env.nub` absent, or its value empty, to keep installing the latest release on every build.

Make sure to commit `.ddev/.env.nub` to version control once you've pinned a version.

## Credits

**Maintained by [MetaSyntactical](https://github.com/MetaSyntactical)**
