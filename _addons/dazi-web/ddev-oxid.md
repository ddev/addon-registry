---
title: "dazi-web/ddev-oxid"
github_url: "https://github.com/dazi-web/ddev-oxid"
description: ""
user: "dazi-web"
repo: "ddev-oxid"
repo_id: 943260880
default_branch: "main"
tag_name: "v0.1.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2025-03-05"
updated_at: "2026-09-30"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/dazi-web/ddev-oxid/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/dazi-web/ddev-oxid/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/dazi-web/ddev-oxid)](https://github.com/dazi-web/ddev-oxid/commits)
[![release](https://img.shields.io/github/v/release/dazi-web/ddev-oxid)](https://github.com/dazi-web/ddev-oxid/releases/latest)

# ddev-oxid

ddev-oxid is a DDEV add-on designed to streamline the installation and configuration of the OXID eShop within a DDEV environment. It automates common setup tasks, making it easier and faster for developers to get a fully configured OXID eShop up and running locally.

## Why use ddev-oxid?

- **Simplified Installation:**  
  Automatically installs the OXID eShop via Composer and sets up the environment with a single interactive command.

- **Automated Configuration:**  
  Runs the necessary OXID CLI commands to configure your shop, including shop setup, demo data installation, administrator creation, license addition, and theme activation.

- **Optimized for DDEV:**  
  Leverages DDEV’s containerized environment, ensuring consistent configuration and dependencies, and uses DDEV environment variables (e.g., for dynamic shop URL generation).

- **Customizable and Interactive:**  
  The add-on provides interactive prompts to select the OXID version, set database parameters (defaulting to DDEV’s standard values), and optionally run additional configuration steps.


## Requirements

- **DDEV:**  
  Version 1.24.10 or above is required.

- **PHP:**  
  Ensure that the PHP version configured in DDEV matches the requirements of the selected OXID eShop version (OXID 7.x needs PHP 8.0+ (7.5 needs 8.3), OXID 6.5 runs on PHP 7.4 to 8.1). The installer filters out versions that are too old or too new for the configured PHP version. Composer runs inside the DDEV web container, so it does not need to be installed on the host.

- **Git:**  
  Git is required for fetching the add-on from the repository.

- **Apache FPM:**  
    This add-on is configured to run with Apache FPM. The document root is set to `htdocs/source` for proper serving of the OXID eShop.


## Quick start

In an empty project directory:

```shell
ddev config --project-type=php --php-version=8.3
ddev add-on get dazi-web/ddev-oxid
ddev restart
ddev install-oxid
```

`ddev install-oxid` guides you through the rest:

1. Pick OXID 7 (default) or 6.
2. Pick a version. Only versions that fit the PHP version of the project are listed, the newest is recommended (just press Enter).
3. Choose the quick installation (demo data, an admin user with a generated password, default theme) or answer the individual questions.

At the end you get the shop URL, the admin URL and the login. The checks run before the first question: a non-empty `htdocs`, a missing `ddev restart` (wrong docroot) or a PHP version that does not fit are reported right away.

Which PHP version to use: OXID 7.5 needs PHP 8.3, 7.2 to 7.4 need 8.2 or newer, 7.1 needs 8.1, 7.0 needs 8.0. OXID 6.5 runs on PHP 7.4 to 8.1.

### Non-interactive installation

`--quick` (or `-y`) accepts all defaults without any question: OXID 7 unless `--major` is given, the newest stable version that fits the project's PHP version, demo data, an admin user `admin@example.com` with a generated password, the default theme. Every flag overrides its default:

```shell
ddev install-oxid --quick
ddev install-oxid --major=7 --version=dev-b-7.4-ce --language=de --demo-data \
  --admin-email=admin@example.com --admin-password=secret --theme=apex
```

Flags: `--quick` / `-y`, `--major` (`6` or `7`), `--version` (e.g. `dev-b-7.4-ce`), `--edition` (`ce` default, `pe`, `ee`), `--language` (OXID 7), `--demo-data` / `--no-demo-data`, `--admin-email`, `--admin-password` (or env `OXID_ADMIN_PASSWORD`), `--no-admin`, `--license-key`, `--theme`, `--shop-url`.

## Tested with

Installation and all commands were tested with OXID CE 7.4 (PHP 8.2) and OXID CE 6.5 (PHP 7.4) on DDEV 1.25. The minimum PHP versions for the other OXID releases are taken from the OXID requirements and have not been tried out, and the PE and EE editions have not been tested.

Note: `--admin-password` and `ddev oxid-admin` pass the password as a command line argument, so it is visible in the process list inside the container while the command runs. This is fine for local development.

## Commands

| Command | Description |
|---|---|
| `ddev install-oxid` | Install and configure OXID eShop 6 or 7 |
| `ddev oxid-cc` | Clear the cache (`source/tmp`) |
| `ddev oxid-views` | Regenerate the database views |
| `ddev oxid-console <args>` | Run `vendor/bin/oe-console` |
| `ddev oxid-module activate\|deactivate <id>` | Activate or deactivate a module |
| `ddev oxid-theme <id>` | Activate a theme |
| `ddev oxid-admin <email> <password>` | Create an admin user |
| `ddev oxid-log [tail-args]` | Follow `source/log/oxideshop.log` (default: last 50 lines, live) |
| `ddev oxid-reset [--demo-data] [-y]` | Drop the database and run the shop setup again |
| `ddev oxid-version` | Show PHP and OXID package versions |
| `ddev oxid-migrate` | Run database migrations |
| `ddev oxid-update` | `composer update`, migrations, views, cache clear |
| `ddev oxid-dev on\|off` | Switch `iDebug` in `config.inc.php` |
| `ddev oxid-module-create <vendor> <id>` | Scaffold a module in `source/modules` |
| `ddev oxid-open [shop\|admin]` | Open shop or admin in the browser (host command) |
| `ddev oxid-db-dump [name]` | Export the DB to `.ddev/db-dumps` with a timestamp (host command) |
| `ddev oxid-db-import [file]` | Import a dump, default: the newest (host command) |

## Tips

- **Mail:** DDEV routes PHP `mail()` to Mailpit (`ddev launch -m`). Keep OXID's mail method on "mail" to see outgoing mails there.
- **Xdebug:** Enable with `ddev xdebug on`; the project root maps to `/var/www/html` in PhpStorm.

## OXID 6 and OXID 7

`ddev install-oxid` asks for the major version (or takes `--major=6|7`). The versions differ, and the add-on handles that:

| | OXID 6.x | OXID 7.x |
|---|---|---|
| PHP | 7.4 to 8.1 (`ddev config --php-version=7.4`) | 8.0+, 7.5 needs 8.3 |
| Shop setup | done by the add-on (DB import, config, views); no `oe:setup:shop` | `oe:setup:shop` |
| Theme | `flow` by default, switch in the admin; `ddev oxid-theme` is not available | `apex` by default, `ddev oxid-theme` works |
| Admin user | `ddev oxid-admin` sets the credentials of the existing admin | creates a new admin user |
| `--language` | ignored (German and English are included) | used |
| Composer | security blocking is switched off for install and `oxid-update`, because the old dependencies have known advisories | normal |

After an OXID 6 installation, the setup directory is removed for security. The two SQL files needed by `ddev oxid-reset` are kept outside the web root in `.ddev/oxid/legacy-sql`.

Commands that do not exist in the installed version say so instead of failing with a stack trace. OXID 6 is no longer maintained, use it for legacy projects only.
