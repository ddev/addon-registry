---
title: colinstillwell/ddev-site-devkit
github_url: https://github.com/colinstillwell/ddev-site-devkit
description: "DDEV addon providing site commands and tools for scaffolding, building, installing, updating and testing projects."
user: colinstillwell
repo: ddev-site-devkit
repo_id: 1061687339
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-09-22
updated_at: 2025-09-29
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/colinstillwell/ddev-site-devkit/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/colinstillwell/ddev-site-devkit/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/colinstillwell/ddev-site-devkit)](https://github.com/colinstillwell/ddev-site-devkit/commits)
[![release](https://img.shields.io/github/v/release/colinstillwell/ddev-site-devkit)](https://github.com/colinstillwell/ddev-site-devkit/releases/latest)

# DDEV Site Devkit

## Overview

DDEV Site Devkit standardises everyday project tasks across multiple repositories while keeping each project in control of its own logic. It adds a set of first-class DDEV commands that orchestrate common "site" workflows such as scaffold, build, install, update and test. The heavy lifting lives in project-owned scripts under `.ddev/site-devkit/site/scripts`, so teams can customise behaviour per project without forking the add-on.

**What you get**
* `site-` commands: each command calls a matching script from your project-owned scripts.
* `devkit` commands: tools you can use directly or from within your project-owned scripts.
* Examples: example scripts live in `.ddev/site-devkit/site/scripts/examples`; copy them into place as needed.
* Safe, CI-friendly defaults: idempotent tasks, sensible exit codes, and pass-through arguments for full control.
* Framework agnostic: works with any tech stack.

## Install or update

```bash
ddev add-on get colinstillwell/ddev-site-devkit
ddev restart
```

After installing or updating, commit the changes this add-on makes under `.ddev`. In most cases these are in `.ddev/site-devkit` and `.ddev/commands`.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev devkit-import-database` | Interactively import a SQL dump file into the project. |
| `ddev devkit-run-script` | Run a script on the host or in the web container. |
| `ddev site-build` | Run build tasks. |
| `ddev site-build-backend` | Run backend build tasks. |
| `ddev site-build-frontend` | Run frontend build tasks. |
| `ddev site-initialisation` | Run initialisation tasks. |
| `ddev site-install` | Run install tasks. |
| `ddev site-mode-development` | Enable development mode. |
| `ddev site-mode-production` | Enable production mode. |
| `ddev site-scaffold` | Run scaffolding tasks. |
| `ddev site-test` | Run test tasks. |
| `ddev site-test-backend` | Run backend test tasks. |
| `ddev site-test-frontend` | Run frontend test tasks. |
| `ddev site-update` | Run update tasks. |
| `ddev site-update-backend` | Run backend update tasks. |
| `ddev site-update-frontend` | Run frontend update tasks. |

## Resources

* [DDEV Documentation for Add-ons](https://ddev.readthedocs.io/en/stable/users/extend/additional-services/)
* [DDEV Add-ons: Creating, maintaining, testing](https://www.youtube.com/watch?v=TmXqQe48iqE) (part of the [DDEV Contributor Live Training](https://ddev.com/blog/contributor-training))
* [Advanced Add-On Techniques](https://ddev.com/blog/advanced-add-on-contributor-training/)
* [DDEV Add-on Maintenance Guide](https://ddev.com/blog/ddev-add-on-maintenance-guide/)

## Contributing

1. Work from an issue. If none exists, create one first.
2. Branch from `main` using `issue/<number>-<short-slug>` in lowercase with hyphens.
3. Make your changes and commit with the prefix `[<number>]`.
4. Add or update tests in `tests` as needed.
5. Update `README.md` as needed.
6. Push your branch and open a pull request on [GitHub](https://github.com/colinstillwell/ddev-site-devkit).

## Testing branch or PR

```bash
# Branch
ddev add-on get https://github.com/colinstillwell/ddev-site-devkit/tarball/<branch>

# Pull request
ddev add-on get https://github.com/colinstillwell/ddev-site-devkit/tarball/refs/pull/<pr-number>/head
```

## Release

Create a [release](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository) on [GitHub](https://github.com/colinstillwell/ddev-site-devkit):
* `main` as the target.
* Follow semantic versioning (`MAJOR.MINOR.PATCH`):
  * `MAJOR`: incompatible changes.
  * `MINOR`: backwards compatible feature additions.
  * `PATCH`: backwards compatible fixes.
* Concise notes.

## Credits

**Contributed and maintained by [@colinstillwell](https://github.com/colinstillwell)**
