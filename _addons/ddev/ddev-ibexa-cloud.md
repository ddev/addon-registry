---
title: ddev/ddev-ibexa-cloud
github_url: https://github.com/ddev/ddev-ibexa-cloud
description: "Pull projects and data from Ibexa Cloud into DDEV"
user: ddev
repo: ddev-ibexa-cloud
repo_id: 869204136
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: official
created_at: 2024-10-07
updated_at: 2025-04-18
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/ddev/ddev-ibexa-cloud/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/ddev/ddev-ibexa-cloud/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/ddev/ddev-ibexa-cloud)](https://github.com/ddev/ddev-ibexa-cloud/commits)
[![release](https://img.shields.io/github/v/release/ddev/ddev-ibexa-cloud)](https://github.com/ddev/ddev-ibexa-cloud/releases/latest)

# DDEV Ibexa Cloud

## Overview

The [Ibexa Cloud](https://www.ibexa.co/products/ibexa-cloud) has its own CLI, instead of using the `platform` CLI.

This add-on provides integration for Ibexa Cloud and adds the `ddev ibexa_cloud` custom command into your [DDEV](https://ddev.com/) project.

## Installation

1. Configure your Ibexa project for DDEV if you haven't already, see [DDEV Ibexa Quickstart](https://ddev.readthedocs.io/en/stable/users/quickstart/#ibexa-dxp)
2. `ddev add-on get ddev/ddev-ibexa-cloud`
3. Configure your `IBEXA_CLI_TOKEN` globally:

   ```bash
   ddev config global --web-environment-add=IBEXA_CLI_TOKEN=nf4amudfn23biyourtoken
   ```

4. (Optional) To use multiple API tokens for different projects, add them to your per-project configuration using the [.ddev/config.local.yaml](https://ddev.readthedocs.io/en/stable/users/configuration/config/#environmental-overrides) file instead. This file is gitignored by default.

   ```yaml
   web_environment:
    - IBEXA_CLI_TOKEN=nf4amudfn23biyourtoken
   ```

5. Configure your `IBEXA_PROJECT`, `IBEXA_ENVIRONMENT`, and `IBEXA_APP` environment variables, for example:

   ```bash
   ddev config --web-environment-add=IBEXA_PROJECT=nf4amudfn23biyourproject,IBEXA_ENVIRONMENT=main,IBEXA_APP=app
   ```

6. `ddev restart`
7. `ddev pull ibexa-cloud`

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev ibexa_cloud` | Run `ibexa_cloud` CLI inside the web container |
| `ddev pull ibexa-cloud` | Pull the configured environment |

## Running Automated Tests Locally

* `IBEXA_CLI_TOKEN`, `IBEXA_PROJECT` and `IBEXA_ENVIRONMENT` should exist in the environment
* `brew tap kaos/shell && brew install bats-assert bats-file`

## Credits

**Contributed and maintained by [@rfay](https://github.com/rfay)**
