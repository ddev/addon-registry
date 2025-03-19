---
title: lpeabody/ddev-drupal-multisite-databases
github_url: https://github.com/lpeabody/ddev-drupal-multisite-databases
description: "Simple DDEV addon which ensures a database is created for every multisite directory that exists in the codebase."
user: lpeabody
repo: ddev-drupal-multisite-databases
repo_id: 950744576
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2025-03-18
updated_at: 2025-03-18
stars: 0
---

# DDEV Drupal Multisite Databases

This is a DDEV addon that ensures a database exists for every valid multisite
directory (i.e. excluding `default` and `settings`).

- The database name is the same as the directory name.
- The database user and password are both `db`.

## Installation

Run the following command to install the addon:

```bash
ddev add-on get lpeabody/ddev-drupal-multisite-databases
```

After the addon is installed, run `ddev restart` to apply the changes.
