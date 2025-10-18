---
title: tyler36/ddev-serverspec
github_url: https://github.com/tyler36/ddev-serverspec
description: "Intergrate serverspec into a DDEV project"
user: tyler36
repo: ddev-serverspec
repo_id: 1076591738
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-10-15
updated_at: 2025-10-17
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-serverspec/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tyler36/ddev-serverspec/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-serverspec)](https://github.com/tyler36/ddev-serverspec/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-serverspec)](https://github.com/tyler36/ddev-serverspec/releases/latest)

# DDEV Serverspec

## Overview

This add-on integrates [Serverspec](https://serverspec.org/) into your [DDEV](https://ddev.com/) project.

This add-on installs the required serverspec packages inside the web container.

## Installation

```bash
ddev add-on get tyler36/ddev-serverspec
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command                | Description          |
| ---------------------- | -------------------- |
| `ddev serverspec-init` | Scaffold tests       |
| `ddev serverspec`      | Run serverspec tests |

### command: serverspec-init

`ddev serverspec-init` starts a wizard to generate required files and scaffold an example test.

The following is an example that generates that confirms a service is listening on port 80.

```shell
$ ddev serverspec-init
Select OS type:
  1) UN*X
  2) Windows
Select number: 1

Select a backend type:
  1) SSH
  2) Exec (local)

Select number: 2
 + spec/
 + spec/localhost/
 + spec/localhost/sample_spec.rb
 + spec/spec_helper.rb
 + Rakefile
 + .rspec
```

### command: serverspec

`ddev serverspec` runs available serverspec tests.

The command accepts flags and arguments, as required.

For example: to stop at the first failing test.

```shell
ddev rspec --fail-fast
```

## Credits

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
