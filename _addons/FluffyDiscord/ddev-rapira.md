---
title: "FluffyDiscord/ddev-rapira"
github_url: "https://github.com/FluffyDiscord/ddev-rapira"
description: "Serve your PHP application with Rapira in DDEV, with DDEV's nginx in front."
user: "FluffyDiscord"
repo: "ddev-rapira"
repo_id: 1359003196
default_branch: "master"
tag_name: "v0.1.1"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-06"
updated_at: "2026-09-06"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/FluffyDiscord/ddev-rapira/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/FluffyDiscord/ddev-rapira/actions/workflows/tests.yml?query=branch%3Amaster)
[![last commit](https://img.shields.io/github/last-commit/FluffyDiscord/ddev-rapira)](https://github.com/FluffyDiscord/ddev-rapira/commits)
[![release](https://img.shields.io/github/v/release/FluffyDiscord/ddev-rapira)](https://github.com/FluffyDiscord/ddev-rapira/releases/latest)

# DDEV Rapira

> [!NOTE]
> This add-on was created with the assistance of AI. It has been verified end-to-end against real DDEV + Docker, but review it before relying on it in your environment.

Serve your PHP application with [Rapira](https://github.com/rapira-rs/rapira) in [DDEV](https://ddev.com/), with DDEV's nginx still in front.

Rapira is a PHP application server written in Rust. In **classic** mode it runs your entry script per request, so it serves any framework unchanged; in **dispatcher** mode it keeps a booted application resident between requests.

It [embeds `libphp.so`](https://rapira.rs/docs/intro/installation) rather than talking to a separate PHP process, so this add-on points it at the image's own PHP. Your project keeps the PHP version, extensions, ini values and `.ddev/php/*.ini` overrides it already had, and `ddev xdebug on`, `ddev xhprof on`, `ddev php`, `ddev composer` and the rest keep working.

## Install

```bash
ddev add-on get FluffyDiscord/ddev-rapira
ddev restart
git add -f .ddev/nginx_full/nginx-site.conf
```

The `-f` matters if you commit `.ddev`: DDEV lists that path in the `.ddev/.gitignore` it generates, and drops it again only on the next start. Commit it in the window between and a fresh clone silently falls back to php-fpm — which still answers 200, so it is a confusing thing to debug.

## Requirements

- DDEV `v1.24.10`+
- PHP **8.4** or **8.5** — Rapira publishes builds for those two only
- Project type **`php`** or **`symfony`** (see [Known limitations](#known-limitations))
- amd64. The arm64 build is wired up but untested

## Trusted proxies

nginx reaches Rapira over plain HTTP on `127.0.0.1`, so the app sees `REMOTE_ADDR=127.0.0.1`, an empty `HTTPS`, and `SERVER_PORT=8000`. The add-on sends `X-Real-IP`, `X-Forwarded-For`, `X-Forwarded-Proto` and `X-Forwarded-Host`; your app has to be told to trust them. For Symfony:

```yaml
# config/packages/framework.yaml
framework:
    trusted_proxies: '127.0.0.1'
    trusted_headers: ['x-forwarded-for', 'x-forwarded-proto', 'x-forwarded-host']
```

Without this, `Request::isSecure()` is false and generated URLs come out as `http://…:8000`.

`X-Forwarded-Port` is deliberately not sent, so leave `x-forwarded-port` out of `trusted_headers`: nginx's `$server_port` is its own listening port, so it would report `80` for a request `ddev-router` terminated as HTTPS, and an app that trusts it builds `https://host:80/` URLs. With the header absent, Symfony derives 443 from the forwarded scheme.

A proxy cannot supply everything `fastcgi_params` did. What changes, measured:

| `$_SERVER` key | Under php-fpm | Under Rapira | What to use instead |
|---|---|---|---|
| `SERVER_NAME` | the request's host | `localhost`, whatever the `Host` header says | `HTTP_HOST`, or a trusted `X-Forwarded-Host` |
| `SERVER_PORT` | 80 | 8000 | Nothing — derive it from the forwarded scheme |
| `HTTPS` | `on` / `off` | empty | Trusted `X-Forwarded-Proto` |
| `REQUEST_SCHEME` | `https` | `http` | Trusted `X-Forwarded-Proto` |
| `REMOTE_ADDR` | the client | `127.0.0.1` | Trusted `X-Real-IP` / `X-Forwarded-For` |
| `SERVER_ADDR`, `DOCUMENT_URI`, `REMOTE_USER`, `REDIRECT_STATUS` | set | absent | Nothing |
| `SERVER_SOFTWARE` | `nginx/<version>` | `Rapira` | — |
| `CONTENT_TYPE` | `''` when absent | absent when absent | `$_SERVER['CONTENT_TYPE'] ?? ''` |

`X-Accel-Redirect` and `X-Accel-Buffering` change hands too: nginx now acts on them itself instead of forwarding them to the client, which is what a resident-worker app wants in production.

## Configuration

| What | Default | How to change |
|---|---|---|
| Rapira config file | `rapira.toml` in the project root, else classic mode on `<docroot>/index.php` | `ddev dotenv set .ddev/.env.web --rapira-config-file=rapira.dev.toml && ddev restart` |
| Docroot (the override's nginx `root`) | your `ddev config --docroot`, set at install | Edit `root` in `.ddev/nginx_full/nginx-site.conf` and `ddev restart` |
| Rapira version | `v0.8.1` | `ARG RAPIRA_VERSION` in `.ddev/web-build/Dockerfile.rapira`, then `ddev restart` |

The add-on does **not** write or manage `rapira.toml` — it is yours, DDEV bind-mounts it, and it stays live. Copy [`example.rapira.toml`](https://github.com/FluffyDiscord/ddev-rapira/blob/master/./example.rapira.toml) to your project root to start from something. One thing that file cannot control: the add-on passes `--listen 127.0.0.1:8000` on the command line, which overrides `[http] listen`, because nginx proxies to that fixed port.

Edits to `.ddev/nginx_full/nginx-site.conf` are replaced on the next `ddev add-on get` (the old file is kept beside it as `nginx-site.conf.ddev-rapira-backup-<epoch>`). Put additive rules in `.ddev/nginx/*.conf`, which the override still includes.

## `ddev rapira-restart`

```bash
ddev rapira-restart
```

In `dispatcher` or `worker` mode the application stays resident, so a code change needs this. In `classic` mode you do not — the entry script re-runs per request. It also starts the daemon if it is stopped, which is the fix for the case below.

## Removal

```bash
ddev add-on remove rapira
git rm --cached .ddev/nginx_full/nginx-site.conf   # if you committed it
ddev restart
```

The `git rm --cached` matters: a force-added file stays tracked, so without it the next branch switch restores an nginx config pointing at a port nothing is listening on.

## Known limitations

- **Project types `php` and `symfony` only.** DDEV renders a different nginx site config per project type; this add-on ships one, derived from the file those two share. Installing on `laravel`, `drupal*`, `magento*`, `typo3`, `shopware6` or `wordpress` would silently replace routing rules those types need, so the install refuses. Changing the project type *after* installing has the same effect without the refusal — `ddev add-on remove rapira` if you do.
- **A container restarted outside `ddev` comes up without Rapira.** DDEV starts extra daemons from the host after the container boots, so a bare `docker restart` leaves nginx and php-fpm running and Rapira stopped; the container still reports healthy because the healthcheck probes php-fpm. Every PHP request 502s until `ddev start` or `ddev rapira-restart`.
- **arm64 is untested.** The build maps to the published `aarch64` release, but nothing has been run on it.

## Resources

- [Rapira documentation](https://rapira.rs/docs/intro/)
- [`fluffydiscord/rapira-symfony-bundle`](https://github.com/FluffyDiscord/rapira-symfony-bundle) — dispatcher-mode Symfony integration
- [DDEV custom nginx / webserver](https://docs.ddev.com/en/stable/users/extend/customization-extendibility/)

## Credits

**Contributed by [@FluffyDiscord](https://github.com/FluffyDiscord)**
