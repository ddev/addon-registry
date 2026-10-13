---
title: "brookemahoney/ddev-assistant-ollama"
github_url: "https://github.com/brookemahoney/ddev-assistant-ollama"
description: "A DDEV add-on to install and run Ollama inside the web container."
user: "brookemahoney"
repo: "ddev-assistant-ollama"
repo_id: 1389878295
default_branch: "main"
tag_name: "1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-26"
updated_at: "2026-09-27"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/brookemahoney/ddev-assistant-ollama/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/brookemahoney/ddev-assistant-ollama/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/brookemahoney/ddev-assistant-ollama)](https://github.com/brookemahoney/ddev-assistant-ollama/commits)
[![release](https://img.shields.io/github/v/release/brookemahoney/ddev-assistant-ollama)](https://github.com/brookemahoney/ddev-assistant-ollama/releases/latest)

# DDEV Assistant Ollama

## Overview

This add-on installs [Ollama](https://ollama.com) inside the web container of your [DDEV](https://ddev.com/) project, so local models are available to your code with no extra setup.

The `ollama` command works in `ddev ssh` and `ddev exec`, and the Ollama API answers on `http://localhost:11434` inside the web container. Models are downloaded once per machine and shared by every project that uses this add-on.

## Requirements

- DDEV >= v1.24.10

## Installation

```bash
ddev add-on get brookemahoney/ddev-assistant-ollama
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev exec ollama run <model>` | Run a model, downloading it first if needed |
| `ddev exec ollama list` | List downloaded models |
| `ddev exec ddev-ollama-start` | Start the server, if it is not already running |
| `ddev describe` | View the Ollama endpoint in the service table |

From inside the web container, the API is at `http://localhost:11434`, or `http://web:11434` from another container in the same project:

```bash
ddev exec ollama run llama3.2
ddev exec curl -s http://localhost:11434/api/version
```

## Configuration

The add-on sets two environment variables, and only as defaults, so a project can override either of them with `web_environment` in `.ddev/config.local.yaml`:

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `OLLAMA_HOST` | `0.0.0.0:11434` | Where the server listens. Point it at a non-local address, such as `host.docker.internal:11434`, to use an Ollama that already runs on the host; the add-on then does not start a server of its own. |
| `OLLAMA_MODELS` | `/mnt/ddev-global-cache/assistant-ollama/models` | Where models are stored. The DDEV global cache is per machine, so a model is downloaded once for all projects. |

```yaml
# .ddev/config.local.yaml
web_environment:
  - OLLAMA_HOST=127.0.0.1:11435
```

Server logs are in `/tmp/ollama.log` inside the web container:

```bash
ddev exec tail -f /tmp/ollama.log
```

### Reaching Ollama from the host

The port is not published by default, because the server is meant for the web container. To reach it from the host as well, publish it on the web service:

```yaml
# .ddev/docker-compose.web-ports.yaml
services:
  web:
    ports:
      - "11434"
```

`ddev describe` then shows the host port to use.

## How it works

- **Installs Ollama into the web image** at `/usr/local/bin/ollama`, with its runtime in `/usr/local/lib/ollama`, so it is on `$PATH` in every shell type and survives `ddev restart` without a per-start copy
- **Starts the server on start** with a `post-start` hook that waits until the API answers, so Ollama is usable as soon as `ddev start` returns. The hook is idempotent, and it fails loudly rather than leaving you with no server
- **Keeps models out of the image and out of the project**, in the DDEV global cache, which is shared across projects and survives `ddev restart`, `ddev rm` and `ddev add-on remove`

## Notes

- The CPU-only Ollama build is installed, so no GPU is used. A web container has no GPU passed through to it, including on a Linux host with an NVIDIA card. Use a sidecar add-on such as [tyler36/ddev-ollama](https://github.com/tyler36/ddev-ollama) if you need GPU acceleration.
- Do not combine this add-on with a sidecar Ollama add-on; two Ollamas in one project will fight over the same port.
- The runtime comes from the [`alpine/ollama`](https://github.com/alpine-docker/ollama) image, which carries the CPU-only runtime in 68 MB rather than the ~1.5 GB of the official Linux release, most of which is GPU libraries. It tracks Ollama releases a week or two behind; `ddev exec ollama --version` shows the installed version.
- `ddev add-on remove brookemahoney/ddev-assistant-ollama` removes the add-on's files, and the next `ddev restart` rebuilds the web image without Ollama. Downloaded models stay in the global cache; remove them with `ddev exec sudo rm -rf /mnt/ddev-global-cache/assistant-ollama`.

## Credits

**Contributed and maintained by [@brookemahoney](https://github.com/brookemahoney)**
