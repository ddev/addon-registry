---
title: "Lullabot/ddev-gitleaks"
github_url: "https://github.com/Lullabot/ddev-gitleaks"
description: "Add gitleaks to your ddev instance and check for secrets on startup."
user: "Lullabot"
repo: "ddev-gitleaks"
repo_id: 1281315994
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-06-26"
updated_at: "2026-06-26"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/Lullabot/ddev-gitleaks/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/Lullabot/ddev-gitleaks/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/Lullabot/ddev-gitleaks)](https://github.com/Lullabot/ddev-gitleaks/commits)
[![release](https://img.shields.io/github/v/release/Lullabot/ddev-gitleaks)](https://github.com/Lullabot/ddev-gitleaks/releases/latest)

# DDEV Gitleaks

## Overview

This add-on installs the [gitleaks](https://github.com/gitleaks/gitleaks) secret
scanner into your [DDEV](https://ddev.com/) project's web container and adds a
`post-start` hook that scans the container environment and project `.env` files
for likely secrets and API keys.

**It warns, it never blocks.** DDEV propagates global `web_environment` into every
project, so a secret set globally (for example `TERMINUS_MACHINE_TOKEN`) becomes
readable by any process in the web container — including AI coding assistants such
as Claude Code. When the scan finds something, it prints a redacted warning at the
end of `ddev start` and exits 0; it never aborts startup. Secret values are always
redacted in the output.

![Example gitleaks secret-scan warning printed at the end of ddev start, listing redacted findings followed by the warning banner](https://raw.githubusercontent.com/Lullabot/ddev-gitleaks/main/images/secret-scan-warning.webp)

## Installation

```bash
ddev add-on get Lullabot/ddev-gitleaks
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

The scan runs automatically on every `ddev start`. To run it on demand:

```bash
ddev exec gitleaks-scan
```

A clean project prints nothing and exits 0. When secrets are detected, a redacted
warning banner is printed. The scan never changes the exit status of `ddev start`.

## What is scanned

- The web container environment (`env`), where DDEV global and project
  `web_environment` values appear.
- `.env`-style files under the project root (`vendor/`, `node_modules/`, `.git/`,
  and `.env.example`/template files are skipped).

Benign DDEV-provided variables are allowlisted by name prefix in
`.ddev/web-build/gitleaks.toml` to avoid false positives. Add project-specific
benign variables there if needed.

## Advanced Customization

Pin a different gitleaks version with the `GITLEAKS_VERSION` build arg in
`.ddev/web-build/Dockerfile.gitleaks` (Renovate keeps the default up to date),
then rebuild:

```bash
ddev debug rebuild
```

## Credits

**Contributed and maintained by [@Lullabot](https://github.com/Lullabot)**
