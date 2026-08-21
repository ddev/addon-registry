---
title: "dustinatx/ddev-share-helper-for-wp"
github_url: "https://github.com/dustinatx/ddev-share-helper-for-wp"
description: "DDEV add-on for WordPress that makes ddev share (cloudflared/ngrok tunnels) work out of the box"
user: "dustinatx"
repo: "ddev-share-helper-for-wp"
repo_id: 1291186960
default_branch: "main"
tag_name: "v1.0.2"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-07-06"
updated_at: "2026-07-07"
workflow_status: "success"
stars: 0
---

# DDEV Share Helper for WordPress

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/dustinatx/ddev-share-helper-for-wp/actions/workflows/tests.yml/badge.svg)](https://github.com/dustinatx/ddev-share-helper-for-wp/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/dustinatx/ddev-share-helper-for-wp)](https://github.com/dustinatx/ddev-share-helper-for-wp/commits)
[![release](https://img.shields.io/github/v/release/dustinatx/ddev-share-helper-for-wp)](https://github.com/dustinatx/ddev-share-helper-for-wp/releases/latest)

A [DDEV](https://ddev.com) add-on for WordPress that makes `ddev share` (cloudflared/ngrok tunnels) work out of the box.

## What it does

Normally, sharing a local WordPress DDEV site over a tunnel breaks: WordPress
generates every URL (assets, links, redirects, REST endpoints) using its
configured local site URL, so visitors hitting the tunnel hostname get
mixed-content errors, broken links, and redirect loops back to the
unreachable local URL.

This add-on installs a single
[must-use plugin](https://developer.wordpress.org/advanced-administration/plugins/mu-plugins/),
`ddev-share-helper-for-wp.php`, that rewrites URLs on the fly for tunneled
requests only:

- Local requests (`Host` = your `*.ddev.site` URL) are completely unaffected.
- Tunneled requests (`Host` = the tunnel hostname) get every generated URL,
  and the final HTML/JSON output, rewritten to the tunnel host — with no
  changes to the database.

A must-use plugin (or "mu-plugin") is a plain PHP file placed in
`wp-content/mu-plugins/`. WordPress loads every file in that directory
automatically, before normal plugins — there is nothing to activate, and it
never shows up as deactivatable in the admin plugin list. That makes it a
good fit here: install the add-on and sharing just works. Removing the
add-on deletes the file, which fully removes the plugin.

## Requirements

- A DDEV project of type `wordpress` with a standard WordPress installation —
  one where the `wp-content/` directory sits alongside `wp-admin/` and
  `wp-includes/`, as in a default install. Composer-based setups that move
  `wp-content` somewhere else (such as Bedrock, which uses `web/app/`) are
  not currently supported.
- Multisite installs aren't currently supported.
- `ddev share` uses [ngrok](https://ngrok.com) by default, which requires a
  free ngrok account and authtoken — see the
  [DDEV sharing docs](https://docs.ddev.com/en/stable/users/topics/sharing/)
  for setup and for alternatives like cloudflared.

## Installation

```bash
ddev add-on get dustinatx/ddev-share-helper-for-wp
ddev restart
```

Then share your site as usual:

```bash
ddev share
```

## Removal

```bash
ddev add-on remove share-helper-for-wp
```

## Issues

If you run into problems, please [file an issue](https://github.com/dustinatx/ddev-share-helper-for-wp/issues).

## Maintainer

@dustinatx
