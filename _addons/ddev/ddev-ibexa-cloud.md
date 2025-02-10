---
title: ddev/ddev-ibexa-cloud
github_url: https://github.com/ddev/ddev-ibexa-cloud
description: "Support `ddev pull ibexa-cloud`"
user: ddev
repo: ddev-ibexa-cloud
categories:
  - official
created_at: 2024-10-07
updated_at: 2024-10-23
stars: 0
---

[![tests](https://github.com/ddev/ddev-ibexa-cloud/actions/workflows/tests.yml/badge.svg)](https://github.com/ddev/ddev-ibexa-cloud/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

The [Ibexa Cloud](https://www.ibexa.co/products/ibexa-cloud) has its own CLI, instead of using the `platform` CLI.

This add-on provides integration for Ibexa Cloud and adds the `ddev ibexa_cloud` custom command.

1. Configure your Ibexa project for DDEV if you haven't already, see [DDEV Ibexa Quickstart](https://ddev.readthedocs.io/en/stable/users/quickstart/#ibexa-dxp)
2. `ddev get ddev/ddev-ibexa-cloud` (# or in DDEV v1.23.5+ `ddev add-on get ddev/ddev-ibexa-cloud`)
3. Configure your IBEXA_CLI_TOKEN globally, `ddev config global --web-environment-add=IBEXA_CLI_TOKEN=nf4amudfn23biyourtoken`
4. Configure your IBEXA_PROJECT, IBEXA_ENVIRONMENT, and IBEXA_APP environment variables, for example `ddev config --web-environment-add=IBEXA_PROJECT=nf4amudfn23biyourproject,IBEXA_ENVIRONMENT=main,IBEXA_APP=app`
5. `ddev restart`
6. `ddev pull ibexa-cloud`

## Running Automated Tests Locally

* `IBEXA_CLI_TOKEN`, `IBEXA_PROJECT` and `IBEXA_ENVIRONMENT` should exist in the environment
* `brew tap kaos/shell && brew install bats-assert bats-file`


**Contributed and maintained by [@rfay](https://github.com/rfay)**

