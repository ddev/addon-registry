---
title: rfay/ddev-drushonhost
github_url: https://github.com/rfay/ddev-drushonhost
description: "DDEV Add-on to Allow using drush on the host"
user: rfay
repo: ddev-drushonhost
type: contrib
created_at: 2023-11-11
updated_at: 2024-07-09
stars: 0
---

[![tests](https://github.com/rfay/ddev-drushonhost/actions/workflows/tests.yml/badge.svg)](https://github.com/rfay/ddev-drushonhost/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

## What is ddev-drushonhost?

In DDEV v1.22.4+, the former facility to use the Drupal `drush` command on the host was removed.  That usage was always fragile, but some people liked and used it. This add-in also permits running PHPUnit/DTT ([Drupal Test Traits](https://gitlab.com/weitzman/drupal-test-traits)) tests on the host instead of running them in the containers. [See this article for more on running DTT on the host](https://selwynpolit.github.io/d9book/dtt#run-tests-on-the-host)

This add-on adds that facility again.

One-time in project:

`ddev get rfay/ddev-drushonhost`

When using drush on the host you need:

```
export IS_DDEV_PROJECT=true
```
OR

In your project `settings.php` make sure the end of the file looks like this:

```php
if (file_exists($app_root . '/' . $site_path . '/settings.local.php')) {
  include $app_root . '/' . $site_path . '/settings.local.php';
}

// Automatically generated include for settings managed by ddev.
$ddev_settings = dirname(__FILE__) . '/settings.ddev.php';
if (getenv('IS_DDEV_PROJECT') == 'true' && is_readable($ddev_settings)) {
  require $ddev_settings;
}
```

and then add this to your `settings.local.php`:

```
putenv("IS_DDEV_PROJECT=true");
```

Running `drush` commands in the containers is faster, however, dependingon your use, you may find it more convenient to run them on the host. [Read more about installing Global Drush and using this add-in](https://selwynpolit.github.io/d9book/drush#global-drush). Note. The host drush version doesn't matter much since it is only used to find the proper `drush` (most likely within /vendor/bin) and call it. This means you always run the version of `drush` that you installed in your project using composer.

From the former documentation:

**Warning:** We don’t recommend using `drush` on your host machine. It’s also mostly irrelevant for Drupal 9+, as you should be using Composer-installed, project-level `drush`.

If you have PHP and Drush installed on your host system and the environment variable `IS_DDEV_PROJECT=true`, you can use Drush to interact with a DDEV project. On the host machine, extra host-side configuration for the database and port in `settings.ddev.php` allow Drush to access the database server. This may not work for all Drush commands because the actual web server environment is not available.

On Drupal 8+, if you want to use `drush uli` on the host (or other Drush commands that require a default URI), you’ll need to set `DRUSH_OPTIONS_URI` on the host. For example, `export DRUSH_OPTIONS_URI=https://mysite.ddev.site`.

**Contributed and maintained by [@rfay](https://github.com/rfay)**

