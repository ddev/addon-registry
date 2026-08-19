---
title: mxr576/ddev-pi
github_url: https://github.com/mxr576/ddev-pi
description: "DDEV addon for pi.dev, which is a minimal terminal coding harness."
user: mxr576
repo: ddev-pi
repo_id: 1250983207
default_branch: main
tag_name: 1.0.0-beta5
ddev_version_constraint: ">= v1.25.2"
dependencies: ["trebormc/ddev-ai-ssh"]
type: contrib
created_at: 2026-05-27
updated_at: 2026-08-18
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/mxr576/ddev-pi/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/mxr576/ddev-pi/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/mxr576/ddev-pi)](https://github.com/mxr576/ddev-pi/commits)
[![release](https://img.shields.io/github/v/release/mxr576/ddev-pi)](https://github.com/mxr576/ddev-pi/releases/latest)

# DDEV add-on for Pi Coding Agent

This DDEV add-on provides a fully containerized, isolated environment for the [Pi Coding Agent](https://pi.dev).

## Architecture & security

- **Full isolation:** The agent runs in a dedicated sidecar container. The host machine's `~/.pi` directory is explicitly *not* mounted, so agent state is scoped strictly to the project and avoids cross-contamination with your host environment.
- **Project access:** The project root is mounted at `/var/www/html` with read/write permissions, allowing Pi to interact with your codebase.
- **Network:** The agent is isolated within the DDEV project network.
- **Sensitive data:** No host-level sensitive directories (e.g., `~/.ssh`) are mounted.
- **Offline by default:** Network-dependent features (`PI_OFFLINE=1`, `PI_SKIP_VERSION_CHECK=1`, `PI_TELEMETRY=0`) are disabled by default to keep the container self-contained and avoid unexpected outbound traffic.

For a detailed diagram and breakdown of the system's components and security boundaries, please see the [Architecture document](https://github.com/mxr576/ddev-pi/blob/main/docs/architecture.md).

> [!WARNING]
> While process isolation is provided by the container, running coding agents on a local read-write workspace carries inherent security risks (such as prompt injection, credential leaks and host-execution escapes). Please see the [Security Advisory](https://github.com/mxr576/ddev-pi/blob/main/docs/SECURITY.md) to understand these risks and learn how to harden your setup.

## Requirements

- DDEV >= v1.25.2
- The [`trebormc/ddev-ai-ssh`](https://github.com/trebormc/ddev-ai-ssh) add-on (installed automatically as a dependency)
- **Optional:** `python3` installed on the host machine (only required if using the experimental host clipboard integration)

## Installation

```bash
ddev add-on get mxr576/ddev-pi
ddev restart
```

## Usage

The Pi agent runs as a DDEV service under the `pi` profile. You must start DDEV with that profile active.

1. Start DDEV with the Pi profile enabled:
   ```bash
   ddev start --profiles=pi
   ```
2. Run Pi commands via the `ddev pi` wrapper:
   ```bash
   ddev pi [command]
   ```

Example:

```bash
ddev pi status
```

> **Note:** The `ddev pi` command is a thin wrapper that runs `ddev exec --service=pi pi "$@"` inside the sidecar container. The container's working directory is `/var/www/html` (your project root).

## Configuration

Default configurations are initialized in your project's `.pi/` directory during installation (`settings.json` is copied from the add-on defaults). Modify `settings.json` there to configure providers and agent preferences.

> **Note:** Re-running `ddev add-on get` will not overwrite an existing `.pi/` directory. Defaults are only copied when `.pi/` is empty.

## Extensibility

This add-on provides several ways to customize the container environment, manage Pi extensions, and easily install AI skills:

- To extend or customize the container environment and DDEV hooks, see the [Extensibility document](https://github.com/mxr576/ddev-pi/blob/main/docs/extensions.md).
- To easily install, manage, and configure AI skills, see the [AI Skills document](https://github.com/mxr576/ddev-pi/blob/main/docs/skills.md).

## Experimental Host Clipboard Integration

This add-on features an experimental cross-platform clipboard integration that securely bridges the containerized Pi Coding Agent to your host system's clipboard.

When you use the `/copy` command (or when the agent writes to the clipboard), the content is written to a temporary pending file inside the shared workspace volume (`.ddev/pi/.clipboard_pending`). A lightweight, background Python daemon (`pi/clipboard-helper.py`) running on your host monitors this file, reads and copies its contents to your host's clipboard using native system utilities, and then immediately deletes the file.

### OS Support & Feedback

Since clipboard tools vary depending on your host OS and available terminal utilities, this feature is currently marked as **experimental**. We welcome your feedback, bugs, and success stories for different operating systems!

- **macOS:** Fully supported and tested using native `pbcopy`.
- **Windows (WSL/native):** Fully supported and tested using `clip.exe` / PowerShell.
- **Linux (X11/Wayland):** Fully supported and tested using `wl-copy`, `xclip`, or `xsel`.

*If clipboard sharing is not working on your host, check the log file at `.ddev/pi/clipboard.log` for debugging details.*

## Uninstalling

```bash
ddev addon remove pi
```

The `.pi/` folder in your project root and the `pi-agent` Docker volume are **not** removed automatically. Clean them up manually if needed:

```bash
# Remove the project Pi configuration
rm -rf .pi/

# Remove the named Docker volume
docker volume rm ddev-<your-project-name>-pi-agent
```

## Credits

Contributed and maintained by [@mxr576](https://github.com/mxr576). The initial proof of concept and early development were sponsored by [Pronovix](https://pronovix.com).
