---
title: tyler36/ddev-copilot
github_url: https://github.com/tyler36/ddev-copilot
description: "Run GitHub Copilot inside your DDEV project."
user: tyler36
repo: ddev-copilot
repo_id: 1086871647
default_branch: main
tag_name: v0.2.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-10-31
updated_at: 2025-10-31
workflow_status: cancelled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-copilot/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tyler36/ddev-copilot/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-copilot)](https://github.com/tyler36/ddev-copilot/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-copilot)](https://github.com/tyler36/ddev-copilot/releases/latest)

# DDEV Copilot

## Overview

This add-on integrates [GitHub Copilot](https://github.com/features/copilot) into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get tyler36/ddev-copilot
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command        | Description   |
| -------------- | ------------- |
| `ddev copilot` | Start Copilot |

@See [GitHub Copilot CLI](https://github.com/github/copilot-cli) for further information.

### Authorization

If you're not currently logged in to GitHub, you'll be prompted to use the /login slash command. Enter this command and follow the on-screen instructions to authenticate.

### Authenticate with a Personal Access Token (PAT)

You can also authenticate using a fine-grained PAT with the "Copilot Requests" permission enabled.

1. Visit <https://github.com/settings/personal-access-tokens/new>
1. Under "Permissions," click "add permissions" and select "Copilot Requests"
1. Generate your token
1. Add the token to your environment via the environment variable `GH_TOKEN` or `GITHUB_TOKEN` (in order of precedence).
  For example:

    ```ini
    # .ddev/.env
    GH_TOKEN=super-secret-github-pat
    ```

> [!CAUTION]
> `.ddev/.env` and your PAT are considered sensitive information. Do *NOT* commit these values to your repository.

## Credits

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
