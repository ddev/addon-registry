---
title: "opixido/ddev-addon-phploy"
github_url: "https://github.com/opixido/ddev-addon-phploy"
description: ""
user: "opixido"
repo: "ddev-addon-phploy"
repo_id: 967426305
default_branch: "master"
tag_name: "0.7"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2025-04-16"
updated_at: "2026-10-02"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/opixido/ddev-addon-phploy/actions/workflows/tests.yml/badge.svg)](https://github.com/opixido/ddev-addon-phploy/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/opixido/ddev-addon-phploy)](https://github.com/opixido/ddev-addon-phploy/commits)
[![release](https://img.shields.io/github/v/release/opixido/ddev-addon-phploy)](https://github.com/opixido/ddev-addon-phploy/releases/latest)

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
```

## 🔧 Actions

If you are using phploy actions like `pre-deploy`, you need to be aware of the context of thoses actions, because they will run inside the phploy container.

So `curl, git, cp, mv, ...` and other basic linux commands will be okay, but if you need actions specific to the `web` container you need to run them inside the web container, so you will have to prefix them with `docker exec ddev-[PROJECT_NAME]-web`

For example to start a `sass` command if `sass` is installed in the web container : 
```ini
pre-deploy[] = "docker exec ddev-[YOUR-PROJECT-NAME]-web sass --style=compressed  src/all.scss:css/all.min.css"
```
