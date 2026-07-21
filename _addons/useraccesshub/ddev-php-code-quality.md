---
title: useraccesshub/ddev-php-code-quality
github_url: https://github.com/useraccesshub/ddev-php-code-quality
description: "DDEV Add-On to add phpcs, phpcbf and phpstan commands to composer based projects."
user: useraccesshub
repo: ddev-php-code-quality
repo_id: 1219223619
default_branch: main
tag_name: 1.0.1
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-04-23
updated_at: 2026-05-01
workflow_status: unknown
stars: 0
---

# DDEV PHP Code Quality

Adds phpcs, phpcbf and phpstan commands to wrap web container commands.

## Installation

```bash
ddev add-on get useraccesshub/ddev-php-code-quality
ddev restart
```

## Post Installation Configuration

This add-on does not include any specific standards so these will need to be included manually.

#### Drupal 11

1. Add the drupal/core-dev package or individual phpcs/phpstan packages:

```bash
ddev composer require --dev drupal/core-dev
```

OR

```bash
ddev composer require --dev drupal/coder
ddev composer require --dev phpstan/phpstan phpstan/extension-installer mglaman/phpstan-drupal phpstan/phpstan-deprecation-rules
```

2. Add a phpcs.xml file at the project root:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<ruleset name="Drupal Project">
  <description>Default PHP CodeSniffer configuration for Drupal.</description>

  <!-- Check these file extensions -->
  <arg name="extensions" value="php,module,inc,install,test,profile,theme"/>
  
  <!-- Use colors in output and show sniff names -->
  <arg name="colors"/>
  <arg value="s"/>

  <!-- Core Drupal standards from the Coder module -->
  <rule ref="Drupal"/>
  <rule ref="DrupalPractice"/>

  <!-- Files or directories to check -->
  <file>web/modules/custom</file>
  <file>web/themes/custom</file>

  <!-- Directories to ignore -->
  <exclude-pattern>*/vendor/*</exclude-pattern>
  <exclude-pattern>*/node_modules/*</exclude-pattern>
  <exclude-pattern>*/dist/*</exclude-pattern>
</ruleset>

```
3. Add a phpstan.neon file at the project root:

```neon
parameters:
  level: 5
  excludePaths:
    - *node_modules/*
    - *vendor/*
    - *.twig
  fileExtensions:
    - php
    - module
    - inc
    - install
    - test
    - profile
    - theme
```

## Usage

Run commands

```bash
ddev phpcs path/to/file/or/directory
ddev phpcbf path/to/file/or/directory
ddev phpstan analyze path/to/file/or/directory
```

Show configuration options:

```bash
ddev phpcs --config-show
```

Get installed code standards:

```bash
ddev phpcs -i
```
