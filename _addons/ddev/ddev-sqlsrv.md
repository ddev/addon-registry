---
title: ddev/ddev-sqlsrv
github_url: https://github.com/ddev/ddev-sqlsrv
description: "MS SQL server add-on for DDEV"
user: ddev
repo: ddev-sqlsrv
repo_id: 618290272
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: official
created_at: 2023-03-24
updated_at: 2025-04-10
stars: 4
---

[![tests](https://github.com/ddev/ddev-sqlsrv/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/ddev/ddev-sqlsrv/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/ddev/ddev-sqlsrv)](https://github.com/ddev/ddev-sqlsrv/commits)
[![release](https://img.shields.io/github/v/release/ddev/ddev-sqlsrv)](https://github.com/ddev/ddev-sqlsrv/releases/latest)

# DDEV SQLSRV

## Overview

This add-on quickly installs the MS SQL server into a DDEV project.
It is based on the [mcr.microsoft.com/mssql/server](https://hub.docker.com/_/microsoft-mssql-server) image.

You have to keep in mind that the [mssql-docker](https://github.com/microsoft/mssql-docker) does not natively work on M1 (arm64).
Some workarounds are described in the following threads:
* [Does not work on Mac M1](https://github.com/microsoft/mssql-docker/issues/734)
* [MSSQL container for aarch64 (arm64) for better performance](https://github.com/microsoft/mssql-docker/issues/802)

## Installation

> [!WARNING]
> Due to lack of upstream support, this add-on can only be used with amd64 machines, and is not usable on arm64 machines like Apple Silicon computers.
> (You can try installing it on Rosetta using `DDEV_IGNORE_ARCH_CHECK=true ddev add-on get ddev/ddev-sqlsrv`)

```bash
ddev add-on get ddev/ddev-sqlsrv
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev sqlcmd` | For Transact-SQL statements, system procedures, and script files |
| `ddev drupal-regex` | For compatibility with Drupal version 9 or higher |

### Drupal Notice

Drupal CMS needs the database function installed that is mimicking the Regex function as Drupal requires. As a one-time setup for Drupal, install the database function by running the following command from your project's directory:

```bash
ddev drupal-regex
```

This script also changes the setting for the following database variables:

* `show advanced options` will be set to 1
* `clr strict security` will be set to 0
* `clr enable` will be set to 1

Drupal also required the [`sqlsrv` module](https://www.drupal.org/project/sqlsrv) to be installed as it is provides the database driver for SQL Server. The module can be installed with composer with the following command:

```bash
ddev composer require drupal/sqlsrv
```

**There is an open issue for Drupal 9.4+ installations. Until merged, you need to apply [patch #4](https://www.drupal.org/project/sqlsrv/issues/3291199#comment-14576456), see [Call to a member function fetchField() on null
](https://www.drupal.org/project/sqlsrv/issues/3291199)**

### Manually enabling MySQL/MariaDB/PostgreSQL

**This addons disables the default database by automatically adding `omit_containers: [db]` in the `config.sqlsrv.yaml`**

If your project needs to use both MariaDB and MS SQL Server databases, you have to remove `#ddev-generated` and
`omit_containers: [db]` from `config.sqlsrv.yaml`.

See [Config Options](https://ddev.readthedocs.io/en/stable/users/configuration/config/) for additional notes.

## Advanced Customization

Use a different port:

```bash
ddev dotenv set .ddev/.env.sqlsrv --mssql-external-port=1434
ddev add-on get ddev/ddev-sqlsrv
ddev restart
```

Make sure to commit the `.ddev/.env.sqlsrv` file to version control.

Or change the password:

```bash
ddev dotenv set .ddev/.env.sqlsrv --mssql-sa-password='myNewPassword'
ddev add-on get ddev/ddev-sqlsrv
ddev restart
```

Make sure to commit the `.ddev/.env.sqlsrv` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `MSSQL_DOCKER_IMAGE` | `--mssql-docker-image` | `mcr.microsoft.com/mssql/server:2022-CU17-ubuntu-22.04` |
| `MSSQL_EXTERNAL_PORT` | `--mssql-external-port` | `1433` |
| `MSSQL_SA_PASSWORD` | `--mssql-sa-password` | `Password12!` |
| `MSSQL_PID` | `--mssql-pid` | `Evaluation` |
| `MSSQL_DB_NAME` | `--mssql-db-name` | `master` |
| `MSSQL_HOST` | `--mssql-host` | `sqlsrv` |
| `MSSQL_COLLATION` | `--mssql-collation` | `LATIN1_GENERAL_100_CI_AS_SC_UTF8` |

## Links with useful information

* [SQL Server docker hub](https://hub.docker.com/_/microsoft-mssql-server)
* [Installing the ODBC driver for SQL Server](https://docs.microsoft.com/en-us/sql/connect/odbc/linux-mac/installing-the-microsoft-odbc-driver-for-sql-server)
* [Installing the ODBC driver for SQL Server Tutorial](https://docs.microsoft.com/en-us/sql/connect/php/installation-tutorial-linux-mac)
* [Installation tutorial for MS drivers for PHP](https://docs.microsoft.com/en-us/sql/connect/php/installation-tutorial-linux-mac)
* [The SQLCMD utility](https://docs.microsoft.com/en-us/sql/tools/sqlcmd-utility)
* [The SQL Server on Linux](https://docs.microsoft.com/en-us/sql/linux/sql-server-linux-overview)
* [The password policy](https://docs.microsoft.com/en-us/sql/relational-databases/security/password-policy)
* [The SQL Server environment variables](https://docs.microsoft.com/en-us/sql/linux/sql-server-linux-configure-environment-variables)
* [Beakerboy's Drupal Regex database function](https://github.com/Beakerboy/drupal-sqlsrv-regex)
* [Drupal's module for the SQL Server](https://www.drupal.org/project/sqlsrv)
* [Github MS drivers for PHP](https://github.com/microsoft/msphpsql)

Note that more advanced techniques are discussed in [DDEV docs](https://ddev.readthedocs.io/en/stable/users/extend/additional-services/).

## Credits

**Contributed and maintained by [@robertoperuzzo](https://github.com/robertoperuzzo) based on the original [ddev-contrib recipe](https://github.com/ddev/ddev-contrib/tree/master/docker-compose-services/sqlsrv) by [drupal-daffie](https://github.com/drupal-daffie)**
