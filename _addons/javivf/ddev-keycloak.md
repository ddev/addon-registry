---
title: javivf/ddev-keycloak
github_url: https://github.com/javivf/ddev-keycloak
description: "Keycloak service for DDEV"
user: javivf
repo: ddev-keycloak
repo_id: 682633589
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-08-24
updated_at: 2025-02-12
stars: 0
---

[![tests](https://github.com/javivf/ddev-keycloak/actions/workflows/tests.yml/badge.svg)](https://github.com/javivf/ddev-keycloak/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## What is ddev-keycloak?

This repository allows you to quickly install [Keycloak](https://keycloak.org) into a [DDEV](https://ddev.readthedocs.io) project using just `ddev get javivf/ddev-keycloak`.

## Installation

1. `ddev get javivf/ddev-keycloak`
2. `ddev restart`

## Explanation

This Keycloak recipe for [DDEV](https://ddev.readthedocs.io) installs a [`.ddev/docker-compose.keycloak.yaml`](https://github.com/javivf/ddev-keycloak/blob/main/docker-compose.keycloak.yaml) using the [`keycloak/keycloak`](https://quay.io/repository/keycloak/keycloak) Docker image.

## Interacting with Keycloak

* The Keycloak instance will listen on TCP port 8080.
* Configure your application to access Keycloak on the host:port `keycloak:8080`.


## Resources

1. [Official documentation](https://keycloak.org)
2. [Official repository](https://github.com/keycloak/keycloak)
