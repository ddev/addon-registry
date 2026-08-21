---
title: "upstreamable/ddev-basin-deploy"
github_url: "https://github.com/upstreamable/ddev-basin-deploy"
description: "Add deployment to ddev projects with Ansistrano"
user: "upstreamable"
repo: "ddev-basin-deploy"
repo_id: 1179123711
default_branch: "main"
tag_name: "1.0.0"
ddev_version_constraint: ">= v1.25.1"
dependencies: ["upstreamable/ddev-basin"]
type: "contrib"
created_at: "2026-03-11"
updated_at: "2026-05-07"
workflow_status: "disabled"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/upstreamable/ddev-basin-deploy/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/upstreamable/ddev-basin-deploy/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/upstreamable/ddev-basin-deploy)](https://github.com/upstreamable/ddev-basin-deploy/commits)
[![release](https://img.shields.io/github/v/release/upstreamable/ddev-basin-deploy)](https://github.com/upstreamable/ddev-basin-deploy/releases/latest)

# DDEV Basin Deploy

## Overview

This add-on provides an [Ansistrano](https://github.com/ansistrano/deploy)
strategy to deploy a [Basin](https://github.com/upstreamable/ddev-basin)
project based on [DDEV](https://ddev.com/).

It deploys a DDEV project on a server that also runs DDEV
using the same configuration but with the hardened images
following the [docs](https://docs.ddev.com/en/stable/users/topics/hosting/).

Only Drupal supported so far.

## Installation

```bash
ddev add-on get upstreamable/ddev-basin-deploy
ddev restart
```

# Requirements

A remote server running DDEV and rsync installed.
Run the following command on the server before the first deploy:
```
ddev config global --instrumentation-opt-in=false --router-bind-all-interfaces --omit-containers=ddev-ssh-agent --use-hardened-images --performance-mode=none --use-letsencrypt --letsencrypt-email=you@example.com
```
To prepare it better for hosting and avoid interactive questions on the
DDEV CLI that would stall the deploy.
Adapt for the letsencrypt email and the instrumentation preferences.

Ensure cron is installed and running if you require cron jobs.

## Usage

Once installed the deployment files need to be generated and adjusted
Run the following command
```
ddev basin deploy:generate
```
And inspect what was created in `.ddev/basin-deploy/`.
The files `before-symlink-tasks.yaml`, `after-update-code.yml` and `playbook.yaml`
are not meant to be edited.

There is subfolders such as `.ddev/basin-deploy/production/` for an
environment with the different configurations such as the domain names
for the environment or the SMTP server to use.

Before running the deployment a remote server is needed.

Add to `.ddev/.env.web` the following variables
```
ANSIBLE_REMOTE_USER=ubuntu
ANSIBLE_REMOTE_HOST=1.1.1.1
```
Replace by the values you would use in a ssh connection such as `ubuntu@1.1.1.1`.
The user need to be able to run DDEV commands.

After the configuration restart the DDEV project and run:
```
ddev basin deploy:release
```
For every deployment. Defaults to the production environment.

# Features

The deploy takes care of running the deploy commands such as `drush deploy`,
configures the traefik proxy for the domains specified and configures the
cron daemon tu run it periodically for the project.

## Credits

**Contributed and maintained by [@upstreamable](https://github.com/upstreamable)**
