---
title: tyler36/ddev-vitest
github_url: https://github.com/tyler36/ddev-vitest
description: "A helper add-on for projects using Vitest"
user: tyler36
repo: ddev-vitest
repo_id: 894807367
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2024-11-27
updated_at: 2025-05-27
workflow_status: disabled
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-vitest/actions/workflows/tests.yml/badge.svg)](https://github.com/tyler36/ddev-vitest/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-vite)](https://github.com/tyler36/ddev-vite/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-vite)](https://github.com/tyler36/ddev-vite/releases/latest)

# ddev-vitest <!-- omit in toc -->

- [What is ddev-vitest?](#what-is-ddev-vitest)
- [Components of the repository](#components-of-the-repository)
- [Getting started](#getting-started)
- [Commands](#commands)
  - [`ddev vitest`](#ddev-vitest)
  - [`ddev vitest-ui`](#ddev-vitest-ui)
- [Auto-start Vitest UI](#auto-start-vitest-ui)
- [Vite](#vite)

## What is ddev-vitest?

`ddev-vitest` is a helper add-on for DDEV that improves the developer experience for projects using Vitest.

[Vitest](https://vitest.dev/) describes itself as a "next generation testing framework", a fast "Vite-native" testing framework.

## Components of the repository

- [config.vitest-ui.yaml](https://github.com/tyler36/ddev-vitest/blob/main/config.vitest-ui.yaml): a configuration file that exposes the Vitest UI port.
- [commands/web/vitest](https://github.com/tyler36/ddev-vitest/blob/main/commands/web/vitest): a helper command for Vitest
- [commands/host/vitest-ui](https://github.com/tyler36/ddev-vitest/blob/main/commands/host/vitest-ui): a helper command for to launch the Vitest UI site.
- [install.yaml](https://github.com/tyler36/ddev-vitest/blob/main/install.yaml): DDEV installation file
- [test.bats](https://github.com/tyler36/ddev-vitest/blob/main/tests/test.bats): test suite to ensure the add-on is working.
- [Github actions setup](https://github.com/tyler36/ddev-vitest/blob/main/.github/workflows/tests.yml): file to automate checks and tests.

## Getting started

This add-on assumes the developer has:

- Installed Vitest via their preferred package manager.

1. Install the add-on and restart DDEV

   ```shell
   ddev add-on get tyler36/ddev-vitest
   ddev restart
   ```

## Commands

### `ddev vitest`

`ddev vitest` is a helper command to run Vitest from the host.
It accepts all flags accepted by vitest.
For example, to see the currently installed version of Vitest:

  ```shell
  ddev vitest --version
  ```

> [!NOTE]
> If you attempt to start Vitest UI via `ddev vitest --ui`, this addon hijacks the command and re-writes it to be compatible with DDEV.

### `ddev vitest-ui`

Use the following command to start the Vitest UI server and launch the site in your default browser:

```shell
ddev vitest-ui -s
```

If the server is already started, use `ddev vitest-ui` to launch the site

## Auto-start Vitest UI

Use DDEV's `post-start` hook to automatically start Vitest UI.

The following snippet starts the UI server and launches the test page.

```yaml
hooks:
  post-start:
    - exec-host: ddev vitest-ui
```

## Vite

Vitest is a great companion to Vite.
For more information about using Vite with DDEV,

- [tyler36/ddev-vite](https://github.com/tyler36/ddev-vite): a DDEV addon for Vite.
- Matthias Andrasch's excellent blog post - [Working with Vite in DDEV - an introduction](https://ddev.com/blog/working-with-vite-in-ddev/).

**Contributed and maintained by `tyler36`**
