---
title: weitzman/ddev-mtk
github_url: https://github.com/weitzman/ddev-mtk
description: "DDEV add-on for the MySQL Toolkit (MTK)"
user: weitzman
repo: ddev-mtk
repo_id: 1009684755
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-06-27
updated_at: 2025-07-28
workflow_status: failure
stars: 2
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/weitzman/ddev-mtk/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/weitzman/ddev-mtk/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/weitzman/ddev-mtk)](https://github.com/weitzman/ddev-mtk/commits)
[![release](https://img.shields.io/github/v/release/weitzman/ddev-mtk)](https://github.com/weitzman/ddev-mtk/releases/latest)

# DDEV MTK

## Overview

This add-on integrates [MySQL Toolkit](https://mtk.skpr.io/) into your [DDEV](https://ddev.com/) project. You want this add-on if your site's database is large, and its getting slow to copy it for local development, or CI testing, or preview environments (e.g. [TugBoat](https://www.tugboatqa.com/)).

DDEV's typical approach for pulling down the database from Production is to copy a database dump and re-import locally. The re-import can be a slow process, even when MySQL is well tuned. MTK imports the database once into a Docker image, so your developers and CI just have to fetch an image and they are off and running. No re-import required. Docker images are already compressed, so the fetch is relatively quick. 

Note that the MTK approach can make site setup fast enough for functional tests using existing site data. See [Drupal Test Traits](https://git.drupalcode.org/project/dtt) for a project that facilitates this.

## Installation

> [!NOTE]
> Installation will not interfere with your current site.

```bash
ddev add-on get weitzman/ddev-mtk
ddev restart
ddev describe
```
Notice that you now have an `mtk` service listed in `ddev describe`. At first, this is an empty and unused placeholder MySQL database image. Read on to learn how you build and publish your site's database image, which will replace the placeholder image.

Run `ddev exec mtk table list db`. You should see a list of table names. This verifies that `mtk` is installed in the web service. 

## Configuration

1. Create a `mtk.yml` file in the root of your project. It can be empty to start. Eventually, populate it as per https://mtk.skpr.io/docs/tutorial#configuration-file, for a slimmed and sanitized database.
1. Run `ddev exec mtk dump db > dump.sql` to generate a SQL dump file.
1. Use the `dump.sql` from above when building and pushing your database image to a Docker registry (e.g. [Docker Hub](https://hub.docker.com/)). See the tutorial at https://mtk.skpr.io/docs/database-image. Do this on the host, not in the container. Remember to push to a _private_ Docker repository.
1. Now that you have published a DB image with your data inside, configure your site to use it.
  1. Edit `.ddev/.env.mtk.web` as follows:
      - Remove the `#ddev-generated` line at the top.
      - Set `MTK_HOSTNAME=mtk`
      - Edit `MTK_USER`, `MTK_PASSWORD`, `MTK_DATABASE` to match whatever your published expects.
      - Set `DDEV_MTK_DOCKER_IMAGE` to the image and tag that you published above. For example, `example/db:latest`
  1. Edit Drupal's settings.php with code like below so that Drupal connects to the `mtk` service instead of the typical `db` service.
      ```php
      if (getenv('IS_DDEV_PROJECT') == 'true') {
      $file_mtk = getenv('DDEV_COMPOSER_ROOT') . '/.ddev/settings.ddev-mtk.php';
      if (file_exists($file_mtk)) {
        include $file_mtk;
       }
      }
      ```
  1. `ddev restart`. Your site is now using the `mtk` service instead of `db`. Make sure the site works by running `ddev drush st` (look for _Drupal bootstrap: Successful_). Run `ddev launch` and to verify that a browser request succeeds.
1. _Optional_. Omit the standard `db` service since your site no longer uses it. `ddev config --omit-containers db`
1. Commit the `.ddev` directory and settings.php change to version control so your teammates start using the `mtk` service.
1. Set up a CI job to refresh your database image on a weekly or nightly basis. The job should push to the same tag every time (e.g. `latest`). 

## CI, Preview Environments, and more.

Consider speeding up other DB consumers by using the image you published above. See https://mtk.skpr.io/docs/database-image#integrate-with-your-cicd-pipeline for a few helpful snippets. Consider using own runners such as ([Bitbucket](https://support.atlassian.com/bitbucket-cloud/docs/runners/), [Gitlab CI](https://docs.gitlab.com/runner/) to highly isolate your DB data.

## Commands

| Command          | Description                                              |
|------------------|----------------------------------------------------------|
| `ddev exec mtk`  | List tables, sanitize tables, dump a database.           |
| `ddev pulldb`    | Refresh your site's database (i.e. the mtk Docker image) |
| `ddev sequelace` | Open your site's DB in the Sequel Ace GUI application    |
| `ddev tableplus` | Open your site's DB in the TablePlus GUI application     |

## Limitations

- Non-functional DDEV commands: 
  - `export-db`, `import-db`. Use Drush sql commands instead.
  - `snapshot`. Not usually needed since you can revert your Docker image.

## Credits

**Contributed and maintained by [@weitzman](https://github.com/weitzman)**
