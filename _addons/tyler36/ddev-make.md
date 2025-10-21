---
title: tyler36/ddev-make
github_url: https://github.com/tyler36/ddev-make
description: "A DDEV helper to run `Makefile`s in the web container."
user: tyler36
repo: ddev-make
repo_id: 962418405
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-04-08
updated_at: 2025-06-09
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-make/actions/workflows/tests.yml/badge.svg)](https://github.com/tyler36/ddev-make/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-make)](https://github.com/tyler36/ddev-make/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-make)](https://github.com/tyler36/ddev-make/releases/latest)

# ddev-make <!-- omit in toc -->

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Credits](#credits)

## Overview

[GNU make](https://www.gnu.org/software/make/) is a tool which controls the generation of executables and other non-source files of a program from the program's source files.

This add-on integrates build-essentials, which contain `make` into your [DDEV](https://ddev.com/) project.

See [Learn Makefile](https://makefiletutorial.com/).

## Installation

```shell
ddev add-on get tyler36/ddev-make
```

## Usage

This add-on provides a helper command:

```shell
ddev make
```

All flags and arguments are passed through to the command.

`ddev make` runs relative to the host working directory.
For example:

Given the following folder structure:

```
project
├── .ddev
└── tests
    └── Makefile
```

`cd tests && ddev make` will run correctly run the `tests/Makefile`.

## Credits

**Contributed and maintained by [`@tyler36`](https://github.com/tyler36)**
