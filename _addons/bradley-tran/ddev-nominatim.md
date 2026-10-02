---
title: "bradley-tran/ddev-nominatim"
github_url: "https://github.com/bradley-tran/ddev-nominatim"
description: "DDEV add-on for OpenStreetMap's Nominatim search engine"
user: "bradley-tran"
repo: "ddev-nominatim"
repo_id: 1375557109
default_branch: "main"
tag_name: "v0.2"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-18"
updated_at: "2026-09-19"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/bradley-tran/ddev-nominatim/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/bradley-tran/ddev-nominatim/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/bradley-tran/ddev-nominatim)](https://github.com/bradley-tran/ddev-nominatim/commits)
[![release](https://img.shields.io/github/v/release/bradley-tran/ddev-nominatim)](https://github.com/bradley-tran/ddev-nominatim/releases/latest)

# DDEV Nominatim

## Overview

This add-on integrates [Nominatim](https://nominatim.org/) (OpenStreetMap search and reverse geocoding engine) into your [DDEV](https://ddev.com/) project using the [`mediagis/nominatim`](https://hub.docker.com/r/mediagis/nominatim) Docker image, along with [Nominatim UI](https://github.com/osm-search/nominatim-ui) for interactive map searching, reverse geocoding, and database inspection.

By default, the add-on downloads and imports the Monaco extract from Geofabrik for quick testing and development. You can easily configure it to import any regional extract or custom OSM data.

## Requirements

- DDEV v1.24.10 or higher

## Installation

```bash
ddev add-on get bradley-tran/ddev-nominatim
ddev restart
```

> [!NOTE]
> Nominatim import can take a really long time, so on the first start, Nominatim downloads and imports the OpenStreetMap extract in the background without blocking `ddev start` or other services. You can follow the import progress using `ddev logs -s nominatim -f`. Once import completes, the API endpoints will automatically become responsive.

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

### Accessing Nominatim

- **Web UI (Host Browser):**
  - HTTP: `http://<projectname>.ddev.site:8765`
  - HTTPS: `https://<projectname>.ddev.site:8744`
  - Quick launch: `ddev nominatim-ui`
- **API (Host Browser / External Tools):**
  - HTTP: `http://<projectname>.ddev.site:8980`
  - HTTPS: `https://<projectname>.ddev.site:8943`
- **Internal (From `web` container or other services):**
  - `http://nominatim:8080`

### API Endpoints

| Endpoint | Description | Example |
| -------- | ----------- | ------- |
| `/status` | Service and database health | `http://<projectname>.ddev.site:8980/status` |
| `/search` | Search by address / query | `http://<projectname>.ddev.site:8980/search?q=avenue+pasteur&format=json` |
| `/reverse` | Reverse geocoding (lat/lon) | `http://<projectname>.ddev.site:8980/reverse?lat=43.7384&lon=7.4246&format=json` |

### Useful Commands

| Command | Description |
| ------- | ----------- |
| `ddev nominatim-ui` | Open the Nominatim Web UI in your browser |
| `ddev nominatim <cmd>` | Run Nominatim CLI commands inside the container |
| `ddev describe` | View service status and exposed ports |
| `ddev logs -s nominatim` | View Nominatim container logs |
| `ddev logs -s nominatim -f` | Follow live import and request logs |
| `ddev logs -s nominatim-ui` | View Nominatim UI web server logs |

### Nominatim CLI

This add-on exposes the `nominatim` command to run CLI commands directly inside the Nominatim container:

```bash
# Check service and database status
ddev nominatim status

# Check database health and consistency
ddev nominatim admin --check-database

# Warm database cache
ddev nominatim admin --warm

# Show Nominatim version
ddev nominatim --version

# View all available CLI commands and help
ddev nominatim --help
```

## Configuration

Configuration is managed via `.ddev/.env.nominatim`. You can configure settings using `ddev dotenv set`:

```bash
# Example 1: Import Germany from URL instead of Monaco
ddev dotenv set .ddev/.env.nominatim --nominatim-pbf-url="https://download.geofabrik.de/europe/germany-latest.osm.pbf"
ddev restart

# Example 2: Import from a local .osm.pbf file (placed in project root or .ddev/)
ddev dotenv set .ddev/.env.nominatim --nominatim-pbf-path="data.osm.pbf"
ddev restart
```

### Available Options

| Variable | Flag | Default | Description |
| -------- | ---- | ------- | ----------- |
| `NOMINATIM_DOCKER_IMAGE` | `--nominatim-docker-image` | `mediagis/nominatim:5.3` | Nominatim Docker image and tag |
| `NOMINATIM_PBF_URL` | `--nominatim-pbf-url` | Monaco extract URL | URL of the `.osm.pbf` file to download and import |
| `NOMINATIM_PBF_PATH` | `--nominatim-pbf-path` | _(empty)_ | Path to a local `.osm.pbf` file (e.g. `data.osm.pbf`, `/mnt/ddev_config/data.osm.pbf`, or `/var/www/html/data.osm.pbf`) |
| `NOMINATIM_IMPORT_STYLE` | `--nominatim-import-style` | `full` | Import detail level: `admin`, `street`, `address`, or `full` |
| `NOMINATIM_PASSWORD` | `--nominatim-password` | `nominatim` | PostgreSQL password for nominatim user |
| `NOMINATIM_THREADS` | `--nominatim-threads` | `2` | Number of threads used during import |
| `NOMINATIM_REPLICATION_URL` | `--nominatim-replication-url` | _(empty)_ | Base URL for live updates from Geofabrik |
| `NOMINATIM_IMPORT_WIKIPEDIA` | `--nominatim-import-wikipedia` | `false` | Import Wikipedia importance dumps for improved ranking |
| `NOMINATIM_HTTP_PORT` | `--nominatim-http-port` | `8980` | Host HTTP port exposed via DDEV router |
| `NOMINATIM_HTTPS_PORT` | `--nominatim-https-port` | `8943` | Host HTTPS port exposed via DDEV router |
| `NOMINATIM_UI_HTTP_PORT` | `--nominatim-ui-http-port` | `8765` | Host HTTP port for Web UI exposed via DDEV router |
| `NOMINATIM_UI_HTTPS_PORT` | `--nominatim-ui-https-port` | `8744` | Host HTTPS port for Web UI exposed via DDEV router |
| `NOMINATIM_UI_VERSION` | `--nominatim-ui-version` | `3.12.0` | Release version of `osm-search/nominatim-ui` |
| `NOMINATIM_UI_PAGE_TITLE` | `--nominatim-ui-page-title` | `Nominatim` | Page title displayed in the Web UI |
| `NOMINATIM_UI_DOCKER_IMAGE` | `--nominatim-ui-docker-image` | `nginx:alpine` | Docker image used to serve Nominatim UI |

### Customizing Nominatim UI

You can customize the UI by creating files in `.ddev/nominatim-ui/`:
- **Theme Configuration:** Create `.ddev/nominatim-ui/config.theme.js` to override frontend settings (e.g., default zoom, map center, tiles).
- **Web Server Configuration:** Create `.ddev/nominatim-ui/nginx.conf` to provide a custom Nginx configuration.
After creating or editing these files, restart DDEV with `ddev restart`.

## Data Persistence & Changing Data Extracts

The PostgreSQL database is persisted in a named Docker volume (`ddev-<projectname>-nominatim-data`).

If you change `NOMINATIM_PBF_URL` / `NOMINATIM_PBF_PATH` or want to re-import data from scratch:

1. Stop DDEV:
   ```bash
   ddev stop
   ```
2. Remove the existing database volume:
   ```bash
   docker volume rm ddev-${DDEV_SITENAME}-nominatim-data
   ```
3. Update your `.ddev/.env.nominatim` configuration:
   ```bash
   # Using a URL:
   ddev dotenv set .ddev/.env.nominatim --nominatim-pbf-url="https://download.geofabrik.de/europe/liechtenstein-latest.osm.pbf"
   # Or using a local file:
   ddev dotenv set .ddev/.env.nominatim --nominatim-pbf-path="data.osm.pbf"
   ```
4. Start DDEV (starts immediately; Nominatim imports in the background):
   ```bash
   ddev start
   ```

## Hardware Considerations

- **Small extracts (Monaco, Liechtenstein):** ~2 GB RAM, 10–20 GB disk, finishes in minutes.
- **Medium extracts (countries like Germany, France):** 16–32 GB RAM, 50–100 GB fast NVMe disk, can take a few hours.
- **Full Planet:** 64–128 GB+ RAM, 1 TB+ NVMe SSD, takes multiple days.

## Credits

**Contributed and maintained by [@bradley-tran](https://github.com/bradley-tran)**
