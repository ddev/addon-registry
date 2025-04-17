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
updated_at: 2025-04-17
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
```

## 🔧 Actions

If you are using phploy actions like `pre-deploy`, you need to be aware of the context of thoses actions, because they will run inside the phploy container.

So `curl, git, cp, mv, ...` and other basic linux commands will be okay, but if you need actions specific to the `web` container you need to run them inside the web container, so you will have to prefix them with `docker exec ddev-[PROJECT_NAME]-web`

For example to start a `sass` command if `sass` is installed in the web container : 
```ini
pre-deploy[] = "docker exec ddev-[YOUR-PROJECT-NAME]-web sass --style=compressed  src/all.scss:css/all.min.css"
```
