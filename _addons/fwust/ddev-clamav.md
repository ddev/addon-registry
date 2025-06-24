---
title: fwust/ddev-clamav
github_url: https://github.com/fwust/ddev-clamav
description: "ClamAV integration for DDEV"
user: fwust
repo: ddev-clamav
repo_id: 710824160
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-10-27
updated_at: 2025-04-25
workflow_status: success
stars: 1
---

# ddev-clamav <!-- omit in toc -->

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/fwust/ddev-clamav/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/fwust/ddev-clamav/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/fwust/ddev-clamav)](https://github.com/fwust/ddev-clamav/commits)
[![release](https://img.shields.io/github/v/release/fwust/ddev-clamav)](https://github.com/fwust/ddev-clamav/releases/latest)

- [Introduction](#introduction)
- [Getting Started](#getting-started)

## Introduction

[ClamAV®](https://www.clamav.net/) is an open-source antivirus engine for detecting trojans, viruses, malware & other malicious threats.

This add-on allows you to run [ClamAV](https://www.clamav.net/) through the DDEV web service.
The docker image used : [tiredofit/clamav](https://github.com/tiredofit/docker-clamav).

## Getting Started

Install the DDEV ClamAV:

For DDEV v1.23.5 or above run

```shell
ddev add-on get fwust/ddev-clamav
```

For earlier versions of DDEV run

```shell
ddev get fwust/ddev-clamav
```

Then restart your project.

```shell
ddev restart
```
