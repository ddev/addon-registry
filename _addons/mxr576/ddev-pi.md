---
title: mxr576/ddev-pi
github_url: https://github.com/mxr576/ddev-pi
description: "DDEV addon for pi.dev, which is a minimal terminal coding harness."
user: mxr576
repo: ddev-pi
repo_id: 1250983207
default_branch: main
tag_name: 1.0.0-alpha3
ddev_version_constraint: ">= v1.25.2"
dependencies: ["trebormc/ddev-ai-ssh"]
type: contrib
created_at: 2026-05-27
updated_at: 2026-06-11
workflow_status: success
stars: 0
---

# DDEV add-on for Pi Coding Agent

This DDEV add-on provides a fully containerized, isolated environment for the [Pi Coding Agent](https://pi.dev).

## Architecture & security

- **Full isolation:** The agent runs in a dedicated sidecar container. The host machine's `~/.pi` directory is explicitly *not* mounted, so agent state is scoped strictly to the project and avoids cross-contamination with your host environment.
- **Project access:** The project root is mounted at `/var/www/html` with read/write permissions, allowing Pi to interact with your codebase.
- **Network:** The agent is isolated within the DDEV project network.
- **Sensitive data:** No host-level sensitive directories (e.g., `~/.ssh`) are mounted.
- **Offline by default:** Network-dependent features (`PI_OFFLINE=1`, `PI_SKIP_VERSION_CHECK=1`, `PI_TELEMETRY=0`) are disabled by default to keep the container self-contained and avoid unexpected outbound traffic.

## Requirements

- DDEV >= v1.25.2
- The [`trebormc/ddev-ai-ssh`](https://github.com/trebormc/ddev-ai-ssh) add-on (installed automatically as a dependency)

## Installation

```bash
ddev addon get mxr576/ddev-pi
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

## Managing extensions

Pi supports two types of extensions:

- **Global extensions:** Install with `ddev pi install <addon-ref>`. These are stored in `.ddev/pi/global/` and are local to your environment. Do not commit this directory to VCS.
- **Project extensions:** Install with `ddev pi install -l <addon-ref>`. These are stored in the project's `.pi/` directory. **Commit this directory to VCS** so that extensions are shared across your team.

To remove extensions, use `ddev pi remove` (global) or `ddev pi remove -l` (project).

## Configuration

Default configurations are initialized in your project's `.pi/` directory during installation (`settings.json` is copied from the add-on defaults). Modify `settings.json` there to configure providers and agent preferences.

> **Note:** Re-running `ddev addon get` will not overwrite an existing `.pi/` directory. Defaults are only copied when `.pi/` is empty.

## Extending DDEV PI

DDEV PI exposes four extensibility seams so addon authors and project maintainers can customize or extend the environment without forking the base addon. Each seam has a different lifecycle, execution context, and activation cost.

### Seam overview

| Seam | Runs as | Activated by | Notes |
|---|---|---|---|
| `build.d/` | `root` (or `pi` via `sudo -u`) | Rebuild | Scripts execute during image build |
| `bashrc.d/` | `pi` (interactive) | Restart | Interactive shell only |
| `entrypoint.d/` | `pi` (with `sudo` available) | Restart | Container startup hooks |
| `global/agent/extensions/` | Pi agent | Restart | Independent command listeners |

### `build.d/` — build-time package installation

Place shell scripts here to install additional packages or make system-level changes during the Docker image build. Scripts run as `root` in lexicographic (sorted) order. Use numeric prefixes to control execution order and to avoid conflicts with existing scripts.

The naming bands are:

- `00–09`: core image setup (reserved)
- `10–49`: runtimes and infrastructure
- `50–89`: tools and utilities
- `90–99`: project-specific overrides

If you need to run a step as the `pi` user inside a build script, prefix the command with `sudo -u ${USERNAME}`.

> **Changes to `build.d/` require a full image rebuild** to take effect. Run `ddev debug rebuild` or equivalent.

### `bashrc.d/` — interactive shell ergonomics

Place scripts here to customize the interactive shell experience inside the PI container. Scripts are sourced in sorted order whenever a user enters the container (e.g., via `ddev ssh --service=pi`).

This seam is **for interactive shells only**. Scripts placed here do not affect PI's autonomous tool execution or entrypoint hooks.

> **A `ddev restart` is sufficient** to pick up changes. No rebuild is needed.

### `entrypoint.d/` — runtime startup hooks

Place scripts here to run initialization logic on every container startup. Scripts execute as the `pi` user in sorted order. `sudo` is available inside these scripts if you need root-level steps.

Hook failures emit a warning but do not crash the container, so a broken hook will not prevent the PI service from starting.

> **A `ddev restart` is sufficient** to pick up changes. No rebuild is needed.

### `global/agent/extensions/` — PI agent command routing

This directory is a mounted path where the PI agent discovers additional command extension files. Extensions are independent listeners and are not processed as an ordered pipeline, so **numeric prefixes have no effect** here.

To avoid filename collisions with other addons, use an addon-specific prefix in every filename, for example:

```
myaddon-summarize.ts
myaddon-review.ts
```

> **A `ddev restart` is sufficient** to pick up new or updated extensions. No rebuild is needed.

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
