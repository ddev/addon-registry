---
title: hanoii/ddev-readme
github_url: https://github.com/hanoii/ddev-readme
description: "An opinionated README formatter"
user: hanoii
repo: ddev-readme
categories:
  - community
created_at: 2024-03-08
updated_at: 2024-11-25
stars: 0
---

[![tests](https://github.com/hanoii/ddev-readme/actions/workflows/tests.yml/badge.svg)](https://github.com/hanoii/ddev-readme/actions/workflows/tests.yml)
![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

<!-- toc -->

- [What is ddev-readme?](#what-is-ddev-readme)
- [TOC](#toc)
- [Install the dev version](#install-the-dev-version)

<!-- tocstop -->

## What is ddev-readme?

This is an opinionated approach to keeping README.md a bit standarized
formatter.

It uses [prettier](https://prettier.io/) and
[markdown-toc](https://www.npmjs.com/package/markdown-toc?activeTab=readme).

Once installed and `ddev restart`, it will start up a process watching for
changes on your README.md. If you wish to disable this you can edit and remove
`#ddev-generated` from `config.readme.yaml` or add
`DDEV_README_WATCH_DISABLED=true` as en environment variable to your project.

## TOC

For you toc generation to work automatically, you have to add the following
somewhere on your README.md:

**&lt;!-- toc --&gt;**

**&lt;!-- tocstop --&gt;**

## Install the dev version

You can always install the latest code:

For DDEV v1.23.5 or above run

```sh
ddev add-on get https://github.com/hanoii/ddev-readme/tarball/main
```

For earlier versions of DDEV run

```sh
ddev get https://github.com/hanoii/ddev-readme/tarball/main
```

**Contributed and maintained by [@hanoii](https://github.com/hanoii)**

