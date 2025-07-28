---
title: Metadrop/ddev-selenium
github_url: https://github.com/Metadrop/ddev-selenium
description: "Provide selenium setup with a default chrome node, with possibilities to add more nodes of several browsers."
user: Metadrop
repo: ddev-selenium
repo_id: 719059276
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-11-15
updated_at: 2025-05-26
workflow_status: disabled
stars: 1
---

[![tests](https://github.com/Metadrop/ddev-selenium/actions/workflows/tests.yml/badge.svg)](https://github.com/Metadrop/ddev-selenium/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)
![GitHub Release](https://img.shields.io/github/v/release/Metadrop/ddev-selenium)

# DDEV selenium

Provide selenium setup with a default chrome node, with possibilities to add more nodes of several browsers.

The main purpose of this addon is allowing developers to have a base setup of selenium
that allows to be extendeds with more selenium nodes like chrome, firefox, or edge. Having
more nodes attached to the same hub not only allows having several browsers, but load balancing
when there are parallel tests being run on the same hub.

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get metadrop/ddev-selenium
```

For earlier versions of DDEV run

```sh
ddev get metadrop/ddev-selenium
```

## Architecture

The addon contains:

- Selenium hub
- Selenium node for chrome

It is possible to add attach additional nodes to the selenium by adding more services to the ddev setup.

### How to add more selenium nodes

To add more nodes, create a custom docker-compose yaml with the node services. Those nodes:

- Must be based on [docker selenium nodes](https://github.com/SeleniumHQ/docker-selenium/#dev-and-beta-on-the-grid).
- Must be attached to the hub container through this configuration:
```yaml
      - SE_EVENT_BUS_HOST=hub
      - SE_EVENT_BUS_PUBLISH_PORT=4442
      - SE_EVENT_BUS_SUBSCRIBE_PORT=4443
```

Example:

```yaml
  chrome:
    image: selenium/node-chrome:4.8.0
    container_name: ddev-${DDEV_SITENAME}-chrome
    shm_size: 256M
    depends_on:
      - hub
    environment:
      - SE_EVENT_BUS_HOST=hub
      - SE_EVENT_BUS_PUBLISH_PORT=4442
      - SE_EVENT_BUS_SUBSCRIBE_PORT=4443
    labels:
      com.ddev.site-name: ${DDEV_SITENAME}
      com.ddev.approot: ${DDEV_APPROOT}
```
