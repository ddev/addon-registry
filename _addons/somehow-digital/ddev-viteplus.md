---
title: somehow-digital/ddev-viteplus
github_url: https://github.com/somehow-digital/ddev-viteplus
description: "DDEV addon that provides a Vite+ service and commands to run inside your DDEV project."
user: somehow-digital
repo: ddev-viteplus
repo_id: 1181005932
default_branch: main
tag_name: v1.0.3
ddev_version_constraint: ">= v1.25.0"
dependencies: []
type: contrib
created_at: 2026-03-13
updated_at: 2026-03-15
workflow_status: success
stars: 2
---

# Vite+ for `DDEV`

DDEV addon that provides [`Vite+`](https://viteplus.dev) commands to run inside your DDEV project.

## What is Vite+?

[Vite+](https://viteplus.dev) is the unified toolchain for the web, providing a complete development environment with Vite, package management, testing, and more.

## Install

In your DDEV project:

```bash
ddev add-on get somehow-digital/ddev-viteplus
ddev restart
```

## Commands

- **`ddev vp`**

## URLs

The dev server container exposes Vite on:

- **HTTPS**: `https://<project>.ddev.site:5173/`

**Internal Container Access**: From other DDEV containers (e.g., PHP/web container), use:
- `http://viteplus:5173/`

## Configuration

### Craft CMS

When using the [Vite Plugin](https://plugins.craftcms.com/vite), configure as follows:

**config/vite.php**
```php
<?php

use craft\helpers\App;

return [
    'useDevServer' => Craft::$app->config->general->devMode,
    'checkDevServer' => true,
    'devServerInternal' => 'http://viteplus:5173',
    'devServerPublic' => App::env('DEFAULT_SITE_URL') . ':5173',
    'serverPublic' => '/dist/',
    'includeReactRefreshShim' => false,
    'includeModulePreloadShim' => false,
    'manifestPath' => '@webroot/dist/.vite/manifest.json',
];
```

**vite.config.js**
```js
export default {
  server: {
    host: '0.0.0.0',
    port: 5173,
    allowedHosts: [
      'viteplus',
    ],
  },
}
```
