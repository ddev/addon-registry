---
title: prioris-dev/ddev-socket-firewall
github_url: https://github.com/prioris-dev/ddev-socket-firewall
description: "DDEV add-on that installs Socket Firewall (sfw) into the web container and routes npm, yarn, pnpm and Composer through it — blocking malicious dependencies in real time, both via the ddev shortcuts and inside ddev ssh."
user: prioris-dev
repo: ddev-socket-firewall
repo_id: 1274673114
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-06-19
updated_at: 2026-06-19
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/prioris-dev/ddev-socket-firewall/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/prioris-dev/ddev-socket-firewall/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/prioris-dev/ddev-socket-firewall)](https://github.com/prioris-dev/ddev-socket-firewall/commits)
[![release](https://img.shields.io/github/v/release/prioris-dev/ddev-socket-firewall)](https://github.com/prioris-dev/ddev-socket-firewall/releases/latest)

# DDEV Socket Firewall

## Overview

This add-on installs [Socket Firewall Free](https://github.com/SocketDev/sfw-free) (`sfw`) into your
[DDEV](https://ddev.com/) project's **web container** and transparently routes your package-manager
commands through it. Socket Firewall blocks known-malicious dependencies in real time, before they
are ever written to disk.

Once installed, the following are wrapped by `sfw` — both **outside** the container (via the `ddev`
shortcuts) and **inside** it (in `ddev ssh` interactive shells):

- `npm`, `yarn`, `pnpm`
- `composer` *(see the note below)*

No API key and no configuration are required.

## Installation

```bash
ddev add-on get prioris-dev/ddev-socket-firewall
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

Use your package managers exactly as before — they are now protected automatically:

| Command | What happens |
| ------- | ------------ |
| `ddev npm install` | Runs `sfw npm install` in the web container |
| `ddev yarn add lodash` | Runs `sfw yarn add lodash` |
| `ddev pnpm install` | Runs `sfw pnpm install` |
| `ddev scomposer install` (alias `ddev sco`) | Runs `sfw composer install` — host-callable Composer through the firewall |
| `ddev sfw <pm> [args]` | Generic passthrough — run any package manager through `sfw` (supports `--verbose`) |
| `ddev ssh` → `npm install` | Aliased to `sfw npm install` in the interactive shell |
| `ddev exec which sfw` | Confirms `sfw` is installed (`/usr/local/bin/sfw`) |

## Debugging

The dedicated shortcuts (`ddev npm`, `ddev sco`, …) run `sfw` quietly. To see Socket Firewall's
diagnostic output (banner, package counts, etc.), use the generic command with `--verbose`:

```bash
ddev sfw --verbose npm install
ddev sfw --verbose composer install
```

You can also set `SFW_VERBOSE=true` in the container environment.

## Notes & limitations

- **Composer is wrapped, but not yet blocked.** Socket Firewall Free's documented, actively-protected
  ecosystems are npm/yarn/pnpm, pip/uv and cargo. `sfw` *runs* any package-manager command, so
  `sfw composer install` works today — but real-time blocking for Packagist is not active yet. The
  wrapper is forward-looking: it starts protecting automatically once Socket enables Composer support.
- **`ddev composer` (host shortcut) is not routed through `sfw`.** `ddev composer` is a built-in DDEV
  Go command and cannot be overridden by a custom command, so it calls Composer directly. For a
  host-callable Composer that *does* go through the firewall, use **`ddev scomposer`** (alias
  **`ddev sco`**) instead — e.g. `ddev sco install`. Inside the container (`ddev ssh`) the `composer`
  alias also routes through `sfw`. Given Composer is not blocked yet (see above), this has no
  functional impact today.

## Advanced Customization

The `sfw` binary is **pinned to a specific version** and its download is **verified against the
published SHA256 digest** for reproducible, tamper-evident builds. The pin lives in
`.ddev/web-build/Dockerfile.socket-firewall`:

```dockerfile
ARG SOCKET_FIREWALL_VERSION=v1.12.0
ARG SOCKET_FIREWALL_SHA256_AMD64=51824f02a242f892c61c42223e05b7e82bb624762337f026afc2ac229ffcade7
ARG SOCKET_FIREWALL_SHA256_ARM64=598c8d19c80832ef5ca7fdb0a9fc35c045847fd02889126a7115c0345a478da3
```

### Upgrading (or changing) the `sfw` version

The version and its two SHA256 checksums are coupled: changing the version **requires** updating both
checksums in the same file, otherwise the build aborts on a checksum mismatch (this is intentional).
Follow these steps:

**1. Pick a version** from the [`sfw-free` releases page](https://github.com/SocketDev/sfw-free/releases),
e.g. `v1.13.0`.

**2. Fetch its Linux checksums** (needs the [`gh` CLI](https://cli.github.com/)):

```bash
gh api repos/SocketDev/sfw-free/releases/tags/v1.13.0 \
  --jq '.assets[] | select(.name|test("linux-(x86_64|arm64)$")) | "\(.name) \(.digest)"'
```

Example output:

```
sfw-free-linux-arm64 sha256:598c8d19...
sfw-free-linux-x86_64 sha256:51824f02...
```

> No `gh`? The same digests are shown on the release page, or compute them yourself with
> `curl -fsSL <asset-url> | sha256sum`.

**3. Edit `.ddev/web-build/Dockerfile.socket-firewall`** — update all three `ARG` lines
(`x86_64` → `AMD64`, `arm64` → `ARM64`, drop the `sha256:` prefix):

```dockerfile
ARG SOCKET_FIREWALL_VERSION=v1.13.0
ARG SOCKET_FIREWALL_SHA256_AMD64=51824f02...
ARG SOCKET_FIREWALL_SHA256_ARM64=598c8d19...
```

**4. Rebuild and verify:**

```bash
ddev restart                         # rebuilds the web image
ddev exec sfw --verbose npm --version   # confirms sfw runs
```

If the checksum is wrong, `ddev restart` fails with `sha256sum: WARNING: 1 computed checksum did NOT
match` — fix the hash and retry.

**Shortcut (skip verification):** if you only want to change the version and not deal with checksums,
set the relevant architecture's checksum to an empty string — the build then warns but proceeds:

```dockerfile
ARG SOCKET_FIREWALL_VERSION=v1.13.0
ARG SOCKET_FIREWALL_SHA256_AMD64=""
ARG SOCKET_FIREWALL_SHA256_ARM64=""
```

**Automation:** [Renovate](https://docs.renovatebot.com/) (configured via `renovate.json`) opens a PR
when a new version ships. The PR's build fails on the now-stale checksums until you update the two
SHA256 lines — an intentional integrity gate that forces a human to confirm the new binary.

## Credits

**Contributed and maintained by [@prioris-dev](https://github.com/prioris-dev)**
