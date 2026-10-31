---
title: "iljapolanskis/ddev-buggregator"
github_url: "https://github.com/iljapolanskis/ddev-buggregator"
description: "Buggregator service for DDEV (similar to Ray, but free)"
user: "iljapolanskis"
repo: "ddev-buggregator"
repo_id: 673257959
default_branch: "main"
tag_name: "1.1.2"
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2023-08-01"
updated_at: "2026-10-30"
workflow_status: "disabled"
stars: 2
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/iljapolanskis/ddev-buggregator/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/iljapolanskis/ddev-buggregator/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/iljapolanskis/ddev-buggregator)](https://github.com/iljapolanskis/ddev-buggregator/commits)
[![release](https://img.shields.io/github/v/release/iljapolanskis/ddev-buggregator)](https://github.com/iljapolanskis/ddev-buggregator/releases/latest)

# DDEV Buggregator

## Overview

This add-on integrates [Buggregator](https://github.com/buggregator/server) into your [DDEV](https://ddev.com/) project. Buggregator is a powerful debugging server that helps with application debugging and monitoring by collecting logs, profiling data, and other debugging information.

## Installation

```bash
ddev add-on get iljapolanskis/ddev-buggregator
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for Buggregator |
| `ddev logs -s buggregator` | Check Buggregator logs |

### Accessing Buggregator UI

Open `https://<project>.ddev.site:8000` in your browser to access the Buggregator web interface.

### Connecting your application

From inside the `web` container Buggregator is reachable by its service name, `buggregator`.
Set these variables in your application's `.env` (see the [Buggregator documentation](https://docs.buggregator.dev/) for framework-specific setup):

| Feature | Port | Configuration |
| ------- | ---- | ------------- |
| [Sentry](https://docs.buggregator.dev/config/sentry.html) | 8000 | `SENTRY_DSN=http://sentry@buggregator:8000/1` |
| [Ray](https://docs.buggregator.dev/config/ray.html) | 8000 | `RAY_HOST=ray@buggregator`, `RAY_PORT=8000` |
| [Inspector](https://docs.buggregator.dev/config/inspector.html) | 8000 | `INSPECTOR_URL=http://inspector@buggregator:8000`, `INSPECTOR_API_KEY=test` |
| [XHProf](https://docs.buggregator.dev/config/xhprof.html) | 8000 | `PROFILER_ENDPOINT=http://profiler@buggregator:8000` |
| [HTTP dumps](https://docs.buggregator.dev/config/http-dumps.html) | 8000 | `http://http-dump@buggregator:8000` |
| [SMTP](https://docs.buggregator.dev/config/smtp.html) | 1025 | `MAILER_DSN=smtp://buggregator:1025` |
| [VarDumper](https://docs.buggregator.dev/config/var-dumper.html) | 9912 | `VAR_DUMPER_FORMAT=server`, `VAR_DUMPER_SERVER=buggregator:9912` |
| [Monolog](https://docs.buggregator.dev/config/monolog.html) | 9913 | `SocketHandler('buggregator:9913')` |

Only the web UI (port 8000) is published through the DDEV router. SMTP, VarDumper and Monolog are plain TCP services and are only reachable from other containers in the project.

## Configuration

Customize the add-on with `ddev dotenv set`, which writes to `.ddev/.env.buggregator`:

```bash
# Pin the Buggregator image version
ddev dotenv set .ddev/.env.buggregator --buggregator-image=ghcr.io/buggregator/server:latest

# Serve the web UI on another router port, e.g. https://<project>.ddev.site:8100
ddev dotenv set .ddev/.env.buggregator --buggregator-http-port=8100

ddev restart
```

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `BUGGREGATOR_IMAGE` | `ghcr.io/buggregator/server:latest` | Docker image to use |
| `BUGGREGATOR_HTTP_PORT` | `8000` | Router port for the HTTPS web UI |

Variables in `.ddev/.env.buggregator` are also passed into the Buggregator container, so any [server setting](https://docs.buggregator.dev/config/server.html) can be configured there too.

## PHP Usage Example

Here's an example of how to send logs to Buggregator from within your PHP application:

First, install the required PHP library:

```bash
composer require monolog/monolog
```

Then use it in your code (playground.php file):

```php
<?php

require_once 'vendor/autoload.php';

use Monolog\Logger;
use Monolog\Handler\SocketHandler;
use Monolog\Formatter\JsonFormatter;

// Create a log channel
$logger = new Logger('myapp');

// Create socket handler for Buggregator
$handler = new SocketHandler('buggregator:9913');
$handler->setFormatter(new JsonFormatter());
$logger->pushHandler($handler);

// Send different log levels to Buggregator
$logger->info('Application started');
$logger->warning('This is a warning message');
$logger->error('An error occurred', ['user_id' => 123, 'action' => 'login']);
$logger->debug('Debug information', ['query' => 'SELECT * FROM users']);
```

Note: This example uses the default Buggregator service hostname (`buggregator`) which is accessible from within the DDEV environment.

```shell
ddev exec php playground.php
```

## Credits

**Contributed and maintained by [@iljapolanskis](https://github.com/iljapolanskis)**

This add-on uses the [Buggregator Server](https://github.com/buggregator/server) Docker image as the base for the debugging service.
