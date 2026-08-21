---
title: "MurzNN/ddev-cursor"
github_url: "https://github.com/MurzNN/ddev-cursor"
description: "DDEV add-on to manage Cursor IDE configuration directory in the container right"
user: "MurzNN"
repo: "ddev-cursor"
repo_id: 1303588200
default_branch: "main"
tag_name: "v1.0.1"
ddev_version_constraint: ""
dependencies: []
type: "contrib"
created_at: "2026-07-17"
updated_at: "2026-07-17"
workflow_status: "unknown"
stars: 1
---

# ddev-cursor

DDEV add-on to manage Cursor IDE configuration directory in the container right.

## Installation

```sh
ddev add-on get MurzNN/ddev-cursor
```

## The issue

When you attach the IDE to the DDEV container, it automatically creates multiple directories in the home 
directory, and downloads all necessary files there, and everything works great!

But there are several issues with this:

1. The directory disappears after restarting the DDEV container, so every time all files have to be downloaded again.

2. Some directories contain user-related settings, which better to have shared between all projects. 
For example, the list of extensions

3. Cache files are stored together with the data, making the directory size huge and not easy to clean up.

## Solution

1. Add-on creates persistent volumes for the related directories to make them not disappear on each restart.

2. User-related settings mounted to `$HOME/.local/state/ddev-cursor` to be shared between all DDEV projects.

3. Cache-related directories mounted to `$HOME/.cache/ddev-cursor` to be easily cleaned up with all other caches.
