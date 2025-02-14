---
title: kevinquillen/ddev-typesense
github_url: https://github.com/kevinquillen/ddev-typesense
description: "Typesense is a modern, privacy-friendly, open source search engine meticulously engineered for performance & ease-of-use. This is an add on for DDEV so you can run and develop integrations for Typesense locally."
user: kevinquillen
repo: ddev-typesense
type: contrib
created_at: 2023-11-06
updated_at: 2025-02-07
stars: 3
---

[![tests](https://github.com/kevinquillen/ddev-typesense/actions/workflows/tests.yml/badge.svg)](https://github.com/kevinquillen/ddev-typesense/actions/workflows/tests.yml)

## Installation

Uses the current stable release of the Typesense Docker image.

With DDEV installed, run this command:

`ddev get kevinquillen/ddev-typesense`

## Configuration

The Typesense container is reached at hostname: "typesense", port: 8108. Outside of the container, you can visit `127.0.0.1:8109/health` in your browser to verify health status.

The default API key for Typesense is `ddev`. You can provide your own by adding to `.ddev/.env` in your project, and adding the `TYPESENSE_API_KEY` variable:

`TYPESENSE_API_KEY=my_api_key_value`

## Admin Dashboard

This DDEV addon also includes the admin dashboard by bfritscher:

https://github.com/bfritscher/typesense-dashboard

The admin dashboard is useful to navigate your collections and schema and debug your search.

You can access the admin dashboard by navigating to this URL in your browser:

`http://typesense.(DDEV_HOSTNAME):8109/#/login`

To login, provide the configured API key, `127.0.0.1` as the hostname, and `8108` as the port. Leave the path blank.

# Drupal and Search API

If you are using Drupal, you can use Search API and the Search API Typesense modules to connect to the running Typesense instance.

**Originally Contributed by [kevinquillen](https://github.com/kevinquillen)**
