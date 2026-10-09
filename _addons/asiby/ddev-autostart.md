---
title: "asiby/ddev-autostart"
github_url: "https://github.com/asiby/ddev-autostart"
description: "Automated system boot and startup management for your DDEV projects."
user: "asiby"
repo: "ddev-autostart"
repo_id: 1404894986
default_branch: "main"
tag_name: "v0.4.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-10-04"
updated_at: "2026-10-08"
workflow_status: "failure"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/asiby/ddev-autostart/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/asiby/ddev-autostart/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/asiby/ddev-autostart)](https://github.com/asiby/ddev-autostart/commits)
[![release](https://img.shields.io/github/v/release/asiby/ddev-autostart)](https://github.com/asiby/ddev-autostart/releases/latest)

# DDEV Autostart

## Overview

Adds a global `ddev autostart` command that starts your DDEV projects automatically when the computer boots. Each project is registered separately, so you choose exactly which ones come up on their own, or register everything you have running with a single `ddev autostart enable --running`.

## Requirements

- DDEV v1.24.10 or later
- Linux with systemd (see [Supported platforms](#supported-platforms))
- `sudo` rights, used once per change to install or remove a boot service
- Your user can run Docker without `sudo` (usually by being in the `docker` group)
- With rootless Docker or Podman, lingering turned on for your user (`sudo loginctl enable-linger $USER`), so they start at boot. `enable` tells you if it's off.

## Installation

Run this from inside any DDEV project (or add `--project <name>` from anywhere):

```bash
ddev add-on get asiby/ddev-autostart
```

No restart is needed. The command is installed into DDEV's global directory (`~/.ddev`).

To update to the latest version, run the same command again from the project you installed it from. If you don't remember which one, `ddev autostart status` tells you at the bottom.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev autostart enable [project...]` | Start projects automatically on boot |
| `ddev autostart enable --running` | Start every project that is running right now on boot |
| `ddev autostart disable [project...]` | Stop starting projects on boot (they keep running) |
| `ddev autostart disable --all` | Stop starting every project on boot |
| `ddev autostart status [project...]` | Show whether projects are registered and whether they started |
| `ddev autostart list` | Show every project and its autostart state |
| `ddev autostart --version` | Show the installed version (include it in bug reports) |

Tab completion works for the actions, their options and project names (after `disable`, only the projects you've registered).

Inside a project folder, the project name is detected automatically:

```bash
cd ~/code/my-site
ddev autostart enable
```

Anywhere else, give one or more project names:

```bash
ddev autostart enable site-a shop blog
```

If any name is wrong, nothing is changed. To make your current setup come back after every reboot, register whatever is running:

```bash
ddev autostart enable --running
```

`--running` asks DDEV which projects are running, so Docker must be up.

`ddev autostart list` gives an overview of all your projects:

```
PROJECT  AUTOSTART  SERVICE   DDEV     APPROOT
blog     disabled   -         stopped  /home/me/code/blog
shop     enabled    failed    stopped  /home/me/code/shop
site-a   enabled    active    running  /home/me/code/site-a
```

- **AUTOSTART**: `enabled`, `disabled`, or `orphaned` (registered, but the project was deleted or renamed)
- **SERVICE**: the boot service's state; `active` means it started successfully this boot
- **DDEV**: whether the project is running right now

`list` also prints a reminder for orphaned or failed registrations.

## Supported platforms

| Platform | Status |
| -------- | ------ |
| Linux with systemd (Ubuntu, Debian, Fedora, Arch, …) | ✅ Supported |
| Linux with OpenRC (Alpine, Gentoo) | Planned |
| macOS (launchd) | Planned |
| Windows (Task Scheduler) | Planned |

## How it works (Linux / systemd)

`enable` writes a system service at `/etc/systemd/system/ddev-autostart-<project>.service`, which is why it needs `sudo`. The service:

- runs `ddev start <project>` as your user, never as root;
- starts after Docker and waits up to about 2 minutes for it to respond;
- runs at boot without anyone needing to log in.

It runs only once, at boot. If Docker is stopped or restarted later, your projects stop with it and stay stopped until you run `ddev start`.

Registrations belong to the Linux user who made them. On a shared machine, other users' registrations are left alone by every command, including uninstalling.

`enable` doesn't start the project now, and `disable` doesn't stop it; they only change what happens at the next boot. To try a project's boot service without rebooting:

```bash
sudo systemctl start ddev-autostart-<project>.service
```

The service records where `ddev` and `docker` are installed, and adds only those folders to the standard system ones. If you move or reinstall either to a different location, run `ddev autostart enable <project>` again to update it. Normal upgrades (apt, Homebrew, `ddev self-upgrade`) keep the same location and need nothing.

## Troubleshooting

**A project didn't start at boot.** `ddev autostart status <project>` shows the last result, and the boot log shows why:

```bash
journalctl -u ddev-autostart-<project>.service -b
```

Common causes: Docker took more than about 2 minutes to start, your user isn't in the `docker` group, lingering is off with rootless Docker or Podman, or the project itself fails to start (try `ddev start <project>` by hand).

**`list` shows a project as `orphaned`.** The project was deleted or renamed while registered, so its boot service fails every time. Remove it with `ddev autostart disable <project>`.

## Uninstalling

Remove the add-on **from the same project you installed it from** (DDEV keeps the add-on's install record in that project). If you don't remember which one, `ddev autostart status` or `ddev autostart list` tells you at the bottom:

```bash
ddev add-on remove ddev-autostart
```

Before deleting the command, uninstalling removes the boot registration of every project, so nothing keeps starting on boot afterwards. It asks for your `sudo` password if needed. Where no password can be entered (for example in a script with no terminal), it leaves the registrations in place and prints the exact commands to remove them.

## Contributing

Each operating system is a plugin in `commands/host/autostart.d/plugins/`. A plugin defines four functions: `plugin_enable`, `plugin_disable`, `plugin_status` and `plugin_list_registered`. See `systemd.sh` for the reference implementation. Shell completion uses `plugin_list_registered` too, so a new plugin gets completion without extra work. Since it runs on every Tab press, it must list only the current user's registrations, never prompt or need `sudo`, and print nothing but its results.

Every file under `commands/host/` must contain `#ddev-generated` so DDEV can update and remove it.

Environment variables useful for development and testing:

| Variable | Effect |
| -------- | ------ |
| `DDEV_AUTOSTART_PLUGIN` | Force a plugin (e.g. `systemd`) instead of detecting the OS |
| `DDEV_AUTOSTART_UNIT_DIR` | Write systemd units to another folder instead of `/etc/systemd/system` |
| `DDEV_AUTOSTART_NONINTERACTIVE` | Never prompt for a `sudo` password; fail with manual steps instead |

The tests use [Bats](https://bats-core.readthedocs.io/): `bats ./tests/test.bats`. Where systemd is running and `sudo` needs no password, they register a real boot service for a temporary test project and remove it afterwards; elsewhere those checks are skipped.

## Credits

**Contributed and maintained by [@asiby](https://github.com/asiby)**
