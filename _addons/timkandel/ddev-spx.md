---
title: timkandel/ddev-spx
github_url: https://github.com/timkandel/ddev-spx
description: "DDEV Addon to install the PHP-SPX performance package "
user: timkandel
repo: ddev-spx
repo_id: 1195168919
default_branch: main
tag_name: v0.1.1
ddev_version_constraint: ">v1.21.4"
dependencies: []
type: contrib
created_at: 2026-03-29
updated_at: 2026-04-02
workflow_status: unknown
stars: 0
---

# DDEV SPX

Provides the [SPX](https://github.com/NoiseByNorthwest/php-spx) extension for PHP.

- Run `ddev add-on get timkandel/ddev-spx`
- Run `ddev restart` - this should install the addons into php
- To enable, run `ddev spx on` and to disable run `ddev spx off`.
- Once enabled, browse to `https://example.ddev.site/?SPX_KEY=dev&SPX_UI_URI=/`
