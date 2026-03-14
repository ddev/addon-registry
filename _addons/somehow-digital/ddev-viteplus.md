---
title: somehow-digital/ddev-viteplus
github_url: https://github.com/somehow-digital/ddev-viteplus
description: "DDEV addon that provides a Vite+ service and commands to run inside your DDEV project."
user: somehow-digital
repo: ddev-viteplus
repo_id: 1181005932
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.25.0"
dependencies: []
type: contrib
created_at: 2026-03-13
updated_at: 2026-03-13
workflow_status: unknown
stars: 1
---

# Vite+ for DDEV

DDEV addon that provides a Vite+ service and commands to run inside your DDEV project.

## What is Vite+?

[Vite+](https://viteplus.dev) is the unified toolchain for the web, providing a complete development environment with Vite, package management, testing, and more.

## Install

In your DDEV project:

```bash
ddev add-on get somehow-digital/ddev-viteplus
ddev restart
```

## Commands

- **`ddev vp <args...>`**

Examples:

```bash
ddev vp install
ddev vp dev
```

## URLs

The dev server container exposes Vite on:

- **HTTPS**: `https://<project>.ddev.site:3000`

**Configuration**: Add this to your `vite.config.js`:

```js
// vite.config.js
export default {
  server: {
    host: '0.0.0.0',
    port: 3000,
    allowedHosts: [
      'viteplus'
    ]
  }
}
```

**Internal Container Access**: From other DDEV containers (e.g., PHP/web container), use:
- `http://viteplus:3000`
