---
title: robertoperuzzo/ddev-pgvector
github_url: https://github.com/robertoperuzzo/ddev-pgvector
description: "The Postgres vector database service for DDEV"
user: robertoperuzzo
repo: ddev-pgvector
repo_id: 1039907838
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-08-18
updated_at: 2025-08-18
workflow_status: disabled
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/robertoperuzzo/ddev-pgvector/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/robertoperuzzo/ddev-pgvector/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/robertoperuzzo/ddev-pgvector)](https://github.com/robertoperuzzo/ddev-pgvector/commits)
[![release](https://img.shields.io/github/v/release/robertoperuzzo/ddev-pgvector)](https://github.com/robertoperuzzo/ddev-pgvector/releases/latest)

# DDEV Pgvector

## Overview

This add-on integrates Pgvector into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get robertoperuzzo/ddev-pgvector
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Pgvector |
| `ddev logs -s pgvector` | Check Pgvector logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.pgvector --pgvector-image-tag="pg17"
ddev add-on get robertoperuzzo/ddev-pgvector
ddev restart
```

Make sure to commit the `.ddev/.env.pgvector` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `PGVECTOR_IMAGE_TAG` | `--pgvector-image-tag` | `pg17` |

## Credits

**Contributed and maintained by [@robertoperuzzo](https://github.com/robertoperuzzo)**
