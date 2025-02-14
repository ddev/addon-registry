---
title: iamntz/ddev-rename-project
github_url: https://github.com/iamntz/ddev-rename-project
description: "ddev addon to rename a project"
user: iamntz
repo: ddev-rename-project
repo_id: 663521436
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-07-07
updated_at: 2024-12-06
stars: 1
---

## Installation:

For DDEV v1.23.5 or above run

```sh
ddev add-on get iamntz/ddev-rename-project
```

For earlier versions of DDEV run

```sh
ddev get iamntz/ddev-rename-project
```

## Usage:

```sh
ddev rename-project my-new-project
```

## Notes:

The command will create a new DB snapshot, it will remove your current project, will create a new project and finally it will import the newly created DB snapshot.
