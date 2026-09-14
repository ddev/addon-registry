---
title: "dearaujoj/ddev-cmsms"
github_url: "https://github.com/dearaujoj/ddev-cmsms"
description: "DDEV add-on: instant disposable CMS Made Simple dev sites for module, plugin, and theme development."
user: "dearaujoj"
repo: "ddev-cmsms"
repo_id: 1312221944
default_branch: "main"
tag_name: "v1.4.0"
ddev_version_constraint: ">= v1.24.0"
dependencies: ["ddev/ddev-phpmyadmin"]
type: "contrib"
created_at: "2026-07-25"
updated_at: "2026-07-27"
workflow_status: "success"
stars: 0
---

# ddev-cmsms

[![tests](https://github.com/dearaujoj/ddev-cmsms/actions/workflows/tests.yml/badge.svg)](https://github.com/dearaujoj/ddev-cmsms/actions/workflows/tests.yml)
[![project](https://img.shields.io/badge/DDEV-Add--on-blue)](https://ddev.com)

A [DDEV](https://ddev.com) add-on that turns a CMS Made Simple extension repository (module, plugin, or theme) into a full, disposable CMSMS dev site. It downloads the official CMSMS installer, drives it headlessly (no browser wizard), symlinks your extension into a running core, and installs it — all from `ddev start`. The core lives in a gitignored `.cmsms/` directory next to your extension source, so the repo you're developing stays clean.

## Requirements

- [DDEV](https://ddev.readthedocs.io/en/stable/) **v1.24.0 or newer** — new to DDEV? Follow the [installation guide](https://ddev.readthedocs.io/en/stable/users/install/ddev-installation/) first; it also installs everything below.
- A Docker provider supported by DDEV ([Docker Desktop, OrbStack, Colima, …](https://ddev.readthedocs.io/en/stable/users/install/docker-installation/)).

Everything else (PHP, MariaDB, the CMSMS core itself) runs inside DDEV's containers — nothing CMSMS-related is installed on your machine.

## Quickstart

```bash
cd my-module/
mkdir -p .cmsms/public
ddev config --project-type=php --docroot=.cmsms/public
ddev add-on get dearaujoj/ddev-cmsms
ddev cmsms setup
ddev start
ddev launch        # frontend; `ddev cmsms admin` opens /admin (admin/admin)
```

Admin credentials default to `admin` / `admin`. Override them with the `CMSMS_ADMIN_USER` / `CMSMS_ADMIN_PASSWORD` environment variables before the first install (see [Configuration](#configuration)).

### Starting from an empty directory

No code yet? The same flow works — `ddev cmsms setup` notices the empty repo
and generates a starter for you (or run `ddev cmsms scaffold` explicitly):

    mkdir my-module && cd my-module
    mkdir -p .cmsms/public
    ddev config --project-type=php --docroot=.cmsms/public
    ddev add-on get dearaujoj/ddev-cmsms
    ddev cmsms setup --type module --name MyModule --yes   # scaffolds MyModule.module.php & co.
    ddev start                                             # installs CMSMS with your new module live

### Workspace mode: several extensions, one site

Developing modules that depend on each other (plus plugins, or the admin
theme that showcases them)? Make the project root a workspace — one
subdirectory per extension, named after it. Subdirectories can be separate
git clones, submodules, or plain folders in a monorepo; the add-on doesn't
care:

    my-suite/
    ├── ModuleA/ModuleA.module.php
    ├── ModuleB/ModuleB.module.php      # GetDependencies() → ModuleA
    └── MyTags/function.mytag.php

    ddev config --project-type=php --docroot=.cmsms/public
    ddev add-on get dearaujoj/ddev-cmsms
    ddev cmsms setup                    # detects the subdirs, proposes the list
    ddev start                          # installs the site with ALL of them linked

The list lives in `.ddev/config.cmsms-project.yaml` as
`CMSMS_EXTENSIONS=module:ModuleA module:ModuleB plugin:MyTags` — **order is
install order**, so list dependencies before dependents. Grow the workspace
with `ddev cmsms add module ModuleC` (scaffolds the subdir if missing), check
everything with `ddev cmsms status`, and package one module with
`ddev cmsms package ModuleA`. A failing entry (missing directory, failed
registration) is warned and skipped — it never breaks `ddev start`. If you
list several `theme:` entries, the last one becomes the active admin theme.

## Commands

| Command | Description |
| --- | --- |
| `ddev cmsms setup [--type module\|plugin\|theme] [--name NAME] [--version X.Y.Z] [--extensions "type:Name …"] [--yes\|-y]` | Detects (or accepts) the extension type/name and CMSMS version, validates them, and writes `.ddev/config.cmsms-project.yaml`. Without `--yes`, prompts interactively for anything not passed as a flag, offering the auto-detected value as the default. `--yes`/`-y` runs fully non-interactively, falling back to auto-detected values for anything not passed. `--extensions` switches to workspace mode (mutually exclusive with `--type`/`--name`); omit it with an empty repo containing extension-named subdirectories and setup detects and proposes the list. |
| `ddev cmsms scaffold [--type T --name N] [--dir NAME] [--yes\|-y]` | Generates a working starter: module (module class + default/admin actions + template + lang), plugin (`function.<name>.php`), or admin theme (a renamed copy of the core's OneEleven — requires an installed site). `--dir NAME` targets `./NAME/` instead of the repo root (for theme, `NAME` must equal the extension name). Refuses to overwrite existing files. `setup` offers this automatically when the repo is empty. |
| `ddev cmsms load <zip> [--use]` | Validates a local CMSMS installer zip (e.g. `~/Downloads/cmsms-2.2.23-install.zip`) and puts it in the cache under its canonical name — the version is read from the zip's own payload, so renamed files and custom/unreleased builds work. `--use` also sets `CMSMS_VERSION` in the project config. |
| `ddev cmsms add <module\|plugin\|theme> <Name> [--yes\|-y]` | Workspace projects only: appends `type:Name` to `CMSMS_EXTENSIONS` in `.ddev/config.cmsms-project.yaml` (no-op with a note if already listed), and offers to scaffold a starter into `./Name/` when that directory doesn't exist yet. |
| `ddev cmsms admin` | Prints the actual admin credentials (`admin`/`admin` unless overridden, read from the running container) and opens `/admin` in your browser. |
| `ddev cmsms status` | Reports the installed CMSMS core version vs. the configured one and the extension type/name; for module projects, also reports whether it's symlinked and registered in the database. |
| `ddev cmsms reinstall [--yes]` | Drops the database and clears the core/installer files (keeping the download cache and `uploads/`), then re-runs the full install. Prompts for confirmation unless `--yes`/`-y` is passed — this is destructive. |
| `ddev cmsms package [<Name>]` | Module projects only. Builds the distributable module XML via the module's own `CreateXMLPackage()`, written to `dist/<Name>-<version>.xml`. `<Name>` is required in workspace mode, and must name a `module:` entry in `CMSMS_EXTENSIONS`. |

Setup validates `--name` against `^[A-Za-z0-9_]+$` and `--version` against a dotted-numeric pattern; invalid values are rejected with an error before anything is written.

The install pipeline itself (fetch → files → db → link) runs automatically via a `post-start` hook (`ddev cmsms-install all`), and each stage is idempotent — safe to re-run individually with `ddev cmsms-install <stage>` for debugging.

## Configuration

Extension identity and target CMSMS version live in `.ddev/config.cmsms-project.yaml`, written by `ddev cmsms setup` (edit it by hand and `ddev restart` if you prefer):

| Variable | Purpose |
| --- | --- |
| `CMSMS_EXTENSIONS` | Workspace mode: ordered space-separated `type:Name` list (see [Workspace mode](#workspace-mode-several-extensions-one-site)); when set it takes precedence over `CMSMS_EXT_TYPE`/`CMSMS_EXT_NAME`. |
| `CMSMS_EXT_TYPE` | `module`, `plugin`, or `theme`. |
| `CMSMS_EXT_NAME` | Extension name (module class name, theme name — not used for plugins). |
| `CMSMS_VERSION` | CMSMS core version to install; must be a key in `cmsms/versions.json`. |

Other environment variables, settable via `.ddev/config.cmsms-project.yaml`'s `web_environment` block or any other DDEV env mechanism:

| Variable | Purpose |
| --- | --- |
| `CMSMS_INSTALLER_URL` | Overrides the resolved download URL — point it at a direct `cmsms-X.Y.Z-install.zip` URL (mirror or local file server) if the default Forge URL is unreachable (a loaded local zip takes precedence). |
| `CMSMS_SAMPLE_CONTENT` | Set to `1` to install CMSMS's demo/sample content instead of the minimal default page. |
| `CMSMS_ADMIN_USER` / `CMSMS_ADMIN_PASSWORD` | Override the admin account created during install (default `admin`/`admin`). |
| `CMSMS_LINK_EXCLUDE` | Space-separated list of top-level repo entries to skip when symlinking (default: `.ddev .cmsms .git .gitignore dist docs`). |
| `CMSMS_INSTALL_VERBOSE` | Set to `1` for verbose per-module logging during the headless DB install. |

### Using a local installer zip

Have the installer on disk (offline work, betas, or cores you built
yourself)? Load it — the version is derived from the zip's payload, and
versions.json is bypassed entirely for that version:

    ddev cmsms load ~/Downloads/cmsms-2.2.23-install.zip --use
    ddev restart && ddev cmsms reinstall --yes   # if the project was already installed

Without `--use`, the zip just lands in `.ddev/cmsms/cache/` and you select
it with `ddev cmsms setup --version 2.2.23`. Power users can also drop a
correctly named `cmsms-<version>-install.zip` into `.ddev/cmsms/cache/` by
hand, or point `CMSMS_INSTALLER_URL` at any `http(s)://` or `file://` URL
reachable from inside the web container. A loaded zip for the configured
version always wins over `CMSMS_INSTALLER_URL`.

## Testing against another core / PHP version

```bash
# another CMSMS core
edit .ddev/config.cmsms-project.yaml   # change CMSMS_VERSION
ddev restart && ddev cmsms reinstall --yes

# another PHP
ddev config --php-version=8.1 && ddev restart
```

Supported core versions are the keys of [`cmsms/versions.json`](https://github.com/dearaujoj/ddev-cmsms/blob/main/cmsms/versions.json), regenerated with `scripts/update-versions.sh` (scrapes the CMSMS Forge file listing). CI verifies 2.2.22 (PHP 8.3 and 8.1), 2.2.21, and 2.2.19 on every push and weekly on a schedule (to catch download-URL rot); other listed versions are expected to work but aren't continuously tested — see [`.github/workflows/tests.yml`](https://github.com/dearaujoj/ddev-cmsms/blob/main/.github/workflows/tests.yml).

## Xdebug

```bash
ddev xdebug on
```

Set a single IDE path mapping from your repository root to `/var/www/html` — the CMSMS core (`.cmsms/public`) lives inside the repo, so this one mapping covers breakpoints in both your extension code and core files.

## Mail

[Mailpit](https://ddev.readthedocs.io/en/stable/users/usage/troubleshooting/#mailpit) is preconfigured — the installed site's mail preferences point at `localhost:1025` (SMTP) out of the box. View captured mail with:

```bash
ddev launch -m
```

## How it works

CMSMS's official installer is a self-extracting phar that runs a browser wizard and explicitly refuses to run non-interactively from the CLI. `ddev-cmsms` sidesteps the wizard entirely: it extracts the phar's bundled core and installer app, then includes the installer's own `app/install/*.php` scripts (schema, base data, admin account, content, system modules) directly inside a small PHP driver that reconstructs the scope and globals they expect — the same steps the wizard's `wizard_step8`/`wizard_step9` perform, without a browser in the loop.

Your extension isn't copied into the core. Modules and themes-at-repo-root get an individual symlink per top-level file/directory into `modules/<Name>/` or `admin/themes/<Name>/`, skipping a configurable exclusion list (`.ddev`, `.cmsms`, `.git`, etc.) — per-entry linking, rather than symlinking the whole repo root in one shot, is what avoids a symlink cycle (the core lives inside `.cmsms/`, which lives inside the repo being linked). Plugins link only their `function.*.php`/`modifier.*.php` files into `plugins/`. A theme that lives in its own subdirectory is simpler: the whole subdirectory gets a single symlink into `admin/themes/<Name>/`, since it can't contain `.cmsms/` itself.

The `.cmsms/` directory (downloaded core, installer scratch space, generated `config.php`) is disposable and gitignored automatically by the add-on's post-install hook, so it never pollutes your extension's repository.

## Troubleshooting

- **Installer download fails**: set `CMSMS_INSTALLER_URL` to a working direct `cmsms-X.Y.Z-install.zip` URL (mirror or local file) and re-run.
- **Install partially completed / left in a bad state**: `ddev cmsms-install all` is safe to re-run — every stage (fetch, files, db, link) checks its own completion marker and skips work already done.
- **Not sure what's installed or linked**: `ddev cmsms status` reports the core version and configured extension in one shot; for module projects it also reports symlink state and DB registration.
- **Need a database GUI**: this add-on depends on [`ddev/ddev-phpmyadmin`](https://github.com/ddev/ddev-phpmyadmin) — run `ddev phpmyadmin` to open it.

## Contributing

PRs welcome — see [CONTRIBUTING.md](https://github.com/dearaujoj/ddev-cmsms/blob/main/CONTRIBUTING.md) for the test
requirements (CI enforces that implementation changes come with tests),
how to run the bats suites locally, and the portability conventions.

## Credits / License

Built for [CMS Made Simple](https://www.cmsmadesimple.org) extension developers on top of [DDEV](https://ddev.com).

Licensed under the [Apache License 2.0](https://github.com/dearaujoj/ddev-cmsms/blob/main/LICENSE).
