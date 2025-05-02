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
updated_at: 2025-05-01
stars: 0
---

[![tests](https://github.com/MurzNN/ddev-pgadmin/actions/workflows/tests.yml/badge.svg)](https://github.com/MurzNN/ddev-pgadmin/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# DDEV pgAdmin add-on<!-- omit in toc -->

This is a [DDEV](https://github.com/ddev/ddev/) add-on provides a pgAdmin service for PostgreSQL databases.

## Installation:

```
ddev add-on get MurzNN/ddev-pgadmin
ddev restart
```

You can run pgAdmin easily with the command after installing this add-on:

```
ddev pgadmin
```

Also, it will be available on the url `https://pgadmin.yourprojectname.ddev.site`. If it asks for the database password, enter `db` there.

> [!TIP]
> For Gitpod: The `ddev pgadmin` command can open a blank page in preview mode, open the link in a new browser tab/window to make it work.

# Tips

- PgAdmin stores the initial configuration from the `server.json` in the internal database on the first start. So, if you make changes in the `servers.json` - delete the `pgadmin-data` Docker volume to apply changes.

**Contributed and maintained by [@MurzNN](https://github.com/MurzNN)**
