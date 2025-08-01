---
title: MurzNN/ddev-grafana
github_url: https://github.com/MurzNN/ddev-grafana
description: "Grafana Stack addon for DDEV: Grafana (GUI), Tempo (tracing, OpenTelemetry), Loki (logs, promtail), Mimir (metrics, prometheus)"
user: MurzNN
repo: ddev-grafana
repo_id: 621377265
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-03-30
updated_at: 2025-07-04
workflow_status: failure
stars: 8
---

[![tests](https://github.com/MurzNN/ddev-grafana/actions/workflows/tests.yml/badge.svg)](https://github.com/MurzNN/ddev-grafana/actions/workflows/tests.yml)
![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# ddev-grafana <!-- omit in toc -->

* [What is ddev-grafana?](#what-is-ddev-grafana)
* [Components of the repository](#components-of-the-repository)
* [Getting started](#getting-started)

## What is ddev-grafana?

This repository provides a Grafana LGTM+ Stack addon to
[DDEV](https://ddev.readthedocs.io).

It contains several components from the Grafana stack:

- **[Grafana](https://grafana.com/oss/grafana/)**: an open source analytics and
  interactive visualization web application. It provides charts, graphs, and
  alerts for the web when connected to supported data sources.

- **[Loki](https://grafana.com/oss/loki/)**: a horizontally scalable, highly
  available, multi-tenant log aggregation solution.

- **[Tempo](https://grafana.com/oss/tempo/)**: an open source, easy-to-use, and
  high-scale distributed tracing backend.

- **[Mimir](https://grafana.com/oss/mimir/)**: an open source, horizontally
  scalable, highly available, multi-tenant TSDB for long-term storage for
  Prometheus.

- **[Alloy](https://grafana.com/oss/alloy-opentelemetry-collector/)**: an open
  source OpenTelemetry collector with built-in Prometheus pipelines and support
  for metrics, logs, traces, and profiles.


## Installation

1. In the DDEV project directory:

  ```sh
  ddev get MurzNN/ddev-grafana
  ```

2. Restart the DDEV instance:

  ```sh
  ddev restart
  ```

3. Open the Grafana web interface via the url:
   https://your-project-name.ddev.site:3000/

4. Configure the OpenTelemetry endpoint in your application
   to the `http://opentelemetry-grafana:4318`.

   By default, the add-on configures the `OTEL_EXPORTER_OTLP_ENDPOINT`
   environment variable with this endpoint, you can disable this in the file
   `.ddev/config.grafana.logs.yaml`.

** Contributed and maintained by [@MurzNN](https://github.com/MurzNN).

## Extra

Integration with popular CMSs and frameworks:

- Drupal CMS: [OpenTelemery
  module](https://www.drupal.org/project/opentelemetry)

- Symfony: [OpenTelemetry Symfony
  auto-instrumentation](https://github.com/opentelemetry-php/contrib-auto-symfony)

- Laravel: [Open Telemetry
  package](https://github.com/spatie/laravel-open-telemetry)
