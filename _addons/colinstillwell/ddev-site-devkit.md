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
updated_at: 2025-10-01
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/colinstillwell/ddev-site-devkit/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/colinstillwell/ddev-site-devkit/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/colinstillwell/ddev-site-devkit)](https://github.com/colinstillwell/ddev-site-devkit/commits)
[![release](https://img.shields.io/github/v/release/colinstillwell/ddev-site-devkit)](https://github.com/colinstillwell/ddev-site-devkit/releases/latest)

# DDEV Site Devkit

## Overview

DDEV Site Devkit standardises everyday project tasks across multiple repositories while keeping each project in control of its own logic.

This add-on adds a set of first-class DDEV commands that orchestrate common site workflows such as scaffolding, authentication, build, synchronisation, installation, and testing, as well as switching between development and production modes. Some workflows provide both frontend and backend variants.

The heavy lifting lives in project-owned scripts under `.ddev/site-devkit/site/scripts`, so teams can customise behaviour per project without forking the add-on.

**What you get**
* `site-` commands: each command calls a matching script from your project-owned scripts.
* `devkit` commands: tools you can use directly or from within your project-owned scripts.
* Example scripts are automatically generated in `.ddev/site-devkit/site/scripts`.
* Safe, CI-friendly defaults: idempotent tasks with strict error handling and sensible exit codes, making them pipeline-safe.
* Framework agnostic: works with any tech stack.

## Install or update

```bash
ddev add-on get colinstillwell/ddev-site-devkit
ddev restart
```

After installing or updating, commit the changes this add-on makes under `.ddev`. In most cases these are in `.ddev/site-devkit` and `.ddev/commands`.

## Usage

There are two types of commands provided by this add-on:
* **`ddev site-*`**: project workflows backed by your own scripts.
* **`ddev devkit-*`**: helper tools you can use directly or inside those scripts.

### `ddev devkit-*` commands

| Command                      | Description                                                                              |
| ---------------------------- | ---------------------------------------------------------------------------------------- |
| `devkit-config-diff`         | Compare config and report keys present in reference but missing in target                |
| `devkit-config-get`          | Get a config value by name from a given format and location                              |
| `devkit-db-import`           | Interactively import an SQL dump into the project database                               |
| `devkit-file-copy`           | Copy a file from source to destination within the project, skipping if it already exists |
| `devkit-minio-create-bucket` | Create a MinIO bucket if it does not exist, and set its policy                           |
| `devkit-script-run`          | Run a script on the host or in a specified service                                       |

### `ddev site-*` commands

| Command                 | Description                    | Examples                                                                |
| ----------------------- | ------------------------------ | ----------------------------------------------------------------------- |
| `site-auth`             | Authentication tasks           | Authenticate SSH keys                                                   |
| `site-build`            | Build tasks                    | Wraps tasks for the backend and frontend                                |
| `site-build-backend`    | Backend build tasks            | Composer install                                                        |
| `site-build-frontend`   | Frontend build tasks           | NPM install                                                             |
| `site-install`          | Installation tasks             | New project installs the application; existing project builds and syncs |
| `site-mode-development` | Enable development mode        | Disable caches, enable verbose logging                                  |
| `site-mode-production`  | Enable production mode         | Enable caches, aggregate CSS and JS                                     |
| `site-scaffold`         | Scaffolding tasks              | Copy required files, set permissions                                    |
| `site-sync`             | Synchronisation tasks          | Wraps tasks for the backend and frontend                                |
| `site-sync-backend`     | Backend synchronisation tasks  | Database import, public files                                           |
| `site-sync-frontend`    | Frontend synchronisation tasks | Images, compiled CSS and JS                                             |
| `site-test`             | Testing tasks                  | Wraps tasks for the backend and frontend                                |
| `site-test-backend`     | Backend testing tasks          | Unit, kernel, integration                                               |
| `site-test-frontend`    | Frontend testing tasks         | Unit, end to end                                                        |

## Customising the generated scripts

When you install this add-on, example scripts are copied into your project at `.ddev/site-devkit/site/scripts`.

Each `ddev site-*` command maps 1:1 to a script in that directory. These scripts are yours to edit and should be committed to your repository.

If your project doesn't need frontend or backend scripts, just leave them unused. They'll reappear on update if deleted.

### How to modify
1. Open the matching script.
2. Remove the `#ddev-generated` line (this prevents the script being replaced on update).
3. Remove the `## Script provided by https://github.com/colinstillwell/ddev-site-devkit.` line.
4. Keep the shebang and safety flags:
   ```bash
   #!/usr/bin/env bash

   # Exit on error; treat unset variables as errors; fail pipelines if any command fails
   set -euo pipefail
   ```
5. Replace the placeholder content with your project's logic.
6. Use the `ddev devkit-*` commands provided by this add-on where useful.

### Update behaviour

When you update the add-on, any script that still has `#ddev-generated` will be overwritten. Once you remove that line, the script is considered project-owned and will not be touched.

To reset a script back to the example, delete it from your project and reinstall the add-on. A fresh copy will be generated.

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
