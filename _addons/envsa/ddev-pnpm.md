---
title: envsa/ddev-pnpm
github_url: https://github.com/envsa/ddev-pnpm
description: "Add pnpm to DDEV"
user: envsa
repo: ddev-pnpm
repo_id: 716421197
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-11-09
updated_at: 2025-03-05
workflow_status: disabled
stars: 6
---

<div align="center">
    <a href="https://ddev.com/">
        <img src="https://raw.githubusercontent.com/ddev/ddev/master/images/ddev-logo.svg" alt="DDEV logo" height="80">
    </a>
    <a href="https://pnpm.io">
        <img src="https://avatars.githubusercontent.com/u/21320719?s=200&v=4" alt="pnpm Logo" height="80">
    </a>
    <h1 align="center">ddev-pnpm</h1>
</div>

[![tests](https://github.com/envsa/ddev-pnpm/actions/workflows/tests.yml/badge.svg)](https://github.com/envsa/ddev-pnpm/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## What is pnpm?

> Fast, disk space efficient package manager

For more information, visit [pnpm.io](https://pnpm.io)

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get envsa/ddev-pnpm
```

For earlier versions of DDEV run

```sh
ddev get envsa/ddev-pnpm
```

Then restart your project

```sh
ddev restart
```

## Usage

```sh
ddev pnpm
```

Please refer to the documentation at [pnpm.io](https://pnpm.io)
