---
title: lpeabody/ddev-ai
github_url: https://github.com/lpeabody/ddev-ai
description: "DDEV add-on to make AI-related services available in the environment."
user: lpeabody
repo: ddev-ai
repo_id: 1018102141
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-07-11
updated_at: 2025-07-11
workflow_status: unknown
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/lpeabody/ddev-ai/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/lpeabody/ddev-ai/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/lpeabody/ddev-ai)](https://github.com/lpeabody/ddev-ai/commits)
[![release](https://img.shields.io/github/v/release/lpeabody/ddev-ai)](https://github.com/lpeabody/ddev-ai/releases/latest)

# DDEV AI

## Overview

This add-on integrates commonly necessary AI services into your [DDEV](https://ddev.com/) project.

- Milvus for Vector Database

## Installation

```bash
ddev add-on get lpeabody/ddev-ai
ddev restart
# You may also need to chown the .ddev directory recursively as some services
# are configured to map data volumes in the DDEV directory.
# sudo chown -R $USER:$USER .ddev/milvus
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command               | Description                                        |
|-----------------------|----------------------------------------------------|
| `ddev describe`       | View service status and used ports for AI services |
| `ddev logs -s milvus` | Check Milvus logs                                  |

### Drupal

1. `composer require drupal/ai_vdb_provider_milvus`
2. `drush en ai_vdb_provider_milvus`
3. Configure the module at `/admin/config/ai/vdb_providers/milvus` with the following settings:
   - Host: `https://milvus`
   - Port: `19530`

**Contributed and maintained by [@lpeabody](https://github.com/lpeabody)**
