---
title: zone1987/ddev-shopware
github_url: https://github.com/zone1987/ddev-shopware
description: ""
user: zone1987
repo: ddev-shopware
repo_id: 1339655611
default_branch: 6.7
tag_name: v6.7.0
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-08-19
updated_at: 2026-08-19
workflow_status: unknown
stars: 0
---

[![tests](https://github.com/zone1987/ddev-shopware/actions/workflows/tests.yml/badge.svg?branch=6.7)](https://github.com/zone1987/ddev-shopware/actions/workflows/tests.yml?query=branch%3A6.7)
[![last commit](https://img.shields.io/github/last-commit/zone1987/ddev-shopware)](https://github.com/zone1987/ddev-shopware/commits)
[![release](https://img.shields.io/github/v/release/zone1987/ddev-shopware)](https://github.com/zone1987/ddev-shopware/releases/latest)

# DDEV Shopware <!-- omit in toc -->

* [What is DDEV Shopware?](#what-is-ddev-shopware)
* [Installation](#installation)
  * [Installing a specific Shopware line](#installing-a-specific-shopware-line)
* [Usage](#usage)
  * [Build everything](#build-everything)
  * [Only your own extensions](#only-your-own-extensions)
  * [Watchers](#watchers)
* [Configuration](#configuration)
  * [Shopware project root](#shopware-project-root)
  * [Ports](#ports)
* [How it works](#how-it-works)
* [Removal](#removal)
* [Resources](#resources)
* [Credits](#credits)

## What is DDEV Shopware?

A [DDEV](https://ddev.com/) add-on that turns the usual
[shopware-cli](https://developer.shopware.com/docs/products/cli/) incantations into short,
memorable commands — and installs the tooling they need.

After `ddev add-on get` + `ddev restart` you can run `ddev build` or `ddev watch` in any Shopware
project, with no per-project setup.

Key features:

* **Short commands** — `ddev build`, `ddev watch`, plus the individual Administration and Storefront
  variants.
* **Batteries included** — `shopware-cli` and [Deployer](https://deployer.org/) are built into the
  web image; the watcher ports are published through `ddev-router`.
* **Scoped builds** — build everything, only your own static extensions, or a named list of plugins.
* **Parallel watchers** — `ddev watch` runs both watchers at once and prints their URLs, while the
  Storefront watcher keeps its interactive TTY for theme and sales-channel selection.
* **Layout-agnostic** — finds your shop whether it lives in the project root or a subdirectory, with
  or without `working_dir`/`composer_root`, from whatever directory you run the command in.

## Installation

```bash
ddev add-on get zone1987/ddev-shopware
ddev restart
```

The `ddev restart` is required: it builds `shopware-cli` into the web image and opens the watcher
ports. Commit the `.ddev` directory to version control afterwards.

### Installing a specific Shopware line

Shopware's build tooling changes between minor versions, so this add-on is maintained in one branch
per Shopware line — there is no `main`. The command above installs the latest release, which tracks
the current line.

Pin an exact release so a later change cannot surprise you:

```bash
ddev add-on get zone1987/ddev-shopware --version v6.7.0
```

Or follow a line's branch to always get its newest state:

```bash
ddev add-on get zone1987/ddev-shopware --version 6.7
```

Releases are versioned `v<shopware-line>.<patch>` — `v6.7.0`, `v6.7.1`, … all target Shopware 6.7.

## Usage

Every command takes an optional scope. Without one, it processes the whole project.

| Command | Description |
| ------- | ----------- |
| `ddev build` | Build Administration and Storefront |
| `ddev watch` | Start both watchers in parallel |
| `ddev admin-build` | Build the Administration |
| `ddev admin-watch` | Start the Administration watcher |
| `ddev storefront-build` | Build the Storefront |
| `ddev storefront-watch` | Start the Storefront watcher |

### Build everything

```bash
ddev build
```

### Only your own extensions

`static` restricts the build to your custom static extensions — much faster during development:

```bash
ddev build static
```

`plugin` narrows it further to a named list:

```bash
ddev build plugin SwagPluginA
ddev build plugin SwagPluginA SwagPluginB
```

Both scopes work with every command, e.g. `ddev admin-watch plugin SwagPluginA`.

### Watchers

```bash
ddev watch
```

Once both dev servers are listening, their URLs are printed:

| Watcher | URL |
| ------- | --- |
| Administration | `https://<project>.ddev.site:5173` |
| Storefront | `https://<project>.ddev.site:9998` |

The Storefront watcher runs in the foreground and keeps its interactive TTY, so its theme and
sales-channel/domain prompts work normally. Press `Ctrl+C` once to stop both watchers.

## Configuration

### Shopware project root

The commands run `shopware-cli` from the Shopware root inside the web container. You do **not**
need `working_dir` or `composer_root` set, and it does not matter whether you run the command from
the project root or after `cd`-ing into the shop yourself.

The root is resolved in this order:

1. `SHOPWARE_PROJECT_ROOT` — explicit override
2. your current directory, or any parent of it
3. `DDEV_COMPOSER_ROOT` — set when your `config.yaml` uses `composer_root`
4. the directory above `DDEV_DOCROOT` — e.g. `shopware/public` → `/var/www/html/shopware`
5. `/var/www/html` — Shopware sits in the project root
6. a one-level scan of `/var/www/html/*`

Every candidate except the explicit override has to prove itself: it only counts if it holds a
`composer.json` requiring a `shopware/*` package. So a wrong guess is never used silently — if
nothing matches, the command fails with a clear message instead of building the wrong tree.

For layouts none of that reaches, set the path yourself:

```bash
ddev dotenv set .ddev/.env.shopware --shopware-project-root=/var/www/html/shopware
ddev restart
```

Commit the `.ddev/.env.shopware` file to version control.

### Ports

The watcher ports and environment live in `.ddev/config.shopware.yaml`. Don't edit that file — it is
overwritten on the next `ddev add-on get`. To change something, override it in your own
`.ddev/config.shopware-local.yaml`, which DDEV merges on top:

```yaml
web_extra_exposed_ports:
  - name: shopware-admin-proxy
    container_port: 5173
    http_port: 6172
    https_port: 6173
```

| Purpose | Container | HTTPS | HTTP |
| ------- | --------- | ----- | ---- |
| Administration (Vite) | 5173 | 5173 | 5172 |
| Storefront proxy | 9998 | 9998 | 8888 |
| Storefront assets | 9999 | 9999 | 8889 |

## How it works

* **`web-build/Dockerfile.shopware`** — installs `shopware-cli` from the FriendsOfShopware package
  repository and a pinned Deployer release into DDEV's web image.
* **`config.shopware.yaml`** — publishes the watcher ports through `ddev-router` and sets the
  environment the watchers expect (`HOST`, `PORT`, `PROXY_URL`, `STOREFRONT_SKIP_SSL_CERT`, …).
* **`commands/web/.shopware-cli-command`** — the shared implementation. It locates the Shopware root
  (verifying each candidate against its `composer.json`), checks that `shopware-cli` is present, and
  translates the scope argument into the matching `shopware-cli project` flags
  (`--only-custom-static-extensions` / `--only-extensions`). The leading dot keeps it out of
  `ddev help`.
* **`commands/web/*`** — thin wrappers around it. `build` chains both builds; `watch` runs the
  Administration watcher in the background, the Storefront watcher in the foreground so it keeps its
  TTY, and prints the URLs once both ports accept connections.

## Removal

```bash
ddev add-on remove shopware
```

Use the add-on **name** (`shopware`), not the repository slug.

## Resources

* [shopware-cli documentation](https://developer.shopware.com/docs/products/cli/)
* [Shopware developer documentation](https://developer.shopware.com/)
* [DDEV documentation for add-ons](https://docs.ddev.com/en/stable/users/extend/additional-services/)

## Credits

**Maintained by [@zone1987](https://github.com/zone1987)**
