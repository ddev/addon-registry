---
title: tyler36/ddev-gitpod-setup
github_url: https://github.com/tyler36/ddev-gitpod-setup
description: "Helper add-on to add Gitpod config."
user: tyler36
repo: ddev-gitpod-setup
repo_id: 840226121
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-08-09
updated_at: 2025-01-09
stars: 1
---

[![tests](https://github.com/tyler36/ddev-gitpod-setup/actions/workflows/tests.yml/badge.svg)](https://github.com/tyler36/ddev-gitpod-setup/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2026.svg)

# ddev-gitpod-setup <!-- omit in toc -->

- [What is ddev-gitpod-setup?](#what-is-ddev-gitpod-setup)
- [Getting started](#getting-started)
- [Customizing](#customizing)
    - [Full control](#full-control)
    - [Project-type customizing](#project-type-customizing)
- [Components of the repository](#components-of-the-repository)
- [Contributing](#contributing)

## What is ddev-gitpod-setup?

ddev-gitpod-setup is a helper add-on to set up Gitpod for a DDEV project.
It does this by generating a `.gitpod.yml` file which includes:

- generic framework-specific tasks
- DDEV base image
- DDEV port settings
- VSCode extensions
- GitHub preferences

@see [Gitpod documentation](https://www.gitpod.io/docs/references/gitpod-yml) for options settings. Most people don't need to change anything after just getting this setup.

## Getting started

1. Install add-on

    For DDEV v1.23.5 or above run

    ```shell
    ddev add-on get tyler36/ddev-gitpod-setup
    ```

    For earlier versions of DDEV run

    ```shell
    ddev get tyler36/ddev-gitpod-setup
    ```

2. Commit files to repository. All that really matters is the .gitpod.yml, but it does no harm to commit all the files the add-on creates.
3. Open the project in Gitpod. You can launch your project many ways with Gitpod, but one is by constructing a URL like `https://gitpod.io/?autostart=true#https://github.com/<org>/<project>` and launching it in your browser.

@see [Gitpod docs](https://www.gitpod.io/docs/introduction/getting-started) to see how this addon based a `.gitpod.yml` file on step 3.

## Customizing

### Full control

You can take over the `.gitpod.yml` and change it as you see fit, but most people don't need to do that.

- Remove `#ddev-generated` from `.gitpod.yml`. (This add-on will no longer manage the file.)
- Make changes, as require.

You are then responsible for all updates to the configuration such as updating the DDEV version.
You will need to manually remove the file if/when you remove this add-on.

### Project-type customizing

Currently this add-on has explicit support for Drupal and Laravel projects, with fallback to generic projects (DDEV's `php` project type). Contributions of support for new project types is welcome.

- Copy an existing `.ddev/gitpod-setup/{project-type}.yml`
- Update `.ddev/gitpod-setup/{project-type}.yml` with tasks your project type needs.
- Run `ddev get tyler36/ddev-gitpod-setup` to generate a new config file.
- Contribute your changes back to this add-on.

This approach works best if you want to:

- customize Gitpod tasks.
- allow this add-on to update the DDEV image, ports and other "standard" DDEV configuration.
- reuse the project-type settings in other projects.

## Components of the repository

- [gitpod-setup](https://github.com/tyler36/ddev-gitpod-setup/blob/main/gitpod-setup/) directory holds the files used to construct the final `.gitpod.yml` file.
- An [install.yaml](https://github.com/tyler36/ddev-gitpod-setup/blob/main/install.yaml) file that describes how to install the service or other component.
- A test suite in [test.bats](https://github.com/tyler36/ddev-gitpod-setup/blob/main/tests/test.bats) that makes sure the service continues to work as expected.
- [Github actions setup](https://github.com/tyler36/ddev-gitpod-setup/blob/main/.github/workflows/tests.yml) so that the tests run automatically when you push to the repository.

## Contributing

PRs are welcome, especially when accompanied with tests.

Project-type files are also welcome. However, please take care to keep them as minimal and un-opinionated.
We all have our preferred setup environment and "essential" extensions and add-ons. They are outside this add-ons scope. @see [customizations](#customizing).

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
