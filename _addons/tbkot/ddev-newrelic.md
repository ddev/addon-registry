---
title: tbkot/ddev-newrelic
github_url: https://github.com/tbkot/ddev-newrelic
description: "New Relic integration for DDEV"
user: tbkot
repo: ddev-newrelic
repo_id: 1044927356
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-08-26
updated_at: 2025-09-09
workflow_status: disabled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tbkot/ddev-newrelic/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tbkot/ddev-newrelic/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tbkot/ddev-newrelic)](https://github.com/tbkot/ddev-newrelic/commits)
[![release](https://img.shields.io/github/v/release/tbkot/ddev-newrelic)](https://github.com/tbkot/ddev-newrelic/releases/latest)

# DDEV New Relic

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)

## Overview

This add-on integrates New Relic into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get tbkot/ddev-newrelic
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

- Create a [New Relic](https://newrelic.com) account. A free account is limited in data ingestion;
  see [New Relic pricing](https://newrelic.com/pricing).
- From the New Relic's control panel, get the Licence key from your Account -> API Keys.<br>
  Use the "Create a key" button and select `Ingest - Licence` as Key type.
- Configure DDEV with the credentials,
  `ddev config global --web-environment-add="NEWRELIC_LICENSE_KEY=<licence_key>, NEWRELIC_APPNAME=<appname>"`.
  The `NEWRELIC_APPNAME` could be any alphanumeric string; New Relic will create a corresponding entity that can be
  found in the UI.<br>
  It’s easiest to do this globally, but you can do the same thing at the project level using
  `ddev config --web-environment-add`. Or create any `.env` file in `.ddev` directory with the corresponding variables.
- `ddev start`
- `ddev newrelic on` to enable, `ddev newrelic off` to disable, `ddev newrelic status` to see status

## Credits

**Contributed and maintained by [@tbkot](https://github.com/tbkot)**
