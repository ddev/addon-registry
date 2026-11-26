---
title: "portrino/ddev-wordpress-multi-version-plugin"
github_url: "https://github.com/portrino/ddev-wordpress-multi-version-plugin"
description: "🧱 Provides a multi-version WordPress environment to develop and test your WordPress plugin."
user: "portrino"
repo: "ddev-wordpress-multi-version-plugin"
repo_id: 1458212515
default_branch: "main"
tag_name: "1.1.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-11-14"
updated_at: "2026-11-25"
workflow_status: "success"
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/portrino/ddev-wordpress-multi-version-plugin/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/portrino/ddev-wordpress-multi-version-plugin/actions/workflows/tests.yml?query=branch%3Amain)
[![release](https://img.shields.io/github/v/release/portrino/ddev-wordpress-multi-version-plugin)](https://github.com/portrino/ddev-wordpress-multi-version-plugin/releases/latest)

# DDEV WordPress Multi-Version Plugin

Testing a WordPress plugin against several WordPress versions in one DDEV
project. Every version gets its own Composer-based instance in `.Build/<label>`,
its own database and its own hostname; the plugin is required from the project
root through a path repository, so a change is visible on every version at once.

| WordPress | Hostname | Directory | Database |
|---|---|---|---|
| 6.4 | `https://64.<project>.ddev.site` | `.Build/64` | `db_64` |
| 7.1 | `https://71.<project>.ddev.site` | `.Build/71` | `db_71` |

## ✨ Features

- [**Parallel WordPress instances**](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/installation.md): one Composer project per version (`roots/wordpress`, wpackagist), single site or multisite
- [**Theme, extra plugins and private repositories**](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/configuration.md#projectsh): configured in `project.sh`, including child themes from the plugin repository
- [**Fixture import**](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/configuration.md#fixtures): SQL dumps, WordPress exports (WXR) and shell scripts on every install
- [**Install hooks**](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/configuration.md#hooks): custom steps before and after Composer, setup and install
- [**Intro page**](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/commands.md#ddev-open-version-path): every version with its release, links and login on the primary URL
- [**Git worktree support**](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/git-worktrees.md): several checkouts side by side without hostname collisions

## 🔥 Installation

> [!IMPORTANT]
> Requires DDEV >= v1.24.10, the webserver type `apache-fpm` or `nginx-fpm`,
> and the plugin in the project root with a `composer.json` that has a `name`
> and `"type": "wordpress-plugin"`. The plugin directory in the instances is
> `extra.installer-name` or the package name without vendor, the same name
> `composer/installers` uses.

```bash
ddev config --webserver-type=apache-fpm   # or nginx-fpm
ddev add-on get portrino/ddev-wordpress-multi-version-plugin
ddev restart
```

The add-on sets the project type to `wordpress` (for DDEV's WP-CLI) with
`disable_settings_management: true` and the docroot `.Build`, and adds
`.Build/` and `.ddev/config.zz-worktree.local.yaml` to `.gitignore`. Commit `.ddev/`.

### Updating

Run `ddev add-on get portrino/ddev-wordpress-multi-version-plugin && ddev restart`
again. Files with a `#ddev-generated` line are overwritten; remove that line from
your copy to keep a customized file.

## 🚀 Quick start

```bash
ddev install all
ddev open
```

That installs every configured version and opens the intro page. Login:
`/wp-admin/` with `admin` / `Password1!`.

## ⚡ Usage

| Command | Does |
|---|---|
| `ddev install [<version>\|all] [--multisite] [-v]` | Install or rebuild one or all instances |
| `ddev <label> <command>` | Run a command in one instance, e.g. `ddev 71 wp plugin list` |
| `ddev all <command>` | Run a command in every instance, e.g. `ddev all wp cache flush` |
| `ddev open [<version>] [<path>]` | Open the intro page or an instance, e.g. `ddev open 71 /wp-admin/` |
| `ddev worktree-init` / `ddev worktree-remove` | Set up / tear down a git worktree checkout |

## 🏗️ Architecture

The plugin stays in the project root. Each WordPress version is a Composer
project of its own in `.Build/<label>/`, which requires the plugin through
symlinks in `packages/<slug>/`:

```text
.
├── composer.json            your plugin ("type": "wordpress-plugin")
├── .ddev/                   the add-on, plus your project.sh, hooks, overrides
└── .Build/                  docroot of the primary URL, gitignored
    ├── index.php            intro page
    ├── 64/
    │   ├── composer.json
    │   ├── packages/<slug>/ -> project root
    │   ├── vendor/
    │   └── public/          docroot of 64.<project>.ddev.site
    │       ├── wp-config.php
    │       ├── wp-source/   WordPress core
    │       └── wp-content/
    └── 71/
```

The webserver picks the instance by the numeric host label and does
WordPress' rewrites itself. `.Build/` isn't shared, so every checkout
(including every git worktree) runs its own `ddev install`. Details:
[instance layout](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/installation.md#instance-layout).

## 📚 Documentation

| Topic | What's inside |
|---|---|
| [Installing instances](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/installation.md) | `ddev install` flags and steps, the instance layout |
| [Commands](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/commands.md) | Running commands per version or in all instances, `ddev open`, quoting |
| [Configuration](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/configuration.md) | Environment variables, `project.sh`, private repositories, fixtures, hooks, webserver configs |
| [Git worktrees](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/docs/git-worktrees.md) | Running several checkouts side by side |

## 🗑️ Remove

```bash
ddev add-on remove wordpress-multi-version-plugin
```

`.Build/` and the databases stay.

## 💎 Credits

Inspired by [ddev-typo3-multi-version-extension](https://github.com/konradmichalik/ddev-typo3-multi-version-extension).

## ⭐ License

[Apache-2.0](https://github.com/portrino/ddev-wordpress-multi-version-plugin/blob/main/LICENSE)
