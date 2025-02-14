---
title: ochorocho/ddev-frankenphp
github_url: https://github.com/ochorocho/ddev-frankenphp
description: "Replaces apache/nginx with FrankenPHP in ddev web container"
user: ochorocho
repo: ddev-frankenphp
repo_id: 755946584
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-02-11
updated_at: 2025-01-23
stars: 11
---

# FrankenPHP Add-On for DDEV

This DDEV add-on for FrankenPHP replaces the currently used webserver
with [FrankenPHP](https://frankenphp.dev/).

FrankenPHP is a modern PHP App Server, written in Go.

## Why?

The WHY!

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get ochorocho/ddev-frankenphp && ddev restart
```

For earlier versions of DDEV run

```sh
ddev get ochorocho/ddev-frankenphp && ddev restart
```

## Configuration

[config.yaml](https://github.com/ochorocho/ddev-frankenphp/blob/main/frankenphp%2Fconfig.yaml)

`SERVER_NAME` will be set depending on the VIRTUAL_HOST used by ddev.
This way all domains configured by ddev will be used by frankenphp as well.

`MERCURE_PUBLIC_URL` will use the `DDEV_PRIMARY_URL` (`$DDEV_PRIMARY_URL/.well-known/mercure`)
