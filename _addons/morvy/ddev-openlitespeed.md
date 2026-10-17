---
title: "morvy/ddev-openlitespeed"
github_url: "https://github.com/morvy/ddev-openlitespeed"
description: "OpenLiteSpeed + lsphp web server for DDEV (LSCache, .htaccess, PHP 5.6–8.5)"
user: "morvy"
repo: "ddev-openlitespeed"
repo_id: 1320285445
default_branch: "master"
tag_name: "v1.1.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-08-02"
updated_at: "2026-10-16"
workflow_status: "disabled"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/morvy/ddev-openlitespeed/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/morvy/ddev-openlitespeed/actions/workflows/tests.yml?query=branch%3Amaster)
[![last commit](https://img.shields.io/github/last-commit/morvy/ddev-openlitespeed)](https://github.com/morvy/ddev-openlitespeed/commits)
[![release](https://img.shields.io/github/v/release/morvy/ddev-openlitespeed)](https://github.com/morvy/ddev-openlitespeed/releases/latest)

# ddev-openlitespeed

Replaces DDEV's default web server (nginx/apache-fpm) with **OpenLiteSpeed** running
**lsphp** (the LiteSpeed SAPI PHP), so you can develop and test against a real
OpenLiteSpeed stack locally.

Built on DDEV's `webserver_type: generic` mechanism (the add-on needs DDEV v1.24.10+), the same one
the official [`ddev-frankenphp`](https://github.com/ddev/ddev-frankenphp) add-on uses.

## What you get

- OpenLiteSpeed serving your project over the normal `https://<project>.ddev.site` URL.
- **lsphp** as the PHP runtime (real LiteSpeed SAPI, not php-fpm behind a proxy) — the
  LiteSpeed SAPI of **DDEV's own PHP build**, so it has the same extensions, `php.ini`,
  `.ddev/php/*.ini` overrides, timezone, `ddev xdebug` and `ddev xhprof` as nginx-fpm /
  apache-fpm (see [docs/php-parity.md](https://github.com/morvy/ddev-openlitespeed/blob/master/docs/php-parity.md)).
- **Correct HTTPS for PHP** behind the DDEV router: `$_SERVER['HTTPS']`, `SERVER_PORT`
  and `REQUEST_SCHEME` reflect the real scheme, so `is_ssl()` / `isSecure()`, redirects
  and absolute-URL generation work (no mixed-content or redirect loops).
- **LSCache** advertised to the [LiteSpeed Cache](https://wordpress.org/plugins/litespeed-cache/)
  plugin (`X-LSCACHE`), so the plugin detects the server and can drive page caching.
- PHP version follows DDEV (`ddev config --php-version=… && ddev restart`).
- Project `.htaccess` is honoured for front-controller rewrites (WordPress, Drupal,
  Laravel, …), like apache-fpm.
- DDEV's own 404 and 403 (directory without an index file) pages, like nginx-fpm /
  apache-fpm (DDEV v1.25.4+); your app's own error responses pass through.
- OLS logs stream to `ddev logs -f`.

## Install

```bash
ddev add-on get morvy/ddev-openlitespeed
ddev restart
```

## PHP versions

Every PHP version DDEV supports (5.6–8.5). Set one with
`ddev config --php-version=<version> && ddev restart`.

lsphp is `php<version>-litespeed` from the same repository as DDEV's `php<version>-fpm`,
installed in the web container. Installing it may move the image's `php<version>-*`
packages to that repository's current patch release (all together).

PHP 8.6 needs a DDEV release newer than v1.25.4 (DDEV refuses to start it before
that); the add-on needs no change for it.

## How it works

Inside the `web` container, two supervised daemons run as the project's (non-root,
code-owning) user:

| Piece | Role |
|-------|------|
| `config.openlitespeed.yaml` | sets `webserver_type: generic`; runs the **lsphp** and **openlitespeed** daemons; exposes 80/443 through `ddev-router` |
| `web-build/Dockerfile.openlitespeed` | installs OpenLiteSpeed and `php${DDEV_PHP_VERSION}-litespeed`; links the SAPI's `php.ini` and `conf.d` to DDEV's fpm ones |
| `web-build/openlitespeed/ddev-openlitespeed-lsphp` | lsphp daemon on `127.0.0.1:9000`; restart-safe (DDEV runs daemons under `bash -c`, so it clears a leftover lsphp before binding) |
| `web-build/openlitespeed/httpd_config.conf` | main OLS config: one vhost, lsphp external app (`autoStart 0`), LSCache module, `:80` (router) + `:443` (DDEV master cert) listeners |
| `web-build/openlitespeed/vhost.conf` | vhost template; docroot injected at start; rewrites `X-Forwarded-Proto` into real HTTPS server vars; serves DDEV's 404/403 pages |
| `web-build/openlitespeed/ddev-openlitespeed-start` | foreground OLS wrapper: renders the vhost, pins the run user, streams logs; restart-safe (clears a leftover OLS and wrapper first) |

Request flow: `ddev-router` terminates public TLS and forwards **plain HTTP to OLS on
:80** (with `X-Forwarded-Proto`); OLS restores the HTTPS server vars and proxies PHP
over LSAPI to lsphp on `127.0.0.1:9000`.

> **Why lsphp is a separate process:** this OLS build has no working suEXEC/cgid helper,
> so OLS cannot fork lsphp itself (`cgidSuEXEC failed`). lsphp runs as its own daemon and
> the external app uses `autoStart 0` (connect-only).

## Known limitations

- **No WebAdmin console.** OpenLiteSpeed's admin console runs its own PHP as an internal
  worker that needs the same (broken) suEXEC helper, so it can't run in this container.
  It's disabled; configure OpenLiteSpeed by editing the files under
  `.ddev/web-build/openlitespeed/` and running `ddev restart`.

## Uninstall

```bash
ddev add-on remove openlitespeed
ddev restart
```

Removes the `#ddev-generated` files and reverts `webserver_type` to your project's
previous value. Upgrading from v1.0.x removes its lsphp sidecar files.

## Releasing

Push a version tag; the `release` workflow runs the full PHP matrix against it and
publishes a GitHub release only if every job passes:

```bash
git tag v1.1.0
git push origin v1.1.0
```

## License

Apache-2.0 — see [LICENSE](https://github.com/morvy/ddev-openlitespeed/blob/master/LICENSE).
