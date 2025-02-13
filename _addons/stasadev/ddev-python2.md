---
title: stasadev/ddev-python2
github_url: https://github.com/stasadev/ddev-python2
description: "Add Python2 for older npm builds with node-gyp"
user: stasadev
repo: ddev-python2
type: contrib
created_at: 2024-08-30
updated_at: 2025-02-08
stars: 7
---

[![tests](https://github.com/stasadev/ddev-python2/actions/workflows/tests.yml/badge.svg)](https://github.com/stasadev/ddev-python2/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

# DDEV Python2 Add-on <!-- omit in toc -->

* [What is ddev-python2?](#what-is-ddev-python2)
* [Installation](#installation)
* [Usage](#usage)

## What is ddev-python2?

This repository allows you to install [legacy Python2](https://www.python.org/doc/sunset-python-2/) inside `ddev-webserver` in a DDEV project.

This is only required if your `npm` uses Python2 to build its dependencies.

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get stasadev/ddev-python2
```

For earlier versions of DDEV run

```sh
ddev get stasadev/ddev-python2
```

Then restart your project

```sh
ddev restart
```

## Usage

After installation, you can run Python2:

- `python2.7` (installed at `/usr/bin/python2.7`)
- `python` (symlink to `python2.7` installed at `/usr/local/bin/python`)

This add-on also adds packages that are normally required for `npm` build, see [config.python2.yaml](https://github.com/stasadev/ddev-python2/blob/main/./config.python2.yaml). Remove or replace the contents of this file if you only need Python2.

**Contributed and maintained by [@stasadev](https://github.com/stasadev)**

