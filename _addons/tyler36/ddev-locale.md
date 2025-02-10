---
title: tyler36/ddev-locale
github_url: https://github.com/tyler36/ddev-locale
description: "Change a DDEV project's timezone and language. "
user: tyler36
repo: ddev-locale
categories:
  - community
created_at: 2024-06-06
updated_at: 2025-01-09
stars: 0
---

[![tests](https://github.com/tyler36/ddev-locale/actions/workflows/tests.yml/badge.svg)](https://github.com/tyler36/ddev-locale/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2026.svg)

# ddev-locale <!-- omit in toc -->

- [What is ddev-locale?](#what-is-ddev-locale)
- [Why ddev-locale?](#why-ddev-locale)
- [Getting started](#getting-started)
- [Configuration](#configuration)
  - [Change timezone](#change-timezone)
  - [Change locale](#change-locale)
- [Components of the repository](#components-of-the-repository)
- [Contributing](#contributing)

## What is ddev-locale?

DDEV-locale is design to setup the DDEV web container for a specific locale.
In additional, it serves as documentation for DDEV internationalization.

DDEV has 2 related locale settings:

- timezone.
- system locale.

This addon, by default, sets both of these to Japan (Japanese).
Why Japan? Japan has a single timezone and a single encoding locale. It is also where my team is based.

## Why ddev-locale?

DDEV allows encoding to be set globally (via `web-environment`).
Timezone, however, is a per-project setting.

As of writing, DDEV includes all language files which significantly inflates the web server container size.
If/when the language files are removed, this addon will be updated to install the language pack.

## Getting started

1. Install this addon.

    ```shell
    ddev addon get tyler36/ddev-locale
    ddev restart
    ```

## Configuration

To change either setting, you need to update `.ddev/config.locale.yaml`.

1. Remove `## ddev-generated` from the file. This will prevent DDEV from updating it.

### Change timezone

1. To change timezones, update the `timezone` line with the TZ identifier for your desired location.
    See [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)

### Change locale

1. To update the language, update the `web_environment` line.

This contains up to 3 parts:

- Language code (<https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes>).
- Territory code (<https://en.wikipedia.org/wiki/ISO_3166-1#Codes>). Separate from language with `_`.

A 3rd code-set value (separated with `.`) is optional.

For example: `LANG=en_AU.UTF-8`

- English language.
- Australian territory.
- `UTF-8` code-set.

## Components of the repository

- `config.locale.yaml`: the settings that are merging with the project's configuration.
- An [install.yaml](https://github.com/tyler36/ddev-locale/blob/main/install.yaml) file that describes how to install the service or other component.
- A test suite in [test.bats](tests/test.bats) that makes sure the service continues to work as expected.
- [Github actions setup](https://github.com/tyler36/ddev-locale/blob/main/.github/workflows/tests.yml) so that the tests run automatically when you push to the repository.

## Contributing

PRs are welcome, especially if they contain tests.

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**

