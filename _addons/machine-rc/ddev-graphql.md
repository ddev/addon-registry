---
title: machine-rc/ddev-graphql
github_url: https://github.com/machine-rc/ddev-graphql
description: "GraphQL service for ddev stack"
user: machine-rc
repo: ddev-graphql
repo_id: 812082669
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-06-07
updated_at: 2024-06-08
stars: 0
---

[![tests](https://github.com/ddev/ddev-addon-template/actions/workflows/tests.yml/badge.svg)](https://github.com/ddev/ddev-addon-template/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# GraphQL ddev Add-on

## Getting started

In the DDEV project directory launch the command:
```sh
ddev get machine-rc/graphql
```
Restart the DDEV instance:
```sh
ddev restart
```
Open the Grafana web interface via the url: https://your-project-name.ddev.site:4000/

## Customizations
Current structure allows support for multiple graphql services. 
To add a new service:
- create a new directory in the `graphql` directory with the service name
- copy the `Dockerfile` from the `graphql/book` directory
- adjust `docker-compose.graphql.yaml` to include the new service by duplicating the `graphql-book` service and changing the service name
  - adjust the `environment` section to expose the new service on a different port
