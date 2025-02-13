---
title: robertoperuzzo/ddev-unstructured
github_url: https://github.com/robertoperuzzo/ddev-unstructured
description: "Unstructured self-hosted service for DDEV"
user: robertoperuzzo
repo: ddev-unstructured
type: contrib
created_at: 2025-01-25
updated_at: 2025-01-27
stars: 0
---

[![tests](https://github.com/robertoperuzzo/ddev-unstructured/actions/workflows/tests.yml/badge.svg)](https://github.com/robertoperuzzo/ddev-unstructured/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

# ddev-unstructured <!-- omit in toc -->

* [What is ddev-unstructured?](#what-is-ddev-add-on-template)
* [Getting started](#getting-started)
* [How to debug in Github Actions](https://github.com/robertoperuzzo/ddev-unstructured/blob/main/./README_DEBUG.md)

## What is ddev-unstructured?

This add-on quickly enable the [Unstructured.io](https://unstructured.io/) service into a [DDEV](https://ddev.readthedocs.io) project.

Just a heads-up—there isn't an Unstructured API image build for Silicon Macs (arm64) yet. 
Instead, we’re using the build for Linux/amd64. For more details, check out issue [#480](https://github.com/Unstructured-IO/unstructured-api/issues/480).

## Getting started

Install the add-on with `ddev add-on get robertoperuzzo/ddev-unstructured`.

**Contributed and maintained by `@robertoperuzzo`** based on the original work of `@roromedia` in the [#ai](https://drupal.slack.com/archives/CDL2YPBNX/p1737148106043569?thread_ts=1737114857.811289&cid=CDL2YPBNX) Slack channel.

