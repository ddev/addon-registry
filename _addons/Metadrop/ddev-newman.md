---
title: Metadrop/ddev-newman
github_url: https://github.com/Metadrop/ddev-newman
description: "Allows running newman tests on ddev setups"
user: Metadrop
repo: ddev-newman
repo_id: 806591742
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-05-27
updated_at: 2025-09-17
workflow_status: failure
stars: 4
---

[![tests](https://github.com/Metadrop/ddev-newman/actions/workflows/tests.yml/badge.svg)](https://github.com/Metadrop/ddev-newman/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)
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

Example using provided exmaple files:

```
ddev newman run postman/collections/example_cache_headers.postman_collection.json -e postman/envs/example_ddev.postman_environment.json

```

To view all the possible command line options, please [check the documentation](https://www.npmjs.com/package/newman#command-line-options).
