---
title: jcandan/ddev-ai-agent
github_url: https://github.com/jcandan/ddev-ai-agent
description: "A DDEV add-on to spin up a simple, opinionated, customizable n8n/supabase AI Agent workflow stack with sensible defaults."
user: jcandan
repo: ddev-ai-agent
repo_id: 1089834277
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-11-04
updated_at: 2025-11-24
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/jcandan/ddev-ai-agent/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/jcandan/ddev-ai-agent/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/jcandan/ddev-ai-agent)](https://github.com/jcandan/ddev-ai-agent/commits)
[![release](https://img.shields.io/github/v/release/jcandan/ddev-ai-agent)](https://github.com/jcandan/ddev-ai-agent/releases/latest)

# DDEV AI Agent

## Overview

A [DDEV](https://ddev.com) add-on to spin up a simple, opinionated, customizable
n8n/supabase AI Agent workflow stack with sensible defaults.

## Getting Started

Create a DDEV project.

```bash
mkdir my-ai-agent
cd my-ai-agent
ddev config --project-name=my-ai-agent --project-type=generic --docroot=.
```

Then, get this add-on and start the project.

```bash
ddev add-on get jcandan/ddev-ai-agent
ddev restart
ddev launch :5678
```

After launch, you may commit the `.ddev` directory to version control.

## Requirements

* DDEV ≥ 1.24.3 (Uses Docker)

[Get Started](https://ddev.com/get-started/)

Install and start using DDEV.

> NOTE: DDEV touts itself as a Docker-based PHP development environment, but it
> really is a docker-compose abstraction that affords us the ability to simplify
> providing you with a customizable AI Agent development environment with
> sensible, opinionated defaults. If you're not interested in PHP, you can
> proceed to use this project undaunted.

## Usage

| Command                       | Description                 |
| ----------------------------- | --------------------------- |
| `ddev start` / `ddev restart` | Launch or rebuild the stack |
| `ddev describe`               | View services and URLs      |
| `ddev logs -s n8n`            | Tail n8n logs               |
| `ddev logs -s supabase-db`    | Tail Postgres logs          |

Access **n8n** at `https://<project>.ddev.site:5678`

### Optional local LLM (Ollama)

This add-on can optionally start a local Ollama LLM service using a DDEV profile.

Start the stack with the `ollama` profile:

```bash
ddev start --profiles ollama
```

From other services in the stack (for example, n8n), you can reach the local LLM
at:

http://ollama:11434

Use that URL as the base for your Ollama / OpenAI-compatible nodes in n8n when
the profile is enabled.

You can check the service connection with:

```bash
ddev exec curl http://ollama:11434/api/version
```

The service will be picked up by subsequent `ddev restart`, but to stop the
`ollama` service, run `ddev stop` and `ddev start` again without the profile.

## Roadmap

See [Roadmap](https://github.com/jcandan/ddev-ai-agent/wiki/Roadmap) for details
about our plans for the future.

## Credits

**Contributed and maintained by [@jcandan](https://github.com/jcandan)**

Inspired by [@coleam00](https://github.com/coleam00)'s [local-ai-packaged](https://github.com/coleam00/local-ai-packaged).
