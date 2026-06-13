---
title: vanWittlaer/ddev-shopware-cli
github_url: https://github.com/vanWittlaer/ddev-shopware-cli
description: "Adds shopware-cli and watcher configurations"
user: vanWittlaer
repo: ddev-shopware-cli
repo_id: 1267265215
default_branch: main
tag_name: 0.0.1
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-06-12
updated_at: 2026-06-12
workflow_status: unknown
stars: 0
---

# ddev-shopware-cli

Installs [`shopware-cli`](https://sw-cli.fos.gg/) into your [DDEV](https://ddev.com) web
container and wires up the Shopware **storefront** and **admin** hot-reload watchers — so
`shopware-cli`, `ddev admin-watch`, and `ddev storefront-watch` all run in the same
environment as `ddev exec`, with no host-side install.

## Install

```bash
ddev add-on get vanWittlaer/ddev-shopware-cli
ddev restart
```

The `ddev restart` is required: this add-on extends the web image (and exposes the
watcher ports), so the container has to rebuild before the `shopware-cli` binary exists.

## Use

```bash
ddev shopware-cli --version                          # any shopware-cli command
ddev admin-watch custom/static-plugins/MyPlugin      # administration watcher (Vite)
ddev storefront-watch                                # storefront watcher
```

Always open the watchers via your **primary DDEV URL** (`https://<project>.ddev.site`),
never `localhost`. The admin-watch URL omits the `/admin` slug.

## What it installs

| File | Purpose |
|------|---------|
| `.ddev/web-build/Dockerfile.shopware-cli` | Copies the `shopware-cli` binary from the official `shopware/shopware-cli:bin` image into the web image. |
| `.ddev/commands/web/shopware-cli` | The `ddev shopware-cli` wrapper. |
| `.ddev/commands/web/admin-watch` | `ddev admin-watch <plugin-dir>` — runs the Vite admin watcher for a plugin. |
| `.ddev/commands/web/storefront-watch` | `ddev storefront-watch` — runs the storefront watcher. |
| `.ddev/config.watcher.yaml` | Watcher env (`PROXY_URL`, `HOST`, …) and the exposed ports. |

These are copied into your project's `.ddev/` directory. Commit them so teammates get the
same commands and watcher setup.

## Ports

Defaults target **Shopware 6.7.4.2+** (Vite admin watcher):

| Port | Purpose |
|------|---------|
| 5173 | Vite admin dev server |
| 9998 | Storefront proxy / hot-reload |
| 9999 | Storefront assets |

## Older Shopware versions

`config.watcher.yaml` ships the 6.7.4.2+ (Vite) configuration. For earlier Shopware,
replace the `web_environment` / `web_extra_exposed_ports` blocks and `ddev restart`:

**Shopware 6.7.0 – 6.7.4.1** (webpack admin watcher on port 9997):

```yaml
web_environment:
    - HOST=0.0.0.0
    - ADMIN_PORT=9997
    - PROXY_URL=${DDEV_PRIMARY_URL}:9998
    - STOREFRONT_SKIP_SSL_CERT=true
web_extra_exposed_ports:
    - name: admin-proxy
      container_port: 9997
    - name: storefront-proxy
      container_port: 9998
    - name: storefront-assets
      container_port: 9999
```

**Shopware 6.5 & 6.6** (note the different variable name `PORT`):

```yaml
web_environment:
    - HOST=0.0.0.0
    - PORT=9997
    - DISABLE_ADMIN_COMPILATION_TYPECHECK=1
    - PROXY_URL=${DDEV_PRIMARY_URL}:9998
    - STOREFRONT_SKIP_SSL_CERT=true
web_extra_exposed_ports:
    - name: admin-proxy
      container_port: 9997
    - name: storefront-proxy
      container_port: 9998
    - name: storefront-assets
      container_port: 9999
```

On these older versions the `admin-watch` command's Vite flags (`--listen :5173`,
`--external-url …:5173`) do not apply — use the webpack admin watcher invocation for
your Shopware version instead.

## Gotchas

- **Use the primary DDEV URL**, not `localhost`.
- **Pre-6.7.4.2** storefront/admin watchers can trigger browser mixed-content warnings;
  allow insecure content for the site.
- Shopware **6.6.8.0–6.6.8.2** are incompatible with the storefront watcher.

## Remove

```bash
ddev add-on remove shopware-cli
ddev restart
```

This deletes the generated files (they carry a `#ddev-generated` marker).

## Credits

Based on the [Using shopware-cli with DDEV](https://notebook.vanwittlaer.de/ddev-for-shopware/using-shopware-cli-with-ddev)
and [Storefront and admin watchers with DDEV](https://notebook.vanwittlaer.de/ddev-for-shopware/storefront-and-admin-watchers-with-ddev) guides.

**Maintained by [@vanWittlaer](https://github.com/vanWittlaer)**
