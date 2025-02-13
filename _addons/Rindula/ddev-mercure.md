---
title: Rindula/ddev-mercure
github_url: https://github.com/Rindula/ddev-mercure
description: "A Mercure Server based on the official docker image from https://mercure.rocks"
user: Rindula
repo: ddev-mercure
type: contrib
created_at: 2022-06-28
updated_at: 2024-10-15
stars: 0
---

[![tests](https://github.com/Rindula/ddev-mercure/actions/workflows/tests.yml/badge.svg)](https://github.com/Rindula/ddev-mercure/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

# Mercure extension for ddev

This extension is based on [dunglas/mercure](https://github.com/dunglas/mercure)s implementation.

To install it, run `ddev get Rindula/ddev-mercure`.

The mercurehost will be available under `mercure.${DDEV_HOSTNAME}`.
So for `example.ddev.site` the url will be `mercure.example.ddev.site`.
From inside the web container, you can reach mercure under `http://mercure:3000`.

