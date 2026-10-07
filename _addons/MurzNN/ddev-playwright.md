---
title: "MurzNN/ddev-playwright"
github_url: "https://github.com/MurzNN/ddev-playwright"
description: " DDEV add-on to install browsers, required for Playwright tests, into the web container "
user: "MurzNN"
repo: "ddev-playwright"
repo_id: 1304882212
default_branch: "main"
tag_name: "v1.1.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-07-18"
updated_at: "2026-08-29"
workflow_status: "success"
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/MurzNN/ddev-playwright/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/MurzNN/ddev-playwright/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/MurzNN/ddev-playwright)](https://github.com/MurzNN/ddev-playwright/commits)
[![release](https://img.shields.io/github/v/release/MurzNN/ddev-playwright)](https://github.com/MurzNN/ddev-playwright/releases/latest)

# DDEV Playwright

## Overview

This add-on installs [Playwright](https://playwright.dev/) and browsers (Chromium, Firefox, WebKit) in the
[DDEV](https://ddev.com/) `web` container so you can run tests against the project's internal network and
dependencies. It also configures X11 forwarding for headed mode (visible browser windows on the host).

## Installation

```bash
ddev add-on get MurzNN/ddev-playwright
```

## What it does

- Installs Playwright system dependencies during the `web` container build (`web-build/Dockerfile.playwright`).
- Runs `npx playwright install` after each container start to keep browsers up to date.
- Mounts a shared host browser cache at `~/.cache/ms-playwright` so multiple DDEV projects reuse the same downloads.
- Configures X11 forwarding for GUI apps launched inside the container (Playwright headed mode, Chrome, etc.).

## X11 / headed mode

By default the add-on:

1. Runs `xhost +local:docker` on the host before start.
2. Bind-mounts `/tmp/.X11-unix` into the `web` container.
3. Sets `DISPLAY=:0` in the container.
4. Sets `LIBGL_ALWAYS_SOFTWARE=1` so Chromium/Mesa do not hang trying to use GPU/DRM inside Docker when no window appears.

### macOS / rootless Docker alternative

When bind-mounting `/tmp/.X11-unix` is not possible (for example rootless Docker on macOS), use the commented
alternative in `.ddev/config.playwright.yaml`:

```yaml
web_environment:
  - DISPLAY=host.docker.internal:0
```

You still need `xhost +local:docker` (or equivalent) on the host so Docker can connect to the X server.

## Credits

**Contributed and maintained by [@MurzNN](https://github.com/MurzNN)**
