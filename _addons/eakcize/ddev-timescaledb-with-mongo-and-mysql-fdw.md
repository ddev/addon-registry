---
title: eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw
github_url: https://github.com/eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw
description: "DDev plugin for timescaleDb (PostregeSQL plugin) which is set to work with mongo and mysql trough foreign data wrappers for psql."
user: eakcize
repo: ddev-timescaledb-with-mongo-and-mysql-fdw
repo_id: 673722538
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-08-02
updated_at: 2025-03-31
workflow_status: disabled
stars: 0
---

[![tests](https://github.com/ddev/ddev-timescale-with-mongo-and-mysql-fdw/actions/workflows/tests.yml/badge.svg)](https://github.com/ddev/ddev-timescale-with-mongo-and-mysql-fdw/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# ddev-timescale-with-mongo-and-mysql-fdw <!-- omit in toc -->

- [What is ddev-timescale-with-mongo-and-mysql-fdw?](#what-is-ddev-timescale-with-mongo-and-mysql-fdw)
- [Components of the repository](#components-of-the-repository)
- [Getting started](#getting-started)

## What is ddev-timescale-with-mongo-and-mysql-fdw?

This repository is a template for providing [DDEV](https://ddev.readthedocs.io) add-ons and services.

In DDEV addons can be installed from the command line using the `ddev get` command, for example, `ddev get ddev/ddev-timescale-with-mongo-and-mysql-fdw` or `ddev get ddev/ddev-drupal9-solr`.

A repository like this one is the way to get started. You can create a new repo from this one by clicking the template button in the top right corner of the page.

![template button](https://raw.githubusercontent.com/eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw/main/images/template-button.png)

## Components of the repository

- The fundamental contents of the add-on service or other component. For example, in this template there is a [docker-compose.timescale-with-mongo-and-mysql-fdw.yaml](https://github.com/eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw/blob/main/docker-compose.timescale-with-mongo-and-mysql-fdw.yaml) file.
- An [install.yaml](https://github.com/eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw/blob/main/install.yaml) file that describes how to install the service or other component.
- A test suite in [test.bats](https://github.com/eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw/blob/main/tests/test.bats) that makes sure the service continues to work as expected.
- [Github actions setup](https://github.com/eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw/blob/main/.github/workflows/tests.yml) so that the tests run automatically when you push to the repository.

## Getting started

1. Choose a good descriptive name for your add-on. It should probably start with "ddev-" and include the basic service or functionality. If it's particular to a specific CMS, perhaps `ddev-<CMS>-servicename`.
2. Create the new template repository by using the template button.
3. Globally replace "timescale-with-mongo-and-mysql-fdw" with the name of your add-on.
4. Add the files that need to be added to a ddev project to the repository. For example, you might remove `docker-compose.timescale-with-mongo-and-mysql-fdw.yaml` with the `docker-compose.*.yaml` for your recipe.
5. Update the install.yaml to give the necessary instructions for installing the add-on.

   - The fundamental line is the `project_files` directive, a list of files to be copied from this repo into the project `.ddev` directory.
   - You can optionally add files to the `global_files` directive as well, which will cause files to be placed in the global `.ddev` directory, `~/.ddev`.
   - Finally, `pre_install_commands` and `post_install_commands` are supported. These can use the host-side environment variables documented [in ddev docs](https://ddev.readthedocs.io/en/stable/users/extend/custom-commands/#environment-variables-provided).

6. Update `tests/test.bats` to provide a reasonable test for your repository. Tests are triggered either by manually executing `bats ./tests/test.bat`, automatically on every push to the repository, or periodically each night. Please make sure to attend to test failures when they happen. Others will be depending on you. `bats` is a simple testing framework that just uses `bash`. To run a Bats test locally, you have to [install bats-core](https://bats-core.readthedocs.io/en/stable/installation.html) first. Then you download your add-on, and finally run `bats ./tests/test.bats` within the root of the uncompressed directory. To learn more about Bats see the [documentation](https://bats-core.readthedocs.io/en/stable/).
7. When everything is working, including the tests, you can push the repository to GitHub.
8. Create a release on GitHub.
9. Test manually with `ddev get <owner/repo>`.
10. You can test PRs with `ddev get https://github.com/<user>/<repo>/tarball/<branch>`
11. Update the README.md to describe the add-on, how to use it, and how to contribute. If there are any manual actions that have to be taken, please explain them. If it requires special configuration of the using project, please explain how to do those. Examples in [ddev/ddev-drupal9-solr](https://github.com/ddev/ddev-drupal9-solr), [ddev/ddev-memcached](https://github.com/eakcize/ddev-timescaledb-with-mongo-and-mysql-fdw/blob/main/github.com/ddev/ddev-memcached), and [ddev/ddev-beanstalkd](https://github.com/ddev/ddev-beanstalkd).
12. Add a good short description to your repo, and add the label "ddev-get". It will immediately be added to the list provided by `ddev get --list --all`.
13. When it has matured you will hopefully want to have it become an "official" maintained add-on. Open an issue in the [ddev queue](https://github.com/ddev/ddev/issues) for that.

Note that more advanced techniques are discussed in [DDEV docs](https://ddev.readthedocs.io/en/latest/users/extend/additional-services/#additional-service-configurations-and-add-ons-for-ddev).

**Contributed and maintained by [@CONTRIBUTOR](https://github.com/CONTRIBUTOR) based on the original [ddev-contrib recipe](https://github.com/ddev/ddev-contrib/tree/master/docker-compose-services/RECIPE) by [@CONTRIBUTOR](https://github.com/CONTRIBUTOR)**

\*\*Originally Contributed by [somebody](https://github.com/somebody) in <https://github.com/ddev/ddev-contrib/>
