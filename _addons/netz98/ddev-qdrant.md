---
title: netz98/ddev-qdrant
github_url: https://github.com/netz98/ddev-qdrant
description: "ddev qdrant"
user: netz98
repo: ddev-qdrant
repo_id: 811850723
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-06-07
updated_at: 2024-07-15
stars: 0
---

# DDEV qdrant Addon

This addon sets up a qdrant instance for your DDEV project. qdrant is a vector similarity search engine and vector database.

## Installation

1. Run `ddev get netz98/ddev-qdrant` to install the addon in your exiting ddev project.
2. `ddev restart` to restart your project.

## Usage

After installation, you can access the qdrant instance by visiting `https://<yourname>.ddev.site:6333`.
For accessing the dashboard you can visit `https://<yourname>.ddev.site:6333/dashboard`

Run `ddev describe` to list your project's services and their URLs.

Different clients can be found  in [qdrant Github](https://github.com/qdrant/qdrant?tab=readme-ov-file#clients)

## Configuration

### qdrant Settings

Settings are in general defined in the file `.ddev/docker-compose.qdrant.yaml` via environment variables.

An overview of the possible configrations can be found here:
https://qdrant.tech/documentation/guides/configuration/#configuration

### Collections

The qdrant collections will be created as folder in  `.ddev/qdrant_data/collections`.

## Logging

qdrant logs are directed to the container's stdout. You can view the logs with `ddev logs -s qdrant`.
