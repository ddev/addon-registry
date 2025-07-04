---
title: claudiu-cristea/ddev-virtuoso
github_url: https://github.com/claudiu-cristea/ddev-virtuoso
description: "Virtuoso triplestore installation for DDEV"
user: claudiu-cristea
repo: ddev-virtuoso
repo_id: 583941282
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2022-12-31
updated_at: 2022-12-31
workflow_status: unknown
stars: 0
---

[![tests](https://github.com/drud/ddev-virtuoso/actions/workflows/tests.yml/badge.svg)](https://github.com/drud/ddev-virtuoso/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2022.svg)

## What is ddev-virtuoso?

This repository is a template for providing [DDEV](https://ddev.readthedocs.io) addons and services.

In ddev v1.19+ addons can be installed from the command line using the `ddev get` command, for example, `ddev get drud/ddev-virtuoso` or `ddev get drud/ddev-drupal9-solr`.

A repository like this one is the way to get started. You can create a new repo from this one by clicking the template button in the top right corner of the page.

![template button](https://raw.githubusercontent.com/claudiu-cristea/ddev-virtuoso/main/images/template-button.png)

## Components of the repository

* The fundamental contents of the add-on service or other component. For example, in this template there is a [docker-compose.virtuoso.yaml](https://github.com/claudiu-cristea/ddev-virtuoso/blob/main/docker-compose.virtuoso.yaml) file.
* An [install.yaml](https://github.com/claudiu-cristea/ddev-virtuoso/blob/main/install.yaml) file that describes how to install the service or other component.
* A test suite in [test.bats](https://github.com/claudiu-cristea/ddev-virtuoso/blob/main/tests/test.bats) that makes sure the service continues to work as expected.
* [Github actions setup](https://github.com/claudiu-cristea/ddev-virtuoso/blob/main/.github/workflows/tests.yml) so that the tests run automatically when you push to the repository.

## Getting started

1. Choose a good descriptive name for your add-on. It should probably start with "ddev-" and include the basic service or functionality. If it's particular to a specific CMS, perhaps `ddev-<CMS>-servicename`.
2. Create the new template repository by using the template button.
3. Globally replace "virtuoso" with the name of your add-on.
4. Add the files that need to be added to a ddev project to the repository. For example, you might remove `docker-composeaddon-template.yaml` with the `docker-compose.*.yaml` for your recipe.
5. Update the install.yaml to give the necessary instructions for installing the add-on.
  * The fundamental line is the `project_files` directive, a list of files to be copied from this repo into the project `.ddev` directory.
  * You can optionally add files to the `global_files` directive as well, which will cause files to be placed in the global `.ddev` directory, `~/.ddev`.
  * Finally, `pre_install_commands` and `post_install_commands` are supported. These can use the host-side environment variables documented [in ddev docs](https://ddev.readthedocs.io/en/stable/users/extend/custom-commands/#environment-variables-provided).
6. Update `tests/test.bats` to provide a reasonable test for the repository. You can run it manually with `bats tests` and it will be run on push and nightly as well. Please make sure to attend to test failures when they happen. Others will be depending on you. `bats` is a simple testing framework that just uses `bash`. You can install it with `brew install bats-core` or [see other techniques](https://bats-core.readthedocs.io/en/stable/installation.html). See [bats tutorial](https://bats-core.readthedocs.io/en/stable/).
7. When everything is working, including the tests, you can push the repository to GitHub.
8. Create a release on GitHub.
9. Test manually with `ddev get <owner/repo>`.
10. Update the README.md to describe the add-on, how to use it, and how to contribute. If there are any manual actions that have to be taken, please explain them. If it requires special configuration of the using project, please explain how to do those. Examples in [drud/ddev-drupal9-solr](https://github.com/drud/ddev-drupal9-solr), [drud/ddev-memcached](https://github.com/drud/ddev-memcached), and [drud/ddev-beanstalkd](https://github.com/drud/ddev-beanstalkd).
11. Add a good short description to your repo, and add the label "ddev-get". It will immediately be added to the list provided by `ddev get --list --all`.
12. When it has matured you will hopefully want to have it become an "official" maintained add-on. Open an issue in the [ddev queue](https://github.com/drud/ddev/issues) for that.

Note that more advanced techniques are discussed in [DDEV docs](https://ddev.readthedocs.io/en/latest/users/extend/additional-services/#additional-service-configurations-and-add-ons-for-ddev).

**Contributed and maintained by [@CONTRIBUTOR](https://github.com/CONTRIBUTOR) based on the original [ddev-contrib recipe](https://github.com/drud/ddev-contrib/tree/master/docker-compose-services/RECIPE) by [@CONTRIBUTOR](https://github.com/CONTRIBUTOR)**

**Originally Contributed by [somebody](https://github.com/somebody) in https://github.com/drud/ddev-contrib/...)
