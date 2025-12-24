---
title: GetDKAN/ddev-dkan
github_url: https://github.com/GetDKAN/ddev-dkan
description: "DDEV add-on for DKAN."
user: GetDKAN
repo: ddev-dkan
repo_id: 514395308
default_branch: main
tag_name: 1.2.7
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2022-07-15
updated_at: 2025-09-19
workflow_status: disabled
stars: 2
---

[![tests](https://github.com/GetDKAN/ddev-dkan/actions/workflows/tests.yml/badge.svg)](https://github.com/GetDKAN/ddev-dkan/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

## What is ddev-dkan?

DDEV-DKAN add-on provides commands and config for development of
DKAN itself and DKAN-enabled Drupal sites in a local docker environment.

## Documentation

Documentation lives in Github Pages: https://getdkan.github.io/ddev-dkan/

## Testing

This addon uses [BATS](https://bats-core.readthedocs.io/en/stable/) for testing.

Get the BATS system by running:
```shell
sh ./setup-bats.sh
```

Now run the tests:
```shell
./tests/bats/bin/bats tests/
```

These tests will run on every push to the Github repo.
