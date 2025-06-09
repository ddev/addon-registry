---
title: tyler36/ddev-storybook
github_url: https://github.com/tyler36/ddev-storybook
description: "Helpers for DDEV and Storybook"
user: tyler36
repo: ddev-storybook
repo_id: 746948704
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2024-01-23
updated_at: 2025-06-09
stars: 8
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-storybook/actions/workflows/tests.yml/badge.svg)](https://github.com/tyler36/ddev-storybook/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-storybook)](https://github.com/tyler36/ddev-storybook/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-storybook)](https://github.com/tyler36/ddev-storybook/releases/latest)

# ddev-storybook <!-- omit in toc -->

- [What is ddev-storybook?](#what-is-ddev-storybook)
- [Components of the repository](#components-of-the-repository)
- [Getting started](#getting-started)
- [Usage](#usage)
    - [Open Storybook](#open-storybook)
    - [Start Storybook server](#start-storybook-server)
- [Contributing](#contributing)

## What is ddev-storybook?

This addon makes it easier to work with Storybook 7 and DDEV.
It does this by:

- exposing Storybook on port 6006.
- adding a helper command.

[Storybook](https://storybook.js.org/) is "a frontend workshop for building UI components".
The addon assumes the developer is correctly installed and configured Storybook using the default port (`6006`).

## Components of the repository

- `config.storybook.yaml`: exposes the default Storybook port.
- `commands/host/storybook`: Helper command

## Getting started

1. Install the add-on and restart DDEV:

```shell
ddev add-on get tyler36/ddev-storybook
ddev restart
```

2. Install storybook if you haven't already. See the [Storybook get started page](https://storybook.js.org/docs/get-started/install) for instructions. E.g.
```shell
ddev exec npx storybook@latest init
```
**Note**: In some cases, Storybook may not launch correctly after the initial installation. To help prevent errors like `Error: spawn xdg-open ENOENT`, try using the `--no-dev` flag during initialization:

```shell
ddev exec npx storybook@latest init --no-dev
```

## Usage

The helper function has 2 basic usage:

### Open Storybook

Use the following command to open Storybook in the default browser.

```shell
ddev storybook
```

### Start Storybook server

Use the following command to start the Storybook server.

```shell
ddev storybook -s
```

Alternatively,

```shell
ddev storybook --start
```

This command will try to run the "storybook" script from `package.json` in NPM.
If a `yarn.lock` file exists, it will run the command with YARN instead.

### Removing the Add-On

```
ddev add-on remove tyler36/ddev-storybook
```

## Contributing

This addon was originally created to work with Drupal and [storybook](https://www.drupal.org/project/storybook) contrib module. It has _not_ been tested in other environments, although assuming Storybook defaults, it should work.

PRs, especially with automated tests are welcome.

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
