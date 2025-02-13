---
title: MrKistic/ddev-beanstalkd
github_url: https://github.com/MrKistic/ddev-beanstalkd
description: "Beanstalkd service add-on for DDEV (with arm64 support)"
user: MrKistic
repo: ddev-beanstalkd
type: contrib
created_at: 2025-01-24
updated_at: 2025-01-24
stars: 0
---

[![tests](https://github.com/mrkistic/ddev-beanstalkd/actions/workflows/tests.yml/badge.svg)](https://github.com/mrkistic/ddev-beanstalkd/actions/workflows/tests.yml)

## What is ddev-beanstalkd?

This add-on allows you to easily install [beanstalkd](https://beanstalkd.github.io/) (with arm64 support) into a [DDEV](https://ddev.readthedocs.io) project.

It configures and installs the `rayyounghong/beanstalkd` docker image to provide a beanstalkd service within your DDEV project.

## Installation

1. Install the addon

    For DDEV v1.23.5 or above run

    ```shell
    ddev add-on get mrkistic/ddev-beanstalkd
    ```

   For earlier versions of DDEV run

    ```shell
    ddev get mrkistic/ddev-beanstalkd
    ```

2. Restart DDEV to start the addon.

   ```shell
   ddev restart
   ```

## Using beanstalkd service

* beanstalkd listens on the default port 11300 on internal hostname `bankstalkd`.
* Configure your application to access beanstalkd on the host:port, i.e. `beanstalkd:11300`.
* To access beanstalkd directly, run `ddev ssh` to connect to the web container, then telnet to the beanstalkd container on port 11300, e.g. `telnet beanstalkd 11300`. 
* Alternatively, expose the listening port with a [custom Docker Compose](https://ddev.readthedocs.io/en/stable/users/extend/custom-compose-files/#docker-composeyaml-examples) file. Then connect directly to your docker container on port 11300, e.g. `telnet docker.lan 11300`. Then you can use a UI console like [aurora](https://github.com/xuri/aurora) to monitor your beanstalkd queue. e.g.:

```yaml
services:
  my_beanstalkd:
    ports:
      - "11300:11300"
```

