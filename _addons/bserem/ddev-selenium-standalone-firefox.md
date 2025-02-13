---
title: bserem/ddev-selenium-standalone-firefox
github_url: https://github.com/bserem/ddev-selenium-standalone-firefox
description: "A DDEV service for running standalone Firefox"
user: bserem
repo: ddev-selenium-standalone-firefox
type: contrib
created_at: 2023-03-17
updated_at: 2023-05-08
stars: 0
---

[![tests](https://github.com/ddev/ddev-selenium-standalone-firefox/actions/workflows/tests.yml/badge.svg)](https://github.com/ddev/ddev-addon-template/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## Introduction

This service can be used with any project type. The examples below are Drupal-specific. Contributions for docs and tests that show this service working with other project types are appreciated.

## Install/Update

1. `ddev get ddev/ddev-selenium-standalone-firefox`
2. Optional. Update the provided .ddev/config.selenium-standalone-firefox.yaml as you see fit(and remove the #ddev-generated line). You can also just override lines in your .ddev/config.yaml
3. Optional. Check config.selenium-standalone-firefox.yaml and docker-compose.selenium-firefox.yaml into your source control.
4. Update by re-running `ddev get ddev/ddev-selenium-standalone-firefox`.

## Use

- Your project is now ready to run FunctionalJavascript and [Nightwatch](https://www.drupal.org/docs/automated-testing/javascript-testing-using-nightwatch) tests from Drupal core, or [Drupal Test Traits](https://gitlab.com/weitzman/drupal-test-traits) (DTT). All these types are tested in this repo. Some examples to try:
  - FunctionalJavascript:
    - Ensure you have the `drupal/core-dev` Composer package or equivalent.
    - `ddev exec -d /var/www/html/web "../vendor/bin/phpunit -v -c ./core/phpunit.xml.dist ./core/modules/system/tests/src/FunctionalJavascript/FrameworkTest.php"`
  - Nightwatch
    - `ddev exec -d /var/www/html/web/core yarn install` (do this once)
    - `ddev exec -d /var/www/html/web/core touch .env` (do this once)
    - `ddev exec -d /var/www/html/web/core yarn test:nightwatch tests/Drupal/Nightwatch/Tests/exampleTest.js`
  - Drupal Test Traits
    - Ensure you have a working site that has the `weitzman/drupal-test-traits` Composer package.
    - `ddev exec -d /var/www/html/web "../vendor/bin/phpunit --bootstrap=../vendor/weitzman/drupal-test-traits/src/bootstrap-fast.php --printer '\Drupal\Tests\Listeners\HtmlOutputPrinter' ../vendor/weitzman/drupal-test-traits/tests/ExampleSelenium2DriverTest.php"`
  - Behat
    - See below for `behat.yml` configuration
    - `composer require drupal/drupal-extension --dev`

## Watching the tests

### The easy way: Use noVNC (built-in)

On your host, browse to https://[DDEV SITE URL]:7900 (password: `secret`) to watch tests run with noVNC (neat!).

This is a no-configuration solution that enables you to quickly see what is going on with your tests.

### Use a local VNC client (try noVNC first!)

If you are using something like behat and want to debug tests when they fail by manually navigating around your site in the Firefox browser included with this addon, you might want to use a VNC client installed on your machine, such as Screen Sharing on macOS (built-in) or TightVNC on Linux and Windows (must be downloaded and installed). This is because with noVNC, you are running a browser (Firefox) inside another browser (whatever browser you use on your local machine), which can be inconvenient-- for example, the keyboard shortcut to reload a page in Firefox will reload your local browser and kick you out of noVNC instead of reloading Firefox, and it may be hard to type a new url in the Firefox address bar due to how your local browser handles keyboard input.

In other words, if you just want to watch the tests, use noVNC.

If you want to use the browser provided by this addon to check out the test results by poking around your site, consider using a local VNC client. To do so, you need to open port 5900.

#### How to open port 5900 for VNC access

1. Open `.ddev/docker-compose.selenium-firefox.yaml`.
2. Uncomment the two lines about `ports` and `5900:5900`.
3. Execute `ddev restart`.

You can now connect to [DDEV SITE URL]:5900 (password: `secret`) in your VNC client.

Note that when using `ports`, only one project at a time can be running with port 5900.

### Behat config example

If you use Behat as a test running, adjust your `behat.yml`:

```yml
  extensions:
    Behat\MinkExtension:
      base_url: http://mysite.ddev.site
      selenium2:
        wd_host: http://selenium-firefox:4444/wd/hub
        capabilities:
          firefox:
```

- Drupal users can swap `Behat\MinkExtension` with `Drupal\MinkExtension` for extra goodness.
- You can remove/comment the `--headless` option to see the browser running in VNC. You can even interact with it when pausing (ie: `When I break` if your test)

## Contribute

- Anyone is welcome to submit a PR to this repo.

## Maintainer

- Contributed and maintained by [@bserem](https://github.com/bserem).
- Sponsored by [Annertech](https://www.annertech.com)
- Forked from [ddev-selenium-standalone-chrome](https://github.com/ddev/ddev-selenium-standalone-chrome) by [@weitzman](https://github.com/weitzman).

