---
title: ckng/ddev-directus-postgres
github_url: https://github.com/ckng/ddev-directus-postgres
description: "Directus with PostgreSQL (and Redis) addon for DDEV"
user: ckng
repo: ddev-directus-postgres
repo_id: 945850419
ddev_version_constraint: ""
dependencies: ["redis"]
type: contrib
created_at: 2025-03-10
updated_at: 2025-03-11
stars: 0
---

[![Version](https://img.shields.io/github/v/release/ckng/ddev-directus-postgres)](https://github.com/ckng/ddev-directus-postgres/releases) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg) [![tests](https://github.com/ckng/ddev-directus-postgres/actions/workflows/tests.yml/badge.svg)](https://github.com/ckng/ddev-directus-postgres/actions/workflows/tests.yml)

## What is this?

This repository allows you to quickly install Directus with PostgreSQL (and Redis) into a [DDEV](https://ddev.readthedocs.io) project using `ddev get ckng/ddev-directus-postgres`.

If you're looking for the SQLite version like Directus Self-Hosted version, check out [ddev-directus](https://github.com/MelaineGerard/ddev-directus).

## What is Directus

[Directus](https://directus.io/) is an Open Source Headless CMS for managing SQL database content. It's a nice tool to use as a backend for all your frontend applications.

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get ddev/ddev-redis && ddev add-on get ckng/ddev-directus-postgres && ddev restart
```

For earlier versions of DDEV run

```sh
ddev get ddev/ddev-redis && ddev get ckng/ddev-directus-postgres && ddev restart
```

## Explanation

This Directus recipe for [DDEV](https://ddev.readthedocs.io) installs a [`.ddev/docker-compose.directus-postgres.yaml`](https://github.com/ckng/ddev-directus-postgres/blob/master/docker-compose.directus-postgres.yaml) using the `directus` Docker image.

## Interacting with Directus

* The Directus instance will listen on HTTP port 8054, and HTTPS port 8055 (the Directus default).
* Configure your application to connect to Directus on the host:port `directus:8055`.
* To reach the Directus admin interface, open [https://your-project.ddev.site:8055/admin](https://your-project.ddev.site:8055/admin) in your browser. (You need to replace `your-project` with your actual project name.)
* Credentials are `admin@ddev.site` and `d1r3ctu5`.

  Your can customize the default credentials in the [docker-compose.directus-postgres.yaml](https://github.com/ckng/ddev-directus-postgres/blob/master/docker-compose.directus-postgres.yaml) by editing the `ADMIN_EMAIL` and `ADMIN_PASSWORD` variables. Or you can remove ADMIN_PASSWORD for it to auto generate.

## How to use this project in a fresh project

```bash
# 1. Create an empty project.
mkdir your-project
cd your-project
# 2. Create a simple Hello, World in this folder.
echo "<?php echo 'Hello, World!'; ?>" > index.php
# 3. Init ddev in the project.
ddev config --project-type php --database postgres:13 --nodejs-version 22
# 4. Add Redis dependency in the project.
ddev add-on get ddev/ddev-redis
# 5. Add Directus in the project.
ddev add-on get ckng/ddev-directus-postgres # or ddev get ckng/ddev-directus-postgres for older versions of DDEV
# 6. Start the project.
ddev start
# Directus should now be started on port 8055 (HTTPS) and port 8054 (HTTP) of your project.
```

Example url :
[https://your-project.ddev.site:8055](https://your-project.ddev.site:8055)
[http://your-project.ddev.site:8054](http://your-project.ddev.site:8054)

## Adding Directus extensions
To add Directus extension locally.

```bash
# 1. Make sure in the extensions folder.
cd your-project/directus/extensions
# 2. Add an extenstion, e.g.
npm require directus-extension-editorjs
# 3. Reload Directus Data Studio app in the browser.
```

## Adding DDev CA certs.
For nodejs to accept DDev mkcert certificates automatically, set the NODE_EXTRA_CA_CERTS per https://github.com/FiloSottile/mkcert#using-the-root-with-nodejs:

```bash
export NODE_EXTRA_CA_CERTS="$(mkcert -CAROOT)/rootCA.pem"
```

Recommend to add to your shell environment.

## Other useful add-ons
* **Mailpit**, built-in mail catcher.
   * Mailpit: https://directus-autoblog.ddev.site:8026
```sh
ddev mailpit
```

* **Adminer**, remember to select PostgreSQL as the database before trying to login.
```sh
ddev get ddev/ddev-adminer && ddev restart
```

## Maintained By

- [CK Ng](https://github.com/ckng)
