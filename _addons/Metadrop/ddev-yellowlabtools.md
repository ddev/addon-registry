---
title: Metadrop/ddev-yellowlabtools
github_url: https://github.com/Metadrop/ddev-yellowlabtools
description: "Ddev addon to run yellowlabs tools cli"
user: Metadrop
repo: ddev-yellowlabtools
repo_id: 1103805997
default_branch: main
tag_name: v1.0.1
ddev_version_constraint: ">= v1.23.0"
dependencies: []
type: contrib
created_at: 2025-11-25
updated_at: 2025-11-26
workflow_status: disabled
stars: 0
---

[![tests](https://github.com/Metadrop/ddev-yellowlabtools/actions/workflows/tests.yml/badge.svg)](https://github.com/Metadrop/ddev-yellowlabtools/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2032.svg)
![GitHub Release](https://img.shields.io/github/v/release/Metadrop/ddev-yellowlabtools)


# ddev-yellowlabtools

A DDEV add-on that provides [YellowLabTools](https://github.com/YellowLabTools/YellowLabTools) for web performance analysis. Includes both CLI and Web UI.

## Installation

```bash
ddev add-on get Metadrop/ddev-yellowlabtools
ddev restart
```

## Usage

### CLI

Analyze any URL:

```bash
ddev yellowlabtools https://example.com
```

Analyze your local DDEV site (this will use the default ddev url):

```bash
ddev yellowlabtools local
```

### Web UI

Access the web interface at:

```
https://<projectname>.ddev.site:9034
```

The Web UI provides a visual interface to run analyses and view detailed reports.

### CLI Options

| Option | Description |
|--------|-------------|
| `--device=<device>` | Device to simulate: `phone`, `tablet` (default: desktop) |
| `--screenshot=<path>` | Save screenshot to path (must end with .png) |
| `--proxy=<host:port>` | Use HTTP proxy |
| `--cookie=<value>` | Set cookie (format: `name=value; domain=.example.com`) |
| `--auth-user=<user>` | HTTP basic auth username |
| `--auth-pass=<pass>` | HTTP basic auth password |
| `--reporter=<format>` | Output format: `json` (default), `xml` |

### Examples

```bash
# Analyze local site with mobile device simulation
ddev yellowlabtools local --device=phone

# Analyze with HTTP basic authentication
ddev yellowlabtools https://staging.example.com --auth-user=admin --auth-pass=secret

# Save output to file
ddev yellowlabtools local > report.json
```

### Output

By default, YellowLabTools outputs JSON to stdout. You can redirect this to a file:

```bash
ddev yellowlabtools https://example.com > analysis.json
```

Or use XML format:

```bash
ddev yellowlabtools https://example.com --reporter=xml > analysis.xml
```

## What is YellowLabTools?

YellowLabTools is an open-source tool that analyzes web page performance and provides detailed reports on:

- Page weight and requests
- JavaScript execution time
- DOM complexity
- Bad practices detection
- Performance scores and recommendations

## Architecture

This add-on creates a dedicated `yellowlabtools` container based on [ousamabenyounes/yellowlabtools](https://github.com/ousamabenyounes/docker-yellowlabtools) with:

- Alpine Linux with Node.js 18
- Chromium browser for headless rendering
- YellowLabTools server (Web UI exposed on port 9034)
- YellowLabTools CLI

The container runs alongside your DDEV services and shares the same network, allowing it to analyze both external URLs and your local DDEV site.

## Requirements

- DDEV v1.23.0 or higher

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

Apache License 2.0
