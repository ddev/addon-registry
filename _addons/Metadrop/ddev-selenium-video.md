---
title: Metadrop/ddev-selenium-video
github_url: https://github.com/Metadrop/ddev-selenium-video
description: ""
user: Metadrop
repo: ddev-selenium-video
repo_id: 986424282
ddev_version_constraint: ">= v1.24.3"
dependencies: ["ddev-selenium", "ddev-aljibe"]
type: contrib
created_at: 2025-05-19
updated_at: 2025-05-19
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/Metadrop/ddev-selenium-video/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/Metadrop/ddev-selenium-video/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/Metadrop/ddev-selenium-video)](https://github.com/Metadrop/ddev-selenium-video/commits)
[![release](https://img.shields.io/github/v/release/Metadrop/ddev-selenium-video)](https://github.com/Metadrop/ddev-selenium-video/releases/latest)

# DDEV Selenium Video

## Overview

This add-on integrates Selenium Video into your [DDEV](https://ddev.com/) project.

## Requirements
- [DDEV Aljibe] (https://www.github.com/metadrop/ddev-aljibe)
- [DDEV selenium] (https://www.github.com/metadrop/ddev-selenium)

## Installation

```bash
ddev add-on get Metadrop/ddev-selenium-video
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command                       | Description                            |
|-------------------------------|----------------------------------------|
| `ddev behat-video <env>`      | Run behat tests with recodings enabled |

## Credits

**Contributed and maintained by [@Metadrop](https://github.com/Metadrop)**
