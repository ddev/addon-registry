---
title: "davereid/ddev-drupalorg-cli"
github_url: "https://github.com/davereid/ddev-drupalorg-cli"
description: "ddev add-on for setting up the Drupal.org CLI"
user: "davereid"
repo: "ddev-drupalorg-cli"
repo_id: 1455822360
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-11-12"
updated_at: "2026-11-12"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/davereid/ddev-drupalorg-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/davereid/ddev-drupalorg-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/davereid/ddev-drupalorg-cli)](https://github.com/davereid/ddev-drupalorg-cli/commits)
[![release](https://img.shields.io/github/v/release/davereid/ddev-drupalorg-cli)](https://github.com/davereid/ddev-drupalorg-cli/releases/latest)

# DDEV drupalorg CLI

## Overview

This add-on installs the [drupalorg CLI](https://github.com/mglaman/drupalorg-cli) in the web container of your [DDEV](https://ddev.com/) project, for working with drupal.org issues, patches and merge requests from inside the container.

## Installation

```bash
ddev add-on get davereid/ddev-drupalorg-cli
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev exec drupalorg list` | List the CLI's commands |
| `ddev exec drupalorg issue:show <nid>` | Show a drupal.org issue |

## Updating the CLI

The latest release is downloaded when the web image is built, and Docker caches it. To get a newer release:

```bash
ddev debug rebuild
ddev restart
```

## Pinning a version

To install a specific release instead of the latest, create `.ddev/docker-compose.drupalorg-cli-version.yaml`:

```yaml
services:
  web:
    build:
      args:
        DRUPALORG_CLI_VERSION: "0.15.1"
```

Then run `ddev restart`.

## Removal

```bash
ddev add-on remove drupalorg-cli
ddev debug rebuild
ddev restart
```

## Releases

Releases are made by [release-please](https://github.com/googleapis/release-please) from [conventional commits](https://www.conventionalcommits.org/) on `main`: `fix:` for a patch release, `feat:` for a minor one, and `feat!:` or a `BREAKING CHANGE:` footer for a major one. It keeps a release pull request open with the next version and changelog; merging it tags the release.

## Credits

**Contributed and maintained by [@davereid](https://github.com/davereid)**
