---
title: credevator/ddev-litellm
github_url: https://github.com/credevator/ddev-litellm
description: "LiteLLM proxy service accessible inside DDEV."
user: credevator
repo: ddev-litellm
repo_id: 1181050593
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-03-13
updated_at: 2026-03-14
workflow_status: failure
stars: 2
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/credevator/ddev-litellm/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/credevator/ddev-litellm/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/credevator/ddev-litellm)](https://github.com/credevator/ddev-litellm/commits)
[![release](https://img.shields.io/github/v/release/credevator/ddev-litellm)](https://github.com/credevator/ddev-litellm/releases/latest)


# ddev-litellm

A [DDEV](https://ddev.com) add-on that runs a [LiteLLM](https://www.litellm.ai/) proxy as a DDEV service, enabling local AI model testing with Drupal's [`ai_provider_litellm`](https://www.drupal.org/project/ai_provider_litellm) module.

## Services

The add-on starts two containers via a single `docker-compose.litellm.yaml`:

| Container | Purpose |
|---|---|
| `ddev-<project>-litellm` | LiteLLM proxy (port 4000) |
| `ddev-<project>-litellm-postgres` | PostgreSQL database required by LiteLLM for key management |

LiteLLM uses [Prisma](https://www.prisma.io/) as its ORM and requires PostgreSQL — key management endpoints (`/key/info`, `/key/generate`) will not work without it.

## Features

- LiteLLM proxy accessible inside DDEV at `http://ddev-<project>-litellm:4000`
- PostgreSQL sidecar started and health-checked automatically before LiteLLM
- Pre-configured to connect to [Ollama](https://ollama.ai) running on the host machine
- Optional vLLM / HuggingFace endpoint support
- On each `ddev start`: waits for LiteLLM readiness, creates a Drupal virtual key in the DB, and auto-sets `ai_provider_litellm.settings.host` via drush

## Requirements

- DDEV ≥ 1.24.10
- [Ollama](https://ollama.ai) installed on your host machine (for Ollama backend)
- At least one Ollama model pulled: `ollama pull llama3.2`

## Installation

```shell
ddev add-on get credevator/ddev-litellm
ddev restart
```

## Usage

### Service URLs

| Context | URL |
|---|---|
| Drupal web container | `http://ddev-<project>-litellm:4000` |
| Host browser | `https://<project>.ddev.site:4001` |

### Commands

```shell
ddev litellm              # Show service status
ddev litellm open         # Open LiteLLM UI in browser
ddev litellm logs         # View service logs
ddev litellm-models       # List available models
```

### Configuring Models

Edit `.ddev/litellm_config.yaml` to add or modify model backends, then run `ddev restart`.

**Ollama** (default, connects to host machine port 11434):
```yaml
- model_name: ollama/llama3.2
  litellm_params:
    model: ollama/llama3.2
    api_base: http://host.docker.internal:11434
```

**vLLM / HuggingFace** (uncomment in `litellm_config.yaml` and set `VLLM_API_BASE`):
```yaml
# In .ddev/config.yaml:
# web_environment:
#   - VLLM_API_BASE=http://host.docker.internal:8000
```

### Configuring Drupal

On each `ddev start` the add-on automatically:
1. Waits for the LiteLLM proxy to be ready
2. Creates a virtual API key `sk-drupal-dev-key` in the LiteLLM database (safe to run repeatedly — no-op if it already exists)
3. Sets `ai_provider_litellm.settings.host` to `http://ddev-<project>-litellm:4000` via drush

You still need to wire up the API key in Drupal once:
1. Go to **Configuration → Key** (`/admin/config/system/keys`)
2. Add a key named `litellm_key` with value `sk-drupal-dev-key`
3. Go to **Configuration → AI → Providers** (`/admin/config/ai/providers/litellm`)
4. Select `litellm_key` under **API Key**

### Keys: master vs. virtual

| Key | Value | Purpose |
|---|---|---|
| Master key | `sk-ddev-litellm` | LiteLLM admin — used to generate virtual keys and call management endpoints |
| Drupal virtual key | `sk-drupal-dev-key` | Used by Drupal for all AI requests — stored in the PostgreSQL DB so `/key/info` resolves it |

To override the master key, add to `.ddev/config.yaml`:

```yaml
web_environment:
  - LITELLM_MASTER_KEY=sk-your-custom-key
```

Note: if you change the master key after first start, update the `key/generate` call in `.ddev/config.ddev-litellm.yaml` to match.

## Troubleshooting

**LiteLLM takes a while to start** — the image is ~2GB. PostgreSQL must be healthy first, then LiteLLM runs Prisma migrations before serving requests. The healthcheck allows up to 5 minutes (`30s interval × 10 retries + 120s start_period`). Check progress with `ddev litellm logs`.

**Cannot reach Ollama** — ensure Ollama is running on the host (`ollama serve`) and the model is pulled (`ollama pull llama3.2`).

**Port 4000 or 4001 conflict** — edit `HTTP_EXPOSE` / `HTTPS_EXPOSE` in `.ddev/docker-compose.litellm.yaml`.

**Linux users** — `host.docker.internal` is injected via `extra_hosts: host.docker.internal:host-gateway`. This is handled automatically; macOS and Windows Docker Desktop provide it natively.

**PostgreSQL data** — key data persists in a Docker named volume (`litellm-postgres-data`). If you need a clean slate: `docker volume rm <project>_litellm-postgres-data` then `ddev restart`.

## Uninstalling

```shell
ddev add-on remove ddev-litellm
ddev restart
```

Note: `.ddev/litellm_config.yaml` is preserved on removal. Delete it manually if no longer needed. The `litellm-postgres-data` Docker volume is also preserved — remove it manually with `docker volume rm <project>_litellm-postgres-data`.

## License

Apache 2.0
