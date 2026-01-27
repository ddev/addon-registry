---
title: tyler36/ddev-site-metrics-cakephp
github_url: https://github.com/tyler36/ddev-site-metrics-cakephp
description: "Setup OpenTelemetry for CakePHP projects"
user: tyler36
repo: ddev-site-metrics-cakephp
repo_id: 999300917
default_branch: main
tag_name: v0.3
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-06-10
updated_at: 2025-06-23
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-site-metrics-cakephp/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tyler36/ddev-site-metrics-cakephp/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-site-metrics-cakephp)](https://github.com/tyler36/ddev-site-metrics-cakephp/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-site-metrics-cakephp)](https://github.com/tyler36/ddev-site-metrics-cakephp/releases/latest)

# DDEV Site Metrics CakePHP

## Overview

This add-on adds Open Telemetry for CakePHP projects. It is designed to integrate with tyler36/ddev-site-metrics.
This is achieved by:

- installing opentelemetry PHP addon.
- configures system-level environmental variables

## Installation

```bash
ddev add-on get tyler36/ddev-site-metrics-cakephp
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

Traces are automatically injected into PHP calls.
This is configured via `.ddev/.env.web`.

```conf
OTEL_PHP_AUTOLOAD_ENABLED="true"
OTEL_SERVICE_NAME="cakephp"
OTEL_METRIC_EXPORTER="none"
OTEL_LOGS_EXPORTER="console"
OTEL_TRACES_EXPORTER="console"
```

- To disable all telemetry, update `.ddev/.env.web` and restart DDEV.

```conf
OTEL_PHP_AUTOLOAD_ENABLED="false"
```

- To send traces to ddev-site-metrics, update `.ddev/.env.web` and restart DDEV.

```conf
OTEL_TRACES_EXPORTER="otlp"
OTEL_EXPORTER_OTLP_ENDPOINT=http://grafana-alloy:4317
```

## Configuration

### Environmental Variables

Environmental variables need to be set early in the process.
This addon uses `.ddev/.env.web` to set them for the container.
Additionally, we set them in the `web` container to prevent leakage into other containers.

## Debugging

- Confirm PHP module is installed.

```shell
$ ddev php --ri opentelemetry
...
opentelemetry hooks => enabled
```

## Credits

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
