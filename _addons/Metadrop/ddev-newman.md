---
title: Metadrop/ddev-newman
github_url: https://github.com/Metadrop/ddev-newman
description: "Allows running newman tests on ddev setups"
user: Metadrop
repo: ddev-newman
type: contrib
created_at: 2024-05-27
updated_at: 2024-10-01
stars: 0
---

[![tests](https://github.com/Metadrop/ddev-newman/actions/workflows/tests.yml/badge.svg)](https://github.com/Metadrop/ddev-newman/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)
![GitHub Release](https://img.shields.io/github/v/release/Metadrop/ddev-newman)


# DDEV newman

Allows running [newman](https://www.npmjs.com/package/newman) on ddev setups. Use it to run postman tests through CLI.

## Installation

Install this addon by running:

```
ddev get metadrop/ddev-newman
```

## Usage

```
ddev newman [args]
```

Example:

```
ddev newman my_collection.json -e environment.json
```

To view all the possible command line options, please [check the documentation](https://www.npmjs.com/package/newman#command-line-options).

