---
title: mxr576/ddev-pi
github_url: https://github.com/mxr576/ddev-pi
description: "DDEV addon for pi.dev, which is a minimal terminal coding harness."
user: mxr576
repo: ddev-pi
repo_id: 1250983207
default_branch: main
tag_name: 1.0.0-alpha1
ddev_version_constraint: ">= v1.25.2"
dependencies: ["trebormc/ddev-ai-ssh"]
type: contrib
created_at: 2026-05-27
updated_at: 2026-05-27
workflow_status: failure
stars: 0
---

# DDEV Add-on for Pi Coding Agent

This DDEV add-on provides a fully containerized, isolated environment for the [Pi Coding Agent](https://pi.dev).

## Architecture & Security
- **Full Isolation:** The agent runs in a dedicated sidecar container. We explicitly do *not* mount your host machine's `~/.pi` directory, ensuring that agent state is scoped strictly to the project and avoids cross-contamination with your host environment.
- **Project Access:** The project root is mounted at `/workspace` with read/write permissions, allowing Pi to interact with your codebase.
- **Network:** The agent is isolated within the DDEV project network.
- **Sensitive Data:** No host-level sensitive directories (e.g., `~/.ssh`) are mounted.

## Installation
```bash
ddev addon get mxr576/ddev-pi
ddev restart
```

## Usage
The Pi agent runs as a DDEV service. To interact with it:

1. Enable the profile: `ddev start --profiles=pi`
2. Run commands: `ddev pi [command]`

Example:
```bash
ddev pi status
```

## Managing Extensions
Pi supports two types of extensions:

- **Global Extensions:** Use `ddev pi install <addon-ref>`. These are stored in `.ddev/pi/global/` and remain local to your environment. Do not commit these to VCS.
- **Project Extensions:** Use `ddev pi install -l <addon-ref>`. These are stored in the project's `.pi/` directory. **Commit this directory to your version control system** so that extensions are shared across your team.

To remove extensions, use the `ddev pi remove` or `ddev pi remove -l` commands respectively.

## Configuration
Default configurations are initialized in your project's `.pi/` directory. You may modify `settings.json` there to configure providers and agent preferences.

## Credits

Contributed and maintained by [@mxr576](https://github.com/mxr576). The initial proof of concept and early development were sponsored by [Pronovix](https://pronovix.com).
