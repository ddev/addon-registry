---
title: backdrop-ops/ddev-backdrop-bee
github_url: https://github.com/backdrop-ops/ddev-backdrop-bee
description: "Install bee for Backdrop sites on ddev."
user: backdrop-ops
repo: ddev-backdrop-bee
repo_id: 873926344
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-10-17
updated_at: 2025-03-05
stars: 1
---

[![tests](https://github.com/backdrop-ops/ddev-backdrop-bee/actions/workflows/tests.yml/badge.svg)](https://github.com/backdrop-ops/ddev-backdrop-bee/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# ddev-backdrop-bee <!-- omit in toc -->

* [What is ddev-backdrop-bee?](#what-is-ddev-backdrop-bee)
* [Installation](#getting-started)

## What is ddev-backdrop-bee?

This is a DDEV addon that will install Bee for use on Backdrop CMS sites.

## What is Backdrop CMS Bee?

Bee is a command line utility for Backdrop CMS. It includes commands that allow developers to interact with Backdrop sites, performing actions like:
- Running cron
- Clearing caches
- Downloading and installing Backdrop
- Downloading, enabling and disabling projects
- Viewing information about a site and/or available projects
See the Release notes and the Changelog for details of changes between versions
[Backdrop CMS Bee Project](https://backdropcms.org/project/bee)
[Github Project for Bee](https://github.com/backdrop-contrib/bee)

## Installation

1. Install this addon
    For DDEV v1.23.5 or above run

    ```sh
    ddev add-on get backdrop-ops/ddev-backdrop-bee
    ```

    For earlier versions of DDEV run

    ```sh
    ddev get backdrop-ops/ddev-backdrop-bee
    ```

2. Restart DDEV with `ddev restart`.
3. Test bee by running `ddev bee status`.


**Contributed and maintained by `@jenlampton` and `@wylbur`**
