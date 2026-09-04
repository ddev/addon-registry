---
title: "bserem/ddev-reg-suit"
github_url: "https://github.com/bserem/ddev-reg-suit"
description: "WIP: ddev addon for reg-suit (docker version of reg-suit is also WIP)"
user: "bserem"
repo: "ddev-reg-suit"
repo_id: 1355944501
default_branch: "main"
tag_name: ""
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-03"
updated_at: "2026-09-03"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)

## DDEV reg-suit

## Overview

[reg-suit](https://github.com/reg-viz/reg-suit) is a visual regression testing
tool. It compares two sets of screenshots, produces an HTML report of the
differences, and — through its plugins — stores the expected images in S3 or GCS
and comments the result on a pull request.

This add-on runs the `reg-suit` CLI in its own container inside a
[DDEV](https://ddev.com) project, so the project itself needs no Node.js
toolchain. No reg-suit config is included; see below for how to generate one.

reg-suit only *compares* images — it does not take screenshots. Pair it with
whatever produces the screenshots for your project (Playwright, Cypress,
Storybook, a `ddev exec` script …).

## Installation

```bash
ddev add-on get bserem/ddev-reg-suit
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

The add-on builds on top of a published `reg-suit` image
(`ghcr.io/bserem/reg-suit` by default). To use your own build, change
`BASE_IMAGE` in `.ddev/docker-compose.reg-suit.yaml` and run `ddev restart`.

## Usage

### Configuration

The reg-suit working directory is `$DDEV_APPROOT/tests/reg-suit`, created by the
installer. Generate a `regconfig.json` there interactively:

```bash
ddev reg-suit init
```

The whole project is mounted into the container, so `reg-keygen-git-hash-plugin`
can read the git history it needs to find the base commit.

A minimal config, comparing a directory of new screenshots against expected ones
stored in S3:

```json
{
  "core": {
    "workingDir": ".reg",
    "actualDir": "actual",
    "thresholdRate": 0,
    "ximgdiff": { "invocationType": "client" }
  },
  "plugins": {
    "reg-keygen-git-hash-plugin": true,
    "reg-publish-s3-plugin": { "bucketName": "your-bucket" }
  }
}
```

### Credentials

Plugin credentials (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`,
`GITHUB_TOKEN`, …) go in `.ddev/.env.reg-suit`, which is read automatically and
should **not** be committed:

```env
AWS_ACCESS_KEY_ID=…
AWS_SECRET_ACCESS_KEY=…
AWS_REGION=eu-west-1
```

### Run

```bash
ddev reg-suit run
```

Any reg-suit subcommand works — `ddev reg-suit compare`, `ddev reg-suit sync-expected`,
`ddev reg-suit prepare -p slack`, `ddev reg-suit run -v` for verbose output.

To run in a different directory than `tests/reg-suit`, set `REGSUIT_WORKDIR`:

```bash
REGSUIT_WORKDIR=/var/www/html ddev reg-suit run
```

### View the report

```bash
ddev reg-suit-report
```

This opens `tests/reg-suit/.reg/index.html` in your default browser. Adjust the
path in the command if you changed `workingDir` in `regconfig.json`.

## Credits

**Contributed and maintained by [@bserem](https://github.com/bserem)**, modelled
on [ddev-backstopjs](https://github.com/ddev/ddev-backstopjs).
