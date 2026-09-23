---
title: "maxwebgr/ddev-zed"
github_url: "https://github.com/maxwebgr/ddev-zed"
description: "Zed editor integration for DDEV projects: tasks, Xdebug listener, ddev zed command, optional ddev-mcp"
user: "maxwebgr"
repo: "ddev-zed"
repo_id: 1382197667
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ">= v1.25.4"
dependencies: []
type: "contrib"
created_at: "2026-09-22"
updated_at: "2026-09-22"
workflow_status: "unknown"
stars: 0
---

[![tests](https://github.com/maxwebgr/ddev-zed/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/maxwebgr/ddev-zed/actions/workflows/tests.yml?query=branch%3Amain)

# DDEV Zed

Zed editor integration for DDEV projects: tasks for everyday `ddev` commands, a ready-to-use Xdebug listener, and a `ddev zed` command to open the project.

## Install

```bash
ddev add-on get maxwebgr/ddev-zed
```

To also connect Zed's Agent Panel to DDEV, opt in to the MCP server (see [DDEV MCP server](#ddev-mcp-server-optional)):

```bash
DDEV_ZED_MCP=true ddev add-on get maxwebgr/ddev-zed
```

## What it installs

| File | Purpose |
|---|---|
| `.zed/tasks.json` | start, stop, restart, describe, launch, mailpit, ssh, logs, xdebug toggle/diagnose, composer install, snapshot |
| `.zed/debug.json` | "DDEV: Listen for Xdebug" on port 9003, `/var/www/html` mapped to `$ZED_WORKTREE_ROOT` |
| `.ddev/commands/host/zed` | `ddev zed` opens the project in Zed; extra arguments are passed through, e.g. `ddev zed -n` for a new window |
| `.zed/settings.json` | optional: registers the [`ddev-mcp`](https://www.npmjs.com/package/ddev-mcp) context server for the Agent Panel (only with `DDEV_ZED_MCP=true`) |

Canonical copies live in `.ddev/zed/`.

## DDEV MCP server (optional)

Setting `DDEV_ZED_MCP=true` during install writes `.zed/settings.json` with a `context_servers` entry that runs [`ddev-mcp`](https://www.npmjs.com/package/ddev-mcp) ([source](https://github.com/codingsasi/ddev-mcp)) via `npx -y ddev-mcp`. Nothing is installed at that moment: `npx` downloads the package the first time Zed starts the server, so Node.js 20+ must be available on your `PATH`.

`ddev-mcp` is a [Model Context Protocol](https://modelcontextprotocol.io/) server that exposes DDEV as tools. Once Zed loads it, the Agent Panel can:

- start, stop, restart the project and read `ddev describe` output
- read container logs
- import, export and snapshot the database
- run any command inside the web container through `ddev_exec` (composer, drush, wp-cli, artisan, phpunit, ...)

so you can ask the agent things like "restart ddev and run the migrations" instead of switching to a terminal. The server runs commands in the project the Agent Panel is open in.

This works with Zed's built-in agent and with external agents such as Claude Code running inside Zed. Zed forwards its configured context servers to external agents, so a Claude Code thread started from the Agent Panel sees the same `ddev_*` tools. The built-in agent needs a model first: Zed's free plan doesn't include one, so add an API key under Configure Providers, use a local Ollama model, or use an external agent with its own account.

Tool calls are subject to the Agent Panel's normal confirmation flow. `ddev-mcp` also blocks commands it classifies as dangerous (for example destructive Platform.sh or database operations) unless `ALLOW_DANGEROUS_COMMANDS=true` is set in the server's `env` block in `.zed/settings.json`.

If `.zed/settings.json` already exists and has no `#ddev-generated` marker, the add-on leaves it alone and prints a message. Copy the `context_servers` block from `.ddev/zed/settings.json` into your file manually. If you skip the opt-in, `.zed/settings.json` is not created and the Agent Panel is unaffected.

## Ownership

Files containing `#ddev-generated` belong to the add-on and are updated on reinstall. Delete that line to take ownership; the add-on will then skip the file on install and keep it on removal. When it skips a file, it lists the template entries your file doesn't have (task and debug labels, or the `ddev-mcp` server) so you can merge them by hand from `.ddev/zed/`.

## Debugging

1. Start "DDEV: Listen for Xdebug" from the debug panel.
2. Run the `ddev: xdebug toggle` task.
3. Load the page.

The listener binds to `0.0.0.0:9003` so the web container can reach it, and maps `/var/www/html` to `$ZED_WORKTREE_ROOT`, the [task variable](https://zed.dev/docs/tasks#variables) Zed resolves to the project root. If breakpoints don't trigger, see the [FAQ](https://github.com/maxwebgr/ddev-zed/blob/main/FAQ.md): the usual causes are a host firewall blocking port 9003 or, rarely, the variable not resolving.

## Keybindings

Zed's `keymap.json` is global, so the add-on doesn't touch it. Example:

```json
[
  {
    "context": "Workspace",
    "bindings": {
      "ctrl-alt-d s": ["task::Spawn", { "task_name": "ddev: start" }],
      "ctrl-alt-d x": ["task::Spawn", { "task_name": "ddev: xdebug toggle" }],
      "ctrl-alt-d h": ["task::Spawn", { "task_name": "ddev: ssh" }]
    }
  }
]
```

## FAQ

See [FAQ.md](https://github.com/maxwebgr/ddev-zed/blob/main/FAQ.md) for firewall setup, path mapping fallbacks, and merging into user-owned files.

## Remove

```bash
ddev add-on remove zed
```

**Contributed and maintained by [@maxwebgr](https://github.com/maxwebgr)**
