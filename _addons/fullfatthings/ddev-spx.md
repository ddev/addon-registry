---
title: fullfatthings/ddev-spx
github_url: https://github.com/fullfatthings/ddev-spx
description: "DDEV Addon to install the PHP-SPX performance package "
user: fullfatthings
repo: ddev-spx
repo_id: 868948247
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-10-07
updated_at: 2025-02-26
stars: 8
---

# DDEV SPX

Provides the [SPX](https://github.com/NoiseByNorthwest/php-spx) extension for PHP.

- Run `ddev get fullfatthings/ddev-spx`
- Run `ddev restart` - this should install the addons into php
- To enable, run `ddev spx on` and to disable run `ddev spx off`.
- Once enabled, browse to `https://example.ddev.site/?SPX_KEY=dev&SPX_UI_URI=/`
