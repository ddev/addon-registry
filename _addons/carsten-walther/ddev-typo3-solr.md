---
title: carsten-walther/ddev-typo3-solr
github_url: https://github.com/carsten-walther/ddev-typo3-solr
description: "TYPO3 solr configuration"
user: carsten-walther
repo: ddev-typo3-solr
repo_id: 691593791
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-09-14
updated_at: 2024-11-08
stars: 0
---

[![tests](https://github.com/carsten-walther/ddev-typo3-solr/actions/workflows/tests.yml/badge.svg)](https://github.com/carsten-walther/ddev-typo3-solr/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)


# ddev-typo3-solr <!-- omit in toc -->

- [What is ddev-typo3-solr?](#what-is-ddev-typo3-solr)
- [Getting started](#getting-started)
- [Using cores](#using-cores)

## What is ddev-typo3-solr?

ddev-typo3-solr provides Solr (Cloud) using a single Solr node, which is sufficient
for local development requirements.

## Getting started

1. Install the addon

   Installation is very simple.

   For DDEV v1.23.5 or above run

   ```shell
   ddev add-on get carsten-walther/ddev-typo3-solr
   ```

   For earlier versions of DDEV run

   ```shell
   ddev get carsten-walther/ddev-typo3-solr
   ```

    This will install the php package `apache-solr-for-typo3/solr`, create the `.ddev/solr` project folder and coping all the available cores from the php package to the solr project folder.

3. Restart DDEV to start the addon.

   ```shell
   ddev restart
   ```

Once up and running, access Solr's UI within your browser by opening
`http://<projectname>.ddev.site:8983`. For example, if the project is named
"myproject" the hostname will be `http://myproject.ddev.site:8983`.

To access the Solr container from DDEV's web container, use  `http://solr:8983`.

## Using cores

Set your desired language core to the language in your sites/XXX/config.yaml.

```yaml
base: 'https://%env(DDEV_SITENAME)%.%env(DDEV_TLD)%/'
...
languages:
  -
    title: English
    enabled: true
    languageId: 0
    base: /
    locale: en_US.UTF-8
    navigationTitle: English
    flag: us
    hreflang: 'en-US'
    websiteTitle: 'TYPO3 Website'
    solr_core_read: core_en
  -
    title: Deutsch
    enabled: true
    languageId: 1
    base: /
    locale: de_DE.UTF-8
    navigationTitle: Deutsch
    flag: de
    hreflang: 'de-DE'
    websiteTitle: 'TYPO3 Website'
    solr_core_read: core_de
  -
    ...
rootPageId: 1
...
solr_enabled_read: true
solr_host_read: '%env(DDEV_SITENAME)%.%env(DDEV_TLD)%'
solr_path_read: /
solr_port_read: '8983'
solr_scheme_read: http
solr_use_write_connection: false
...
websiteTitle: 'TYPO3 Website'
```

You are able to use the DDEV constants in yout local environment. Keep in mind to replace them in production to the coorect one.

**Maintainer**
- [@carsten-walther](https://github.com/carsten-walther)

**Contributers**
- [@erik-konrad](https://github.com/erik-konrad)
