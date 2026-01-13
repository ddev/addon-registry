---
title: lussoluca/ddev-typesense
github_url: https://github.com/lussoluca/ddev-typesense
description: "Typesense is a modern, privacy-friendly, open source search engine meticulously engineered for performance & ease-of-use. This is an add on for DDEV so you can run and develop integrations for Typesense locally. "
user: lussoluca
repo: ddev-typesense
repo_id: 947892943
default_branch: main
tag_name: v0.0.2
ddev_version_constraint: ">= v1.24.2"
dependencies: []
type: contrib
created_at: 2025-03-13
updated_at: 2025-06-03
workflow_status: disabled
stars: 0
---

[![tests](https://github.com/lussoluca/ddev-typesense/actions/workflows/tests.yml/badge.svg)](https://github.com/lussoluca/ddev-typesense/actions/workflows/tests.yml)

## Installation

Uses the current stable release of the Typesense Docker image.

With DDEV installed, run this command:

`ddev get lussoluca/ddev-typesense`

## Configuration

The Typesense container is reached at hostname: `https://(DDEV_HOSTNAME)`, port: 8109.

The default API key for Typesense is `ddev`. You can provide your own by adding to `.ddev/.env` in your project, and adding the `TYPESENSE_API_KEY` variable:

`TYPESENSE_API_KEY=my_api_key_value`

## Admin Dashboard

This DDEV addon also includes the admin dashboard by bfritscher:

https://github.com/bfritscher/typesense-dashboard

The admin dashboard is useful to navigate your collections and schema and debug your search.

You can access the admin dashboard by navigating to this URL in your browser:

`https://(DDEV_HOSTNAME):8111`

# Drupal and Search API

If you are using Drupal, you can use [Search API](https://www.drupal.org/project/search_api) and the [Search API Typesense](https://www.drupal.org/project/search_api_typesense) modules to connect to the running Typesense instance.

> Originally Contributed by [kevinquillen](https://github.com/kevinquillen), then adapted to work with latest Typesense server over HTTPS.
