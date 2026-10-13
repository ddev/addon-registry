---
title: "brookemahoney/ddev-vtracer"
github_url: "https://github.com/brookemahoney/ddev-vtracer"
description: "A DDEV add-on that installs the [vtracer](https://github.com/visioncortex/vtracer) CLI binary into the DDEV web container."
user: "brookemahoney"
repo: "ddev-vtracer"
repo_id: 1416075681
default_branch: "main"
tag_name: "1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-10-12"
updated_at: "2026-10-12"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/brookemahoney/ddev-vtracer/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/brookemahoney/ddev-vtracer/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/brookemahoney/ddev-vtracer)](https://github.com/brookemahoney/ddev-vtracer/commits)
[![release](https://img.shields.io/github/v/release/brookemahoney/ddev-vtracer)](https://github.com/brookemahoney/ddev-vtracer/releases/latest)

# DDEV Vtracer

## Overview

This add-on installs the [vtracer](https://github.com/visioncortex/vtracer) CLI binary into your [DDEV](https://ddev.com/) web container.

## Installation

```bash
ddev add-on get brookemahoney/ddev-vtracer
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

Once installed and restarted, run `vtracer` inside the web container:

```bash
ddev vtracer --version
ddev vtracer --help
```

## Architecture Support

- **amd64** (`x86_64`)
- **arm64** (`aarch64`)

The add-on selects the correct pre-built musl binary based on Docker's `TARGETARCH` build argument.

## Version

Pinned to `1.0.0-alpha.4` (latest available release at creation time).

## Credits

**Contributed and maintained by [@brookemahoney](https://github.com/brookemahoney)**
