---
title: DiffyWebsite/ddev-diffy
github_url: https://github.com/DiffyWebsite/ddev-diffy
description: "Diffy integration with DDEV"
user: DiffyWebsite
repo: ddev-diffy
repo_id: 832381294
default_branch: main
tag_name: 0.0.15
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2024-07-22
updated_at: 2026-02-27
workflow_status: success
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/DiffyWebsite/ddev-diffy/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/DiffyWebsite/ddev-diffy/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/DiffyWebsite/ddev-diffy)](https://github.com/DiffyWebsite/ddev-diffy/commits)
[![release](https://img.shields.io/github/v/release/DiffyWebsite/ddev-diffy)](https://github.com/DiffyWebsite/ddev-diffy/releases/latest)

# DDEV Diffy

## Overview

[Diffy](https://diffy.website) is a visual regression testing service. It is primarily tailored to testing Drupal and WordPress websites.

This repository is a recipe for [DDEV](https://ddev.com/) to have a local worker that will allow taking screenshots from local website and upload them to Diffy for further comparisons.

## Installation

```bash
ddev add-on get DiffyWebsite/ddev-diffy
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

Register an account in Diffy and create an [API key](https://app.diffy.website/#/keys).

Once you have your container ready after `ddev restart`, we need to provide API key and project ID to `.env` file. For that go to `.ddev/diffy-worker`, copy existing example file `cp .env.example .env` and edit `.env` file to provide your credentials:

```bash
cp .ddev/diffy-worker/.env.example .ddev/diffy-worker/.env
ddev dotenv set .ddev/diffy-worker/.env --diffy-api-key=XXX --diffy-project-id=XXX
```

To run the screenshots do `ddev screenshot`. It will produce a URL to your screenshots uploaded to Diffy. Next you can compare them to your any other sets (production, staging, baseline etc.).

Remember to check our [documentation page](https://docs.diffy.website/features/local-development/ddev-add-on).

| Command | Description |
| ------- | ----------- |
| `ddev screenshot` | Launch Diffy screenshot |
| `ddev describe` | View service status for Diffy |
| `ddev logs -s diffy` | Check Diffy logs |

## Credits

**Contributed and maintained by [@ygerasimov](https://github.com/ygerasimov)**
