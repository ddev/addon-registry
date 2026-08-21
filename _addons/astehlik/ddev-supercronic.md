---
title: "astehlik/ddev-supercronic"
github_url: "https://github.com/astehlik/ddev-supercronic"
description: "DDEV addon to run scheduled tasks with supercronic, a cron-compatible job runner for containers"
user: "astehlik"
repo: "ddev-supercronic"
repo_id: 1275315473
default_branch: "main"
tag_name: "v0.2.46"
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: "contrib"
created_at: "2026-06-20"
updated_at: "2026-06-21"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/astehlik/ddev-supercronic/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/astehlik/ddev-supercronic/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/astehlik/ddev-supercronic)](https://github.com/astehlik/ddev-supercronic/commits)
[![release](https://img.shields.io/github/v/release/astehlik/ddev-supercronic)](https://github.com/astehlik/ddev-supercronic/releases/latest)

# DDEV Supercronic

A [DDEV](https://ddev.com) add-on that runs scheduled cron jobs in the web
container using [Supercronic](https://github.com/aptible/supercronic).

Unlike the system `cron` daemon, Supercronic inherits the full container
environment — DDEV variables like `IS_DDEV_PROJECT`, database credentials, and
`TYPO3_CONTEXT` are automatically available to every job without any inline
exports.

## Installation

```bash
ddev add-on get astehlik/ddev-supercronic
ddev restart
```

## Usage

Add one or more `*.cron` files to `.ddev/web-build/`. Each file uses standard
crontab syntax (no username column).

On `ddev start`, all `.ddev/web-build/*.cron` files are concatenated into a
single crontab and executed by Supercronic inside the web container.

An example job is provided at `.ddev/web-build/time.cron.example`. Rename it to
activate it:

```bash
mv .ddev/web-build/time.cron.example .ddev/web-build/time.cron
ddev restart
```

After at least a minute, `./time.log` should appear in your project root
containing the current time written by the example job.

## Examples

Because Supercronic inherits the container environment, `IS_DDEV_PROJECT` and
similar variables do not need to be set explicitly in cron files.

### TYPO3 scheduler

```cron
* * * * * cd /var/www/html && vendor/bin/typo3 scheduler:run
```

Feel free to contribute own working examples for other systems.

## Debugging

- Use [crontab guru](https://crontab.guru/) to build schedule expressions.
- Inspect the active crontab with `ddev exec cat /etc/supercronic/crontab`.
- Connect to the web container with `ddev ssh` to run commands manually.
- To run jobs on local time instead of UTC, set timezone in `.ddev/config.yaml`.

## Migrating from ddev-cron

This add-on is a drop-in replacement for
[ddev/ddev-cron](https://github.com/ddev/ddev-cron). Both use the same
`.ddev/web-build/*.cron` file format and location, so existing cron files work
without any changes.

```bash
ddev add-on remove ddev/ddev-cron
ddev add-on get astehlik/ddev-supercronic
ddev restart
```

Because Supercronic inherits the full container environment, any
`IS_DDEV_PROJECT=true` or other explicit env var prefixes in your cron commands
are no longer necessary, though leaving them in place is harmless.

## Removal

```bash
ddev add-on remove astehlik/ddev-supercronic
```

## Versioning

The add-on version tracks the Supercronic version it ships. A GitHub Actions
workflow checks for new Supercronic releases every Monday and opens a pull
request with the update.
