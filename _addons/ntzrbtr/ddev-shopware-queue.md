---
title: "ntzrbtr/ddev-shopware-queue"
github_url: "https://github.com/ntzrbtr/ddev-shopware-queue"
description: "Run Shopware message queue and scheduler inside ddev"
user: "ntzrbtr"
repo: "ddev-shopware-queue"
repo_id: 1344971680
default_branch: "main"
tag_name: "1.0.2"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-08-24"
updated_at: "2026-08-25"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/ntzrbtr/ddev-shopware-queue/actions/workflows/tests.yml/badge.svg)](https://github.com/ntzrbtr/ddev-shopware-queue/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/ntzrbtr/ddev-shopware-queue)](https://github.com/ntzrbtr/ddev-shopware-queue/commits)
[![release](https://img.shields.io/github/v/release/ntzrbtr/ddev-shopware-queue)](https://github.com/ntzrbtr/ddev-shopware-queue/releases/latest)

# DDEV Shopware Queues

## Overview

This add-on runs the Shopware 6 message queue consumer and the scheduled task
runner inside the DDEV web container, so you don't have to start them
manually or rely on the (deprecated) Admin Worker.

See the official [Message Queue](https://developer.shopware.com/docs/guides/hosting/infrastructure/message-queue.html)
and [Scheduled Task](https://developer.shopware.com/docs/guides/hosting/infrastructure/scheduled-task.html)
documentation for more details.

## Installation

```shell
ddev add-on get ntzrbtr/ddev-shopware-queue
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

It's recommended to disable Shopware's Admin Worker once this add-on is
running, so messages aren't processed twice, in `config/packages/shopware.yaml`:

```yaml
shopware:
  admin_worker:
    enable_admin_worker: false
```

## Usage

Both processes run automatically as background daemons inside the web
container. No further configuration is required. To see their output, check
the web container logs:

```shell
$ ddev logs -s web
...
2026-01-01 09:12:23,817 INFO spawned: 'shopware-queue' with pid 1666
2026-01-01 09:12:23,818 INFO spawned: 'shopware-scheduler' with pid 1667
2026-01-01 09:12:27,571 INFO success: shopware-queue entered RUNNING state, process has stayed up for > than 3 seconds (startsecs)
2026-01-01 09:12:27,572 INFO success: shopware-scheduler entered RUNNING state, process has stayed up for > than 3 seconds (startsecs)
...
```

You can also inspect the daemons directly:

```shell
ddev exec supervisorctl status
```

## Configuration

`.ddev/config.shopware-queue.yaml` starts two daemons:

- **`shopware-queue`**: runs `bin/console messenger:consume --all`
  to process queued messages.
  - `--time-limit=60` : restarts the worker every 60 seconds to avoid memory
    leaks from long-running PHP processes.
  - `--memory-limit=512M` : restarts the worker if it exceeds 512MB of memory.
  - `--all` : runs all transports. Adjust this if your project defines custom
    or you just want to run some transports.
- **`shopware-scheduler`**: runs `bin/console scheduled-task:run`, which
  blocks and dispatches scheduled tasks (cache warmers, cleanup jobs, etc.)
  as they become due, without needing a separate cron entry.

Both commands are wrapped in a `while true; do ...; sleep 1; done` loop so
they're automatically restarted if they exit (e.g. after hitting
`--time-limit` or `--memory-limit`).

To configure this add-on, remove `#ddev-generated` from
`.ddev/config.shopware-queue.yaml` and edit the file as required — for
example to also consume the `failed` transport (see the commented-out
example in the file).

## Credits

**Contributed and maintained by [ntzrbtr](https://github.com/ntzrbtr)**
