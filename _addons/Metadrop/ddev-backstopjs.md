---
title: Metadrop/ddev-backstopjs
github_url: https://github.com/Metadrop/ddev-backstopjs
description: "ddev addon to provide a backstopjs container for visual regression testing in Metadrop Aljibe"
user: Metadrop
repo: ddev-backstopjs
repo_id: 709536102
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-10-24
updated_at: 2025-09-17
workflow_status: disabled
stars: 8
---

[![tests](https://github.com/Metadrop/ddev-backstopjs/actions/workflows/tests.yml/badge.svg)](https://github.com/Metadrop/ddev-backstopjs/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

* [What is DDEV Backstopjs Add-on?](#what-is-ddev-backstopjs-add-on)
* [Getting started](#getting-started)
* [Using backstopjs](#using-backstopjs)
  * [Configuration](#configuration)
  * [Run tests](#run-tests)
  * [View test results](#view-test-results)
* [Changes to the original docker image](#changes-to-the-original-docker-image)
* [Advanced](#advanced)
  * [How to add additional hostnames?](#how-to-add-additional-hostnames)
  * [Change backstop tests directory](#change-backstop-tests-directory)


## What is DDEV Backstopjs Add-on?

This is a ddev-addon for [backstop.js](https://github.com/garris/BackstopJS) based on [mmunz](https://github.com/mmunz/ddev-backstopjs) addon but optimized for Aljibe. Backstopjs is a visual regression testing tool.
Backstop is executed in a docker container based on the official [backstopjs docker image](https://hub.docker.com/r/backstopjs/backstopjs).

This addon just provides the basics to run backstopjs and some basic tests, but those need to be adapted to your needs.

## Getting started

Install this addon

For DDEV v1.23.5 or above run

```shell
ddev add-on get Metadrop/ddev-backstopjs
```

For earlier versions of DDEV run

```shell
ddev get Metadrop/ddev-backstopjs
```

After that you need to restart the ddev project:

```shell
ddev restart
```

**Note: If you haven't downloaded the backstopjs base image before, then it will be downloaded when ddev is restarted.
The backstopjs/backstopjs is about 2.6GB, so this may take some time.**


## Using backstopjs

### Configuration

By default, the backstop tests are expected in $DDEV_APPDIR/tests/backstopjs/<environemnt_folder>.

This add-on provide some tests inside "local" environment folder ($DDEV_APPDIR/tests/backstopjs/local). This can be taken as a base to add more tests or provide your own backstop.js or backstop.json configs there.

Hint: have a look at the example from mmunz [backstopjs-config](https://github.com/mmunz/backstopjs-config)

Alternatively you can create a simple backstop.json config with:

```shell
ddev backstop init
```

### Run tests

After the config was created it is time to run the tests.

Create reference screenshots:

```shell
ddev backstop <environment> reference
```

Create test images and compare to reference screenshots:

```shell
ddev backstop <environment> test
```

Where <environment> is the environment folder name, or 'local' if not especified.

If your config file is not 'backstop.json' you need to use the --config argument, e.g. --config=backstop.js

### View test results

The backstop commands 'backstop remote' and 'backstop openReport' do not work in this setup.

But there is a host command that will open the latest test report in your default browser:

```shell
ddev backstopjs-report <environment>
```

Alternatively open the generated HTML-Report with your browser, e.g.:

```shell
open tests/backstopjs/<environment>/backstop_data/_mytestproject_/html_report/index.html
```

### Commands aliases

Command ```ddev backstopjs``` can be called as:
 - ```ddev backstop```
 - ```ddev bkjs```

Command ```ddev backstopjs-report``` can be called as:
 - ```ddev backstop-report```
 - ```ddev bjsr```

## Changes to the original docker image

The backstopjs docker image is extended with some functions using a custom docker build, see [Dockerfile](https://github.com/Metadrop/ddev-backstopjs/blob/main/backstopBuild/Dockerfile)
and uses a custom [entrypoint](https://github.com/Metadrop/ddev-backstopjs/blob/main/backstopBuild/entrypoint.sh).

In the Dockerfile the following is added/changed:

- add the custom entrypoint.sh to the image
- delete the default 'node' user with uid 1000 and add current ddev user
- install the [minimist](https://www.npmjs.com/package/minimist) npm package globally. This is not needed by default
  but very handy to parse command line args for more complex custom backstopjs configs.

The entrypoint is responsible for:

- add /etc/hosts entries for all hosts configured in the ddev web container automatically
- add sleep command to keep the container running

## Advanced

### How to add additional hostnames?

If you want to test hosts not configured in the web container, you need to use external_links in
[docker-compose.backstop.yaml](https://github.com/Metadrop/ddev-backstopjs/blob/main/docker-compose.backstop.yaml).

See: [ddev FAQ: Can different projects communicate with each other?](https://ddev.readthedocs.io/en/latest/users/usage/faq/#features-requirements)

Do not forget to remove the #ddev-generated line!

### Change backstop tests directory
Per default the backstop directory containing backstop config etc. is expected in your project directory (besides the
.ddev folder) in the directory *tests/backstop*.

If you want to change that edit the file [docker-compose.backstop.yaml](https://github.com/Metadrop/ddev-backstopjs/blob/main/docker-compose.backstop.yaml) and
change the line in volumes to the path you want to use, move the files to the new directory and restart ddev.

Make sure to remove the #ddev-generated line from the file to prevent ddev from making changes to it.
