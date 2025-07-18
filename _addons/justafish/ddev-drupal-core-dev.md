---
title: justafish/ddev-drupal-core-dev
github_url: https://github.com/justafish/ddev-drupal-core-dev
description: "ddev addon for core development"
user: justafish
repo: ddev-drupal-core-dev
repo_id: 706601727
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-10-18
updated_at: 2025-06-12
workflow_status: success
stars: 28
---

# ddev-core-dev

This is a DDEV addon for doing Drupal core development.

We're in #ddev-for-core-dev on [Drupal Slack](https://www.drupal.org/community/contributor-guide/reference-information/talk/tools/slack) (but please try and keep work
and feature requests in Issues where it's visible to all 🙏)

## Installation
```
git clone https://git.drupalcode.org/project/drupal.git drupal
cd drupal
ddev config --omit-containers=db --disable-settings-management
ddev composer install
ddev add-on get justafish/ddev-drupal-core-dev

# See included commands
ddev drupal list
```

The `drupal` command is an extension of core's [`drupal`](https://git.drupalcode.org/project/drupal/-/blob/11.x/core/scripts/drupal?ref_type=heads)
command. This allows you to perform some basic tasks without needing to install
`drush` which will alter your composer dependencies.

## Examples
```
# Install drupal
# Run "ddev drupal install" to see all available options
ddev drupal install standard

# Run PHPUnit tests
ddev phpunit core/modules/announcements_feed

# Run Nightwatch tests (currently only runs on Chrome)
ddev nightwatch --tag core
```

## Nightwatch Examples

You can watch Nightwatch running in real time at https://drupal.ddev.site:7900
for Chrome and https://drupal.ddev.site:7901 for Firefox. The password is
"secret". YMMV using Firefox as core tests don't currently run on it.

Only core tests
```
ddev nightwatch --tag core
```

Skip running core tests
```
ddev nightwatch --skiptags core
```

Run a single test
```
ddev nightwatch tests/Drupal/Nightwatch/Tests/exampleTest.js
```

a11y tests for both the admin and default themes
```
ddev nightwatch --tag a11y
```

a11y tests for the admin theme only
```
ddev nightwatch --tag a11y:admin
```

a11y tests for the default theme only
```
ddev nightwatch --tag a11y:default
```

a11y test for a custom theme used as the default theme
```
ddev nightwatch --tag a11y:default --defaultTheme bartik
```

a11y test for a custom admin theme
```
ddev nightwatch --tag a11y:admin --adminTheme seven
```

## Core Linting

This will run static tests against core standards.

```
ddev drupal lint:phpstan
ddev drupal lint:phpcs
ddev drupal lint:js
ddev drupal lint:css
ddev drupal lint:cspell
# CSpell against only modified files
ddev drupal lint:cspell --modified-only
```

You can run all linting with `ddev drupal lint`, or with fail-fast turned on:
`ddev drupal lint --stop-on-failure`
