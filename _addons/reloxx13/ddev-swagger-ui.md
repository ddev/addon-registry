---
title: reloxx13/ddev-swagger-ui
github_url: https://github.com/reloxx13/ddev-swagger-ui
description: "Swagger UI add-on for DDEV"
user: reloxx13
repo: ddev-swagger-ui
repo_id: 693512516
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-09-19
updated_at: 2024-04-29
stars: 1
---

[![tests](https://github.com/reloxx13/ddev-swagger-ui/actions/workflows/tests.yml/badge.svg)](https://github.com/reloxx13/ddev-swagger-ui/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## What is ddev-swagger-ui?

This repository provides [Swagger UI](https://swagger.io/tools/swagger-ui/) as a [DDEV](https://ddev.com/) add-on to serve a web-ui for a [OpenAPI/Swagger](https://swagger.io/resources/open-api/) (api documentation) file.

This addon is based on the [Swagger UI Docker Image](https://hub.docker.com/r/swaggerapi/swagger-ui).

Advanced documentation can be found in the [Swagger UI Docker Docs](https://github.com/swagger-api/swagger-ui/blob/master/docs/usage/installation.md#docker).

DDEV docs for [Additional Service Configurations & Add-ons](https://ddev.readthedocs.io/en/latest/users/extend/additional-services/)


## Using ddev-swagger-ui within a project
### Dependencies

Make sure you have [DDEV v1.22.1+ installed](https://ddev.readthedocs.io/en/latest/users/install/ddev-installation/)

### Install
1. `cd` into your project directory
2. Run `ddev config` and answer the questions as appropriate
3. Run `ddev get reloxx13/ddev-swagger-ui` and answer the questions as appropriate
4. Run `ddev start` or `ddev restart`

### Upgrade

To upgrade your version of ddev-swagger-ui, repeat the `ddev get reloxx13/ddev-swagger-ui` to get the latest release. To see the installed version, `ddev get --installed`.

## Notes

## What does it do right now?

* Provides the following swagger-ui-docker-provided environmental variables:
   * PORT=8080
   * BASE_URL=/
   * SWAGGER_JSON=/var/www/html/swagger.json
   * SWAGGER_JSON_URL=
