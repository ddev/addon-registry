---
title: "jfeid/ddev-zed"
github_url: "https://github.com/jfeid/ddev-zed"
description: "Zed editor integration for DDEV projects: tasks, Xdebug listener, ddev zed command, optional ddev-mcp"
user: "jfeid"
repo: "ddev-zed"
repo_id: 1382197667
default_branch: "main"
tag_name: "v1.1.2"
ddev_version_constraint: ">= v1.25.4"
dependencies: []
type: "contrib"
created_at: "2026-09-22"
updated_at: "2026-09-28"
workflow_status: "success"
stars: 0
---

[![tests](https://github.com/jfeid/ddev-zed/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/jfeid/ddev-zed/actions/workflows/tests.yml?query=branch%3Amain)

# DDEV Zed

Zed editor integration for DDEV projects: tasks for everyday `ddev` commands, a ready-to-use Xdebug listener, and a `ddev zed` command to open the project.

## Install

```bash
ddev add-on get jfeid/ddev-zed
```

To also connect Zed's Agent Panel to DDEV, opt in to the MCP server (see [DDEV MCP server](#ddev-mcp-server-optional)):

```bash
DDEV_ZED_MCP=true ddev add-on get jfeid/ddev-zed
```

## Requirements

- DDEV v1.25.4 or newer.
- Zed with the **PHP** extension. The extension provides the Xdebug debug adapter; without it the debug picker shows "No matches" for the add-on's config. Its default language server, phpactor, needs `php` on the machine running Zed. DDEV projects usually have PHP only in the container, so either switch the project to intelephense in `.zed/settings.json`:

  ```json
  { "languages": { "PHP": { "language_servers": ["intelephense", "!phpactor"] } } }
  ```

  or install a host PHP (`sudo apt install php-cli`).
- Node.js 20+ with npm on the machine running Zed, only if you opt in to the MCP server. On Ubuntu that means `sudo apt install nodejs npm`, since the `nodejs` package alone has no `npx`.

## What it installs

| File | Purpose |
|---|---|
| `.zed/tasks.json` | start, stop, restart, describe, launch, mailpit, ssh, logs, xdebug toggle/diagnose, composer install, snapshot |
| `.zed/debug.json` | "DDEV: Listen for Xdebug" on port 9003, `/var/www/html` mapped to `$ZED_WORKTREE_ROOT` |
| `.ddev/commands/host/zed` | `ddev zed` opens the project in Zed; extra arguments are passed through, e.g. `ddev zed -n` for a new window |
| `.zed/settings.json` | optional: registers the [`ddev-mcp`](https://www.npmjs.com/package/ddev-mcp) context server for the Agent Panel (opt in with `DDEV_ZED_MCP=true`; stays on for later installs until you delete the file) |

Canonical copies live in `.ddev/zed/`.

## DDEV MCP server (optional)

Setting `DDEV_ZED_MCP=true` during install writes `.zed/settings.json` with a `context_servers` entry that runs [`ddev-mcp`](https://www.npmjs.com/package/ddev-mcp) ([source](https://github.com/codingsasi/ddev-mcp)) via `npx -y ddev-mcp`. Nothing is installed at that moment: `npx` downloads the package the first time Zed starts the server, so Node.js 20+ with npm must be available on your `PATH`.

On Windows the installer writes a different entry, depending on where it runs:

- **Inside WSL:** `wsl.exe -d <distro> --cd <project path> npx -y ddev-mcp`. Zed for Windows starts MCP servers on the Windows side even for projects opened through WSL, where neither `npx` nor `ddev` exists; `wsl.exe` runs the server inside your distro instead. The distro name and project path are filled in at install time, so this is the one entry to regenerate with `ddev add-on get` if you move the project (see the [FAQ](https://github.com/jfeid/ddev-zed/blob/main/FAQ.md#i-moved-or-renamed-the-project)).
- **On Windows itself:** `cmd /c "set PWD=%CD%&& npx -y ddev-mcp"`. `ddev-mcp` takes the project folder from the `PWD` variable and crashes on start without it. Zed starts the server in the project folder, and `cmd` copies that folder into `PWD`. The file holds no machine-specific path.

On Linux and macOS the plain entry needs nothing extra: Zed starts the server in the project folder through `sh -c`, and the shell sets `PWD` itself.

`ddev-mcp` is a [Model Context Protocol](https://modelcontextprotocol.io/) server that exposes DDEV as tools. Once Zed loads it, the Agent Panel can:

- start, stop, restart the project and read `ddev describe` output
- read container logs
- import, export and snapshot the database
- run any command inside the web container through `ddev_exec` (composer, drush, wp-cli, artisan, phpunit, ...)

so you can ask the agent things like "restart ddev and run the migrations" instead of switching to a terminal. The server runs commands in the project the Agent Panel is open in.

This works with Zed's built-in agent and with external agents such as Claude Code running inside Zed. Zed forwards its configured context servers to external agents, so a Claude Code thread started from the Agent Panel sees the same `ddev_*` tools. Exception: for projects opened through WSL, Zed doesn't forward the server to Claude Agent threads; the built-in agent sees it. The [FAQ](https://github.com/jfeid/ddev-zed/blob/main/FAQ.md#is-the-agent-actually-using-ddev-mcp) has a one-file workaround for Claude Agent. The built-in agent needs a model first: Zed's free plan doesn't include one, so add an API key under Configure Providers, use a local Ollama model, or use an external agent with its own account.

Tool calls are subject to the Agent Panel's normal confirmation flow. `ddev-mcp` also blocks commands it classifies as dangerous (for example destructive Platform.sh or database operations) unless `ALLOW_DANGEROUS_COMMANDS=true` is set in the server's `env` block in `.zed/settings.json`.

MCP stays on once enabled: later `ddev add-on get` or `ddev add-on update` runs refresh the generated `.zed/settings.json` even without `DDEV_ZED_MCP=true`, so fixes to the entry reach you. To turn it off, delete `.zed/settings.json` or remove the add-on. If you never opt in, the file isn't created and the Agent Panel is unaffected.

If `.zed/settings.json` already exists and has no `#ddev-generated` marker, the add-on leaves it alone and checks your `ddev-mcp` entry. When the entry is missing, or uses a different command than your platform needs, the installer prints the right one; paste its `"ddev-mcp"` block into your `context_servers`. Don't copy it from `.ddev/zed/settings.json`: that template is the Linux and macOS form, which doesn't start on Windows or under WSL.

## Ownership

Files containing `#ddev-generated` belong to the add-on and are updated on reinstall. Delete that line to take ownership; the add-on will then skip the file on install and keep it on removal. When it skips a file, it prints the entries your file is missing, ready to paste (see the [FAQ](https://github.com/jfeid/ddev-zed/blob/main/FAQ.md#the-add-on-skipped-one-of-my-files)).

## Debugging

1. Start "DDEV: Listen for Xdebug" from the debug panel.
2. Run the `ddev: xdebug toggle` task.
3. Load the page.

The listener binds to `0.0.0.0:9003` so the web container can reach it, and maps `/var/www/html` to `$ZED_WORKTREE_ROOT`, the [task variable](https://zed.dev/docs/tasks#variables) Zed resolves to the project root. If breakpoints don't trigger, see the [FAQ](https://github.com/jfeid/ddev-zed/blob/main/FAQ.md): the usual causes are a host firewall blocking port 9003 or, rarely, the variable not resolving.

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

## Windows

DDEV's Windows installer offers three modes. The add-on was tested on Windows 11 with Zed 1.21 and DDEV v1.25.4 (September 2026) in two of them:

| DDEV mode | Tested |
|---|---|
| Docker CE inside WSL2 (DDEV's recommendation) | yes |
| Docker Desktop / Rancher Desktop with WSL2 | no; expected to behave like the mode above, since `ddev` and the project live in the WSL distro either way |
| Traditional Windows (Docker Desktop, PowerShell or Git Bash) | yes, with Docker Desktop's Hyper-V backend |

**WSL2 modes.** Install and run DDEV inside WSL. Open the project from Zed for Windows with "Open Remote" and pick your WSL distro, or run `ddev zed` from the WSL shell: Zed's installer puts a WSL-aware `zed` wrapper on the Windows `PATH`, which WSL interop exposes. Zed asks you to trust the project on first open. What works:

- `ddev zed` and `ddev zed -n`.
- Tasks. They run in a WSL shell, so `ddev` is found without any PATH changes.
- The MCP server, through the `wsl.exe` entry the installer writes under WSL. Zed's built-in agent lists its tools; Claude Agent threads need a project `.mcp.json` over WSL (see [DDEV MCP server](#ddev-mcp-server-optional)).
- Xdebug, **after one WSL setting**. Out of the box the listener fails with "Connection to TCP DAP timeout", a Zed bug in how it reaches debug adapters inside WSL. Disabling IPv6 in WSL works around it: breakpoints then hit normally. See the [FAQ](https://github.com/jfeid/ddev-zed/blob/main/FAQ.md#the-debugger-fails-with-connection-to-tcp-dap-timeout-wsl2) for the setting and its trade-off.

**Traditional Windows.** The project lives on a Windows drive and Zed opens it as a normal local folder. DDEV runs the add-on's install script and `ddev zed` through Git Bash, so Git for Windows is required, even if you type commands in PowerShell. What works, including from a project path with non-ASCII characters:

- Install, and the skip message for user-owned files.
- `ddev zed` and `ddev zed -n`.
- Tasks. Zed runs them through PowerShell.
- Xdebug, with no extra setup. Windows Defender Firewall may ask to allow Node.js (the debug adapter) the first time; allow it.
- The MCP server, through the `cmd` entry the installer writes on Windows, however Zed is started. Tested with Zed's tool list and a Claude Agent tool call.

## FAQ

See [FAQ.md](https://github.com/jfeid/ddev-zed/blob/main/FAQ.md) for firewall setup, path mapping fallbacks, and merging into user-owned files.

## Remove

```bash
ddev add-on remove zed
```

**Contributed and maintained by [@jfeid](https://github.com/jfeid)**
