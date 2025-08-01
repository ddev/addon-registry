---
title: froboy/ddev-tsh
github_url: https://github.com/froboy/ddev-tsh
description: "DDEV integration with Teleport and tsh."
user: froboy
repo: ddev-tsh
repo_id: 1025009597
ddev_version_constraint: ">= v1.24.3"
dependencies: ["ddev-kubernetes"]
type: contrib
created_at: 2025-07-23
updated_at: 2025-07-25
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/froboy/ddev-tsh/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/froboy/ddev-tsh/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/froboy/ddev-tsh)](https://github.com/froboy/ddev-tsh/commits)
[![release](https://img.shields.io/github/v/release/froboy/ddev-tsh)](https://github.com/froboy/ddev-tsh/releases/latest)

# DDEV Tsh

## Overview

This add-on integrates [Teleport](https://goteleport.com/) and its command line tool, [tsh](https://goteleport.com/docs/connect-your-client/tsh/) into your [DDEV](https://ddev.com/) project.

Teleport is an open-source platform that provides secure access to SSH, Kubernetes, databases, and internal applications. The `tsh` command allows users to connect to infrastructure via Teleport and manage their resources. If your infrastructure is behind a Teleport proxy, this add-on allows you to:

- use `tsh` commands directly within your DDEV project, and 
- configure [drush aliases](https://www.drush.org/latest/site-aliases/#kubernetes) to connect to remote environments.

## Installation

It is recommended to use this add-on in conjunction with the [DDEV Kubernetes add-on](https://github.com/MurzNN/ddev-kubernetes).

```shell
 ddev add-on get MurzNN/ddev-kubernetes
```

During installation, you will be asked to input your Teleport username, proxy, cluster, and k8s cluster (if applicable). You can skip any if you don't need them or don't know them yet.

```bash
ddev add-on get froboy/ddev-tsh
ddev restart
```

After installation, adjust any files for your project needs. Remove `#ddev-generated` from any files that you edit, as they will be overwritten on the next `ddev restart`.

## Usage

| Command          | Description                                                                       |
|------------------|-----------------------------------------------------------------------------------|
| `ddev tsh`       | Passthrough for the [tsh](https://goteleport.com/docs/reference/cli/tsh/) command |
| `ddev tsh-login` | Log in to a Teleport server and connect to a k8s cluster.                         |

## Advanced Customization

If you are using drush aliases, you will want to set up a drush alias using `kubectl` to connect to your Teleport cluster. 

For example, in `drush/sites/example.site.yml`:
```yaml
my-site-k8s:
  kubectl:
    namespace: 'my-site'
    resource: 'deploy/my-site'
    container: 'php'
    tty: true
    interactive: true
  paths:
    drush-script: /path/to/project/vendor/bin/drush
  uri: https://my-site.example.com/
  ```

Please note that `tty` and `interactive` are set to `true` to allow for interactive commands via `tsh` in the container.

## Credits

**Contributed and maintained by [@froboy](https://github.com/froboy) for [ImageX](https://imagexmedia.com/)**
