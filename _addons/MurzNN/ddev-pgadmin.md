---
title: MurzNN/ddev-pgadmin
github_url: https://github.com/MurzNN/ddev-pgadmin
description: "pgAdmin Add-on For DDEV: PostgreSQL database management web interface"
user: MurzNN
repo: ddev-pgadmin
repo_id: 834866824
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-07-28
updated_at: 2024-07-29
stars: 0
---

[![tests](https://github.com/MurzNN/ddev-pgadmin/actions/workflows/tests.yml/badge.svg)](https://github.com/MurzNN/ddev-pgadmin/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# ddev-pgadmin <!-- omit in toc -->

* [What is ddev-pgadmin?](#what-is-ddev-pgadmin)

## What is ddev-pgadmin?

This add-on provides a pgAdmin service for [DDEV](https://github.com/ddev/ddev/).

Installation:

```
ddev get MurzNN/ddev-pgadmin
ddev restart
```

You can run pgAdmin easily with the command after installing this add-on:

```
ddev pgadmin
```

Also, it will be available on the url `https://pgadmin.yourprojectname.ddev.site`. If it asks for the database password, enter `db` there.

> [!TIP]
> For Gitpod: The `ddev pgadmin` command can open a blank page in preview mode, open the link in a new browser tab/window to make it work.

**Contributed and maintained by [@MurzNN](https://github.com/MurzNN)**
