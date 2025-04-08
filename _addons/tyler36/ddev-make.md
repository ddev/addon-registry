---
title: tyler36/ddev-make
github_url: https://github.com/tyler36/ddev-make
description: "A DDEV helper to run `Makefile`s in the web container."
user: tyler36
repo: ddev-make
repo_id: 962418405
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2025-04-08
updated_at: 2025-04-08
stars: 0
---

[![tests](https://github.com/tyler36/ddev-make/actions/workflows/tests.yml/badge.svg)](https://github.com/tyler36/ddev-make/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

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
