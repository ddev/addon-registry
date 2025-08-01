---
title: stasadev/ddev-python2
github_url: https://github.com/stasadev/ddev-python2
description: "Python 2 for older npm builds with node-gyp inside DDEV"
user: stasadev
repo: ddev-python2
repo_id: 849742328
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-08-30
updated_at: 2025-05-25
workflow_status: success
stars: 8
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/stasadev/ddev-python2/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/stasadev/ddev-python2/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/stasadev/ddev-python2)](https://github.com/stasadev/ddev-python2/commits)
[![release](https://img.shields.io/github/v/release/stasadev/ddev-python2)](https://github.com/stasadev/ddev-python2/releases/latest)

# DDEV Python 2

## Overview

[Python 2](https://www.python.org/doc/sunset-python-2/) has reached its end of life.

This add-on integrates legacy Python 2 inside the `ddev-webserver` into your [DDEV](https://ddev.com/) project.

It is only needed if your `npm` setup requires Python 2 to build dependencies.

## Installation

For DDEV v1.23.5 or above run:

```bash
ddev add-on get stasadev/ddev-python2
ddev restart
```

For earlier versions of DDEV run:

```bash
ddev get stasadev/ddev-python2
ddev restart
```

## Usage

After installation, you can run Python 2:

- `python2.7` (installed at `/usr/bin/python2.7`)
- `python` (symlink to `python2.7` installed at `/usr/local/bin/python`)

This add-on also adds packages that are normally required for `npm` build, see [config.python2.yaml](https://github.com/stasadev/ddev-python2/blob/main/./config.python2.yaml). Remove or replace the contents of this file if you only need Python 2.

## Credits

**Contributed and maintained by [@stasadev](https://github.com/stasadev)**
