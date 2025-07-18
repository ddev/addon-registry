---
title: rfay/ddev-php-patch-build
github_url: https://github.com/rfay/ddev-php-patch-build
description: "Build a patch version of PHP for use with DDEV"
user: rfay
repo: ddev-php-patch-build
repo_id: 692610647
ddev_version_constraint: ">= v1.23.5"
dependencies: []
type: contrib
created_at: 2023-09-17
updated_at: 2025-06-14
workflow_status: failure
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/rfay/ddev-php-patch-build/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/rfay/ddev-php-patch-build/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/rfay/ddev-php-patch-build)](https://github.com/rfay/ddev-php-patch-build/commits)
[![release](https://img.shields.io/github/v/release/rfay/ddev-php-patch-build)](https://github.com/rfay/ddev-php-patch-build/releases/latest)

# Experimental Add-on to use PHP patch versions: ddev-php-patch-build

[DDEV](https://ddev.com) users periodically ask if they can use a specific PHP patch version with their projects, either to test/compare for a particular bug or to match a production version exactly.

This **experimental** add-on tries to address that need, although there are a number of caveats:

1. The build done using this process uses the technique from [static-php-cli](https://github.com/crazywhalecc/static-php-cli) and full details are available there. It supports all the versions supported in the upstream repository, currently PHP 7.3-8.4.
2. The build does not match the build done using the official DDEV php packages (which actually come from [deb.sury.org](https://deb.sury.org/).
3. It currently does not provide xdebug, see [extension support](https://static-php.dev/en/guide/extension-notes.html).
4. The resultant PHP binaries built here do not have the exact same extensions as the official DDEV PHP binaries.
5. Building the binaries happens in the build phase of `ddev start`, and it takes a long time on the first `ddev start` or whenever you change versions. My tests on Gitpod, with a great internet connection, took about 8-9 minutes. It can be really annoying, and a better way to build would be an improvement.
6. If you want to see the build process as it proceeds, you can use `ddev debug refresh` or `DDEV_VERBOSE=true ddev start`.
7. When you use this add-on, your DDEV `php_version` setting is ignored.
8. To see what's going on during the build, use `ddev debug rebuild -s web --cache`
9. Your mileage may vary.

* The [static-php-ci](https://github.com/crazywhalecc/static-php-cli) repository provides a relatively easy way to build static PHP binaries (the CLI and `php-fpm`) which can be used to replace the ones installed in `ddev-webserver`.
* The provided Dockerfile.php-patch-build does the building.
* The provided `config.php-patch-build` provides post-start hooks that install the built binaries.

## Installation

```bash
ddev add-on get rfay/ddev-php-patch-build
ddev restart
```

You can choose a different PHP version, the command below creates a `.ddev/.env.php-patch-build` file that you can commit:

1. `ddev dotenv set .ddev/.env.php-patch-build --static-php-version=8.0.10`
2. `ddev add-on get rfay/ddev-php-patch-build`
3. `ddev restart`

## Components of the add-on

* `config.php-yaml`
* `Dockerfile.php-patch-build`
* A test suite in [test.bats](https://github.com/rfay/ddev-php-patch-build/blob/main/tests/test.bats) that makes sure the service continues to work as expected.
* [Github actions setup](https://github.com/rfay/ddev-php-patch-build/blob/main/.github/workflows/tests.yml) so that the tests run automatically when you push to the repository.


**Contributed and maintained by [@rfay](https://github.com/rfay)**
