---
title: opixido/ddev-addon-phploy
github_url: https://github.com/opixido/ddev-addon-phploy
description: ""
user: opixido
repo: ddev-addon-phploy
repo_id: 967426305
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2025-04-16
updated_at: 2025-04-16
stars: 0
---

# DDEV Add-on: phploy

This DDEV add-on adds a `ddev phploy` command to handle deployments through [phploy](https://github.com/banago/phploy), it includes its own docker container to avoid PHP versions issues.

## 🔧 Installation

```bash
ddev add-on get https://github.com/opixido/ddev-addon-phploy/tarball/master
ddev restart
```

## 🔧 Usage

Just like phploy, but prefixed with `ddev`

```bash
ddev phploy -s server -l
...
